/*
=============================================================================
                        🌐 GO HTTP SERVER TUTORIAL
=============================================================================

📚 CORE CONCEPT:
HTTP servers in Go handle incoming requests and send responses. The net/http
package provides everything needed to build robust web servers and APIs.

🔑 KEY FEATURES:
• Simple HTTP server setup
• Route handling and multiplexing
• Middleware support
• JSON API endpoints
• Static file serving

💡 REAL-WORLD ANALOGY:
HTTP Server = Restaurant
- Routes = Menu items (what you can order)
- Handlers = Kitchen staff (who prepares your order)
- Middleware = Waiters (who process requests before/after)
- Response = Your prepared meal

🎯 WHY LEARN HTTP SERVERS?
• Build REST APIs and web services
• Create microservices
• Serve web applications
• Handle real-time communication

=============================================================================
*/

package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"
)

// 📊 DATA STRUCTURES: For API responses
type User struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

type APIResponse struct {
	Success bool        `json:"success"`
	Data    interface{} `json:"data,omitempty"`
	Message string      `json:"message,omitempty"`
	Error   string      `json:"error,omitempty"`
}

// 💾 IN-MEMORY DATA STORE: Simple storage for demo
var users = []User{
	{ID: 1, Name: "John Doe", Email: "john@example.com"},
	{ID: 2, Name: "Jane Smith", Email: "jane@example.com"},
	{ID: 3, Name: "Bob Johnson", Email: "bob@example.com"},
}

var nextUserID = 4

// 🎯 BASIC HANDLERS: Simple request handlers
func homeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "🏠 Welcome to Go HTTP Server!\n")
	fmt.Fprintf(w, "📅 Current time: %s\n", time.Now().Format(time.RFC3339))
	fmt.Fprintf(w, "🔗 Available endpoints:\n")
	fmt.Fprintf(w, "  GET  /users     - List all users\n")
	fmt.Fprintf(w, "  GET  /users/1   - Get user by ID\n")
	fmt.Fprintf(w, "  POST /users     - Create new user\n")
	fmt.Fprintf(w, "  PUT  /users/1   - Update user\n")
	fmt.Fprintf(w, "  DELETE /users/1 - Delete user\n")
}

