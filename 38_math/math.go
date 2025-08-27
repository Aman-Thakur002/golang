/*
=============================================================================
                           🔢 GO MATH TUTORIAL
=============================================================================

📚 CORE CONCEPT:
Go's math package provides basic constants and mathematical functions
for floating-point operations. It includes trigonometric, logarithmic,
and other mathematical functions.

🔑 KEY FEATURES:
• Mathematical constants (Pi, E, etc.)
• Basic arithmetic functions
• Trigonometric functions
• Logarithmic and exponential functions
• Rounding and comparison functions

💡 REAL-WORLD ANALOGY:
Math Package = Scientific Calculator
- Constants = Pre-stored values (π, e)
- Functions = Calculator buttons (sin, cos, log)
- Precision = Calculator's decimal accuracy
- Special values = Error indicators (NaN, Inf)

🎯 WHY LEARN MATH?
• Scientific and engineering calculations
• Graphics and game programming
• Statistical analysis
• Financial calculations

=============================================================================
*/

package main

import (
	"fmt"
	"math"
	"math/big"
)

func main() {
	fmt.Println("🔢 MATH TUTORIAL")
	fmt.Println("================")

	// 🎯 DEMO 1: Mathematical Constants
	fmt.Println("\n🎯 DEMO 1: Mathematical Constants")
	fmt.Println("=================================")

	constants := map[string]float64{
		"Pi (π)":           math.Pi,
		"E (Euler's)":      math.E,
		"Phi (Golden)":     math.Phi,
		"Sqrt2":            math.Sqrt2,
		"SqrtE":            math.SqrtE,
		"SqrtPi":           math.SqrtPi,
		"SqrtPhi":          math.SqrtPhi,
		"Ln2":              math.Ln2,
		"Log2E":            math.Log2E,
		"Ln10":             math.Ln10,
		"Log10E":           math.Log10E,
	}

	fmt.Println("Important mathematical constants:")
	for name, value := range constants {
		fmt.Printf("  %-15s: %.10f\n", name, value)
	}

	// Special values
	fmt.Println("\nSpecial floating-point values:")
	fmt.Printf("  Positive infinity: %v\n", math.Inf(1))
	fmt.Printf("  Negative infinity: %v\n", math.Inf(-1))
	fmt.Printf("  Not a Number:      %v\n", math.NaN())

	// 🎯 DEMO 2: Basic Arithmetic Functions
	fmt.Println("\n🎯 DEMO 2: Basic Arithmetic")
	fmt.Println("===========================")

	x, y := 16.0, 4.0
	fmt.Printf("x = %.1f, y = %.1f\n", x, y)

	fmt.Printf("Abs(-7.5):     %.1f\n", math.Abs(-7.5))
	fmt.Printf("Max(x, y):     %.1f\n", math.Max(x, y))
	fmt.Printf("Min(x, y):     %.1f\n", math.Min(x, y))
	fmt.Printf("Pow(x, 2):     %.1f\n", math.Pow(x, 2))
	fmt.Printf("Pow(x, 0.5):   %.1f\n", math.Pow(x, 0.5))
	fmt.Printf("Sqrt(x):       %.1f\n", math.Sqrt(x))
	fmt.Printf("Cbrt(27):      %.1f\n", math.Cbrt(27))

	// Modulo and remainder
	fmt.Printf("Mod(17, 5):    %.1f\n", math.Mod(17, 5))
	fmt.Printf("Remainder(17, 5): %.1f\n", math.Remainder(17, 5))

	// 🎯 DEMO 3: Rounding Functions
	fmt.Println("\n🎯 DEMO 3: Rounding Functions")
	fmt.Println("=============================")

	values := []float64{3.2, 3.7, -2.3, -2.8, 4.5, -4.5}
	
	fmt.Println("Value  | Floor | Ceil  | Round | Trunc")
	fmt.Println("-------|-------|-------|-------|-------")
	for _, v := range values {
		fmt.Printf("%6.1f | %5.1f | %5.1f | %5.1f | %5.1f\n",
			v, math.Floor(v), math.Ceil(v), math.Round(v), math.Trunc(v))
	}

	// Round to even (banker's rounding)
	fmt.Println("\nRound to even examples:")
	evenValues := []float64{0.5, 1.5, 2.5, 3.5, 4.5}
	for _, v := range evenValues {
		fmt.Printf("RoundToEven(%.1f): %.1f\n", v, math.RoundToEven(v))
	}

	// 🎯 DEMO 4: Trigonometric Functions
	fmt.Println("\n🎯 DEMO 4: Trigonometric Functions")
	fmt.Println("==================================")

	angles := []float64{0, math.Pi/6, math.Pi/4, math.Pi/3, math.Pi/2, math.Pi}
	
	fmt.Println("Angle (rad) | Degrees |   Sin   |   Cos   |   Tan")
	fmt.Println("------------|---------|---------|---------|----------")
	for _, angle := range angles {
		degrees := angle * 180 / math.Pi
		sin := math.Sin(angle)
		cos := math.Cos(angle)
		tan := math.Tan(angle)
		fmt.Printf("%11.4f | %7.1f | %7.4f | %7.4f | %8.4f\n",
			angle, degrees, sin, cos, tan)
	}

	// Inverse trigonometric functions
	fmt.Println("\nInverse trigonometric functions:")
	fmt.Printf("Asin(0.5):  %.4f rad (%.1f°)\n", math.Asin(0.5), math.Asin(0.5)*180/math.Pi)
	fmt.Printf("Acos(0.5):  %.4f rad (%.1f°)\n", math.Acos(0.5), math.Acos(0.5)*180/math.Pi)
	fmt.Printf("Atan(1):    %.4f rad (%.1f°)\n", math.Atan(1), math.Atan(1)*180/math.Pi)
	fmt.Printf("Atan2(1,1): %.4f rad (%.1f°)\n", math.Atan2(1, 1), math.Atan2(1, 1)*180/math.Pi)

	// 🎯 DEMO 5: Logarithmic and Exponential Functions
	fmt.Println("\n🎯 DEMO 5: Logarithmic and Exponential")
	fmt.Println("======================================")

	testValues := []float64{1, 2, math.E, 10, 100}
	
	fmt.Println("Value |   Exp   |   Log   |  Log2   |  Log10")
	fmt.Println("------|---------|---------|---------|----------")
	for _, v := range testValues {
		exp := math.Exp(v)
		log := math.Log(v)
		log2 := math.Log2(v)
		log10 := math.Log10(v)
		fmt.Printf("%5.1f | %7.2f | %7.4f | %7.4f | %8.4f\n",
			v, exp, log, log2, log10)
	}

	// Special exponential functions
	fmt.Printf("\nSpecial exponential functions:\n")
	fmt.Printf("Exp2(3):     %.1f (2^3)\n", math.Exp2(3))
	fmt.Printf("Expm1(0.1):  %.6f (e^0.1 - 1)\n", math.Expm1(0.1))
	fmt.Printf("Log1p(0.1):  %.6f (ln(1 + 0.1))\n", math.Log1p(0.1))

	// 🎯 DEMO 6: Hyperbolic Functions
	fmt.Println("\n🎯 DEMO 6: Hyperbolic Functions")
	fmt.Println("===============================")

	x = 1.0
	fmt.Printf("For x = %.1f:\n", x)
	fmt.Printf("Sinh(x):  %.6f\n", math.Sinh(x))
	fmt.Printf("Cosh(x):  %.6f\n", math.Cosh(x))
	fmt.Printf("Tanh(x):  %.6f\n", math.Tanh(x))

	// Inverse hyperbolic functions
	fmt.Printf("Asinh(x): %.6f\n", math.Asinh(x))
	fmt.Printf("Acosh(x): %.6f\n", math.Acosh(math.Cosh(x)))
	fmt.Printf("Atanh(0.5): %.6f\n", math.Atanh(0.5))

	// 🎯 DEMO 7: Special Functions and Utilities
	fmt.Println("\n🎯 DEMO 7: Special Functions")
	fmt.Println("============================")

	// Gamma function
	fmt.Printf("Gamma(5):    %.1f (4! = %d)\n", math.Gamma(5), 4*3*2*1)
	fmt.Printf("Gamma(0.5):  %.6f (√π)\n", math.Gamma(0.5))

	// Error function
	fmt.Printf("Erf(1):      %.6f\n", math.Erf(1))
	fmt.Printf("Erfc(1):     %.6f\n", math.Erfc(1))

	// Bessel functions
	fmt.Printf("J0(1):       %.6f\n", math.J0(1))
	fmt.Printf("J1(1):       %.6f\n", math.J1(1))
	fmt.Printf("Y0(1):       %.6f\n", math.Y0(1))
	fmt.Printf("Y1(1):       %.6f\n", math.Y1(1))

	// 🎯 DEMO 8: Float Utilities
	fmt.Println("\n🎯 DEMO 8: Float Utilities")
	fmt.Println("==========================")

	testFloat := 123.456
	fmt.Printf("Original value: %.3f\n", testFloat)

	// Decompose float
	frac := math.Modf(testFloat)
	integer, fraction := math.Modf(testFloat)
	fmt.Printf("Modf: integer=%.0f, fraction=%.3f\n", integer, fraction)

	// Frexp and Ldexp
	mantissa, exponent := math.Frexp(testFloat)
	fmt.Printf("Frexp: mantissa=%.6f, exponent=%d\n", mantissa, exponent)
	reconstructed := math.Ldexp(mantissa, exponent)
	fmt.Printf("Ldexp reconstruction: %.3f\n", reconstructed)

	// Check for special values
	specialValues := []float64{1.0, math.Inf(1), math.Inf(-1), math.NaN()}
	for _, v := range specialValues {
		fmt.Printf("Value %v: IsInf=%t, IsNaN=%t, Signbit=%t\n",
			v, math.IsInf(v, 0), math.IsNaN(v), math.Signbit(v))
	}

	// 🎯 DEMO 9: Practical Examples
	fmt.Println("\n🎯 DEMO 9: Practical Examples")
	fmt.Println("=============================")

	// Distance between two points
	x1, y1 := 3.0, 4.0
	x2, y2 := 6.0, 8.0
	distance := math.Sqrt(math.Pow(x2-x1, 2) + math.Pow(y2-y1, 2))
	fmt.Printf("Distance between (%.1f,%.1f) and (%.1f,%.1f): %.2f\n",
		x1, y1, x2, y2, distance)

	// Area of a circle
	radius := 5.0
	area := math.Pi * math.Pow(radius, 2)
	circumference := 2 * math.Pi * radius
	fmt.Printf("Circle with radius %.1f: area=%.2f, circumference=%.2f\n",
		radius, area, circumference)

	// Compound interest calculation
	principal := 1000.0
	rate := 0.05 // 5%
	time := 10.0 // years
	compound := principal * math.Pow(1+rate, time)
	fmt.Printf("Compound interest: $%.2f after %.0f years at %.1f%%\n",
		compound, time, rate*100)

	// Convert degrees to radians and vice versa
	degrees := 45.0
	radians := degrees * math.Pi / 180
	backToDegrees := radians * 180 / math.Pi
	fmt.Printf("%.0f degrees = %.4f radians = %.0f degrees\n",
		degrees, radians, backToDegrees)

	// 🎯 DEMO 10: Big Numbers (math/big package)
	fmt.Println("\n🎯 DEMO 10: Big Numbers")
	fmt.Println("=======================")

	// Big integers
	bigInt1 := big.NewInt(123456789)
	bigInt2 := big.NewInt(987654321)
	result := new(big.Int)
	result.Mul(bigInt1, bigInt2)
	fmt.Printf("Big integer multiplication: %s × %s = %s\n",
		bigInt1.String(), bigInt2.String(), result.String())

	// Big floats
	bigFloat1 := big.NewFloat(math.Pi)
	bigFloat1.SetPrec(100) // Set precision to 100 bits
	fmt.Printf("High precision π: %s\n", bigFloat1.Text('f', 30))

	// Factorial using big integers
	factorial := big.NewInt(1)
	for i := 1; i <= 20; i++ {
		factorial.Mul(factorial, big.NewInt(int64(i)))
	}
	fmt.Printf("20! = %s\n", factorial.String())

	fmt.Println("\n✨ All math demos completed!")
}

