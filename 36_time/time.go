/*
=============================================================================
                           ⏰ GO TIME TUTORIAL
=============================================================================

📚 CORE CONCEPT:
Go's time package provides functionality for measuring and displaying time.
It handles time zones, formatting, parsing, and arithmetic operations.

🔑 KEY FEATURES:
• Time creation and manipulation
• Formatting and parsing
• Time zones and locations
• Duration calculations
• Timers and tickers

💡 REAL-WORLD ANALOGY:
Time Package = Swiss Watch
- Time = Current moment on the watch
- Duration = Time interval measurement
- Location = Time zone setting
- Format = How you display the time
- Timer = Alarm functionality

🎯 WHY LEARN TIME?
• Handle timestamps in applications
• Schedule tasks and events
• Measure performance and duration
• Work with different time zones

=============================================================================
*/

package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("⏰ TIME TUTORIAL")
	fmt.Println("================")

	// 🎯 DEMO 1: Creating Time Values
	fmt.Println("\n🎯 DEMO 1: Creating Time Values")
	fmt.Println("===============================")

	// Current time
	now := time.Now()
	fmt.Printf("Current time: %v\n", now)
	fmt.Printf("UTC time: %v\n", now.UTC())

	// Specific time
	specificTime := time.Date(2023, time.December, 25, 15, 30, 45, 0, time.UTC)
	fmt.Printf("Christmas 2023: %v\n", specificTime)

	// Unix timestamp
	unixTime := time.Unix(1640995200, 0) // Jan 1, 2022 00:00:00 UTC
	fmt.Printf("Unix timestamp: %v\n", unixTime)

	// Parse time from string
	parsed, err := time.Parse("2006-01-02 15:04:05", "2023-12-01 14:30:25")
	if err != nil {
		fmt.Printf("Parse error: %v\n", err)
	} else {
		fmt.Printf("Parsed time: %v\n", parsed)
	}

	// 🎯 DEMO 2: Time Formatting
	fmt.Println("\n🎯 DEMO 2: Time Formatting")
	fmt.Println("==========================")

	now = time.Now()

	// Go's reference time: Mon Jan 2 15:04:05 MST 2006 (Unix: 1136239445)
	formats := map[string]string{
		"RFC3339":     time.RFC3339,
		"RFC822":      time.RFC822,
		"Kitchen":     time.Kitchen,
		"Stamp":       time.Stamp,
		"Custom 1":    "2006-01-02 15:04:05",
		"Custom 2":    "Jan 2, 2006 at 3:04 PM",
		"Custom 3":    "Monday, January 2, 2006",
		"ISO 8601":    "2006-01-02T15:04:05Z07:00",
		"US Format":   "01/02/2006 03:04:05 PM",
		"Unix":        "Mon Jan _2 15:04:05 MST 2006",
	}

	fmt.Println("Time formatting examples:")
	for name, format := range formats {
		formatted := now.Format(format)
		fmt.Printf("  %-12s: %s\n", name, formatted)
	}

	// 🎯 DEMO 3: Time Parsing
	fmt.Println("\n🎯 DEMO 3: Time Parsing")
	fmt.Println("=======================")

	timeStrings := []struct {
		input  string
		layout string
	}{
		{"2023-12-01 14:30:25", "2006-01-02 15:04:05"},
		{"Dec 1, 2023", "Jan 2, 2006"},
		{"01/12/2023", "01/02/2006"},
		{"2023-12-01T14:30:25Z", time.RFC3339},
		{"Fri, 01 Dec 2023 14:30:25 GMT", time.RFC1123},
	}

	fmt.Println("Parsing time strings:")
	for _, ts := range timeStrings {
		parsed, err := time.Parse(ts.layout, ts.input)
		if err != nil {
			fmt.Printf("  ❌ '%s': %v\n", ts.input, err)
		} else {
			fmt.Printf("  ✅ '%s' → %v\n", ts.input, parsed)
		}
	}

	// 🎯 DEMO 4: Duration Operations
	fmt.Println("\n🎯 DEMO 4: Duration Operations")
	fmt.Println("==============================")

	// Creating durations
	durations := []time.Duration{
		time.Nanosecond,
		time.Microsecond,
		time.Millisecond,
		time.Second,
		time.Minute,
		time.Hour,
		24 * time.Hour, // 1 day
		7 * 24 * time.Hour, // 1 week
	}

	fmt.Println("Duration examples:")
	for _, d := range durations {
		fmt.Printf("  %v = %v nanoseconds\n", d, d.Nanoseconds())
	}

	// Duration arithmetic
	start := time.Now()
	time.Sleep(100 * time.Millisecond) // Simulate work
	elapsed := time.Since(start)
	fmt.Printf("\nElapsed time: %v\n", elapsed)

	// Duration parsing
	customDuration, err := time.ParseDuration("2h45m30s")
	if err != nil {
		fmt.Printf("Duration parse error: %v\n", err)
	} else {
		fmt.Printf("Parsed duration: %v\n", customDuration)
		fmt.Printf("In hours: %.2f\n", customDuration.Hours())
		fmt.Printf("In minutes: %.2f\n", customDuration.Minutes())
		fmt.Printf("In seconds: %.2f\n", customDuration.Seconds())
	}

	// 🎯 DEMO 5: Time Arithmetic
	fmt.Println("\n🎯 DEMO 5: Time Arithmetic")
	fmt.Println("==========================")

	baseTime := time.Date(2023, 12, 1, 12, 0, 0, 0, time.UTC)
	fmt.Printf("Base time: %v\n", baseTime)

	// Adding durations
	operations := []struct {
		name     string
		duration time.Duration
	}{
		{"Add 1 hour", time.Hour},
		{"Add 30 minutes", 30 * time.Minute},
		{"Add 1 day", 24 * time.Hour},
		{"Add 1 week", 7 * 24 * time.Hour},
		{"Subtract 2 hours", -2 * time.Hour},
	}

	for _, op := range operations {
		result := baseTime.Add(op.duration)
		fmt.Printf("  %s: %v\n", op.name, result.Format("2006-01-02 15:04:05"))
	}

	// Time difference
	time1 := time.Date(2023, 12, 1, 10, 0, 0, 0, time.UTC)
	time2 := time.Date(2023, 12, 1, 15, 30, 0, 0, time.UTC)
	diff := time2.Sub(time1)
	fmt.Printf("\nTime difference: %v\n", diff)
	fmt.Printf("Difference in hours: %.2f\n", diff.Hours())

	// 🎯 DEMO 6: Time Zones
	fmt.Println("\n🎯 DEMO 6: Time Zones")
	fmt.Println("=====================")

	// Load different time zones
	locations := []string{
		"UTC",
		"America/New_York",
		"Europe/London",
		"Asia/Tokyo",
		"Australia/Sydney",
	}

	currentTime := time.Now()
	fmt.Printf("Current time in different zones:\n")

	for _, locName := range locations {
		loc, err := time.LoadLocation(locName)
		if err != nil {
			fmt.Printf("  ❌ %s: %v\n", locName, err)
			continue
		}
		
		localTime := currentTime.In(loc)
		fmt.Printf("  %-20s: %s\n", locName, localTime.Format("2006-01-02 15:04:05 MST"))
	}

	// 🎯 DEMO 7: Time Comparisons
	fmt.Println("\n🎯 DEMO 7: Time Comparisons")
	fmt.Println("===========================")

	t1 := time.Date(2023, 12, 1, 10, 0, 0, 0, time.UTC)
	t2 := time.Date(2023, 12, 1, 15, 0, 0, 0, time.UTC)
	t3 := time.Date(2023, 12, 1, 10, 0, 0, 0, time.UTC)

	fmt.Printf("t1: %v\n", t1.Format("15:04:05"))
	fmt.Printf("t2: %v\n", t2.Format("15:04:05"))
	fmt.Printf("t3: %v\n", t3.Format("15:04:05"))
	fmt.Println()

	fmt.Printf("t1.Before(t2): %t\n", t1.Before(t2))
	fmt.Printf("t1.After(t2): %t\n", t1.After(t2))
	fmt.Printf("t1.Equal(t3): %t\n", t1.Equal(t3))
	fmt.Printf("t2.After(t1): %t\n", t2.After(t1))

	// 🎯 DEMO 8: Time Components
	fmt.Println("\n🎯 DEMO 8: Time Components")
	fmt.Println("==========================")

	sampleTime := time.Date(2023, 12, 25, 15, 30, 45, 123456789, time.UTC)
	fmt.Printf("Sample time: %v\n", sampleTime)
	fmt.Println()

	fmt.Printf("Year: %d\n", sampleTime.Year())
	fmt.Printf("Month: %s (%d)\n", sampleTime.Month(), int(sampleTime.Month()))
	fmt.Printf("Day: %d\n", sampleTime.Day())
	fmt.Printf("Hour: %d\n", sampleTime.Hour())
	fmt.Printf("Minute: %d\n", sampleTime.Minute())
	fmt.Printf("Second: %d\n", sampleTime.Second())
	fmt.Printf("Nanosecond: %d\n", sampleTime.Nanosecond())
	fmt.Printf("Weekday: %s\n", sampleTime.Weekday())
	fmt.Printf("Year day: %d\n", sampleTime.YearDay())

	// ISO week
	year, week := sampleTime.ISOWeek()
	fmt.Printf("ISO Week: %d-W%02d\n", year, week)

	// 🎯 DEMO 9: Practical Examples
	fmt.Println("\n🎯 DEMO 9: Practical Examples")
	fmt.Println("=============================")

	// Age calculation
	birthDate := time.Date(1990, 5, 15, 0, 0, 0, 0, time.UTC)
	age := time.Since(birthDate)
	years := int(age.Hours() / 24 / 365.25)
	fmt.Printf("Age calculation: %d years old\n", years)

	// Business days calculation (simplified)
	startDate := time.Date(2023, 12, 1, 0, 0, 0, 0, time.UTC) // Friday
	businessDays := 0
	current := startDate

	fmt.Printf("Business days from %s:\n", startDate.Format("2006-01-02"))
	for i := 0; i < 10; i++ {
		weekday := current.Weekday()
		if weekday != time.Saturday && weekday != time.Sunday {
			businessDays++
			fmt.Printf("  Day %d: %s (%s) - Business day #%d\n", 
				i+1, current.Format("2006-01-02"), weekday, businessDays)
		} else {
			fmt.Printf("  Day %d: %s (%s) - Weekend\n", 
				i+1, current.Format("2006-01-02"), weekday)
		}
		current = current.Add(24 * time.Hour)
	}

	// Time until next occurrence
	now = time.Now()
	nextYear := time.Date(now.Year()+1, 1, 1, 0, 0, 0, 0, now.Location())
	timeUntilNewYear := time.Until(nextYear)
	fmt.Printf("\nTime until next New Year: %v\n", timeUntilNewYear)
	fmt.Printf("Days: %.0f\n", timeUntilNewYear.Hours()/24)

	// 🎯 DEMO 10: Performance Measurement
	fmt.Println("\n🎯 DEMO 10: Performance Measurement")
	fmt.Println("===================================")

	// Measure function execution time
	start = time.Now()
	
	// Simulate some work
	sum := 0
	for i := 0; i < 1000000; i++ {
		sum += i
	}
	
	duration := time.Since(start)
	fmt.Printf("Loop execution time: %v\n", duration)
	fmt.Printf("Result: %d\n", sum)

	// Multiple measurements
	fmt.Println("\nMultiple measurements:")
	for i := 0; i < 3; i++ {
		start := time.Now()
		time.Sleep(50 * time.Millisecond)
		elapsed := time.Since(start)
		fmt.Printf("  Measurement %d: %v\n", i+1, elapsed)
	}

	fmt.Println("\n✨ All time demos completed!")
}

