/*
=============================================================================
                        🧪 GO TESTING EXAMPLES - TEST FILE
=============================================================================

This file contains actual test implementations for the functions in testing.go
Run with: go test -v
*/

package main

import (
	"errors"
	"testing"
)

// 🧮 CALCULATOR TESTS
func TestCalculatorAdd(t *testing.T) {
	calc := Calculator{}
	
	tests := []struct {
		name     string
		a, b     int
		expected int
	}{
		{"positive numbers", 5, 3, 8},
		{"negative numbers", -2, -3, -5},
		{"mixed numbers", -5, 10, 5},
		{"zero", 0, 5, 5},
	}
	
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := calc.Add(tt.a, tt.b)
			if result != tt.expected {
				t.Errorf("Add(%d, %d) = %d; want %d", tt.a, tt.b, result, tt.expected)
			}
		})
	}
}

func TestCalculatorDivide(t *testing.T) {
	calc := Calculator{}
	
	// Test successful division
	result, err := calc.Divide(10, 2)
	if err != nil {
		t.Errorf("Divide(10, 2) returned unexpected error: %v", err)
	}
	if result != 5 {
		t.Errorf("Divide(10, 2) = %d; want 5", result)
	}
	
	// Test division by zero
	_, err = calc.Divide(10, 0)
	if err == nil {
		t.Error("Divide(10, 0) should return error, got nil")
	}
	
	expectedError := "division by zero"
	if err.Error() != expectedError {
		t.Errorf("Divide(10, 0) error = %q; want %q", err.Error(), expectedError)
	}
}

// 📝 STRING UTILITIES TESTS
func TestReverseString(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{"simple string", "hello", "olleh"},
		{"empty string", "", ""},
		{"single character", "a", "a"},
		{"palindrome", "racecar", "racecar"},
		{"with spaces", "hello world", "dlrow olleh"},
	}
	
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := ReverseString(tt.input)
			if result != tt.expected {
				t.Errorf("ReverseString(%q) = %q; want %q", tt.input, result, tt.expected)
			}
		})
	}
}

func TestIsPalindrome(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected bool
	}{
		{"simple palindrome", "racecar", true},
		{"not palindrome", "hello", false},
		{"empty string", "", true},
		{"single character", "a", true},
		{"case insensitive", "RaceCar", true},
		{"with spaces", "race car", false}, // Note: this implementation doesn't handle spaces
	}
	
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := IsPalindrome(tt.input)
			if result != tt.expected {
				t.Errorf("IsPalindrome(%q) = %t; want %t", tt.input, result, tt.expected)
			}
		})
	}
}

func TestCountWords(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected int
	}{
		{"simple sentence", "hello world", 2},
		{"empty string", "", 0},
		{"single word", "hello", 1},
		{"multiple spaces", "hello    world", 2},
		{"leading/trailing spaces", "  hello world  ", 2},
		{"only spaces", "   ", 0},
	}
	
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := CountWords(tt.input)
			if result != tt.expected {
				t.Errorf("CountWords(%q) = %d; want %d", tt.input, result, tt.expected)
			}
		})
	}
}

// 👤 USER TESTS
func TestUserIsAdult(t *testing.T) {
	tests := []struct {
		name     string
		age      int
		expected bool
	}{
		{"adult", 25, true},
		{"exactly 18", 18, true},
		{"minor", 17, false},
		{"child", 10, false},
		{"elderly", 80, true},
	}
	
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			user := User{Age: tt.age}
			result := user.IsAdult()
			if result != tt.expected {
				t.Errorf("User{Age: %d}.IsAdult() = %t; want %t", tt.age, result, tt.expected)
			}
		})
	}
}

func TestUserGetDisplayName(t *testing.T) {
	tests := []struct {
		name     string
		userName string
		expected string
	}{
		{"with name", "John Doe", "John Doe"},
		{"empty name", "", "Anonymous"},
		{"whitespace name", "   ", "   "}, // Note: this implementation doesn't trim
	}
	
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			user := User{Name: tt.userName}
			result := user.GetDisplayName()
			if result != tt.expected {
				t.Errorf("User{Name: %q}.GetDisplayName() = %q; want %q", tt.userName, result, tt.expected)
			}
		})
	}
}

func TestUserValidate(t *testing.T) {
	tests := []struct {
		name        string
		user        User
		expectError bool
		errorMsg    string
	}{
		{
			name: "valid user",
			user: User{Name: "John", Email: "john@example.com", Age: 25},
			expectError: false,
		},
		{
			name: "missing name",
			user: User{Email: "john@example.com", Age: 25},
			expectError: true,
			errorMsg: "name is required",
		},
		{
			name: "missing email",
			user: User{Name: "John", Age: 25},
			expectError: true,
			errorMsg: "email is required",
		},
		{
			name: "negative age",
			user: User{Name: "John", Email: "john@example.com", Age: -5},
			expectError: true,
			errorMsg: "age cannot be negative",
		},
	}
	
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.user.Validate()
			
			if tt.expectError {
				if err == nil {
					t.Errorf("User.Validate() should return error, got nil")
					return
				}
				if err.Error() != tt.errorMsg {
					t.Errorf("User.Validate() error = %q; want %q", err.Error(), tt.errorMsg)
				}
			} else {
				if err != nil {
					t.Errorf("User.Validate() should not return error, got %v", err)
				}
			}
		})
	}
}

