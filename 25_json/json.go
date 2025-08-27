/*
=============================================================================
                           📄 GO JSON TUTORIAL
=============================================================================

📚 CORE CONCEPT:
JSON (JavaScript Object Notation) is a lightweight data format. Go's json
package provides powerful tools for encoding (marshaling) and decoding
(unmarshaling) JSON data.

🔑 KEY FEATURES:
• Marshal: Go struct → JSON string
• Unmarshal: JSON string → Go struct
• Struct tags for field mapping
• Custom JSON handling with interfaces

💡 REAL-WORLD ANALOGY:
JSON = Universal Language Translator
- Marshal = Translate your thoughts to universal language
- Unmarshal = Understand universal language as your thoughts
- Struct tags = Translation rules and preferences

🎯 WHY USE JSON?
• API communication (REST, GraphQL)
• Configuration files
• Data storage and exchange
• Web frontend-backend communication

=============================================================================
*/

package main

import (
	"encoding/json"
	"fmt"
	"time"
)

// 👤 BASIC STRUCT: Simple user data
type User struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Age      int    `json:"age"`
	IsActive bool   `json:"is_active"`
}

// 🏢 NESTED STRUCT: Complex data with embedded structs
type Company struct {
	ID      int    `json:"id"`
	Name    string `json:"name"`
	Address struct {
		Street  string `json:"street"`
		City    string `json:"city"`
		Country string `json:"country"`
		ZipCode string `json:"zip_code"`
	} `json:"address"`
	Employees []User `json:"employees"`
	Founded   time.Time `json:"founded"`
}

// 🎯 STRUCT TAGS: Different JSON field behaviors
type Product struct {
	ID          int     `json:"id"`
	Name        string  `json:"name"`
	Price       float64 `json:"price"`
	Description string  `json:"description,omitempty"` // Omit if empty
	InStock     bool    `json:"in_stock"`
	Tags        []string `json:"tags,omitempty"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"-"` // Never include in JSON
	internal    string   // Unexported fields are ignored
}

// 🎨 CUSTOM JSON: Implementing json.Marshaler and json.Unmarshaler
type Color struct {
	R, G, B uint8
}

// Custom JSON marshaling
func (c Color) MarshalJSON() ([]byte, error) {
	hex := fmt.Sprintf("#%02x%02x%02x", c.R, c.G, c.B)
	return json.Marshal(hex)
}

// Custom JSON unmarshaling
func (c *Color) UnmarshalJSON(data []byte) error {
	var hex string
	if err := json.Unmarshal(data, &hex); err != nil {
		return err
	}
	
	if len(hex) != 7 || hex[0] != '#' {
		return fmt.Errorf("invalid color format: %s", hex)
	}
	
	_, err := fmt.Sscanf(hex, "#%02x%02x%02x", &c.R, &c.G, &c.B)
	return err
}