/*
=============================================================================
                              📝 LEARNING NOTES
=============================================================================

🔢 MATHEMATICAL CONSTANTS:
┌─────────────────────────────────────────────────────────────────────────┐
│ math.Pi      = 3.14159265358979323846...  // π                          │
│ math.E       = 2.71828182845904523536...  // e (Euler's number)         │
│ math.Phi     = 1.61803398874989484820...  // φ (Golden ratio)           │
│ math.Sqrt2   = 1.41421356237309504880...  // √2                         │
│ math.SqrtE   = 1.64872127070012814684...  // √e                         │
│ math.SqrtPi  = 1.77245385090551602729...  // √π                         │
│ math.SqrtPhi = 1.27201964951406896425...  // √φ                         │
│ math.Ln2     = 0.69314718055994530941...  // ln(2)                      │
│ math.Log2E   = 1.44269504088896338700...  // log₂(e)                    │
│ math.Ln10    = 2.30258509299404568401...  // ln(10)                     │
│ math.Log10E  = 0.43429448190325182765...  // log₁₀(e)                   │
└─────────────────────────────────────────────────────────────────────────┘

🔧 BASIC ARITHMETIC:
┌─────────────────────────────────────────────────────────────────────────┐
│ math.Abs(x)           // Absolute value                                 │
│ math.Max(x, y)        // Maximum of two values                          │
│ math.Min(x, y)        // Minimum of two values                          │
│ math.Pow(x, y)        // x raised to power y                            │
│ math.Sqrt(x)          // Square root                                    │
│ math.Cbrt(x)          // Cube root                                      │
│ math.Hypot(x, y)      // √(x² + y²) without overflow                    │
│ math.Mod(x, y)        // Floating-point remainder of x/y                │
│ math.Remainder(x, y)  // IEEE 754 remainder                             │
└─────────────────────────────────────────────────────────────────────────┘

📐 ROUNDING FUNCTIONS:
┌─────────────────────────────────────────────────────────────────────────┐
│ math.Floor(x)         // Largest integer ≤ x                            │
│ math.Ceil(x)          // Smallest integer ≥ x                           │
│ math.Round(x)         // Round to nearest integer (away from zero)      │
│ math.RoundToEven(x)   // Round to nearest even integer                  │
│ math.Trunc(x)         // Integer part (toward zero)                     │
└─────────────────────────────────────────────────────────────────────────┘

📊 TRIGONOMETRIC FUNCTIONS:
┌─────────────────────────────────────────────────────────────────────────┐
│ // Basic trigonometric functions (radians)                              │
│ math.Sin(x), math.Cos(x), math.Tan(x)                                   │
│                                                                         │
│ // Inverse trigonometric functions                                      │
│ math.Asin(x), math.Acos(x), math.Atan(x)                                │
│ math.Atan2(y, x)      // atan(y/x) with correct quadrant               │
│                                                                         │
│ // Hyperbolic functions                                                 │
│ math.Sinh(x), math.Cosh(x), math.Tanh(x)                                │
│ math.Asinh(x), math.Acosh(x), math.Atanh(x)                             │
│                                                                         │
│ // Degree/radian conversion                                             │
│ radians = degrees * math.Pi / 180                                       │
│ degrees = radians * 180 / math.Pi                                       │
└─────────────────────────────────────────────────────────────────────────┘

📈 EXPONENTIAL & LOGARITHMIC:
┌─────────────────────────────────────────────────────────────────────────┐
│ // Exponential functions                                                │
│ math.Exp(x)           // e^x                                            │
│ math.Exp2(x)          // 2^x                                            │
│ math.Expm1(x)         // e^x - 1 (accurate for small x)                 │
│                                                                         │
│ // Logarithmic functions                                                │
│ math.Log(x)           // Natural logarithm (ln)                         │
│ math.Log2(x)          // Base-2 logarithm                               │
│ math.Log10(x)         // Base-10 logarithm                              │
│ math.Log1p(x)         // ln(1 + x) (accurate for small x)               │
└─────────────────────────────────────────────────────────────────────────┘

🔍 FLOAT UTILITIES:
┌─────────────────────────────────────────────────────────────────────────┐
│ // Float decomposition                                                  │
│ integer, fraction := math.Modf(x)    // Split into integer and fraction │
│ mantissa, exp := math.Frexp(x)       // x = mantissa × 2^exp            │
│ result := math.Ldexp(mantissa, exp)  // Reconstruct from mantissa/exp   │
│                                                                         │
│ // Special value checks                                                 │
│ math.IsInf(x, sign)   // Check if infinite (sign: 1=+∞, -1=-∞, 0=both) │
│ math.IsNaN(x)         // Check if Not a Number                          │
│ math.Signbit(x)       // Check if sign bit is set                       │
│                                                                         │
│ // Special values                                                       │
│ math.Inf(1)           // Positive infinity                              │
│ math.Inf(-1)          // Negative infinity                              │
│ math.NaN()            // Not a Number                                    │
└─────────────────────────────────────────────────────────────────────────┘

🎯 PRACTICAL FORMULAS:
┌─────────────────────────────────────────────────────────────────────────┐
│ // Distance between points                                              │
│ distance := math.Sqrt(math.Pow(x2-x1, 2) + math.Pow(y2-y1, 2))         │
│                                                                         │
│ // Circle calculations                                                  │
│ area := math.Pi * math.Pow(radius, 2)                                   │
│ circumference := 2 * math.Pi * radius                                   │
│                                                                         │
│ // Compound interest                                                    │
│ amount := principal * math.Pow(1 + rate, time)                          │
│                                                                         │
│ // Degree/radian conversion                                             │
│ radians := degrees * math.Pi / 180                                      │
│ degrees := radians * 180 / math.Pi                                      │
│                                                                         │
│ // Pythagorean theorem                                                  │
│ hypotenuse := math.Hypot(a, b)  // More accurate than Sqrt(a²+b²)       │
└─────────────────────────────────────────────────────────────────────────┘

💡 BEST PRACTICES:
• Use math.Hypot() instead of math.Sqrt(x*x + y*y) to avoid overflow
• Use math.Expm1() and math.Log1p() for better accuracy with small values
• Check for special values (NaN, Inf) when needed
• Use appropriate precision for your use case
• Consider math/big package for arbitrary precision

🚨 COMMON MISTAKES:
❌ Forgetting to convert degrees to radians for trig functions
❌ Not checking for domain errors (e.g., sqrt of negative numbers)
❌ Using == to compare floating-point results
❌ Not handling special values (NaN, Inf)
❌ Precision loss in floating-point calculations

⚡ PERFORMANCE TIPS:
• Cache frequently used constants
• Use integer arithmetic when possible
• Consider lookup tables for expensive functions
• Use math.Pow(x, 2) instead of x*x only when needed
• Profile math-heavy code for bottlenecks

🎯 WHEN TO USE MATH/BIG:
• Financial calculations requiring exact precision
• Cryptographic operations with large numbers
• Scientific computing with arbitrary precision
• When standard float64 precision is insufficient

=============================================================================
*/