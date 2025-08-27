/*
=============================================================================
                           🧪 GO TESTING TUTORIAL
=============================================================================

📚 CORE CONCEPT:
Go has built-in testing support through the testing package. Tests are
written as functions that start with "Test" and take *testing.T parameter.

🔑 KEY FEATURES:
• Unit tests with testing.T
• Benchmarks with testing.B
• Examples with testing.Example
• Table-driven tests
• Test coverage analysis

💡 REAL-WORLD ANALOGY:
Testing = Quality Control in Factory
- Unit tests = Testing individual components
- Benchmarks = Performance measurements
- Coverage = How much of the product was tested
- Table tests = Testing multiple scenarios efficiently

🎯 WHY WRITE TESTS?
• Catch bugs early in development
• Ensure code works as expected
• Enable safe refactoring
• Document expected behavior

=============================================================================
*/

package main

import (
	"errors"
	"fmt"
	"strings"
)

// 🧮 CALCULATOR: Functions to test
type Calculator struct{}

func (c Calculator) Add(a, b int) int {
	return a + b
}

func (c Calculator) Subtract(a, b int) int {
	return a - b
}

func (c Calculator) Multiply(a, b int) int {
	return a * b
}

func (c Calculator) Divide(a, b int) (int, error) {
	if b == 0 {
		return 0, errors.New("division by zero")
	}
	return a / b, nil
}

