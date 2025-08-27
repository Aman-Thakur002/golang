/*
=============================================================================
                           🎯 GO CONTEXT TUTORIAL
=============================================================================

📚 CORE CONCEPT:
Context carries deadlines, cancellation signals, and request-scoped values
across API boundaries. It's essential for managing goroutine lifecycles
and request timeouts.

🔑 KEY FEATURES:
• Cancellation propagation
• Timeout and deadline management
• Request-scoped value passing
• Graceful shutdown patterns

💡 REAL-WORLD ANALOGY:
Context = Project Manager
- Cancellation = "Stop all work on this project"
- Timeout = "This project has a deadline"
- Values = "Here's the project information everyone needs"
- Done channel = "Project status updates"

🎯 WHY USE CONTEXT?
• Cancel long-running operations
• Set timeouts for network requests
• Pass request-scoped data (user ID, trace ID)
• Coordinate goroutine cleanup

=============================================================================
*/

package main

import (
	"context"
	"fmt"
	"math/rand"
	"time"
)

// 🔑 CONTEXT KEYS: Type-safe keys for context values
type contextKey string

const (
	userIDKey    contextKey = "userID"
	requestIDKey contextKey = "requestID"
)

// 🎯 DEMO FUNCTIONS: Simulate different operations

// 📡 NETWORK REQUEST: Simulates HTTP request with timeout
func fetchData(ctx context.Context, url string) (string, error) {
	// Create a channel to receive the result
	resultChan := make(chan string, 1)
	errorChan := make(chan error, 1)

	// Start the "network request" in a goroutine
	go func() {
		// Simulate network delay
		delay := time.Duration(rand.Intn(3000)) * time.Millisecond
		fmt.Printf("📡 Fetching %s (will take %v)\n", url, delay)
		
		time.Sleep(delay)
		
		// Simulate successful response
		resultChan <- fmt.Sprintf("Data from %s", url)
	}()

	// Wait for either completion or cancellation
	select {
	case result := <-resultChan:
		return result, nil
	case err := <-errorChan:
		return "", err
	case <-ctx.Done():
		return "", ctx.Err() // Returns context.Canceled or context.DeadlineExceeded
	}
}

// 🔄 WORKER FUNCTION: Long-running task that respects cancellation
func worker(ctx context.Context, id int) {
	fmt.Printf("🔄 Worker %d starting\n", id)
	
	for i := 0; i < 10; i++ {
		select {
		case <-ctx.Done():
			fmt.Printf("🔄 Worker %d cancelled at step %d: %v\n", id, i, ctx.Err())
			return
		default:
			fmt.Printf("🔄 Worker %d: step %d\n", id, i)
			time.Sleep(500 * time.Millisecond)
		}
	}
	
	fmt.Printf("🔄 Worker %d completed\n", id)
}

// 📊 DATABASE QUERY: Simulates database operation with context
func queryDatabase(ctx context.Context, query string) ([]string, error) {
	// Extract user ID from context
	userID, ok := ctx.Value(userIDKey).(string)
	if !ok {
		userID = "unknown"
	}
	
	fmt.Printf("📊 Executing query for user %s: %s\n", userID, query)
	
	// Simulate database processing time
	select {
	case <-time.After(1 * time.Second):
		return []string{"result1", "result2", "result3"}, nil
	case <-ctx.Done():
		fmt.Printf("📊 Database query cancelled: %v\n", ctx.Err())
		return nil, ctx.Err()
	}
}

// 🌐 HTTP REQUEST HANDLER: Simulates web request processing
func handleRequest(ctx context.Context) {
	// Add request ID to context
	requestID := fmt.Sprintf("req-%d", rand.Intn(10000))
	ctx = context.WithValue(ctx, requestIDKey, requestID)
	
	fmt.Printf("🌐 Handling request %s\n", requestID)
	
	// Simulate multiple operations
	results, err := queryDatabase(ctx, "SELECT * FROM users")
	if err != nil {
		fmt.Printf("🌐 Request %s failed: %v\n", requestID, err)
		return
	}
	
	fmt.Printf("🌐 Request %s completed with %d results\n", requestID, len(results))
}

