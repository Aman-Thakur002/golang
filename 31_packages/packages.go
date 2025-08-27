/*
=============================================================================
                           📦 GO PACKAGES TUTORIAL
=============================================================================

📚 CORE CONCEPT:
Packages are Go's way of organizing and reusing code. Every Go file belongs
to a package, and packages group related functionality together.

🔑 KEY FEATURES:
• Code organization and modularity
• Exported vs unexported identifiers
• Package initialization
• Import paths and aliases

💡 REAL-WORLD ANALOGY:
Package = Library Section
- Package name = Section label (Fiction, Science, etc.)
- Exported functions = Books available for checkout
- Unexported functions = Internal cataloging system
- Import = Getting a library card for that section

🎯 WHY USE PACKAGES?
• Code reusability across projects
• Namespace management
• Encapsulation and data hiding
• Better code organization

=============================================================================
*/

package main

import (
	"fmt"
	"math"
	"strings"
	"time"
)

// 📦 PACKAGE CONCEPTS DEMONSTRATION

// 🎯 EXPORTED IDENTIFIERS: Start with uppercase
type User struct {
	ID   int    // Exported field
	Name string // Exported field
	age  int    // unexported field (private)
}

// Exported method
func (u User) GetInfo() string {
	return fmt.Sprintf("User: %s (ID: %d)", u.Name, u.ID)
}

// unexported method (private)
func (u User) getAge() int {
	return u.age
}

// Exported function
func CreateUser(id int, name string, age int) User {
	return User{
		ID:   id,
		Name: name,
		age:  age, // Can access unexported field within same package
	}
}

// unexported function (private)
func validateUser(u User) bool {
	return u.Name != "" && u.ID > 0
}

// 🎯 PACKAGE VARIABLES AND CONSTANTS
var (
	// Exported variables
	Version = "1.0.0"
	Author  = "Go Developer"
	
	// unexported variables
	initialized = false
	userCount   = 0
)

// Exported constants
const (
	MaxUsers     = 1000
	DefaultLimit = 10
)

// unexported constants
const (
	internalKey = "secret"
	bufferSize  = 1024
)

// 🎯 PACKAGE INITIALIZATION
func init() {
	fmt.Println("📦 Package initializing...")
	initialized = true
	fmt.Printf("📦 Package %s v%s by %s initialized\n", "main", Version, Author)
}

func main() {
	fmt.Println("📦 PACKAGES TUTORIAL")
	fmt.Println("====================")

	// 🎯 DEMO 1: Using Standard Library Packages
	fmt.Println("\n🎯 DEMO 1: Standard Library Packages")
	fmt.Println("====================================")

	// math package
	fmt.Printf("📐 math.Pi = %.6f\n", math.Pi)
	fmt.Printf("📐 math.Sqrt(16) = %.2f\n", math.Sqrt(16))
	fmt.Printf("📐 math.Max(10, 20) = %.0f\n", math.Max(10, 20))

	// strings package
	text := "Hello, Go Packages!"
	fmt.Printf("📝 strings.ToUpper(%q) = %q\n", text, strings.ToUpper(text))
	fmt.Printf("📝 strings.Contains(%q, \"Go\") = %t\n", text, strings.Contains(text, "Go"))
	fmt.Printf("📝 strings.Split(%q, \" \") = %v\n", text, strings.Split(text, " "))

	// time package
	now := time.Now()
	fmt.Printf("⏰ time.Now() = %v\n", now.Format("2006-01-02 15:04:05"))
	fmt.Printf("⏰ Unix timestamp = %d\n", now.Unix())

	// 🎯 DEMO 2: Package Visibility
	fmt.Println("\n🎯 DEMO 2: Package Visibility")
	fmt.Println("=============================")

	// Create user using exported function
	user := CreateUser(1, "John Doe", 25)
	fmt.Printf("✅ Created user: %s\n", user.GetInfo())

	// Access exported fields
	fmt.Printf("✅ User ID: %d\n", user.ID)
	fmt.Printf("✅ User Name: %s\n", user.Name)

	// Cannot access unexported field directly
	// fmt.Printf("❌ User age: %d\n", user.age) // This would cause compile error

	// Can access unexported items within same package
	fmt.Printf("✅ User validation: %t\n", validateUser(user))
	fmt.Printf("✅ User age (via unexported method): %d\n", user.getAge())

	// 🎯 DEMO 3: Package Variables and Constants
	fmt.Println("\n🎯 DEMO 3: Package Variables & Constants")
	fmt.Println("=======================================")

	fmt.Printf("📋 Package Version: %s\n", Version)
	fmt.Printf("📋 Package Author: %s\n", Author)
	fmt.Printf("📋 Max Users: %d\n", MaxUsers)
	fmt.Printf("📋 Default Limit: %d\n", DefaultLimit)
	fmt.Printf("📋 Initialized: %t\n", initialized)

	// Modify exported package variable
	userCount++
	fmt.Printf("📋 User Count: %d\n", userCount)

	// 🎯 DEMO 4: Import Aliases
	fmt.Println("\n🎯 DEMO 4: Import Aliases")
	fmt.Println("=========================")

	// Example of how you would use import aliases:
	fmt.Println("📝 Import alias examples:")
	fmt.Println(`  import m "math"           // Alias 'm' for math`)
	fmt.Println(`  import . "fmt"            // Dot import (not recommended)`)
	fmt.Println(`  import _ "image/png"      // Blank import (for side effects)`)
	fmt.Println(`  import "path/to/package"  // Standard import`)

	// 🎯 DEMO 5: Package Best Practices
	fmt.Println("\n🎯 DEMO 5: Package Best Practices")
	fmt.Println("=================================")

	fmt.Println("✅ Package naming conventions:")
	fmt.Println("  • Use lowercase, single words")
	fmt.Println("  • Avoid underscores or mixed caps")
	fmt.Println("  • Make names descriptive but concise")
	fmt.Println("  • Examples: http, json, strings, time")

	fmt.Println("\n✅ Exported identifier conventions:")
	fmt.Println("  • Start with uppercase letter")
	fmt.Println("  • Use clear, descriptive names")
	fmt.Println("  • Document exported functions")

	fmt.Println("\n✅ Package organization:")
	fmt.Println("  • Group related functionality")
	fmt.Println("  • Keep packages focused and cohesive")
	fmt.Println("  • Minimize dependencies between packages")

	fmt.Println("\n✨ All package demos completed!")
}

