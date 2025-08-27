/*
=============================================================================
                        ❌ GO ERROR HANDLING TUTORIAL
=============================================================================

📚 CORE CONCEPT:
Go uses explicit error handling with the built-in error interface.
Errors are values that can be returned, checked, and handled gracefully.

🔑 KEY FEATURES:
• error interface for custom errors
• Multiple return values (result, error)
• Error wrapping and unwrapping
• Custom error types
• Error handling patterns

💡 REAL-WORLD ANALOGY:
Error Handling = Medical Diagnosis
- Symptoms = Error messages
- Diagnosis = Error types
- Treatment = Error handling
- Prevention = Input validation

🎯 WHY EXPLICIT ERROR HANDLING?
• Forces you to think about failure cases
• Makes error paths visible in code
• Enables graceful degradation
• Improves program reliability

=============================================================================
*/

package main

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

// 🎯 CUSTOM ERROR TYPES: Structured error information

// Simple custom error
type ValidationError struct {
	Field   string
	Message string
}

func (e ValidationError) Error() string {
	return fmt.Sprintf("validation error in field '%s': %s", e.Field, e.Message)
}

// Rich custom error with context
type DatabaseError struct {
	Operation string
	Table     string
	Err       error
}

func (e DatabaseError) Error() string {
	return fmt.Sprintf("database error during %s on table %s: %v", e.Operation, e.Table, e.Err)
}

func (e DatabaseError) Unwrap() error {
	return e.Err
}

// Error with error code
type APIError struct {
	Code    int
	Message string
	Details map[string]interface{}
}

func (e APIError) Error() string {
	return fmt.Sprintf("API error %d: %s", e.Code, e.Message)
}

// 📊 BUSINESS LOGIC: Functions that can fail

// Simple function with error
func divide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, errors.New("division by zero")
	}
	return a / b, nil
}

// Function with custom error
func validateAge(age int) error {
	if age < 0 {
		return ValidationError{
			Field:   "age",
			Message: "cannot be negative",
		}
	}
	if age > 150 {
		return ValidationError{
			Field:   "age",
			Message: "cannot be greater than 150",
		}
	}
	return nil
}

// Function with error wrapping
func parseAndValidateAge(ageStr string) (int, error) {
	age, err := strconv.Atoi(ageStr)
	if err != nil {
		return 0, fmt.Errorf("failed to parse age '%s': %w", ageStr, err)
	}
	
	if err := validateAge(age); err != nil {
		return 0, fmt.Errorf("age validation failed: %w", err)
	}
	
	return age, nil
}

// Simulated database operation
func getUserFromDB(userID int) (string, error) {
	if userID <= 0 {
		return "", DatabaseError{
			Operation: "SELECT",
			Table:     "users",
			Err:       errors.New("invalid user ID"),
		}
	}
	
	// Simulate user not found
	if userID == 999 {
		return "", DatabaseError{
			Operation: "SELECT",
			Table:     "users",
			Err:       errors.New("user not found"),
		}
	}
	
	return fmt.Sprintf("User%d", userID), nil
}

// API call simulation
func callExternalAPI(endpoint string) (string, error) {
	if endpoint == "" {
		return "", APIError{
			Code:    400,
			Message: "Bad Request",
			Details: map[string]interface{}{
				"field": "endpoint",
				"issue": "cannot be empty",
			},
		}
	}
	
	if endpoint == "timeout" {
		return "", APIError{
			Code:    408,
			Message: "Request Timeout",
			Details: map[string]interface{}{
				"timeout": "30s",
				"retry":   true,
			},
		}
	}
	
	return fmt.Sprintf("Response from %s", endpoint), nil
}

// 🔄 ERROR HANDLING PATTERNS

// Pattern 1: Early return
func processUser(userID int, ageStr string) (string, error) {
	// Get user from database
	username, err := getUserFromDB(userID)
	if err != nil {
		return "", fmt.Errorf("failed to get user: %w", err)
	}
	
	// Parse and validate age
	age, err := parseAndValidateAge(ageStr)
	if err != nil {
		return "", fmt.Errorf("failed to process age: %w", err)
	}
	
	return fmt.Sprintf("User: %s, Age: %d", username, age), nil
}

