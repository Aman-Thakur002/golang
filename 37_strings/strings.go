/*
=============================================================================
                           📝 GO STRINGS TUTORIAL
=============================================================================

📚 CORE CONCEPT:
Go's strings package provides utilities for manipulating UTF-8 encoded strings.
Strings in Go are immutable, and the package offers efficient operations
for searching, replacing, and transforming text.

🔑 KEY FEATURES:
• String manipulation and transformation
• Searching and pattern matching
• Case conversion and trimming
• String building and formatting
• UTF-8 support

💡 REAL-WORLD ANALOGY:
Strings Package = Text Editor Toolkit
- Contains = Find function
- Replace = Find & Replace
- Split = Cut text into pieces
- Join = Paste pieces together
- Trim = Remove extra spaces

🎯 WHY LEARN STRINGS?
• Text processing and validation
• Data parsing and formatting
• User input sanitization
• Log processing and analysis

=============================================================================
*/

package main

import (
	"fmt"
	"strings"
	"unicode"
)

func main() {
	fmt.Println("📝 STRINGS TUTORIAL")
	fmt.Println("===================")

	// 🎯 DEMO 1: Basic String Operations
	fmt.Println("\n🎯 DEMO 1: Basic String Operations")
	fmt.Println("==================================")

	text := "Hello, Go Programming World!"
	fmt.Printf("Original: %q\n", text)

	// Length and basic info
	fmt.Printf("Length: %d characters\n", len(text))
	fmt.Printf("Contains 'Go': %t\n", strings.Contains(text, "Go"))
	fmt.Printf("Contains 'Python': %t\n", strings.Contains(text, "Python"))
	fmt.Printf("Starts with 'Hello': %t\n", strings.HasPrefix(text, "Hello"))
	fmt.Printf("Ends with 'World!': %t\n", strings.HasSuffix(text, "World!"))

	// Case operations
	fmt.Printf("Uppercase: %q\n", strings.ToUpper(text))
	fmt.Printf("Lowercase: %q\n", strings.ToLower(text))
	fmt.Printf("Title case: %q\n", strings.Title(text))

	// 🎯 DEMO 2: String Searching and Indexing
	fmt.Println("\n🎯 DEMO 2: String Searching")
	fmt.Println("===========================")

	searchText := "Go is great, Go is fast, Go is simple"
	fmt.Printf("Text: %q\n", searchText)

	// Find positions
	fmt.Printf("Index of 'Go': %d\n", strings.Index(searchText, "Go"))
	fmt.Printf("Last index of 'Go': %d\n", strings.LastIndex(searchText, "Go"))
	fmt.Printf("Index of 'Python': %d\n", strings.Index(searchText, "Python")) // -1 if not found

	// Count occurrences
	fmt.Printf("Count of 'Go': %d\n", strings.Count(searchText, "Go"))
	fmt.Printf("Count of 'is': %d\n", strings.Count(searchText, "is"))

	// Case-insensitive operations
	fmt.Printf("Contains 'GREAT' (case-insensitive): %t\n", 
		strings.Contains(strings.ToLower(searchText), strings.ToLower("GREAT")))

	// 🎯 DEMO 3: String Splitting and Joining
	fmt.Println("\n🎯 DEMO 3: Splitting and Joining")
	fmt.Println("================================")

	sentence := "apple,banana,cherry,date,elderberry"
	fmt.Printf("Original: %q\n", sentence)

	// Split by delimiter
	fruits := strings.Split(sentence, ",")
	fmt.Printf("Split by comma: %v\n", fruits)
	fmt.Printf("Number of fruits: %d\n", len(fruits))

	// Split with limit
	limited := strings.SplitN(sentence, ",", 3)
	fmt.Printf("Split with limit 3: %v\n", limited)

	// Split by whitespace
	text2 := "  apple   banana    cherry  "
	words := strings.Fields(text2)
	fmt.Printf("Fields from %q: %v\n", text2, words)

	// Join back together
	joined := strings.Join(fruits, " | ")
	fmt.Printf("Joined with ' | ': %q\n", joined)

	// Join with different separators
	separators := []string{", ", " - ", " -> ", "\n"}
	for _, sep := range separators {
		result := strings.Join(fruits[:3], sep)
		fmt.Printf("Joined with %q: %q\n", sep, result)
	}

	// 🎯 DEMO 4: String Replacement
	fmt.Println("\n🎯 DEMO 4: String Replacement")
	fmt.Println("=============================")

	original := "I love Python. Python is great. Python rocks!"
	fmt.Printf("Original: %q\n", original)

	// Replace all occurrences
	replaced := strings.ReplaceAll(original, "Python", "Go")
	fmt.Printf("Replace all 'Python' with 'Go': %q\n", replaced)

	// Replace with limit
	limitedReplace := strings.Replace(original, "Python", "Go", 2)
	fmt.Printf("Replace first 2 'Python' with 'Go': %q\n", limitedReplace)

	// Multiple replacements using Replacer
	replacer := strings.NewReplacer(
		"Python", "Go",
		"great", "awesome",
		"rocks", "is fantastic",
	)
	multiReplaced := replacer.Replace(original)
	fmt.Printf("Multiple replacements: %q\n", multiReplaced)

	// 🎯 DEMO 5: String Trimming and Cleaning
	fmt.Println("\n🎯 DEMO 5: Trimming and Cleaning")
	fmt.Println("================================")

	messyText := "   \t\n  Hello, World!  \n\t   "
	fmt.Printf("Messy text: %q\n", messyText)

	// Trim whitespace
	fmt.Printf("TrimSpace: %q\n", strings.TrimSpace(messyText))

	// Trim specific characters
	bracketText := "[[Hello, World!]]"
	fmt.Printf("Original: %q\n", bracketText)
	fmt.Printf("Trim '[' and ']': %q\n", strings.Trim(bracketText, "[]"))

	// Trim from left or right only
	fmt.Printf("TrimLeft '[': %q\n", strings.TrimLeft(bracketText, "["))
	fmt.Printf("TrimRight ']': %q\n", strings.TrimRight(bracketText, "]"))

	// Trim prefix/suffix
	urlText := "https://example.com/path"
	fmt.Printf("Original URL: %q\n", urlText)
	fmt.Printf("TrimPrefix 'https://': %q\n", strings.TrimPrefix(urlText, "https://"))
	fmt.Printf("TrimSuffix '/path': %q\n", strings.TrimSuffix(urlText, "/path"))

	// 🎯 DEMO 6: String Building
	fmt.Println("\n🎯 DEMO 6: String Building")
	fmt.Println("==========================")

	// Using strings.Builder for efficient string building
	var builder strings.Builder
	
	fmt.Println("Building a string efficiently:")
	words2 := []string{"Go", "is", "an", "awesome", "programming", "language"}
	
	for i, word := range words2 {
		if i > 0 {
			builder.WriteString(" ")
		}
		builder.WriteString(word)
	}
	
	result := builder.String()
	fmt.Printf("Built string: %q\n", result)
	fmt.Printf("Builder length: %d\n", builder.Len())

	// Reset and reuse builder
	builder.Reset()
	builder.WriteString("New content after reset")
	fmt.Printf("After reset: %q\n", builder.String())

	// 🎯 DEMO 7: String Comparison
	fmt.Println("\n🎯 DEMO 7: String Comparison")
	fmt.Println("============================")

	strings1 := []string{"apple", "Apple", "APPLE", "banana", "Banana"}
	
	fmt.Println("String comparisons:")
	for i, s1 := range strings1 {
		for j, s2 := range strings1 {
			if i < j {
				fmt.Printf("%-8s == %-8s: %t\n", s1, s2, s1 == s2)
				fmt.Printf("%-8s == %-8s (case-insensitive): %t\n", 
					s1, s2, strings.EqualFold(s1, s2))
			}
		}
	}

	// String comparison with Compare
	fmt.Println("\nString ordering:")
	testStrings := []string{"apple", "banana", "cherry"}
	for i := 0; i < len(testStrings)-1; i++ {
		cmp := strings.Compare(testStrings[i], testStrings[i+1])
		var relation string
		switch {
		case cmp < 0:
			relation = "comes before"
		case cmp > 0:
			relation = "comes after"
		default:
			relation = "equals"
		}
		fmt.Printf("%q %s %q\n", testStrings[i], relation, testStrings[i+1])
	}

	// 🎯 DEMO 8: Advanced String Operations
	fmt.Println("\n🎯 DEMO 8: Advanced Operations")
	fmt.Println("==============================")

	// Repeat strings
	fmt.Printf("Repeat 'Go! ' 5 times: %q\n", strings.Repeat("Go! ", 5))

	// Map function for character transformation
	rot13 := func(r rune) rune {
		if r >= 'a' && r <= 'z' {
			return 'a' + (r-'a'+13)%26
		}
		if r >= 'A' && r <= 'Z' {
			return 'A' + (r-'A'+13)%26
		}
		return r
	}
	
	secret := "Hello, World!"
	encoded := strings.Map(rot13, secret)
	decoded := strings.Map(rot13, encoded) // ROT13 is its own inverse
	fmt.Printf("Original: %q\n", secret)
	fmt.Printf("ROT13 encoded: %q\n", encoded)
	fmt.Printf("ROT13 decoded: %q\n", decoded)

	// Custom character filtering
	removeDigits := func(r rune) rune {
		if unicode.IsDigit(r) {
			return -1 // Remove character
		}
		return r
	}
	
	mixedText := "Hello123World456!"
	filtered := strings.Map(removeDigits, mixedText)
	fmt.Printf("Remove digits from %q: %q\n", mixedText, filtered)

	// 🎯 DEMO 9: Practical Examples
	fmt.Println("\n🎯 DEMO 9: Practical Examples")
	fmt.Println("=============================")

	// Email validation (simple)
	emails := []string{
		"user@example.com",
		"invalid.email",
		"test@domain.co.uk",
		"@invalid.com",
	}

	fmt.Println("Simple email validation:")
	for _, email := range emails {
		isValid := strings.Contains(email, "@") && 
				  strings.Contains(email, ".") &&
				  !strings.HasPrefix(email, "@") &&
				  !strings.HasSuffix(email, "@")
		status := "❌"
		if isValid {
			status = "✅"
		}
		fmt.Printf("  %s %s\n", status, email)
	}

	// URL path extraction
	urls := []string{
		"https://example.com/api/users",
		"http://localhost:8080/admin/dashboard",
		"https://api.github.com/repos/golang/go",
	}

	fmt.Println("\nURL path extraction:")
	for _, url := range urls {
		// Simple path extraction
		if idx := strings.Index(url, "://"); idx != -1 {
			remaining := url[idx+3:]
			if pathIdx := strings.Index(remaining, "/"); pathIdx != -1 {
				path := remaining[pathIdx:]
				fmt.Printf("  URL: %s → Path: %s\n", url, path)
			}
		}
	}

	// CSV parsing (simple)
	csvData := "John,25,Engineer\nJane,30,Designer\nBob,35,Manager"
	fmt.Println("\nSimple CSV parsing:")
	lines := strings.Split(csvData, "\n")
	for i, line := range lines {
		fields := strings.Split(line, ",")
		if len(fields) >= 3 {
			fmt.Printf("  Row %d: Name=%s, Age=%s, Job=%s\n", 
				i+1, fields[0], fields[1], fields[2])
		}
	}

	fmt.Println("\n✨ All string demos completed!")
}