/*
=============================================================================
                              📝 LEARNING NOTES
=============================================================================

📦 PACKAGE STRUCTURE:
┌─────────────────────────────────────────────────────────────────────────┐
│ // Package declaration (must be first non-comment line)                 │
│ package packagename                                                     │
│                                                                         │
│ // Import statements                                                    │
│ import (                                                                │
│     "fmt"                                                               │
│     "strings"                                                           │
│ )                                                                       │
│                                                                         │
│ // Package-level declarations                                           │
│ var packageVar = "value"                                                │
│ const PackageConst = 42                                                 │
│                                                                         │
│ // Functions, types, etc.                                               │
│ func ExportedFunction() { }                                             │
│ func unexportedFunction() { }                                           │
└─────────────────────────────────────────────────────────────────────────┘

🔍 VISIBILITY RULES:
┌─────────────────┬─────────────────┬─────────────────────────────────────┐
│   Identifier    │   Visibility    │           Access                    │
├─────────────────┼─────────────────┼─────────────────────────────────────┤
│ Uppercase       │ Exported        │ Accessible from other packages      │
│ lowercase       │ unexported      │ Only within same package            │
└─────────────────┴─────────────────┴─────────────────────────────────────┘

📥 IMPORT PATTERNS:
┌─────────────────────────────────────────────────────────────────────────┐
│ // Standard import                                                      │
│ import "fmt"                                                            │
│ fmt.Println("Hello")                                                    │
│                                                                         │
│ // Import with alias                                                    │
│ import f "fmt"                                                          │
│ f.Println("Hello")                                                      │
│                                                                         │
│ // Dot import (brings names into current namespace)                     │
│ import . "fmt"                                                          │
│ Println("Hello") // Not recommended                                     │
│                                                                         │
│ // Blank import (for side effects only)                                 │
│ import _ "image/png" // Registers PNG decoder                           │
│                                                                         │
│ // Multiple imports                                                     │
│ import (                                                                │
│     "fmt"                                                               │
│     "strings"                                                           │
│     "time"                                                              │
│ )                                                                       │
└─────────────────────────────────────────────────────────────────────────┘

🎯 PACKAGE INITIALIZATION:
┌─────────────────────────────────────────────────────────────────────────┐
│ // Package-level variables are initialized first                        │
│ var config = loadConfig()                                               │
│                                                                         │
│ // init() functions run after variable initialization                   │
│ func init() {                                                           │
│     // Initialization code                                              │
│     setupLogging()                                                      │
│     validateConfig()                                                    │
│ }                                                                       │
│                                                                         │
│ // Multiple init() functions are allowed                                │
│ func init() {                                                           │
│     // Another initialization step                                      │
│ }                                                                       │
│                                                                         │
│ // Initialization order:                                                │
│ // 1. Package-level variables (in dependency order)                     │
│ // 2. init() functions (in source order)                                │
│ // 3. main() function (if main package)                                 │
└─────────────────────────────────────────────────────────────────────────┘

💡 PACKAGE NAMING CONVENTIONS:
┌─────────────────────────────────────────────────────────────────────────┐
│ ✅ Good package names:                                                   │
│ • http, json, xml, sql                                                  │
│ • strings, bytes, time                                                  │
│ • crypto, hash, rand                                                    │
│                                                                         │
│ ❌ Avoid:                                                                │
│ • util, common, base (too generic)                                      │
│ • myPackage, my_package (wrong case/underscores)                        │
│ • verylongpackagename (too long)                                        │
│                                                                         │
│ 📝 Guidelines:                                                           │
│ • Use lowercase only                                                    │
│ • Single word preferred                                                 │
│ • Descriptive but concise                                               │
│ • No underscores or mixed caps                                          │
└─────────────────────────────────────────────────────────────────────────┘

🚨 COMMON MISTAKES:
❌ Using dot imports (import . "package")
❌ Circular imports between packages
❌ Making everything exported
❌ Poor package organization
❌ Generic package names (util, helper)

🔧 BEST PRACTICES:
• Keep packages focused and cohesive
• Minimize exported surface area
• Use clear, descriptive names
• Document exported identifiers
• Avoid circular dependencies
• Group related functionality
• Use init() sparingly

🎯 STANDARD LIBRARY HIGHLIGHTS:
┌─────────────────┬─────────────────────────────────────────────────────────┐
│    Package      │                Purpose                                  │
├─────────────────┼─────────────────────────────────────────────────────────┤
│ fmt             │ Formatted I/O (printing, scanning)                     │
│ strings         │ String manipulation utilities                           │
│ strconv         │ String conversions (to/from other types)               │
│ time            │ Time and date operations                                │
│ math            │ Mathematical functions and constants                    │
│ os              │ Operating system interface                              │
│ io              │ I/O primitives and utilities                            │
│ net/http        │ HTTP client and server                                  │
│ encoding/json   │ JSON encoding and decoding                              │
│ database/sql    │ Database access interface                               │
└─────────────────┴─────────────────────────────────────────────────────────┘

=============================================================================
*/