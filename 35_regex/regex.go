/*
=============================================================================
                        🔍 GO REGULAR EXPRESSIONS TUTORIAL
=============================================================================

📚 CORE CONCEPT:
Regular expressions (regex) are patterns used to match character combinations
in strings. Go's regexp package provides powerful pattern matching capabilities.

🔑 KEY FEATURES:
• Pattern matching and searching
• Text replacement and substitution
• String validation
• Data extraction from text

💡 REAL-WORLD ANALOGY:
Regex = Search Filter
- Pattern = Search criteria (like "find all phone numbers")
- Match = Items that meet the criteria
- Groups = Specific parts you want to extract
- Flags = Search options (case-sensitive, etc.)

🎯 WHY USE REGEX?
• Validate input formats (email, phone, etc.)
• Extract data from unstructured text
• Find and replace complex patterns
• Parse log files and data

=============================================================================
*/

package main

import (
	"fmt"
	"regexp"
	"strings"
)

func main() {
	fmt.Println("🔍 REGULAR EXPRESSIONS TUTORIAL")
	fmt.Println("===============================")

	// 🎯 DEMO 1: Basic Pattern Matching
	fmt.Println("\n🎯 DEMO 1: Basic Pattern Matching")
	fmt.Println("=================================")

	text := "The quick brown fox jumps over the lazy dog"
	
	// Simple string matching
	matched, _ := regexp.MatchString("fox", text)
	fmt.Printf("Text contains 'fox': %t\n", matched)

	// Case-insensitive matching
	matched, _ = regexp.MatchString("(?i)FOX", text)
	fmt.Printf("Text contains 'FOX' (case-insensitive): %t\n", matched)

	// Word boundary matching
	re := regexp.MustCompile(`\bthe\b`)
	matches := re.FindAllString(text, -1)
	fmt.Printf("Word 'the' found %d times: %v\n", len(matches), matches)

	// 🎯 DEMO 2: Email Validation
	fmt.Println("\n🎯 DEMO 2: Email Validation")
	fmt.Println("===========================")

	emailPattern := `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`
	emailRegex := regexp.MustCompile(emailPattern)

	emails := []string{
		"user@example.com",
		"test.email+tag@domain.co.uk",
		"invalid.email",
		"@invalid.com",
		"user@.com",
		"valid123@test-domain.org",
	}

	fmt.Println("Email validation results:")
	for _, email := range emails {
		isValid := emailRegex.MatchString(email)
		status := "❌ Invalid"
		if isValid {
			status = "✅ Valid"
		}
		fmt.Printf("  %-30s %s\n", email, status)
	}

	// 🎯 DEMO 3: Phone Number Extraction
	fmt.Println("\n🎯 DEMO 3: Phone Number Extraction")
	fmt.Println("==================================")

	phoneText := `
	Contact us at (555) 123-4567 or call 555.987.6543.
	International: +1-800-555-0199
	Mobile: 555 111 2222
	`

	// Different phone number patterns
	phonePatterns := []string{
		`\(\d{3}\) \d{3}-\d{4}`,           // (555) 123-4567
		`\d{3}\.\d{3}\.\d{4}`,             // 555.987.6543
		`\+1-\d{3}-\d{3}-\d{4}`,           // +1-800-555-0199
		`\d{3} \d{3} \d{4}`,               // 555 111 2222
	}

	fmt.Println("Found phone numbers:")
	for _, pattern := range phonePatterns {
		re := regexp.MustCompile(pattern)
		matches := re.FindAllString(phoneText, -1)
		for _, match := range matches {
			fmt.Printf("  📞 %s\n", match)
		}
	}

	// 🎯 DEMO 4: Data Extraction with Groups
	fmt.Println("\n🎯 DEMO 4: Data Extraction with Groups")
	fmt.Println("======================================")

	logEntry := "2023-12-01 14:30:25 [ERROR] Failed to connect to database: connection timeout"
	
	// Pattern with named groups
	logPattern := `(?P<date>\d{4}-\d{2}-\d{2}) (?P<time>\d{2}:\d{2}:\d{2}) \[(?P<level>\w+)\] (?P<message>.*)`
	logRegex := regexp.MustCompile(logPattern)

	matches := logRegex.FindStringSubmatch(logEntry)
	if len(matches) > 0 {
		fmt.Println("Parsed log entry:")
		names := logRegex.SubexpNames()
		for i, match := range matches {
			if i > 0 && names[i] != "" {
				fmt.Printf("  %s: %s\n", names[i], match)
			}
		}
	}

	// 🎯 DEMO 5: Text Replacement
	fmt.Println("\n🎯 DEMO 5: Text Replacement")
	fmt.Println("===========================")

	originalText := "The price is $19.99 and the discount is $5.00"
	
	// Replace all dollar amounts
	priceRegex := regexp.MustCompile(`\$(\d+\.\d{2})`)
	
	// Simple replacement
	replaced := priceRegex.ReplaceAllString(originalText, "€$1")
	fmt.Printf("Original: %s\n", originalText)
	fmt.Printf("Replaced: %s\n", replaced)

	// Complex replacement with function
	complexReplaced := priceRegex.ReplaceAllStringFunc(originalText, func(match string) string {
		// Remove $ and convert to float, then format as EUR
		amount := strings.TrimPrefix(match, "$")
		return fmt.Sprintf("€%s EUR", amount)
	})
	fmt.Printf("Complex:  %s\n", complexReplaced)

	// 🎯 DEMO 6: URL Parsing
	fmt.Println("\n🎯 DEMO 6: URL Parsing")
	fmt.Println("======================")

	urls := []string{
		"https://www.example.com/path/to/page?param=value",
		"http://subdomain.test.org:8080/api/v1/users",
		"ftp://files.company.com/documents/file.pdf",
	}

	urlPattern := `(?P<protocol>\w+)://(?P<host>[^:/]+)(?::(?P<port>\d+))?(?P<path>/[^?]*)?(?:\?(?P<query>.*))?`
	urlRegex := regexp.MustCompile(urlPattern)

	for _, url := range urls {
		fmt.Printf("\nParsing URL: %s\n", url)
		matches := urlRegex.FindStringSubmatch(url)
		if len(matches) > 0 {
			names := urlRegex.SubexpNames()
			for i, match := range matches {
				if i > 0 && names[i] != "" && match != "" {
					fmt.Printf("  %s: %s\n", names[i], match)
				}
			}
		}
	}

	// 🎯 DEMO 7: Word and Character Classes
	fmt.Println("\n🎯 DEMO 7: Character Classes")
	fmt.Println("============================")

	testText := "Hello123 World! @#$ test_case 2023-12-01"

	patterns := map[string]string{
		`\d+`:           "Digits",
		`\w+`:           "Word characters",
		`\s+`:           "Whitespace",
		`[A-Z]+`:        "Uppercase letters",
		`[a-z]+`:        "Lowercase letters",
		`[!@#$%^&*]+`:   "Special characters",
		`\b\w{4}\b`:     "4-letter words",
		`\d{4}-\d{2}-\d{2}`: "Date format",
	}

	for pattern, description := range patterns {
		re := regexp.MustCompile(pattern)
		matches := re.FindAllString(testText, -1)
		fmt.Printf("%-20s: %v\n", description, matches)
	}

	// 🎯 DEMO 8: Validation Functions
	fmt.Println("\n🎯 DEMO 8: Common Validation Functions")
	fmt.Println("=====================================")

	testData := map[string][]string{
		"emails": {
			"valid@example.com",
			"invalid.email",
			"test@domain.co.uk",
		},
		"phones": {
			"(555) 123-4567",
			"555-1234",
			"+1-800-555-0199",
		},
		"urls": {
			"https://www.example.com",
			"not-a-url",
			"http://localhost:8080",
		},
		"ips": {
			"192.168.1.1",
			"256.1.1.1",
			"10.0.0.1",
		},
	}

	validators := map[string]*regexp.Regexp{
		"emails": regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`),
		"phones": regexp.MustCompile(`^(\+1-)?(\(?\d{3}\)?[-.\s]?)?\d{3}[-.\s]?\d{4}$`),
		"urls":   regexp.MustCompile(`^https?://[^\s/$.?#].[^\s]*$`),
		"ips":    regexp.MustCompile(`^((25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)$`),
	}

	for category, items := range testData {
		fmt.Printf("\n%s validation:\n", strings.Title(category))
		validator := validators[category]
		for _, item := range items {
			isValid := validator.MatchString(item)
			status := "❌"
			if isValid {
				status = "✅"
			}
			fmt.Printf("  %s %-25s\n", status, item)
		}
	}

	// 🎯 DEMO 9: Performance Comparison
	fmt.Println("\n🎯 DEMO 9: Performance Tips")
	fmt.Println("===========================")

	fmt.Println("💡 Regex Performance Tips:")
	fmt.Println("  • Compile regex once, use many times")
	fmt.Println("  • Use MustCompile for patterns known at compile time")
	fmt.Println("  • Avoid complex nested groups when possible")
	fmt.Println("  • Use non-capturing groups (?:...) when you don't need the match")
	fmt.Println("  • Consider strings package for simple operations")

	// Example: Compiled vs non-compiled
	pattern := `\d+`
	text_sample := "There are 123 numbers and 456 more numbers"

	// Compiled (efficient for multiple uses)
	compiledRegex := regexp.MustCompile(pattern)
	matches1 := compiledRegex.FindAllString(text_sample, -1)
	fmt.Printf("\nCompiled regex found: %v\n", matches1)

	// Non-compiled (less efficient for repeated use)
	matches2, _ := regexp.FindAllString(pattern, text_sample, -1)
	fmt.Printf("Direct regex found: %v\n", matches2)

	fmt.Println("\n✨ All regex demos completed!")
}