func aboutHandler(w http.ResponseWriter, r *http.Request) {
	response := APIResponse{
		Success: true,
		Data: map[string]string{
			"name":        "Go HTTP Server Tutorial",
			"version":     "1.0.0",
			"description": "Learning HTTP server development in Go",
		},
	}
	
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

// 👥 USER HANDLERS: CRUD operations for users
func usersHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	
	switch r.Method {
	case http.MethodGet:
		handleGetUsers(w, r)
	case http.MethodPost:
		handleCreateUser(w, r)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func userHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	
	// Extract user ID from URL path
	path := strings.TrimPrefix(r.URL.Path, "/users/")
	userID, err := strconv.Atoi(path)
	if err != nil {
		response := APIResponse{
			Success: false,
			Error:   "Invalid user ID",
		}
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(response)
		return
	}
	
	switch r.Method {
	case http.MethodGet:
		handleGetUser(w, r, userID)
	case http.MethodPut:
		handleUpdateUser(w, r, userID)
	case http.MethodDelete:
		handleDeleteUser(w, r, userID)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func handleGetUsers(w http.ResponseWriter, r *http.Request) {
	response := APIResponse{
		Success: true,
		Data:    users,
		Message: fmt.Sprintf("Found %d users", len(users)),
	}
	json.NewEncoder(w).Encode(response)
}

func handleGetUser(w http.ResponseWriter, r *http.Request, userID int) {
	for _, user := range users {
		if user.ID == userID {
			response := APIResponse{
				Success: true,
				Data:    user,
			}
			json.NewEncoder(w).Encode(response)
			return
		}
	}
	
	response := APIResponse{
		Success: false,
		Error:   "User not found",
	}
	w.WriteHeader(http.StatusNotFound)
	json.NewEncoder(w).Encode(response)
}

func handleCreateUser(w http.ResponseWriter, r *http.Request) {
	var newUser User
	err := json.NewDecoder(r.Body).Decode(&newUser)
	if err != nil {
		response := APIResponse{
			Success: false,
			Error:   "Invalid JSON data",
		}
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(response)
		return
	}
	
	// Assign new ID
	newUser.ID = nextUserID
	nextUserID++
	
	// Add to users slice
	users = append(users, newUser)
	
	response := APIResponse{
		Success: true,
		Data:    newUser,
		Message: "User created successfully",
	}
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(response)
}

func handleUpdateUser(w http.ResponseWriter, r *http.Request, userID int) {
	var updatedUser User
	err := json.NewDecoder(r.Body).Decode(&updatedUser)
	if err != nil {
		response := APIResponse{
			Success: false,
			Error:   "Invalid JSON data",
		}
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(response)
		return
	}
	
	// Find and update user
	for i, user := range users {
		if user.ID == userID {
			updatedUser.ID = userID // Preserve ID
			users[i] = updatedUser
			
			response := APIResponse{
				Success: true,
				Data:    updatedUser,
				Message: "User updated successfully",
			}
			json.NewEncoder(w).Encode(response)
			return
		}
	}
	
	response := APIResponse{
		Success: false,
		Error:   "User not found",
	}
	w.WriteHeader(http.StatusNotFound)
	json.NewEncoder(w).Encode(response)
}

func handleDeleteUser(w http.ResponseWriter, r *http.Request, userID int) {
	for i, user := range users {
		if user.ID == userID {
			// Remove user from slice
			users = append(users[:i], users[i+1:]...)
			
			response := APIResponse{
				Success: true,
				Message: "User deleted successfully",
			}
			json.NewEncoder(w).Encode(response)
			return
		}
	}
	
	response := APIResponse{
		Success: false,
		Error:   "User not found",
	}
	w.WriteHeader(http.StatusNotFound)
	json.NewEncoder(w).Encode(response)
}

// 🔧 MIDDLEWARE: Functions that wrap handlers
func loggingMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		
		// Call the next handler
		next(w, r)
		
		// Log the request
		duration := time.Since(start)
		log.Printf("📝 %s %s - %v", r.Method, r.URL.Path, duration)
	}
}

func corsMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Set CORS headers
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
		
		// Handle preflight requests
		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusOK)
			return
		}
		
		// Call the next handler
		next(w, r)
	}
}

func authMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Simple API key authentication (for demo)
		apiKey := r.Header.Get("X-API-Key")
		if apiKey != "demo-api-key" && r.URL.Path != "/" && r.URL.Path != "/about" {
			response := APIResponse{
				Success: false,
				Error:   "Invalid or missing API key",
			}
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusUnauthorized)
			json.NewEncoder(w).Encode(response)
			return
		}
		
		// Call the next handler
		next(w, r)
	}
}

// 🎯 CUSTOM MULTIPLEXER: Route handling
func setupRoutes() *http.ServeMux {
	mux := http.NewServeMux()
	
	// Apply middleware to handlers
	mux.HandleFunc("/", corsMiddleware(loggingMiddleware(authMiddleware(homeHandler))))
	mux.HandleFunc("/about", corsMiddleware(loggingMiddleware(aboutHandler)))
	mux.HandleFunc("/users", corsMiddleware(loggingMiddleware(authMiddleware(usersHandler))))
	
	// Handle user-specific routes
	mux.HandleFunc("/users/", corsMiddleware(loggingMiddleware(authMiddleware(userHandler))))
	
	return mux
}

