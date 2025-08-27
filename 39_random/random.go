/*
=============================================================================
                           🎲 GO RANDOM TUTORIAL
=============================================================================

📚 CORE CONCEPT:
Go's math/rand package provides pseudorandom number generation.
It includes functions for generating random integers, floats, and
selecting random elements from collections.

🔑 KEY FEATURES:
• Pseudorandom number generation
• Seeding for reproducible results
• Various distributions (uniform, normal)
• Random selection from collections
• Shuffling algorithms

💡 REAL-WORLD ANALOGY:
Random Package = Dice and Card Deck
- Seed = Starting position of shuffled deck
- Rand.Int() = Rolling a die
- Shuffle = Mixing cards
- Source = The randomness mechanism

🎯 WHY USE RANDOM?
• Simulations and modeling
• Game development
• Testing with random data
• Cryptographic applications (with crypto/rand)

=============================================================================
*/

package main

import (
	"crypto/rand"
	"fmt"
	"math/big"
	mathRand "math/rand"
	"time"
)

func main() {
	fmt.Println("🎲 RANDOM TUTORIAL")
	fmt.Println("==================")

	// 🎯 DEMO 1: Basic Random Numbers
	fmt.Println("\n🎯 DEMO 1: Basic Random Numbers")
	fmt.Println("===============================")

	// Seed the random number generator
	mathRand.Seed(time.Now().UnixNano())

	fmt.Println("Random integers:")
	for i := 0; i < 5; i++ {
		fmt.Printf("  Random int: %d\n", mathRand.Int())
	}

	fmt.Println("\nRandom integers in range [0, 100):")
	for i := 0; i < 5; i++ {
		fmt.Printf("  Random int [0,100): %d\n", mathRand.Intn(100))
	}

	fmt.Println("\nRandom floats [0.0, 1.0):")
	for i := 0; i < 5; i++ {
		fmt.Printf("  Random float: %.6f\n", mathRand.Float64())
	}

	// 🎯 DEMO 2: Random Numbers with Custom Ranges
	fmt.Println("\n🎯 DEMO 2: Custom Ranges")
	fmt.Println("========================")

	// Random integers in custom range [min, max]
	min, max := 10, 50
	fmt.Printf("Random integers in range [%d, %d]:\n", min, max)
	for i := 0; i < 5; i++ {
		randomInt := mathRand.Intn(max-min+1) + min
		fmt.Printf("  %d\n", randomInt)
	}

	// Random floats in custom range
	minFloat, maxFloat := 1.5, 10.5
	fmt.Printf("\nRandom floats in range [%.1f, %.1f]:\n", minFloat, maxFloat)
	for i := 0; i < 5; i++ {
		randomFloat := mathRand.Float64()*(maxFloat-minFloat) + minFloat
		fmt.Printf("  %.3f\n", randomFloat)
	}

	// 🎯 DEMO 3: Random Selection from Collections
	fmt.Println("\n🎯 DEMO 3: Random Selection")
	fmt.Println("===========================")

	// Random selection from slice
	colors := []string{"red", "green", "blue", "yellow", "purple", "orange"}
	fmt.Println("Random colors:")
	for i := 0; i < 5; i++ {
		randomColor := colors[mathRand.Intn(len(colors))]
		fmt.Printf("  %s\n", randomColor)
	}

	// Random selection from map
	fruits := map[string]int{
		"apple":  5,
		"banana": 3,
		"orange": 8,
		"grape":  12,
	}

	// Convert map keys to slice for random selection
	fruitNames := make([]string, 0, len(fruits))
	for name := range fruits {
		fruitNames = append(fruitNames, name)
	}

	fmt.Println("\nRandom fruits:")
	for i := 0; i < 3; i++ {
		randomFruit := fruitNames[mathRand.Intn(len(fruitNames))]
		count := fruits[randomFruit]
		fmt.Printf("  %s: %d\n", randomFruit, count)
	}

	// 🎯 DEMO 4: Shuffling
	fmt.Println("\n🎯 DEMO 4: Shuffling")
	fmt.Println("====================")

	// Shuffle a slice
	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Printf("Original: %v\n", numbers)

	// Fisher-Yates shuffle
	mathRand.Shuffle(len(numbers), func(i, j int) {
		numbers[i], numbers[j] = numbers[j], numbers[i]
	})
	fmt.Printf("Shuffled: %v\n", numbers)

	// Shuffle strings
	words := []string{"hello", "world", "go", "programming", "random"}
	fmt.Printf("\nOriginal words: %v\n", words)
	mathRand.Shuffle(len(words), func(i, j int) {
		words[i], words[j] = words[j], words[i]
	})
	fmt.Printf("Shuffled words: %v\n", words)

	// 🎯 DEMO 5: Random Booleans and Choices
	fmt.Println("\n🎯 DEMO 5: Random Booleans and Choices")
	fmt.Println("======================================")

	// Random booleans
	fmt.Println("Random booleans (50/50 chance):")
	for i := 0; i < 10; i++ {
		randomBool := mathRand.Intn(2) == 1
		fmt.Printf("  %t", randomBool)
	}
	fmt.Println()

	// Weighted random choice
	fmt.Println("\nWeighted random choice (70% true, 30% false):")
	for i := 0; i < 10; i++ {
		weightedBool := mathRand.Float64() < 0.7
		fmt.Printf("  %t", weightedBool)
	}
	fmt.Println()

	// 🎯 DEMO 6: Random Strings
	fmt.Println("\n🎯 DEMO 6: Random Strings")
	fmt.Println("=========================")

	// Generate random string
	charset := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	
	fmt.Println("Random strings:")
	for i := 0; i < 5; i++ {
		length := 8
		randomString := make([]byte, length)
		for j := range randomString {
			randomString[j] = charset[mathRand.Intn(len(charset))]
		}
		fmt.Printf("  %s\n", string(randomString))
	}

	// Random passwords
	fmt.Println("\nRandom passwords:")
	for i := 0; i < 3; i++ {
		password := generatePassword(12)
		fmt.Printf("  %s\n", password)
	}

	// 🎯 DEMO 7: Seeded Random for Reproducibility
	fmt.Println("\n🎯 DEMO 7: Seeded Random")
	fmt.Println("========================")

	// Same seed produces same sequence
	seed := int64(12345)
	
	fmt.Printf("Sequence 1 (seed %d):\n", seed)
	mathRand.Seed(seed)
	for i := 0; i < 5; i++ {
		fmt.Printf("  %d", mathRand.Intn(100))
	}
	fmt.Println()

	fmt.Printf("Sequence 2 (same seed %d):\n", seed)
	mathRand.Seed(seed)
	for i := 0; i < 5; i++ {
		fmt.Printf("  %d", mathRand.Intn(100))
	}
	fmt.Println()

	// 🎯 DEMO 8: Normal Distribution
	fmt.Println("\n🎯 DEMO 8: Normal Distribution")
	fmt.Println("==============================")

	fmt.Println("Normal distribution (mean=0, stddev=1):")
	for i := 0; i < 10; i++ {
		normal := mathRand.NormFloat64()
		fmt.Printf("  %.3f", normal)
	}
	fmt.Println()

	// Custom normal distribution
	mean, stddev := 100.0, 15.0
	fmt.Printf("\nCustom normal (mean=%.0f, stddev=%.0f):\n", mean, stddev)
	for i := 0; i < 10; i++ {
		customNormal := mathRand.NormFloat64()*stddev + mean
		fmt.Printf("  %.1f", customNormal)
	}
	fmt.Println()

	// 🎯 DEMO 9: Exponential Distribution
	fmt.Println("\n🎯 DEMO 9: Exponential Distribution")
	fmt.Println("===================================")

	rate := 2.0
	fmt.Printf("Exponential distribution (rate=%.1f):\n", rate)
	for i := 0; i < 10; i++ {
		exponential := mathRand.ExpFloat64() / rate
		fmt.Printf("  %.3f", exponential)
	}
	fmt.Println()

	// 🎯 DEMO 10: Cryptographically Secure Random
	fmt.Println("\n🎯 DEMO 10: Cryptographically Secure Random")
	fmt.Println("===========================================")

	fmt.Println("Cryptographically secure random numbers:")
	for i := 0; i < 5; i++ {
		// Generate random number in range [0, 100)
		n, err := rand.Int(rand.Reader, big.NewInt(100))
		if err != nil {
			fmt.Printf("Error: %v\n", err)
			continue
		}
		fmt.Printf("  %s\n", n.String())
	}

	// Secure random bytes
	fmt.Println("\nSecure random bytes (hex):")
	for i := 0; i < 3; i++ {
		bytes := make([]byte, 8)
		_, err := rand.Read(bytes)
		if err != nil {
			fmt.Printf("Error: %v\n", err)
			continue
		}
		fmt.Printf("  %x\n", bytes)
	}

	fmt.Println("\n✨ All random demos completed!")
}