func main() {
	fmt.Println("🎯 CONTEXT TUTORIAL")
	fmt.Println("===================")

	// 🎯 DEMO 1: Basic Context with Cancellation
	fmt.Println("\n🎯 DEMO 1: Basic Cancellation")
	fmt.Println("=============================")

	// Create a cancellable context
	ctx, cancel := context.WithCancel(context.Background())
	
	// Start a worker
	go worker(ctx, 1)
	
	// Let it run for a bit, then cancel
	time.Sleep(2 * time.Second)
	fmt.Println("🚫 Cancelling context...")
	cancel()
	
	// Give time to see the cancellation
	time.Sleep(1 * time.Second)

	// 🎯 DEMO 2: Context with Timeout
	fmt.Println("\n🎯 DEMO 2: Context with Timeout")
	fmt.Println("===============================")

	// Create context with 2-second timeout
	ctx, cancel = context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel() // Always call cancel to free resources

	// Try to fetch data (might timeout)
	data, err := fetchData(ctx, "https://api.example.com/data")
	if err != nil {
		fmt.Printf("❌ Fetch failed: %v\n", err)
	} else {
		fmt.Printf("✅ Fetch succeeded: %s\n", data)
	}

	// 🎯 DEMO 3: Context with Deadline
	fmt.Println("\n🎯 DEMO 3: Context with Deadline")
	fmt.Println("================================")

	// Create context with specific deadline
	deadline := time.Now().Add(1500 * time.Millisecond)
	ctx, cancel = context.WithDeadline(context.Background(), deadline)
	defer cancel()

	fmt.Printf("⏰ Deadline set for: %v\n", deadline.Format("15:04:05.000"))

	// Try another fetch
	data, err = fetchData(ctx, "https://api.example.com/slow-endpoint")
	if err != nil {
		fmt.Printf("❌ Fetch failed: %v\n", err)
	} else {
		fmt.Printf("✅ Fetch succeeded: %s\n", data)
	}

	// 🎯 DEMO 4: Context with Values
	fmt.Println("\n🎯 DEMO 4: Context with Values")
	fmt.Println("==============================")

	// Create context with user information
	ctx = context.WithValue(context.Background(), userIDKey, "user123")
	ctx = context.WithValue(ctx, requestIDKey, "req456")

	// Use context in request handling
	handleRequest(ctx)

	// 🎯 DEMO 5: Multiple Workers with Shared Cancellation
	fmt.Println("\n🎯 DEMO 5: Multiple Workers")
	fmt.Println("===========================")

	ctx, cancel = context.WithCancel(context.Background())

	// Start multiple workers
	for i := 1; i <= 3; i++ {
		go worker(ctx, i)
	}

	// Let them run, then cancel all
	time.Sleep(3 * time.Second)
	fmt.Println("🚫 Cancelling all workers...")
	cancel()

	// Wait for cleanup
	time.Sleep(1 * time.Second)

	// 🎯 DEMO 6: Context Chain (Parent-Child Relationship)
	fmt.Println("\n🎯 DEMO 6: Context Chain")
	fmt.Println("========================")

	// Parent context with timeout
	parentCtx, parentCancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer parentCancel()

	// Child context with shorter timeout
	childCtx, childCancel := context.WithTimeout(parentCtx, 2*time.Second)
	defer childCancel()

	// The child will timeout first
	fmt.Println("⏰ Parent timeout: 5s, Child timeout: 2s")
	
	start := time.Now()
	data, err = fetchData(childCtx, "https://api.example.com/data")
	duration := time.Since(start)
	
	if err != nil {
		fmt.Printf("❌ Child context failed after %v: %v\n", duration, err)
	} else {
		fmt.Printf("✅ Child context succeeded after %v: %s\n", duration, data)
	}

	// 🎯 DEMO 7: Context Best Practices
	fmt.Println("\n🎯 DEMO 7: Best Practices")
	fmt.Println("=========================")

	// ✅ GOOD: Always check context in loops
	ctx, cancel = context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	fmt.Println("✅ Demonstrating proper context checking in loop:")
	for i := 0; i < 10; i++ {
		select {
		case <-ctx.Done():
			fmt.Printf("✅ Loop cancelled at iteration %d: %v\n", i, ctx.Err())
			goto cleanup
		default:
			fmt.Printf("✅ Loop iteration %d\n", i)
			time.Sleep(500 * time.Millisecond)
		}
	}

cleanup:
	// 🎯 DEMO 8: Context Error Types
	fmt.Println("\n🎯 DEMO 8: Context Error Types")
	fmt.Println("==============================")

	// Cancelled context
	ctx, cancel = context.WithCancel(context.Background())
	cancel() // Cancel immediately

	if ctx.Err() == context.Canceled {
		fmt.Println("✅ Context was cancelled")
	}

	// Timeout context
	ctx, cancel = context.WithTimeout(context.Background(), 1*time.Nanosecond)
	defer cancel()
	time.Sleep(1 * time.Millisecond) // Ensure timeout

	if ctx.Err() == context.DeadlineExceeded {
		fmt.Println("✅ Context deadline exceeded")
	}

	fmt.Println("\n✨ All context demos completed!")
}