func main() {
	fmt.Println("📄 JSON TUTORIAL")
	fmt.Println("================")

	// 🎯 DEMO 1: Basic Marshal (Go → JSON)
	fmt.Println("\n🎯 DEMO 1: Basic Marshal")
	fmt.Println("========================")

	user := User{
		ID:       1,
		Name:     "John Doe",
		Email:    "john@example.com",
		Age:      30,
		IsActive: true,
	}

	// Marshal to JSON
	jsonData, err := json.Marshal(user)
	if err != nil {
		fmt.Printf("❌ Marshal error: %v\n", err)
		return
	}

	fmt.Printf("📤 Marshaled JSON:\n%s\n", string(jsonData))

	// Pretty print JSON
	prettyJSON, err := json.MarshalIndent(user, "", "  ")
	if err != nil {
		fmt.Printf("❌ Pretty marshal error: %v\n", err)
		return
	}

	fmt.Printf("🎨 Pretty JSON:\n%s\n", string(prettyJSON))

	// 🎯 DEMO 2: Basic Unmarshal (JSON → Go)
	fmt.Println("\n🎯 DEMO 2: Basic Unmarshal")
	fmt.Println("==========================")

	jsonString := `{
		"id": 2,
		"name": "Jane Smith",
		"email": "jane@example.com",
		"age": 25,
		"is_active": false
	}`

	var user2 User
	err = json.Unmarshal([]byte(jsonString), &user2)
	if err != nil {
		fmt.Printf("❌ Unmarshal error: %v\n", err)
		return
	}

	fmt.Printf("📥 Unmarshaled User: %+v\n", user2)

	// 🎯 DEMO 3: Nested Structures
	fmt.Println("\n🎯 DEMO 3: Nested Structures")
	fmt.Println("============================")

	company := Company{
		ID:   1,
		Name: "Tech Corp",
		Address: struct {
			Street  string `json:"street"`
			City    string `json:"city"`
			Country string `json:"country"`
			ZipCode string `json:"zip_code"`
		}{
			Street:  "123 Tech Street",
			City:    "San Francisco",
			Country: "USA",
			ZipCode: "94105",
		},
		Employees: []User{user, user2},
		Founded:   time.Date(2020, 1, 15, 0, 0, 0, 0, time.UTC),
	}

	companyJSON, err := json.MarshalIndent(company, "", "  ")
	if err != nil {
		fmt.Printf("❌ Company marshal error: %v\n", err)
		return
	}

	fmt.Printf("🏢 Company JSON:\n%s\n", string(companyJSON))

	// 🎯 DEMO 4: Struct Tags and Omitempty
	fmt.Println("\n🎯 DEMO 4: Struct Tags")
	fmt.Println("======================")

	product1 := Product{
		ID:        1,
		Name:      "Laptop",
		Price:     999.99,
		Description: "High-performance laptop",
		InStock:   true,
		Tags:      []string{"electronics", "computers"},
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		internal:  "secret", // Won't be included
	}

	product1JSON, err := json.MarshalIndent(product1, "", "  ")
	if err != nil {
		fmt.Printf("❌ Product marshal error: %v\n", err)
		return
	}

	fmt.Printf("📦 Product with all fields:\n%s\n", string(product1JSON))

	// Product with empty optional fields
	product2 := Product{
		ID:        2,
		Name:      "Mouse",
		Price:     29.99,
		InStock:   true,
		CreatedAt: time.Now(),
		// Description and Tags are empty (will be omitted)
	}

	product2JSON, err := json.MarshalIndent(product2, "", "  ")
	if err != nil {
		fmt.Printf("❌ Product2 marshal error: %v\n", err)
		return
	}

	fmt.Printf("📦 Product with omitted fields:\n%s\n", string(product2JSON))

	// 🎯 DEMO 5: Working with Maps and Interfaces
	fmt.Println("\n🎯 DEMO 5: Maps and Interfaces")
	fmt.Println("==============================")

	// JSON to map[string]interface{}
	jsonData2 := `{
		"name": "Dynamic Data",
		"count": 42,
		"active": true,
		"tags": ["go", "json", "tutorial"],
		"metadata": {
			"version": "1.0",
			"author": "Go Developer"
		}
	}`

	var data map[string]interface{}
	err = json.Unmarshal([]byte(jsonData2), &data)
	if err != nil {
		fmt.Printf("❌ Map unmarshal error: %v\n", err)
		return
	}

	fmt.Printf("🗺️ Dynamic data: %+v\n", data)

	// Access nested data
	if metadata, ok := data["metadata"].(map[string]interface{}); ok {
		fmt.Printf("📝 Author: %v\n", metadata["author"])
	}

	// Marshal map back to JSON
	mapJSON, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		fmt.Printf("❌ Map marshal error: %v\n", err)
		return
	}

	fmt.Printf("🗺️ Map as JSON:\n%s\n", string(mapJSON))

	// 🎯 DEMO 6: Custom JSON Marshaling
	fmt.Println("\n🎯 DEMO 6: Custom JSON")
	fmt.Println("======================")

	colors := []Color{
		{255, 0, 0},   // Red
		{0, 255, 0},   // Green
		{0, 0, 255},   // Blue
		{255, 255, 0}, // Yellow
	}

	colorsJSON, err := json.MarshalIndent(colors, "", "  ")
	if err != nil {
		fmt.Printf("❌ Colors marshal error: %v\n", err)
		return
	}

	fmt.Printf("🎨 Colors as JSON:\n%s\n", string(colorsJSON))

	// Unmarshal custom JSON
	colorJSON := `["#ff0000", "#00ff00", "#0000ff"]`
	var unmarshaledColors []Color
	err = json.Unmarshal([]byte(colorJSON), &unmarshaledColors)
	if err != nil {
		fmt.Printf("❌ Colors unmarshal error: %v\n", err)
		return
	}

	fmt.Printf("🎨 Unmarshaled colors: %+v\n", unmarshaledColors)

	// 🎯 DEMO 7: JSON Validation and Error Handling
	fmt.Println("\n🎯 DEMO 7: Error Handling")
	fmt.Println("=========================")

	invalidJSON := `{
		"id": "not_a_number",
		"name": "Invalid User",
		"email": "invalid@email"
	}`

	var invalidUser User
	err = json.Unmarshal([]byte(invalidJSON), &invalidUser)
	if err != nil {
		fmt.Printf("❌ Expected error: %v\n", err)
		
		// Type assertion to get more details
		if jsonErr, ok := err.(*json.UnmarshalTypeError); ok {
			fmt.Printf("🔍 Field: %s, Expected: %s, Got: %s\n", 
				jsonErr.Field, jsonErr.Type, jsonErr.Value)
		}
	}

	fmt.Println("\n✨ All JSON demos completed!")
}