// Pattern 2: Error accumulation
func validateUserData(name, email, ageStr string) []error {
	var errors []error
	
	if name == "" {
		errors = append(errors, ValidationError{Field: "name", Message: "is required"})
	}
	
	if email == "" {
		errors = append(errors, ValidationError{Field: "email", Message: "is required"})
	} else if !strings.Contains(email, "@") {
		errors = append(errors, ValidationError{Field: "email", Message: "invalid format"})
	}
	
	if ageStr == "" {
		errors = append(errors, ValidationError{Field: "age", Message: "is required"})
	} else {
		if _, err := parseAndValidateAge(ageStr); err != nil {
			errors = append(errors, err)
		}
	}
	
	return errors
}

// Pattern 3: Error handling with recovery
func safeOperation(data []string, index int) (string, error) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Printf("🚨 Recovered from panic: %v\n", r)
		}
	}()
	
	if index < 0 || index >= len(data) {
		return "", errors.New("index out of bounds")
	}
	
	return data[index], nil
}

// 🎯 ERROR CHECKING UTILITIES
func isValidationError(err error) bool {
	var validationErr ValidationError
	return errors.As(err, &validationErr)
}

func isDatabaseError(err error) bool {
	var dbErr DatabaseError
	return errors.As(err, &dbErr)
}

func getAPIErrorCode(err error) (int, bool) {
	var apiErr APIError
	if errors.As(err, &apiErr) {
		return apiErr.Code, true
	}
	return 0, false
}

func main() {
	fmt.Println("❌ ERROR HANDLING TUTORIAL")
	fmt.Println("===========================")

	// 🎯 DEMO 1: Basic Error Handling
	fmt.Println("\n🎯 DEMO 1: Basic Error Handling")
	fmt.Println("===============================")

	result, err := divide(10, 2)
	if err != nil {
		fmt.Printf("❌ Division failed: %v\n", err)
	} else {
		fmt.Printf("✅ Division result: %.2f\n", result)
	}

	result, err = divide(10, 0)
	if err != nil {
		fmt.Printf("❌ Division failed: %v\n", err)
	} else {
		fmt.Printf("✅ Division result: %.2f\n", result)
	}

	// 🎯 DEMO 2: Custom Error Types
	fmt.Println("\n🎯 DEMO 2: Custom Error Types")
	fmt.Println("=============================")

	ages := []int{25, -5, 200}
	for _, age := range ages {
		if err := validateAge(age); err != nil {
			fmt.Printf("❌ Age %d: %v\n", age, err)
		} else {
			fmt.Printf("✅ Age %d: valid\n", age)
		}
	}

	// 🎯 DEMO 3: Error Wrapping
	fmt.Println("\n🎯 DEMO 3: Error Wrapping")
	fmt.Println("=========================")

	ageStrings := []string{"25", "abc", "-10"}
	for _, ageStr := range ageStrings {
		age, err := parseAndValidateAge(ageStr)
		if err != nil {
			fmt.Printf("❌ Age string '%s': %v\n", ageStr, err)
			
			// Check if it's a validation error
			if isValidationError(err) {
				fmt.Printf("   🔍 This is a validation error\n")
			}
		} else {
			fmt.Printf("✅ Age string '%s': %d\n", ageStr, age)
		}
	}

	// 🎯 DEMO 4: Database Errors
	fmt.Println("\n🎯 DEMO 4: Database Errors")
	fmt.Println("==========================")

	userIDs := []int{1, -1, 999}
	for _, userID := range userIDs {
		user, err := getUserFromDB(userID)
		if err != nil {
			fmt.Printf("❌ User ID %d: %v\n", userID, err)
			
			if isDatabaseError(err) {
				fmt.Printf("   🔍 This is a database error\n")
			}
		} else {
			fmt.Printf("✅ User ID %d: %s\n", userID, user)
		}
	}

	// 🎯 DEMO 5: API Errors
	fmt.Println("\n🎯 DEMO 5: API Errors")
	fmt.Println("=====================")

	endpoints := []string{"users", "", "timeout"}
	for _, endpoint := range endpoints {
		response, err := callExternalAPI(endpoint)
		if err != nil {
			fmt.Printf("❌ Endpoint '%s': %v\n", endpoint, err)
			
			if code, ok := getAPIErrorCode(err); ok {
				fmt.Printf("   🔍 API Error Code: %d\n", code)
			}
		} else {
			fmt.Printf("✅ Endpoint '%s': %s\n", endpoint, response)
		}
	}

	// 🎯 DEMO 6: Error Handling Patterns
	fmt.Println("\n🎯 DEMO 6: Error Handling Patterns")
	fmt.Println("==================================")

	// Early return pattern
	result_str, err := processUser(1, "25")
	if err != nil {
		fmt.Printf("❌ Process user failed: %v\n", err)
	} else {
		fmt.Printf("✅ Process user result: %s\n", result_str)
	}

	// Error accumulation pattern
	fmt.Println("\n📋 Validating user data:")
	validationErrors := validateUserData("", "invalid-email", "abc")
	if len(validationErrors) > 0 {
		fmt.Printf("❌ Validation failed with %d errors:\n", len(validationErrors))
		for i, err := range validationErrors {
			fmt.Printf("   %d. %v\n", i+1, err)
		}
	} else {
		fmt.Println("✅ Validation passed")
	}

	// 🎯 DEMO 7: Error Unwrapping
	fmt.Println("\n🎯 DEMO 7: Error Unwrapping")
	fmt.Println("===========================")

	_, err = parseAndValidateAge("abc")
	if err != nil {
		fmt.Printf("❌ Original error: %v\n", err)
		
		// Unwrap to find the root cause
		unwrapped := errors.Unwrap(err)
		if unwrapped != nil {
			fmt.Printf("🔍 Unwrapped error: %v\n", unwrapped)
		}
		
		// Check if it contains a specific error
		var numErr *strconv.NumError
		if errors.As(err, &numErr) {
			fmt.Printf("🔍 Contains NumError: %v\n", numErr)
		}
	}

	// 🎯 DEMO 8: Error Comparison
	fmt.Println("\n🎯 DEMO 8: Error Comparison")
	fmt.Println("===========================")

	err1 := errors.New("test error")
	err2 := errors.New("test error")
	err3 := err1

	fmt.Printf("err1 == err2: %t (different instances)\n", err1 == err2)
	fmt.Printf("err1 == err3: %t (same instance)\n", err1 == err3)
	fmt.Printf("errors.Is(err1, err2): %t\n", errors.Is(err1, err2))

	// Using sentinel errors
	var ErrNotFound = errors.New("not found")
	
	testErr := fmt.Errorf("user lookup failed: %w", ErrNotFound)
	fmt.Printf("errors.Is(testErr, ErrNotFound): %t\n", errors.Is(testErr, ErrNotFound))

	fmt.Println("\n✨ All error handling demos completed!")
}