/*
=============================================================================
                              📝 LEARNING NOTES
=============================================================================

🎯 CONTEXT TYPES:
┌─────────────────────────────────────────────────────────────────────────┐
│ // Background context (root)                                            │
│ ctx := context.Background()                                             │
│                                                                         │
│ // TODO context (placeholder)                                           │
│ ctx := context.TODO()                                                   │
│                                                                         │
│ // Cancellable context                                                  │
│ ctx, cancel := context.WithCancel(parent)                               │
│                                                                         │
│ // Context with timeout                                                 │
│ ctx, cancel := context.WithTimeout(parent, 5*time.Second)               │
│                                                                         │
│ // Context with deadline                                                │
│ deadline := time.Now().Add(5*time.Second)                               │
│ ctx, cancel := context.WithDeadline(parent, deadline)                   │
│                                                                         │
│ // Context with value                                                   │
│ ctx := context.WithValue(parent, key, value)                            │
└─────────────────────────────────────────────────────────────────────────┘

🔄 CONTEXT USAGE PATTERNS:
┌─────────────────────────────────────────────────────────────────────────┐
│ // Function signature                                                   │
│ func DoWork(ctx context.Context, data string) error {                   │
│     // Always accept context as first parameter                        │
│ }                                                                       │
│                                                                         │
│ // Check for cancellation                                               │
│ select {                                                                │
│ case <-ctx.Done():                                                      │
│     return ctx.Err()                                                    │
│ default:                                                                │
│     // Continue work                                                    │
│ }                                                                       │
│                                                                         │
│ // In loops                                                             │
│ for i := 0; i < n; i++ {                                                │
│     select {                                                            │
│     case <-ctx.Done():                                                  │
│         return ctx.Err()                                                │
│     default:                                                            │
│         // Do work                                                      │
│     }                                                                   │
│ }                                                                       │
└─────────────────────────────────────────────────────────────────────────┘

🔑 CONTEXT VALUES:
┌─────────────────────────────────────────────────────────────────────────┐
│ // Define typed keys                                                    │
│ type contextKey string                                                  │
│ const userIDKey contextKey = "userID"                                   │
│                                                                         │
│ // Set value                                                            │
│ ctx := context.WithValue(parent, userIDKey, "user123")                  │
│                                                                         │
│ // Get value                                                            │
│ if userID, ok := ctx.Value(userIDKey).(string); ok {                    │
│     // Use userID                                                       │
│ }                                                                       │
│                                                                         │
│ // Helper function                                                      │
│ func GetUserID(ctx context.Context) string {                            │
│     if userID, ok := ctx.Value(userIDKey).(string); ok {                │
│         return userID                                                   │
│     }                                                                   │
│     return ""                                                           │
│ }                                                                       │
└─────────────────────────────────────────────────────────────────────────┘

⚡ CONTEXT ERRORS:
┌─────────────────┬─────────────────────────────────────────────────────────┐
│     Error       │                Description                              │
├─────────────────┼─────────────────────────────────────────────────────────┤
│ context.Canceled│ Context was cancelled via cancel()                      │
│ context.DeadlineExceeded │ Context timeout/deadline reached          │
└─────────────────┴─────────────────────────────────────────────────────────┘

🚨 COMMON MISTAKES:
❌ Not calling cancel() (resource leak)
❌ Using context.Value for non-request-scoped data
❌ Not checking ctx.Done() in long operations
❌ Passing nil context (use context.Background())
❌ Using context for optional parameters

💡 BEST PRACTICES:
┌─────────────────────────────────────────────────────────────────────────┐
│ // Always accept context as first parameter                             │
│ func MyFunction(ctx context.Context, other params) error                │
│                                                                         │
│ // Always call cancel to free resources                                 │
│ ctx, cancel := context.WithTimeout(parent, timeout)                     │
│ defer cancel()                                                          │
│                                                                         │
│ // Check cancellation in long operations                                │
│ for {                                                                   │
│     select {                                                            │
│     case <-ctx.Done():                                                  │
│         return ctx.Err()                                                │
│     default:                                                            │
│         // Do work                                                      │
│     }                                                                   │
│ }                                                                       │
│                                                                         │
│ // Use typed keys for context values                                    │
│ type contextKey string                                                  │
│ const myKey contextKey = "myKey"                                        │
│                                                                         │
│ // Don't store contexts in structs                                      │
│ // Pass them as function parameters instead                             │
└─────────────────────────────────────────────────────────────────────────┘

🎯 REAL-WORLD PATTERNS:
┌─────────────────────────────────────────────────────────────────────────┐
│ // HTTP request with timeout                                            │
│ func makeRequest(ctx context.Context, url string) (*http.Response, error) { │
│     req, err := http.NewRequestWithContext(ctx, "GET", url, nil)         │
│     if err != nil {                                                     │
│         return nil, err                                                 │
│     }                                                                   │
│     return http.DefaultClient.Do(req)                                   │
│ }                                                                       │
│                                                                         │
│ // Database query with context                                          │
│ func queryDB(ctx context.Context, query string) (*sql.Rows, error) {    │
│     return db.QueryContext(ctx, query)                                  │
│ }                                                                       │
│                                                                         │
│ // Graceful shutdown                                                    │
│ func (s *Server) Shutdown(ctx context.Context) error {                  │
│     // Stop accepting new requests                                      │
│     s.listener.Close()                                                  │
│                                                                         │
│     // Wait for existing requests to complete                           │
│     done := make(chan struct{})                                         │
│     go func() {                                                         │
│         s.wg.Wait()                                                     │
│         close(done)                                                     │
│     }()                                                                 │
│                                                                         │
│     select {                                                            │
│     case <-done:                                                        │
│         return nil                                                      │
│     case <-ctx.Done():                                                  │
│         return ctx.Err()                                                │
│     }                                                                   │
│ }                                                                       │
└─────────────────────────────────────────────────────────────────────────┘

🎯 WHEN TO USE EACH CONTEXT TYPE:
• Background(): Root context for main, init, tests
• TODO(): When you're not sure which context to use
• WithCancel(): When you need to cancel operations
• WithTimeout(): When operations have time limits
• WithDeadline(): When you have a specific end time
• WithValue(): For request-scoped data (user ID, trace ID)

⚡ PERFORMANCE CONSIDERATIONS:
• Context operations are fast (nanoseconds)
• Context values use interface{} (type assertions needed)
• Deep context chains can impact performance
• Don't overuse context values
• Prefer passing data as function parameters when possible

=============================================================================
*/