/*
=============================================================================
                              📝 LEARNING NOTES
=============================================================================

📄 JSON OPERATIONS:
┌─────────────────┬─────────────────┬─────────────────────────────────────┐
│   Operation     │     Function    │           Purpose                   │
├─────────────────┼─────────────────┼─────────────────────────────────────┤
│ Marshal         │ json.Marshal()  │ Go struct → JSON bytes             │
│ MarshalIndent   │ json.MarshalIndent() │ Go struct → Pretty JSON    │
│ Unmarshal       │ json.Unmarshal()│ JSON bytes → Go struct             │
│ NewEncoder      │ json.NewEncoder()│ Stream encoding to io.Writer      │
│ NewDecoder      │ json.NewDecoder()│ Stream decoding from io.Reader    │
└─────────────────┴─────────────────┴─────────────────────────────────────┘

🏷️ STRUCT TAGS:
┌─────────────────────────────────────────────────────────────────────────┐
│ type User struct {                                                      │
│     ID       int    `json:"id"`              // Map to "id"             │
│     Name     string `json:"name,omitempty"`  // Omit if empty           │
│     Email    string `json:"email"`           // Map to "email"          │
│     Password string `json:"-"`               // Never include           │
│     Age      int    `json:"age,string"`      // Convert to string       │
│ }                                                                       │
└─────────────────────────────────────────────────────────────────────────┘

🎯 COMMON STRUCT TAG OPTIONS:
• `json:"field_name"` - Custom field name
• `json:",omitempty"` - Omit if zero value
• `json:"-"` - Never include in JSON
• `json:",string"` - Encode as string
• No tag - Use field name as-is (if exported)

🔄 TYPE MAPPING:
┌─────────────────┬─────────────────┬─────────────────────────────────────┐
│   Go Type       │   JSON Type     │           Notes                     │
├─────────────────┼─────────────────┼─────────────────────────────────────┤
│ bool            │ boolean         │ true/false                          │
│ int, int64      │ number          │ Integer values                      │
│ float64         │ number          │ Floating point                      │
│ string          │ string          │ UTF-8 text                          │
│ []T             │ array           │ Array of type T                     │
│ map[string]T    │ object          │ Object with string keys             │
│ struct          │ object          │ Object with field mapping           │
│ *T              │ null or T       │ Pointer types                       │
│ interface{}     │ any             │ Dynamic type                        │
│ time.Time       │ string          │ RFC3339 format                      │
└─────────────────┴─────────────────┴─────────────────────────────────────┘

🎨 CUSTOM JSON INTERFACES:
┌─────────────────────────────────────────────────────────────────────────┐
│ // Custom marshaling                                                    │
│ type Marshaler interface {                                              │
│     MarshalJSON() ([]byte, error)                                       │
│ }                                                                       │
│                                                                         │
│ // Custom unmarshaling                                                  │
│ type Unmarshaler interface {                                            │
│     UnmarshalJSON([]byte) error                                         │
│ }                                                                       │
│                                                                         │
│ // Text marshaling (for map keys, etc.)                                 │
│ type TextMarshaler interface {                                          │
│     MarshalText() ([]byte, error)                                       │
│ }                                                                       │
│                                                                         │
│ type TextUnmarshaler interface {                                        │
│     UnmarshalText([]byte) error                                         │
│ }                                                                       │
└─────────────────────────────────────────────────────────────────────────┘

🚨 COMMON ERRORS:
❌ Unmarshaling into nil pointer
❌ Type mismatches (string vs number)
❌ Unexported struct fields (won't be marshaled)
❌ Circular references (causes infinite loop)
❌ Invalid JSON syntax

💡 BEST PRACTICES:
• Use struct tags for API compatibility
• Handle errors from Marshal/Unmarshal
• Use omitempty for optional fields
• Validate JSON structure before unmarshaling
• Use json.RawMessage for delayed parsing
• Consider using json.NewEncoder/Decoder for streams

🎯 REAL-WORLD PATTERNS:
┌─────────────────────────────────────────────────────────────────────────┐
│ // API Response wrapper                                                 │
│ type APIResponse struct {                                               │
│     Success bool        `json:"success"`                                │
│     Data    interface{} `json:"data,omitempty"`                         │
│     Error   string      `json:"error,omitempty"`                        │
│ }                                                                       │
│                                                                         │
│ // Configuration file                                                   │
│ type Config struct {                                                    │
│     Database struct {                                                   │
│         Host     string `json:"host"`                                   │
│         Port     int    `json:"port"`                                   │
│         Username string `json:"username"`                               │
│         Password string `json:"-"` // Don't serialize passwords         │
│     } `json:"database"`                                                 │
│ }                                                                       │
│                                                                         │
│ // Flexible data handling                                               │
│ var data map[string]interface{}                                         │
│ json.Unmarshal(jsonBytes, &data)                                        │
└─────────────────────────────────────────────────────────────────────────┘

⚡ PERFORMANCE TIPS:
• Use json.NewEncoder/Decoder for large data streams
• Reuse structs to reduce allocations
• Use json.RawMessage to defer parsing
• Consider alternative libraries (easyjson, ffjson) for high performance
• Profile your JSON operations in production

=============================================================================
*/