// 🔧 UTILITY FUNCTION: Generate random password
func generatePassword(length int) string {
	lowercase := "abcdefghijklmnopqrstuvwxyz"
	uppercase := "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	digits := "0123456789"
	special := "!@#$%^&*"
	
	allChars := lowercase + uppercase + digits + special
	
	password := make([]byte, length)
	
	// Ensure at least one character from each category
	password[0] = lowercase[mathRand.Intn(len(lowercase))]
	password[1] = uppercase[mathRand.Intn(len(uppercase))]
	password[2] = digits[mathRand.Intn(len(digits))]
	password[3] = special[mathRand.Intn(len(special))]
	
	// Fill the rest randomly
	for i := 4; i < length; i++ {
		password[i] = allChars[mathRand.Intn(len(allChars))]
	}
	
	// Shuffle the password
	mathRand.Shuffle(len(password), func(i, j int) {
		password[i], password[j] = password[j], password[i]
	})
	
	return string(password)
}

/*
=============================================================================
                              📝 LEARNING NOTES
=============================================================================

🎲 BASIC RANDOM FUNCTIONS:
┌─────────────────────────────────────────────────────────────────────────┐
│ // Integer random numbers                                               │
│ rand.Int()           // Random int (full range)                         │
│ rand.Intn(n)         // Random int [0, n)                               │
│ rand.Int31()         // Random int31                                    │
│ rand.Int31n(n)       // Random int31 [0, n)                             │
│ rand.Int63()         // Random int64                                    │
│ rand.Int63n(n)       // Random int64 [0, n)                             │
│                                                                         │
│ // Float random numbers                                                 │
│ rand.Float32()       // Random float32 [0.0, 1.0)                       │
│ rand.Float64()       // Random float64 [0.0, 1.0)                       │
│                                                                         │
│ // Other types                                                          │
│ rand.Uint32()        // Random uint32                                   │
│ rand.Uint64()        // Random uint64                                   │
└─────────────────────────────────────────────────────────────────────────┘

🌱 SEEDING:
┌─────────────────────────────────────────────────────────────────────────┐
│ // Seed with current time (common pattern)                              │
│ rand.Seed(time.Now().UnixNano())                                        │
│                                                                         │
│ // Seed with fixed value (for reproducible results)                     │
│ rand.Seed(12345)                                                        │
│                                                                         │
│ // Create custom source                                                 │
│ source := rand.NewSource(time.Now().UnixNano())                         │
│ rng := rand.New(source)                                                 │
└─────────────────────────────────────────────────────────────────────────┘

📊 DISTRIBUTIONS:
┌─────────────────────────────────────────────────────────────────────────┐
│ // Normal (Gaussian) distribution                                       │
│ rand.NormFloat64()   // Mean=0, StdDev=1                                │
│ customNormal := rand.NormFloat64()*stddev + mean                        │
│                                                                         │
│ // Exponential distribution                                             │
│ rand.ExpFloat64()    // Rate=1                                          │
│ customExp := rand.ExpFloat64() / rate                                   │
│                                                                         │
│ // Uniform distribution                                                 │
│ rand.Float64()       // [0.0, 1.0)                                      │
│ custom := rand.Float64()*(max-min) + min                                │
└─────────────────────────────────────────────────────────────────────────┘

🔀 SHUFFLING:
┌─────────────────────────────────────────────────────────────────────────┐
│ // Shuffle slice                                                        │
│ rand.Shuffle(len(slice), func(i, j int) {                               │
│     slice[i], slice[j] = slice[j], slice[i]                             │
│ })                                                                      │
│                                                                         │
│ // Manual Fisher-Yates shuffle                                          │
│ for i := len(slice) - 1; i > 0; i-- {                                   │
│     j := rand.Intn(i + 1)                                               │
│     slice[i], slice[j] = slice[j], slice[i]                             │
│ }                                                                       │
└─────────────────────────────────────────────────────────────────────────┘

🔐 CRYPTOGRAPHICALLY SECURE:
┌─────────────────────────────────────────────────────────────────────────┐
│ import "crypto/rand"                                                    │
│                                                                         │
│ // Random bytes                                                         │
│ bytes := make([]byte, 32)                                               │
│ _, err := rand.Read(bytes)                                              │
│                                                                         │
│ // Random big integer                                                   │
│ n, err := rand.Int(rand.Reader, big.NewInt(100))                        │
│                                                                         │
│ // Random prime                                                         │
│ prime, err := rand.Prime(rand.Reader, 1024)                             │
└─────────────────────────────────────────────────────────────────────────┘

🎯 COMMON PATTERNS:
┌─────────────────────────────────────────────────────────────────────────┐
│ // Random element from slice                                            │
│ element := slice[rand.Intn(len(slice))]                                 │
│                                                                         │
│ // Random boolean                                                       │
│ randomBool := rand.Intn(2) == 1                                         │
│                                                                         │
│ // Weighted random                                                      │
│ weighted := rand.Float64() < probability                                │
│                                                                         │
│ // Random string                                                        │
│ charset := "abcdefghijklmnopqrstuvwxyz"                                  │
│ result := make([]byte, length)                                          │
│ for i := range result {                                                 │
│     result[i] = charset[rand.Intn(len(charset))]                        │
│ }                                                                       │
│                                                                         │
│ // Random range [min, max]                                              │
│ randomInt := rand.Intn(max-min+1) + min                                 │
│ randomFloat := rand.Float64()*(max-min) + min                           │
└─────────────────────────────────────────────────────────────────────────┘

💡 BEST PRACTICES:
• Seed the generator once at program start
• Use crypto/rand for security-sensitive applications
• Create separate rand.Rand instances for concurrent use
• Use appropriate distribution for your use case
• Consider performance implications of random operations

🚨 COMMON MISTAKES:
❌ Not seeding the random generator
❌ Using math/rand for cryptographic purposes
❌ Seeding multiple times unnecessarily
❌ Not handling errors from crypto/rand
❌ Using global rand in concurrent programs without synchronization

⚡ PERFORMANCE TIPS:
• Create dedicated rand.Rand instances for hot paths
• Cache random values when appropriate
• Use simpler distributions when possible
• Consider lookup tables for complex distributions
• Profile random-heavy code

🎯 WHEN TO USE EACH:
• math/rand: Simulations, games, testing, non-security applications
• crypto/rand: Passwords, tokens, cryptographic keys, security-sensitive data
• Custom sources: When you need reproducible sequences or specific properties

🔒 SECURITY CONSIDERATIONS:
• Never use math/rand for passwords, tokens, or keys
• Always use crypto/rand for security-sensitive applications
• Be aware that math/rand is predictable if seed is known
• Consider entropy sources for seeding

=============================================================================
*/