/*
=============================================================================
                              📝 LEARNING NOTES
=============================================================================

⏰ TIME CREATION:
┌─────────────────────────────────────────────────────────────────────────┐
│ // Current time                                                         │
│ now := time.Now()                                                       │
│ utc := time.Now().UTC()                                                 │
│                                                                         │
│ // Specific time                                                        │
│ t := time.Date(2023, 12, 25, 15, 30, 45, 0, time.UTC)                  │
│                                                                         │
│ // From Unix timestamp                                                  │
│ t := time.Unix(1640995200, 0)                                           │
│ t := time.UnixMilli(1640995200000)                                      │
│ t := time.UnixMicro(1640995200000000)                                   │
│                                                                         │
│ // Parse from string                                                    │
│ t, err := time.Parse("2006-01-02 15:04:05", "2023-12-01 14:30:25")     │
│ t, err := time.ParseInLocation(layout, value, location)                 │
└─────────────────────────────────────────────────────────────────────────┘

📅 TIME FORMATTING:
┌─────────────────────────────────────────────────────────────────────────┐
│ // Go's reference time: Mon Jan 2 15:04:05 MST 2006                     │
│ // Unix timestamp: 1136239445                                           │
│                                                                         │
│ // Common formats                                                       │
│ t.Format("2006-01-02 15:04:05")        // 2023-12-01 14:30:25          │
│ t.Format("Jan 2, 2006")                // Dec 1, 2023                   │
│ t.Format("15:04:05")                   // 14:30:25                      │
│ t.Format("3:04 PM")                    // 2:30 PM                       │
│ t.Format("Monday, January 2, 2006")    // Friday, December 1, 2023      │
│                                                                         │
│ // Predefined constants                                                 │
│ t.Format(time.RFC3339)     // 2023-12-01T14:30:25Z                     │
│ t.Format(time.RFC822)      // 01 Dec 23 14:30 UTC                      │
│ t.Format(time.Kitchen)     // 2:30PM                                    │
│ t.Format(time.Stamp)       // Dec  1 14:30:25                          │
└─────────────────────────────────────────────────────────────────────────┘

⏱️ DURATION OPERATIONS:
┌─────────────────────────────────────────────────────────────────────────┐
│ // Duration constants                                                   │
│ time.Nanosecond, time.Microsecond, time.Millisecond                     │
│ time.Second, time.Minute, time.Hour                                     │
│                                                                         │
│ // Custom durations                                                     │
│ d := 5 * time.Second                                                    │
│ d := 2*time.Hour + 30*time.Minute                                       │
│                                                                         │
│ // Parse duration                                                       │
│ d, err := time.ParseDuration("2h45m30s")                                │
│ d, err := time.ParseDuration("100ms")                                   │
│                                                                         │
│ // Duration methods                                                     │
│ d.Hours()        // Duration in hours                                   │
│ d.Minutes()      // Duration in minutes                                 │
│ d.Seconds()      // Duration in seconds                                 │
│ d.Milliseconds() // Duration in milliseconds                            │
│ d.Nanoseconds()  // Duration in nanoseconds                             │
└─────────────────────────────────────────────────────────────────────────┘

🔢 TIME ARITHMETIC:
┌─────────────────────────────────────────────────────────────────────────┐
│ // Add/subtract duration                                                │
│ future := t.Add(2 * time.Hour)                                          │
│ past := t.Add(-30 * time.Minute)                                        │
│                                                                         │
│ // Time difference                                                      │
│ diff := t2.Sub(t1)              // Returns duration                     │
│ elapsed := time.Since(start)     // Time since start                    │
│ remaining := time.Until(deadline) // Time until deadline                │
│                                                                         │
│ // Add date components                                                  │
│ nextMonth := t.AddDate(0, 1, 0)  // Add 1 month                        │
│ nextYear := t.AddDate(1, 0, 0)   // Add 1 year                         │
│ tomorrow := t.AddDate(0, 0, 1)   // Add 1 day                          │
└─────────────────────────────────────────────────────────────────────────┘

🌍 TIME ZONES:
┌─────────────────────────────────────────────────────────────────────────┐
│ // Load location                                                        │
│ loc, err := time.LoadLocation("America/New_York")                       │
│ loc, err := time.LoadLocation("Europe/London")                          │
│ loc, err := time.LoadLocation("Asia/Tokyo")                             │
│                                                                         │
│ // Common locations                                                     │
│ time.UTC           // UTC timezone                                      │
│ time.Local         // System local timezone                             │
│                                                                         │
│ // Convert timezone                                                     │
│ localTime := utcTime.In(loc)                                            │
│ utcTime := localTime.UTC()                                              │
│                                                                         │
│ // Create time in specific timezone                                     │
│ t := time.Date(2023, 12, 1, 15, 30, 0, 0, loc)                         │
└─────────────────────────────────────────────────────────────────────────┘

🔍 TIME COMPARISONS:
┌─────────────────────────────────────────────────────────────────────────┐
│ // Comparison methods                                                   │
│ t1.Before(t2)     // true if t1 is before t2                           │
│ t1.After(t2)      // true if t1 is after t2                            │
│ t1.Equal(t2)      // true if t1 equals t2                              │
│                                                                         │
│ // Check if time is zero                                                │
│ t.IsZero()        // true if t is zero time                             │
│                                                                         │
│ // Unix timestamp comparison                                            │
│ t1.Unix() < t2.Unix()                                                   │
└─────────────────────────────────────────────────────────────────────────┘

📊 TIME COMPONENTS:
┌─────────────────────────────────────────────────────────────────────────┐
│ // Date components                                                      │
│ t.Year()          // 2023                                               │
│ t.Month()         // December (or 12)                                   │
│ t.Day()           // 1-31                                               │
│ t.Weekday()       // Sunday, Monday, etc.                               │
│ t.YearDay()       // 1-366 (day of year)                                │
│                                                                         │
│ // Time components                                                      │
│ t.Hour()          // 0-23                                               │
│ t.Minute()        // 0-59                                               │
│ t.Second()        // 0-59                                               │
│ t.Nanosecond()    // 0-999999999                                        │
│                                                                         │
│ // ISO week                                                             │
│ year, week := t.ISOWeek()                                               │
│                                                                         │
│ // Date only                                                            │
│ year, month, day := t.Date()                                            │
│ hour, min, sec := t.Clock()                                             │
└─────────────────────────────────────────────────────────────────────────┘

💡 BEST PRACTICES:
• Always handle time zone conversions explicitly
• Use UTC for storage, local time for display
• Be careful with daylight saving time transitions
• Use time.Time for timestamps, not strings
• Consider leap years and leap seconds
• Use monotonic time for measuring durations

🚨 COMMON MISTAKES:
❌ Ignoring time zones in comparisons
❌ Using local time for storage
❌ Not handling parsing errors
❌ Assuming all days have 24 hours (DST)
❌ Using time.Sleep in production for delays

⚡ PERFORMANCE TIPS:
• Cache time.Location objects
• Use time.Now() sparingly in tight loops
• Consider monotonic time for benchmarks
• Use time.Unix() for timestamp comparisons
• Avoid frequent time formatting in hot paths

🎯 REAL-WORLD PATTERNS:
• Store timestamps in UTC
• Display times in user's timezone
• Use ISO 8601 format for APIs
• Implement timeout patterns with context
• Use time.Ticker for periodic tasks
• Measure performance with time.Since()

=============================================================================
*/