// 📝 STRING UTILITIES: More functions to test
func ReverseString(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

func IsPalindrome(s string) bool {
	s = strings.ToLower(s)
	return s == ReverseString(s)
}

func CountWords(s string) int {
	if strings.TrimSpace(s) == "" {
		return 0
	}
	return len(strings.Fields(s))
}

// 👤 USER STRUCT: For testing struct methods
type User struct {
	ID    int
	Name  string
	Email string
	Age   int
}

func (u User) IsAdult() bool {
	return u.Age >= 18
}

func (u User) GetDisplayName() string {
	if u.Name == "" {
		return "Anonymous"
	}
	return u.Name
}

func (u User) Validate() error {
	if u.Name == "" {
		return errors.New("name is required")
	}
	if u.Email == "" {
		return errors.New("email is required")
	}
	if u.Age < 0 {
		return errors.New("age cannot be negative")
	}
	return nil
}

// 🔍 SEARCH FUNCTION: For benchmark testing
func LinearSearch(slice []int, target int) int {
	for i, v := range slice {
		if v == target {
			return i
		}
	}
	return -1
}

func BinarySearch(slice []int, target int) int {
	left, right := 0, len(slice)-1
	
	for left <= right {
		mid := (left + right) / 2
		if slice[mid] == target {
			return mid
		} else if slice[mid] < target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return -1
}

func main() {
	fmt.Println("🧪 TESTING TUTORIAL")
	fmt.Println("===================")
	
	fmt.Println("\n📝 This file contains functions to be tested.")
	fmt.Println("The actual tests are in testing_test.go")
	fmt.Println("\nTo run tests:")
	fmt.Println("  go test")
	fmt.Println("  go test -v                    # Verbose output")
	fmt.Println("  go test -cover               # With coverage")
	fmt.Println("  go test -bench=.             # Run benchmarks")
	fmt.Println("  go test -run TestCalculator  # Run specific test")
	
	fmt.Println("\n🧮 Calculator Demo:")
	calc := Calculator{}
	fmt.Printf("  Add(5, 3) = %d\n", calc.Add(5, 3))
	fmt.Printf("  Subtract(10, 4) = %d\n", calc.Subtract(10, 4))
	fmt.Printf("  Multiply(6, 7) = %d\n", calc.Multiply(6, 7))
	
	result, err := calc.Divide(15, 3)
	if err != nil {
		fmt.Printf("  Divide(15, 3) = Error: %v\n", err)
	} else {
		fmt.Printf("  Divide(15, 3) = %d\n", result)
	}
	
	fmt.Println("\n📝 String Utilities Demo:")
	fmt.Printf("  ReverseString('hello') = '%s'\n", ReverseString("hello"))
	fmt.Printf("  IsPalindrome('racecar') = %t\n", IsPalindrome("racecar"))
	fmt.Printf("  CountWords('Hello Go World') = %d\n", CountWords("Hello Go World"))
	
	fmt.Println("\n👤 User Demo:")
	user := User{ID: 1, Name: "John Doe", Email: "john@example.com", Age: 25}
	fmt.Printf("  User: %+v\n", user)
	fmt.Printf("  IsAdult() = %t\n", user.IsAdult())
	fmt.Printf("  GetDisplayName() = '%s'\n", user.GetDisplayName())
	fmt.Printf("  Validate() = %v\n", user.Validate())
	
	fmt.Println("\n🔍 Search Demo:")
	numbers := []int{1, 3, 5, 7, 9, 11, 13, 15}
	target := 7
	fmt.Printf("  LinearSearch(%v, %d) = %d\n", numbers, target, LinearSearch(numbers, target))
	fmt.Printf("  BinarySearch(%v, %d) = %d\n", numbers, target, BinarySearch(numbers, target))
	
	fmt.Println("\n✨ Run 'go test' to see the tests in action!")
}

/*
=============================================================================
                              📝 LEARNING NOTES
=============================================================================

🧪 TEST FILE STRUCTURE:
┌─────────────────────────────────────────────────────────────────────────┐
│ // File: calculator_test.go                                             │
│ package main                                                            │
│                                                                         │
│ import (                                                                │
│     "testing"                                                           │
│ )                                                                       │
│                                                                         │
│ func TestAdd(t *testing.T) {                                            │
│     calc := Calculator{}                                                │
│     result := calc.Add(2, 3)                                            │
│     expected := 5                                                       │
│                                                                         │
│     if result != expected {                                             │
│         t.Errorf("Add(2, 3) = %d; want %d", result, expected)           │
│     }                                                                   │
│ }                                                                       │
└─────────────────────────────────────────────────────────────────────────┘

🎯 TEST FUNCTION PATTERNS:
┌─────────────────────────────────────────────────────────────────────────┐
│ // Basic test                                                           │
│ func TestFunctionName(t *testing.T) {                                   │
│     // Arrange                                                          │
│     input := "test"                                                     │
│     expected := "expected"                                              │
│                                                                         │
│     // Act                                                              │
│     result := FunctionToTest(input)                                     │
│                                                                         │
│     // Assert                                                           │
│     if result != expected {                                             │
│         t.Errorf("got %v, want %v", result, expected)                   │
│     }                                                                   │
│ }                                                                       │
│                                                                         │
│ // Table-driven test                                                    │
│ func TestFunction(t *testing.T) {                                       │
│     tests := []struct {                                                 │
│         name     string                                                 │
│         input    string                                                 │
│         expected string                                                 │
│     }{                                                                  │
│         {"case1", "input1", "expected1"},                               │
│         {"case2", "input2", "expected2"},                               │
│     }                                                                   │
│                                                                         │
│     for _, tt := range tests {                                          │
│         t.Run(tt.name, func(t *testing.T) {                             │
│             result := FunctionToTest(tt.input)                          │
│             if result != tt.expected {                                  │
│                 t.Errorf("got %v, want %v", result, tt.expected)        │
│             }                                                           │
│         })                                                              │
│     }                                                                   │
│ }                                                                       │
└─────────────────────────────────────────────────────────────────────────┘

📊 TESTING METHODS:
┌─────────────────┬─────────────────────────────────────────────────────────┐
│    Method       │                Description                              │
├─────────────────┼─────────────────────────────────────────────────────────┤
│ t.Error()       │ Report error but continue test                          │
│ t.Errorf()      │ Report formatted error but continue                     │
│ t.Fatal()       │ Report error and stop test immediately                  │
│ t.Fatalf()      │ Report formatted error and stop                         │
│ t.Log()         │ Log information (only shown with -v)                    │
│ t.Logf()        │ Log formatted information                               │
│ t.Skip()        │ Skip this test                                          │
│ t.Skipf()       │ Skip with formatted message                             │
│ t.Parallel()    │ Run test in parallel with others                        │
└─────────────────┴─────────────────────────────────────────────────────────┘

🏃 BENCHMARK TESTS:
┌─────────────────────────────────────────────────────────────────────────┐
│ func BenchmarkFunction(b *testing.B) {                                  │
│     // Setup code (not timed)                                           │
│     data := setupTestData()                                             │
│                                                                         │
│     b.ResetTimer() // Reset timer after setup                           │
│                                                                         │
│     for i := 0; i < b.N; i++ {                                          │
│         // Code to benchmark                                            │
│         FunctionToBenchmark(data)                                       │
│     }                                                                   │
│ }                                                                       │
│                                                                         │
│ // Run with: go test -bench=.                                           │
│ // Output: BenchmarkFunction-8  1000000  1234 ns/op                     │
└─────────────────────────────────────────────────────────────────────────┘

📋 EXAMPLE TESTS:
┌─────────────────────────────────────────────────────────────────────────┐
│ func ExampleFunction() {                                                │
│     result := Function("input")                                         │
│     fmt.Println(result)                                                 │
│     // Output: expected output                                          │
│ }                                                                       │
│                                                                         │
│ func ExampleFunction_variant() {                                        │
│     result := Function("different input")                               │
│     fmt.Println(result)                                                 │
│     // Output: different expected output                                │
│ }                                                                       │
└─────────────────────────────────────────────────────────────────────────┘

🎯 TEST COMMANDS:
┌─────────────────────────────────────────────────────────────────────────┐
│ go test                          # Run all tests                        │
│ go test -v                       # Verbose output                       │
│ go test -cover                   # Show coverage                        │
│ go test -coverprofile=cover.out  # Generate coverage file               │
│ go test -bench=.                 # Run benchmarks                       │
│ go test -run TestName            # Run specific test                    │
│ go test -short                   # Skip long-running tests              │
│ go test -race                    # Enable race detector                 │
│ go test -timeout 30s             # Set test timeout                     │
│ go test ./...                    # Test all packages                    │
└─────────────────────────────────────────────────────────────────────────┘

💡 TESTING BEST PRACTICES:
┌─────────────────────────────────────────────────────────────────────────┐
│ // Use table-driven tests for multiple cases                            │
│ tests := []struct {                                                     │
│     name string                                                         │
│     // test fields                                                      │
│ }{                                                                      │
│     // test cases                                                       │
│ }                                                                       │
│                                                                         │
│ // Use t.Helper() in test helper functions                              │
│ func assertEqual(t *testing.T, got, want interface{}) {                 │
│     t.Helper()                                                          │
│     if got != want {                                                    │
│         t.Errorf("got %v, want %v", got, want)                          │
│     }                                                                   │
│ }                                                                       │
│                                                                         │
│ // Use setup and teardown                                               │
│ func TestMain(m *testing.M) {                                           │
│     setup()                                                             │
│     code := m.Run()                                                     │
│     teardown()                                                          │
│     os.Exit(code)                                                       │
│ }                                                                       │
│                                                                         │
│ // Test error cases                                                     │
│ result, err := FunctionThatCanFail()                                    │
│ if err == nil {                                                         │
│     t.Error("expected error, got nil")                                  │
│ }                                                                       │
└─────────────────────────────────────────────────────────────────────────┘

🚨 COMMON TESTING MISTAKES:
❌ Not testing error cases
❌ Tests that depend on external resources
❌ Tests that depend on specific timing
❌ Not using table-driven tests for similar cases
❌ Testing implementation instead of behavior
❌ Not cleaning up resources in tests

🎯 WHAT TO TEST:
• Public functions and methods
• Edge cases and error conditions
• Different input combinations
• Boundary values
• Integration between components

🎯 WHAT NOT TO TEST:
• Private functions (test through public interface)
• Third-party library code
• Simple getters/setters without logic
• Generated code

⚡ PERFORMANCE TESTING:
• Use benchmarks for performance-critical code
• Test with realistic data sizes
• Use b.ResetTimer() after setup
• Run benchmarks multiple times
• Compare before/after optimizations

=============================================================================
*/