func main() {
	fmt.Println("🌐 HTTP SERVER TUTORIAL")
	fmt.Println("=======================")

	// 🎯 DEMO: Complete HTTP Server
	fmt.Println("\n🎯 Starting HTTP Server")
	fmt.Println("=======================")

	// Setup routes
	mux := setupRoutes()
	
	// Create server with custom configuration
	server := &http.Server{
		Addr:         ":8080",
		Handler:      mux,
		ReadTimeout:  15 * time.Second,
		WriteTimeout: 15 * time.Second,
		IdleTimeout:  60 * time.Second,
	}

	fmt.Println("🚀 Server starting on http://localhost:8080")
	fmt.Println("📋 Available endpoints:")
	fmt.Println("  GET    http://localhost:8080/")
	fmt.Println("  GET    http://localhost:8080/about")
	fmt.Println("  GET    http://localhost:8080/users")
	fmt.Println("  POST   http://localhost:8080/users")
	fmt.Println("  GET    http://localhost:8080/users/1")
	fmt.Println("  PUT    http://localhost:8080/users/1")
	fmt.Println("  DELETE http://localhost:8080/users/1")
	fmt.Println()
	fmt.Println("🔑 API Key required for protected endpoints: X-API-Key: demo-api-key")
	fmt.Println()
	fmt.Println("📝 Example curl commands:")
	fmt.Println(`  curl http://localhost:8080/`)
	fmt.Println(`  curl -H "X-API-Key: demo-api-key" http://localhost:8080/users`)
	fmt.Println(`  curl -X POST -H "X-API-Key: demo-api-key" -H "Content-Type: application/json" \`)
	fmt.Println(`       -d '{"name":"New User","email":"new@example.com"}' \`)
	fmt.Println(`       http://localhost:8080/users`)
	fmt.Println()
	fmt.Println("⏹️  Press Ctrl+C to stop the server")

	// Start server
	log.Fatal(server.ListenAndServe())
}

/*
=============================================================================
                              📝 LEARNING NOTES
=============================================================================

🌐 HTTP SERVER BASICS:
┌─────────────────────────────────────────────────────────────────────────┐
│ // Simple server                                                        │
│ http.HandleFunc("/", handler)                                           │
│ log.Fatal(http.ListenAndServe(":8080", nil))                            │
│                                                                         │
│ // Custom server                                                        │
│ server := &http.Server{                                                 │
│     Addr:    ":8080",                                                   │
│     Handler: mux,                                                       │
│     ReadTimeout:  15 * time.Second,                                     │
│     WriteTimeout: 15 * time.Second,                                     │
│ }                                                                       │
│ log.Fatal(server.ListenAndServe())                                      │
└─────────────────────────────────────────────────────────────────────────┘

🎯 HANDLER PATTERNS:
┌─────────────────────────────────────────────────────────────────────────┐
│ // Function handler                                                     │
│ func handler(w http.ResponseWriter, r *http.Request) {                   │
│     fmt.Fprintf(w, "Hello, World!")                                     │
│ }                                                                       │
│                                                                         │
│ // Method-based routing                                                 │
│ func apiHandler(w http.ResponseWriter, r *http.Request) {               │
│     switch r.Method {                                                   │
│     case http.MethodGet:                                                │
│         // Handle GET                                                   │
│     case http.MethodPost:                                               │
│         // Handle POST                                                  │
│     default:                                                            │
│         http.Error(w, "Method not allowed", 405)                        │
│     }                                                                   │
│ }                                                                       │
└─────────────────────────────────────────────────────────────────────────┘

📊 HTTP RESPONSE PATTERNS:
┌─────────────────────────────────────────────────────────────────────────┐
│ // Plain text response                                                  │
│ fmt.Fprintf(w, "Hello, World!")                                         │
│                                                                         │
│ // JSON response                                                        │
│ w.Header().Set("Content-Type", "application/json")                      │
│ json.NewEncoder(w).Encode(data)                                         │
│                                                                         │
│ // Error response                                                       │
│ http.Error(w, "Not found", http.StatusNotFound)                         │
│                                                                         │
│ // Custom status code                                                   │
│ w.WriteHeader(http.StatusCreated)                                       │
│ json.NewEncoder(w).Encode(response)                                     │
└─────────────────────────────────────────────────────────────────────────┘

🔧 MIDDLEWARE PATTERN:
┌─────────────────────────────────────────────────────────────────────────┐
│ func middleware(next http.HandlerFunc) http.HandlerFunc {               │
│     return func(w http.ResponseWriter, r *http.Request) {               │
│         // Before request processing                                    │
│         log.Printf("Request: %s %s", r.Method, r.URL.Path)              │
│                                                                         │
│         // Call next handler                                            │
│         next(w, r)                                                      │
│                                                                         │
│         // After request processing                                     │
│         log.Println("Request completed")                                │
│     }                                                                   │
│ }                                                                       │
│                                                                         │
│ // Usage                                                                │
│ http.HandleFunc("/api", middleware(apiHandler))                         │
└─────────────────────────────────────────────────────────────────────────┘

🗂️ ROUTING STRATEGIES:
┌─────────────────────────────────────────────────────────────────────────┐
│ // Built-in ServeMux                                                    │
│ mux := http.NewServeMux()                                               │
│ mux.HandleFunc("/users", usersHandler)                                  │
│ mux.HandleFunc("/users/", userHandler) // Trailing slash for subtree    │
│                                                                         │
│ // Pattern matching                                                     │
│ mux.HandleFunc("/api/v1/", apiV1Handler)                                │
│ mux.HandleFunc("/static/", http.StripPrefix("/static/",                 │
│     http.FileServer(http.Dir("./static/"))))                            │
│                                                                         │
│ // Third-party routers (gorilla/mux, chi, gin)                         │
│ // r := mux.NewRouter()                                                 │
│ // r.HandleFunc("/users/{id}", getUserHandler).Methods("GET")           │
└─────────────────────────────────────────────────────────────────────────┘

📝 REQUEST HANDLING:
┌─────────────────────────────────────────────────────────────────────────┐
│ // Read request body                                                    │
│ body, err := io.ReadAll(r.Body)                                         │
│ defer r.Body.Close()                                                    │
│                                                                         │
│ // Parse JSON                                                           │
│ var data MyStruct                                                       │
│ err := json.NewDecoder(r.Body).Decode(&data)                            │
│                                                                         │
│ // Get query parameters                                                 │
│ id := r.URL.Query().Get("id")                                           │
│ values := r.URL.Query()["tags"]                                         │
│                                                                         │
│ // Get form values                                                      │
│ r.ParseForm()                                                           │
│ name := r.FormValue("name")                                             │
│                                                                         │
│ // Get headers                                                          │
│ auth := r.Header.Get("Authorization")                                   │
└─────────────────────────────────────────────────────────────────────────┘

🚨 ERROR HANDLING:
┌─────────────────────────────────────────────────────────────────────────┐
│ // Standard error response                                              │
│ http.Error(w, "Internal Server Error", http.StatusInternalServerError)  │
│                                                                         │
│ // JSON error response                                                  │
│ type ErrorResponse struct {                                             │
│     Error   string `json:"error"`                                       │
│     Code    int    `json:"code"`                                        │
│     Message string `json:"message"`                                     │
│ }                                                                       │
│                                                                         │
│ func sendError(w http.ResponseWriter, err string, code int) {            │
│     w.Header().Set("Content-Type", "application/json")                  │
│     w.WriteHeader(code)                                                 │
│     json.NewEncoder(w).Encode(ErrorResponse{                            │
│         Error:   err,                                                   │
│         Code:    code,                                                  │
│         Message: http.StatusText(code),                                 │
│     })                                                                  │
│ }                                                                       │
└─────────────────────────────────────────────────────────────────────────┘

💡 BEST PRACTICES:
• Always set appropriate Content-Type headers
• Use proper HTTP status codes
• Implement proper error handling
• Add request logging and monitoring
• Use middleware for cross-cutting concerns
• Set timeouts on your server
• Validate input data
• Handle graceful shutdown

🎯 REAL-WORLD PATTERNS:
┌─────────────────────────────────────────────────────────────────────────┐
│ // API versioning                                                       │
│ mux.HandleFunc("/api/v1/users", v1UsersHandler)                         │
│ mux.HandleFunc("/api/v2/users", v2UsersHandler)                         │
│                                                                         │
│ // Health check endpoint                                                │
│ mux.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) { │
│     w.WriteHeader(http.StatusOK)                                        │
│     fmt.Fprintf(w, "OK")                                                │
│ })                                                                      │
│                                                                         │
│ // Static file serving                                                  │
│ fs := http.FileServer(http.Dir("./static/"))                            │
│ mux.Handle("/static/", http.StripPrefix("/static/", fs))                │
│                                                                         │
│ // Graceful shutdown                                                    │
│ c := make(chan os.Signal, 1)                                            │
│ signal.Notify(c, os.Interrupt)                                          │
│ go func() {                                                             │
│     <-c                                                                 │
│     ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second) │
│     defer cancel()                                                      │
│     server.Shutdown(ctx)                                                │
│ }()                                                                     │
└─────────────────────────────────────────────────────────────────────────┘

⚡ PERFORMANCE TIPS:
• Use connection pooling
• Implement proper caching
• Use compression middleware
• Set appropriate buffer sizes
• Consider using reverse proxy (nginx)
• Monitor memory usage and goroutine leaks
• Use profiling tools (pprof)

=============================================================================
*/