/*
=============================================================================
                              📝 LEARNING NOTES
=============================================================================

🔍 REGEX BASICS:
┌─────────────────────────────────────────────────────────────────────────┐
│ // Compile regex pattern                                                │
│ re := regexp.MustCompile(`pattern`)                                     │
│ re, err := regexp.Compile(`pattern`)                                    │
│                                                                         │
│ // Basic matching                                                       │
│ matched := re.MatchString("text")                                       │
│ matched, err := regexp.MatchString(`pattern`, "text")                   │
│                                                                         │
│ // Find matches                                                         │
│ match := re.FindString("text")           // First match                 │
│ matches := re.FindAllString("text", -1)  // All matches                 │
│ matches := re.FindAllString("text", 2)   // First 2 matches             │
└─────────────────────────────────────────────────────────────────────────┘

🎯 COMMON PATTERNS:
┌─────────────────┬─────────────────────────────────────────────────────────┐
│   Pattern       │                Description                              │
├─────────────────┼─────────────────────────────────────────────────────────┤
│ .               │ Any character except newline                            │
│ \d              │ Any digit (0-9)                                         │
│ \w              │ Any word character (a-z, A-Z, 0-9, _)                   │
│ \s              │ Any whitespace character                                │
│ ^               │ Start of string                                         │
│ $               │ End of string                                           │
│ *               │ Zero or more of preceding                               │
│ +               │ One or more of preceding                                │
│ ?               │ Zero or one of preceding                                │
│ {n}             │ Exactly n of preceding                                  │
│ {n,m}           │ Between n and m of preceding                            │
│ [abc]           │ Any of a, b, or c                                       │
│ [a-z]           │ Any lowercase letter                                    │
│ [^abc]          │ Any character except a, b, or c                         │
│ \b              │ Word boundary                                           │
│ |               │ OR operator                                             │
│ ()              │ Capturing group                                         │
│ (?:...)         │ Non-capturing group                                     │
│ (?P<name>...)   │ Named capturing group                                   │
└─────────────────┴─────────────────────────────────────────────────────────┘

📋 REGEX METHODS:
┌─────────────────────────────────────────────────────────────────────────┐
│ // Matching                                                             │
│ re.MatchString(s)                    // Returns bool                    │
│ re.Match([]byte)                     // Returns bool                    │
│                                                                         │
│ // Finding                                                              │
│ re.FindString(s)                     // First match as string           │
│ re.FindAllString(s, n)               // All matches as []string         │
│ re.FindStringSubmatch(s)             // Match with submatches           │
│ re.FindAllStringSubmatch(s, n)       // All matches with submatches     │
│                                                                         │
│ // Replacing                                                            │
│ re.ReplaceAllString(s, repl)         // Replace with string             │
│ re.ReplaceAllStringFunc(s, func)     // Replace with function result    │
│                                                                         │
│ // Splitting                                                            │
│ re.Split(s, n)                       // Split string by pattern         │
└─────────────────────────────────────────────────────────────────────────┘

🔧 ADVANCED FEATURES:
┌─────────────────────────────────────────────────────────────────────────┐
│ // Flags (use with (?flags) at start of pattern)                        │
│ (?i)        // Case insensitive                                         │
│ (?m)        // Multiline mode (^ and $ match line boundaries)           │
│ (?s)        // Dot matches newline                                      │
│                                                                         │
│ // Named groups                                                         │
│ pattern := `(?P<year>\d{4})-(?P<month>\d{2})-(?P<day>\d{2})`            │
│ matches := re.FindStringSubmatch(text)                                  │
│ names := re.SubexpNames()                                               │
│                                                                         │
│ // Lookahead/Lookbehind (limited support)                               │
│ (?=...)     // Positive lookahead                                       │
│ (?!...)     // Negative lookahead                                       │
└─────────────────────────────────────────────────────────────────────────┘

📝 COMMON VALIDATION PATTERNS:
┌─────────────────────────────────────────────────────────────────────────┐
│ // Email                                                                │
│ ^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$                       │
│                                                                         │
│ // Phone (US format)                                                    │
│ ^\(?([0-9]{3})\)?[-. ]?([0-9]{3})[-. ]?([0-9]{4})$                      │
│                                                                         │
│ // URL                                                                  │
│ ^https?://[^\s/$.?#].[^\s]*$                                            │
│                                                                         │
│ // IP Address                                                           │
│ ^((25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)$ │
│                                                                         │
│ // Date (YYYY-MM-DD)                                                    │
│ ^\d{4}-\d{2}-\d{2}$                                                     │
│                                                                         │
│ // Credit Card                                                          │
│ ^(?:4[0-9]{12}(?:[0-9]{3})?|5[1-5][0-9]{14}|3[47][0-9]{13})$           │
└─────────────────────────────────────────────────────────────────────────┘

💡 BEST PRACTICES:
• Compile regex once, use multiple times
• Use raw strings (backticks) for patterns
• Use MustCompile for patterns known at compile time
• Test regex patterns thoroughly
• Use named groups for complex patterns
• Consider performance for large texts
• Escape special characters when needed

🚨 COMMON MISTAKES:
❌ Not escaping special characters
❌ Greedy vs non-greedy matching confusion
❌ Forgetting to handle empty matches
❌ Overcomplicating simple string operations
❌ Not validating regex compilation errors

⚡ PERFORMANCE TIPS:
• Use strings package for simple operations
• Compile regex patterns once
• Use non-capturing groups (?:...) when possible
• Avoid complex nested patterns
• Consider alternatives for very simple patterns
• Profile regex-heavy code

🎯 WHEN TO USE REGEX:
✅ Pattern validation (email, phone, etc.)
✅ Data extraction from text
✅ Complex find and replace operations
✅ Log parsing and analysis
✅ Input sanitization

❌ Simple string operations (use strings package)
❌ Parsing structured data (use proper parsers)
❌ Performance-critical simple matching

=============================================================================
*/