/*
=============================================================================
                              📝 LEARNING NOTES
=============================================================================

📝 STRING BASICS:
┌─────────────────────────────────────────────────────────────────────────┐
│ // String properties                                                    │
│ len(s)                    // Length in bytes                            │
│ strings.Contains(s, sub)  // Check if s contains substring              │
│ strings.HasPrefix(s, pre) // Check if s starts with prefix              │
│ strings.HasSuffix(s, suf) // Check if s ends with suffix                │
│                                                                         │
│ // Case conversion                                                      │
│ strings.ToUpper(s)        // Convert to uppercase                       │
│ strings.ToLower(s)        // Convert to lowercase                       │
│ strings.Title(s)          // Convert to title case                      │
│ strings.ToTitle(s)        // Convert to title case (Unicode)            │
└─────────────────────────────────────────────────────────────────────────┘

🔍 STRING SEARCHING:
┌─────────────────────────────────────────────────────────────────────────┐
│ // Find positions                                                       │
│ strings.Index(s, sub)     // First occurrence index (-1 if not found)   │
│ strings.LastIndex(s, sub) // Last occurrence index                      │
│ strings.IndexAny(s, chars)// Index of any character in chars            │
│ strings.IndexFunc(s, f)   // Index of first rune satisfying f           │
│                                                                         │
│ // Count occurrences                                                    │
│ strings.Count(s, sub)     // Number of non-overlapping instances        │
│                                                                         │
│ // Case-insensitive                                                     │
│ strings.EqualFold(s1, s2) // Case-insensitive equality                  │
│ strings.Compare(s1, s2)   // Lexicographic comparison                    │
└─────────────────────────────────────────────────────────────────────────┘

✂️ STRING SPLITTING:
┌─────────────────────────────────────────────────────────────────────────┐
│ // Split operations                                                     │
│ strings.Split(s, sep)     // Split by separator                         │
│ strings.SplitN(s, sep, n) // Split with maximum n parts                 │
│ strings.SplitAfter(s, sep)// Split after separator                      │
│ strings.Fields(s)         // Split by whitespace                        │
│ strings.FieldsFunc(s, f)  // Split by function                          │
│                                                                         │
│ // Join operations                                                      │
│ strings.Join(slice, sep)  // Join slice with separator                  │
└─────────────────────────────────────────────────────────────────────────┘

🔄 STRING REPLACEMENT:
┌─────────────────────────────────────────────────────────────────────────┐
│ // Replace operations                                                   │
│ strings.Replace(s, old, new, n)    // Replace first n occurrences       │
│ strings.ReplaceAll(s, old, new)    // Replace all occurrences           │
│                                                                         │
│ // Multiple replacements                                                │
│ r := strings.NewReplacer("old1", "new1", "old2", "new2")               │
│ result := r.Replace(s)                                                  │
│                                                                         │
│ // Character mapping                                                    │
│ strings.Map(mapping, s)   // Apply function to each rune                │
└─────────────────────────────────────────────────────────────────────────┘

✨ STRING TRIMMING:
┌─────────────────────────────────────────────────────────────────────────┐
│ // Whitespace trimming                                                  │
│ strings.TrimSpace(s)      // Remove leading/trailing whitespace         │
│                                                                         │
│ // Character trimming                                                   │
│ strings.Trim(s, cutset)   // Remove leading/trailing characters         │
│ strings.TrimLeft(s, cutset)  // Remove leading characters               │
│ strings.TrimRight(s, cutset) // Remove trailing characters              │
│                                                                         │
│ // Prefix/suffix trimming                                               │
│ strings.TrimPrefix(s, prefix) // Remove prefix if present               │
│ strings.TrimSuffix(s, suffix) // Remove suffix if present               │
│                                                                         │
│ // Function-based trimming                                              │
│ strings.TrimFunc(s, f)    // Trim characters satisfying function        │
│ strings.TrimLeftFunc(s, f)   // Trim left with function                 │
│ strings.TrimRightFunc(s, f)  // Trim right with function                │
└─────────────────────────────────────────────────────────────────────────┘

🏗️ STRING BUILDING:
┌─────────────────────────────────────────────────────────────────────────┐
│ // Efficient string building                                            │
│ var builder strings.Builder                                             │
│ builder.WriteString("hello")                                            │
│ builder.WriteString(" world")                                           │
│ result := builder.String()                                              │
│                                                                         │
│ // Builder methods                                                      │
│ builder.Len()             // Current length                             │
│ builder.Cap()             // Current capacity                           │
│ builder.Reset()           // Reset to empty                             │
│ builder.Grow(n)           // Grow capacity by n bytes                   │
│                                                                         │
│ // Other operations                                                     │
│ strings.Repeat(s, count)  // Repeat string count times                  │
└─────────────────────────────────────────────────────────────────────────┘

💡 BEST PRACTICES:
┌─────────────────────────────────────────────────────────────────────────┐
│ // Use strings.Builder for concatenation                                │
│ var b strings.Builder                                                   │
│ for _, s := range parts {                                               │
│     b.WriteString(s)                                                    │
│ }                                                                       │
│ result := b.String()                                                    │
│                                                                         │
│ // Use strings.Join for slices                                          │
│ result := strings.Join(parts, separator)                                │
│                                                                         │
│ // Cache strings.Replacer for multiple uses                             │
│ replacer := strings.NewReplacer("old1", "new1", "old2", "new2")        │
│ // Use replacer multiple times                                          │
│                                                                         │
│ // Use strings.Fields for whitespace splitting                          │
│ words := strings.Fields(text)  // Better than Split(" ")               │
└─────────────────────────────────────────────────────────────────────────┘

🚨 COMMON MISTAKES:
❌ Using + for string concatenation in loops (inefficient)
❌ Not checking if Index returns -1
❌ Assuming len(s) equals number of characters (it's bytes)
❌ Using strings.Title for proper title casing (use cases package)
❌ Not handling empty strings in Split operations

⚡ PERFORMANCE TIPS:
• Use strings.Builder for multiple concatenations
• Use strings.Join instead of repeated concatenation
• Cache compiled strings.Replacer for reuse
• Use strings.Fields instead of Split with space
• Consider bytes package for byte slice operations
• Use strings.Contains before strings.Index when only checking existence

🎯 REAL-WORLD PATTERNS:
┌─────────────────────────────────────────────────────────────────────────┐
│ // Input validation                                                     │
│ func isValidEmail(email string) bool {                                  │
│     return strings.Contains(email, "@") &&                              │
│            strings.Contains(email, ".") &&                              │
│            !strings.HasPrefix(email, "@")                               │
│ }                                                                       │
│                                                                         │
│ // CSV parsing                                                          │
│ func parseCSV(data string) [][]string {                                 │
│     lines := strings.Split(data, "\n")                                  │
│     var result [][]string                                               │
│     for _, line := range lines {                                        │
│         if line = strings.TrimSpace(line); line != "" {                 │
│             result = append(result, strings.Split(line, ","))           │
│         }                                                               │
│     }                                                                   │
│     return result                                                       │
│ }                                                                       │
│                                                                         │
│ // Template processing                                                  │
│ func processTemplate(template string, vars map[string]string) string {  │
│     result := template                                                  │
│     for key, value := range vars {                                      │
│         placeholder := "{{" + key + "}}"                                │
│         result = strings.ReplaceAll(result, placeholder, value)         │
│     }                                                                   │
│     return result                                                       │
│ }                                                                       │
└─────────────────────────────────────────────────────────────────────────┘

=============================================================================
*/