/*
=============================================================================
                              📝 LEARNING NOTES
=============================================================================

❌ ERROR INTERFACE:
┌─────────────────────────────────────────────────────────────────────────┐
│ type error interface {                                                  │
│     Error() string                                                      │
│ }                                                                       │
│                                                                         │
│ // Any type that implements Error() string is an error                  │
│ type MyError struct {                                                   │
│     Message string                                                      │
│ }                                                                       │
│                                                                         │
│ func (e MyError) Error() string {                                       │
│     return e.Message                                                    │
│ }                                                                       │
└─────────────────────────────────────────────────────────────────────────┘

🎯 ERROR HANDLING PATTERNS:
┌─────────────────────────────────────────────────────────────────────────┐
│ // Basic pattern                                                        │
│ result, err := someFunction()                                           │
│ if err != nil {                                                         │
│     // Handle error                                                     │
│     return err                                                          │
│ }                                                                       │
│ // Use result                                                           │
│                                                                         │
│ // Error wrapping                                                       │
│ if err != nil {                                                         │
│     return fmt.Errorf("operation failed: %w", err)                      │
│ }                                                                       │
│                                                                         │
│ // Multiple error handling                                              │
│ if err := step1(); err != nil {                                         │
│     return fmt.Errorf("step1 failed: %w", err)                          │
│ }                                                                       │
│ if err := step2(); err != nil {                                         │
│     return fmt.Errorf("step2 failed: %w", err)                          │
│ }                                                                       │
└─────────────────────────────────────────────────────────────────────────┘

🔧 ERROR CREATION:
┌─────────────────────────────────────────────────────────────────────────┐
│ // Simple error                                                         │
│ err := errors.New("something went wrong")                               │
│                                                                         │
│ // Formatted error                                                      │
│ err := fmt.Errorf("failed to process %s: %v", name, originalErr)        │
│                                                                         │
│ // Error wrapping (Go 1.13+)                                            │
│ err := fmt.Errorf("operation failed: %w", originalErr)                  │
│                                                                         │
│ // Custom error type                                                    │
│ type ValidationError struct {                                           │
│     Field string                                                        │
│     Value interface{}                                                   │
│ }                                                                       │
│                                                                         │
│ func (e ValidationError) Error() string {                               │
│     return fmt.Sprintf("invalid %s: %v", e.Field, e.Value)              │
│ }                                                                       │
└─────────────────────────────────────────────────────────────────────────┘

🔍 ERROR INSPECTION:
┌─────────────────────────────────────────────────────────────────────────┐
│ // Check if error is a specific value                                   │
│ if errors.Is(err, ErrNotFound) {                                        │
│     // Handle not found error                                           │
│ }                                                                       │
│                                                                         │
│ // Check if error is a specific type                                    │
│ var validationErr ValidationError                                       │
│ if errors.As(err, &validationErr) {                                     │
│     // Handle validation error                                          │
│     fmt.Printf("Field: %s", validationErr.Field)                        │
│ }                                                                       │
│                                                                         │
│ // Unwrap error                                                         │
│ unwrapped := errors.Unwrap(err)                                         │
│ if unwrapped != nil {                                                   │
│     // Handle unwrapped error                                           │
│ }                                                                       │
└─────────────────────────────────────────────────────────────────────────┘

💡 BEST PRACTICES:
┌─────────────────────────────────────────────────────────────────────────┐
│ // Use descriptive error messages                                       │
│ return fmt.Errorf("failed to connect to database %s: %w", dbName, err)  │
│                                                                         │
│ // Don't ignore errors                                                  │
│ if err != nil {                                                         │
│     log.Printf("warning: %v", err) // At minimum, log it               │
│ }                                                                       │
│                                                                         │
│ // Use sentinel errors for expected conditions                          │
│ var ErrNotFound = errors.New("not found")                               │
│                                                                         │
│ // Wrap errors to add context                                           │
│ if err != nil {                                                         │
│     return fmt.Errorf("user validation failed: %w", err)                │
│ }                                                                       │
│                                                                         │
│ // Handle errors at the right level                                     │
│ // - Low level: wrap and return                                         │
│ // - High level: handle and recover                                     │
└─────────────────────────────────────────────────────────────────────────┘

🚨 COMMON MISTAKES:
❌ Ignoring errors: _, err := someFunc(); // Don't do this
❌ Generic error messages: errors.New("error")
❌ Not wrapping errors: return err (loses context)
❌ Handling errors too early (should bubble up)
❌ Using panic for recoverable errors

🎯 ERROR TYPES BY USE CASE:
┌─────────────────┬─────────────────────────────────────────────────────────┐
│   Use Case      │                Pattern                                  │
├─────────────────┼─────────────────────────────────────────────────────────┤
│ Simple failure  │ errors.New("message")                                   │
│ Formatted error │ fmt.Errorf("failed to %s: %v", action, err)             │
│ Wrapped error   │ fmt.Errorf("context: %w", err)                          │
│ Custom error    │ type MyError struct { ... }                             │
│ Sentinel error  │ var ErrNotFound = errors.New("not found")               │
│ Temporary error │ type TemporaryError interface { Temporary() bool }       │
└─────────────────┴─────────────────────────────────────────────────────────┘

⚡ PERFORMANCE CONSIDERATIONS:
• Error creation is relatively expensive
• Use error variables for common errors
• Don't create errors in hot paths unnecessarily
• Consider error pools for high-frequency errors
• Stack traces in custom errors can be expensive

🎯 WHEN TO USE EACH APPROACH:
• errors.New(): Simple, static error messages
• fmt.Errorf(): Dynamic error messages with context
• Custom types: Rich error information, error codes
• Wrapping: Adding context while preserving original error
• Sentinel errors: Expected conditions that callers check for

=============================================================================
*/