// 🔍 SEARCH TESTS
func TestLinearSearch(t *testing.T) {
	slice := []int{1, 3, 5, 7, 9, 11, 13}
	
	tests := []struct {
		name     string
		target   int
		expected int
	}{
		{"found at beginning", 1, 0},
		{"found in middle", 7, 3},
		{"found at end", 13, 6},
		{"not found", 4, -1},
		{"not found - too large", 20, -1},
	}
	
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := LinearSearch(slice, tt.target)
			if result != tt.expected {
				t.Errorf("LinearSearch(%v, %d) = %d; want %d", slice, tt.target, result, tt.expected)
			}
		})
	}
}

func TestBinarySearch(t *testing.T) {
	slice := []int{1, 3, 5, 7, 9, 11, 13}
	
	tests := []struct {
		name     string
		target   int
		expected int
	}{
		{"found at beginning", 1, 0},
		{"found in middle", 7, 3},
		{"found at end", 13, 6},
		{"not found", 4, -1},
		{"not found - too large", 20, -1},
	}
	
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := BinarySearch(slice, tt.target)
			if result != tt.expected {
				t.Errorf("BinarySearch(%v, %d) = %d; want %d", slice, tt.target, result, tt.expected)
			}
		})
	}
}

// 🏃 BENCHMARK TESTS
func BenchmarkLinearSearch(b *testing.B) {
	slice := make([]int, 1000)
	for i := range slice {
		slice[i] = i * 2
	}
	target := 500
	
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		LinearSearch(slice, target)
	}
}

func BenchmarkBinarySearch(b *testing.B) {
	slice := make([]int, 1000)
	for i := range slice {
		slice[i] = i * 2
	}
	target := 500
	
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		BinarySearch(slice, target)
	}
}

func BenchmarkReverseString(b *testing.B) {
	input := "This is a test string for benchmarking reverse function"
	
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		ReverseString(input)
	}
}

// 📋 EXAMPLE TESTS (for documentation)
func ExampleCalculator_Add() {
	calc := Calculator{}
	result := calc.Add(5, 3)
	fmt.Println(result)
	// Output: 8
}

func ExampleReverseString() {
	result := ReverseString("hello")
	fmt.Println(result)
	// Output: olleh
}

func ExampleIsPalindrome() {
	result := IsPalindrome("racecar")
	fmt.Println(result)
	// Output: true
}

func ExampleUser_IsAdult() {
	user := User{Age: 25}
	result := user.IsAdult()
	fmt.Println(result)
	// Output: true
}

// 🧪 TEST HELPER FUNCTIONS
func assertEqual(t *testing.T, got, want interface{}) {
	t.Helper() // This marks the function as a test helper
	if got != want {
		t.Errorf("got %v, want %v", got, want)
	}
}

func assertError(t *testing.T, got error, wantErr bool, wantMsg string) {
	t.Helper()
	if wantErr {
		if got == nil {
			t.Error("expected error, got nil")
			return
		}
		if got.Error() != wantMsg {
			t.Errorf("error message = %q; want %q", got.Error(), wantMsg)
		}
	} else {
		if got != nil {
			t.Errorf("unexpected error: %v", got)
		}
	}
}

// Example using helper functions
func TestCalculatorWithHelpers(t *testing.T) {
	calc := Calculator{}
	
	// Test Add
	result := calc.Add(5, 3)
	assertEqual(t, result, 8)
	
	// Test Divide success
	result, err := calc.Divide(10, 2)
	assertError(t, err, false, "")
	assertEqual(t, result, 5)
	
	// Test Divide error
	_, err = calc.Divide(10, 0)
	assertError(t, err, true, "division by zero")
}

/*
=============================================================================
                              📝 TEST OUTPUT EXAMPLES
=============================================================================

🎯 RUNNING TESTS:
$ go test
PASS
ok      your-package    0.002s

$ go test -v
=== RUN   TestCalculatorAdd
=== RUN   TestCalculatorAdd/positive_numbers
=== RUN   TestCalculatorAdd/negative_numbers
=== RUN   TestCalculatorAdd/mixed_numbers
=== RUN   TestCalculatorAdd/zero
--- PASS: TestCalculatorAdd (0.00s)
    --- PASS: TestCalculatorAdd/positive_numbers (0.00s)
    --- PASS: TestCalculatorAdd/negative_numbers (0.00s)
    --- PASS: TestCalculatorAdd/mixed_numbers (0.00s)
    --- PASS: TestCalculatorAdd/zero (0.00s)
PASS
ok      your-package    0.002s

🏃 RUNNING BENCHMARKS:
$ go test -bench=.
BenchmarkLinearSearch-8     1000000      1234 ns/op
BenchmarkBinarySearch-8    10000000       123 ns/op
BenchmarkReverseString-8    5000000       345 ns/op
PASS
ok      your-package    4.567s

📊 COVERAGE REPORT:
$ go test -cover
PASS
coverage: 85.7% of statements
ok      your-package    0.002s

$ go test -coverprofile=coverage.out
$ go tool cover -html=coverage.out

=============================================================================
*/