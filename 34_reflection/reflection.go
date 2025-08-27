/*
=============================================================================
                        🔍 GO REFLECTION TUTORIAL
=============================================================================

📚 CORE CONCEPT:
Reflection allows programs to examine their own structure at runtime.
Go's reflect package provides the ability to inspect types and values
dynamically.

🔑 KEY FEATURES:
• Runtime type inspection
• Dynamic value manipulation
• Struct field and method discovery
• Interface{} value examination

💡 REAL-WORLD ANALOGY:
Reflection = X-Ray Machine
- TypeOf = Identifying bone structure
- ValueOf = Seeing internal organs
- Field access = Examining specific body parts
- Method calls = Testing reflexes

🎯 WHY USE REFLECTION?
• Generic programming before generics
• Serialization/deserialization
• ORM and database mapping
• Testing and debugging tools

=============================================================================
*/

package main

import (
	"fmt"
	"reflect"
	// "strconv"
)

// 📊 SAMPLE TYPES FOR REFLECTION
type Person struct {
	Name    string `json:"name" validate:"required"`
	Age     int    `json:"age" validate:"min=0,max=150"`
	Email   string `json:"email" validate:"email"`
	private string // unexported field
}

func (p Person) GetInfo() string {
	return fmt.Sprintf("%s (%d years old)", p.Name, p.Age)
}

func (p Person) IsAdult() bool {
	return p.Age >= 18
}

func (p *Person) SetName(name string) {
	p.Name = name
}

type Animal interface {
	Speak() string
}

type Dog struct {
	Name string
}

func (d Dog) Speak() string {
	return "Woof!"
}

func main() {
	fmt.Println("🔍 REFLECTION TUTORIAL")
	fmt.Println("======================")

	// 🎯 DEMO 1: Basic Type and Value Inspection
	fmt.Println("\n🎯 DEMO 1: Basic Type and Value")
	fmt.Println("===============================")

	var x int = 42
	var s string = "hello"
	var f float64 = 3.14

	values := []interface{}{x, s, f}
	for _, v := range values {
		t := reflect.TypeOf(v)
		val := reflect.ValueOf(v)
		
		fmt.Printf("Value: %v, Type: %v, Kind: %v\n", v, t, t.Kind())
		fmt.Printf("  Reflect Value: %v, Type: %v\n", val, val.Type())
		fmt.Printf("  Can Set: %t, Can Interface: %t\n\n", val.CanSet(), val.CanInterface())
	}

	// 🎯 DEMO 2: Struct Reflection
	fmt.Println("🎯 DEMO 2: Struct Reflection")
	fmt.Println("============================")

	person := Person{
		Name:    "John Doe",
		Age:     30,
		Email:   "john@example.com",
		private: "secret",
	}

	t := reflect.TypeOf(person)
	v := reflect.ValueOf(person)

	fmt.Printf("Struct Type: %v\n", t)
	fmt.Printf("Number of fields: %d\n", t.NumField())
	fmt.Printf("Number of methods: %d\n", t.NumMethod())

	// Iterate through fields
	fmt.Println("\n📋 Fields:")
	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		value := v.Field(i)
		
		fmt.Printf("  Field %d: %s\n", i, field.Name)
		fmt.Printf("    Type: %v\n", field.Type)
		fmt.Printf("    Tag: %q\n", field.Tag)
		fmt.Printf("    JSON tag: %q\n", field.Tag.Get("json"))
		fmt.Printf("    Validate tag: %q\n", field.Tag.Get("validate"))
		
		if value.CanInterface() {
			fmt.Printf("    Value: %v\n", value.Interface())
		} else {
			fmt.Printf("    Value: <unexported>\n")
		}
		fmt.Println()
	}

	// Iterate through methods
	fmt.Println("🔧 Methods:")
	for i := 0; i < t.NumMethod(); i++ {
		method := t.Method(i)
		fmt.Printf("  Method %d: %s\n", i, method.Name)
		fmt.Printf("    Type: %v\n", method.Type)
		fmt.Printf("    Func: %v\n", method.Func)
		fmt.Println()
	}

	// 🎯 DEMO 3: Method Invocation
	fmt.Println("🎯 DEMO 3: Method Invocation")
	fmt.Println("============================")

	// Call method by name
	methodName := "GetInfo"
	method := v.MethodByName(methodName)
	if method.IsValid() {
		result := method.Call(nil)
		fmt.Printf("Called %s(): %v\n", methodName, result[0].Interface())
	}

	// Call method with parameters (need pointer for SetName)
	ptrValue := reflect.ValueOf(&person)
	setNameMethod := ptrValue.MethodByName("SetName")
	if setNameMethod.IsValid() {
		args := []reflect.Value{reflect.ValueOf("Jane Smith")}
		setNameMethod.Call(args)
		fmt.Printf("After SetName: %s\n", person.Name)
	}

	// 🎯 DEMO 4: Slice and Array Reflection
	fmt.Println("\n🎯 DEMO 4: Slice and Array Reflection")
	fmt.Println("=====================================")

	numbers := []int{1, 2, 3, 4, 5}
	sliceValue := reflect.ValueOf(numbers)
	sliceType := reflect.TypeOf(numbers)

	fmt.Printf("Slice Type: %v\n", sliceType)
	fmt.Printf("Element Type: %v\n", sliceType.Elem())
	fmt.Printf("Length: %d\n", sliceValue.Len())
	fmt.Printf("Capacity: %d\n", sliceValue.Cap())

	fmt.Println("Elements:")
	for i := 0; i < sliceValue.Len(); i++ {
		elem := sliceValue.Index(i)
		fmt.Printf("  [%d]: %v\n", i, elem.Interface())
	}

	// 🎯 DEMO 5: Map Reflection
	fmt.Println("\n🎯 DEMO 5: Map Reflection")
	fmt.Println("=========================")

	data := map[string]int{
		"apple":  5,
		"banana": 3,
		"orange": 8,
	}

	mapValue := reflect.ValueOf(data)
	mapType := reflect.TypeOf(data)

	fmt.Printf("Map Type: %v\n", mapType)
	fmt.Printf("Key Type: %v\n", mapType.Key())
	fmt.Printf("Value Type: %v\n", mapType.Elem())
	fmt.Printf("Length: %d\n", mapValue.Len())

	fmt.Println("Key-Value pairs:")
	for _, key := range mapValue.MapKeys() {
		value := mapValue.MapIndex(key)
		fmt.Printf("  %v: %v\n", key.Interface(), value.Interface())
	}

	// 🎯 DEMO 6: Interface Reflection
	fmt.Println("\n🎯 DEMO 6: Interface Reflection")
	fmt.Println("===============================")

	var animal Animal = Dog{Name: "Buddy"}
	
	interfaceValue := reflect.ValueOf(animal)
	interfaceType := reflect.TypeOf(animal)

	fmt.Printf("Interface Type: %v\n", interfaceType)
	fmt.Printf("Underlying Type: %v\n", interfaceValue.Type())
	fmt.Printf("Underlying Value: %v\n", interfaceValue.Interface())

	// Check if it implements the interface
	animalType := reflect.TypeOf((*Animal)(nil)).Elem()
	fmt.Printf("Implements Animal interface: %t\n", interfaceValue.Type().Implements(animalType))

	// 🎯 DEMO 7: Dynamic Value Modification
	fmt.Println("\n🎯 DEMO 7: Dynamic Value Modification")
	fmt.Println("=====================================")

	// Create a modifiable value
	modifiablePerson := Person{Name: "Original", Age: 25}
	ptrVal := reflect.ValueOf(&modifiablePerson)
	structVal := ptrVal.Elem() // Get the struct value from pointer

	fmt.Printf("Before modification: %+v\n", modifiablePerson)

	// Modify field by name
	nameField := structVal.FieldByName("Name")
	if nameField.IsValid() && nameField.CanSet() {
		nameField.SetString("Modified Name")
	}

	ageField := structVal.FieldByName("Age")
	if ageField.IsValid() && ageField.CanSet() {
		ageField.SetInt(35)
	}

	fmt.Printf("After modification: %+v\n", modifiablePerson)

	// 🎯 DEMO 8: Type Creation and Instantiation
	fmt.Println("\n🎯 DEMO 8: Dynamic Type Creation")
	fmt.Println("================================")

	// Create new instance of a type
	personType := reflect.TypeOf(Person{})
	newPersonValue := reflect.New(personType) // Creates *Person
	newPersonStruct := newPersonValue.Elem()  // Dereference to get Person

	// Set field values
	newPersonStruct.FieldByName("Name").SetString("Dynamic Person")
	newPersonStruct.FieldByName("Age").SetInt(40)
	newPersonStruct.FieldByName("Email").SetString("dynamic@example.com")

	// Get the actual struct
	dynamicPerson := newPersonValue.Interface().(*Person)
	fmt.Printf("Dynamically created person: %+v\n", *dynamicPerson)

	// 🎯 DEMO 9: Practical Example - Generic JSON-like Serializer
	fmt.Println("\n🎯 DEMO 9: Generic Serializer")
	fmt.Println("=============================")

	serialized := serialize(person)
	fmt.Printf("Serialized: %s\n", serialized)

	fmt.Println("\n✨ All reflection demos completed!")
}

// 🔧 UTILITY FUNCTION: Simple serializer using reflection
func serialize(v interface{}) string {
	val := reflect.ValueOf(v)
	typ := reflect.TypeOf(v)

	if typ.Kind() != reflect.Struct {
		return fmt.Sprintf("%v", v)
	}

	result := "{"
	for i := 0; i < val.NumField(); i++ {
		field := typ.Field(i)
		value := val.Field(i)

		if !value.CanInterface() {
			continue // Skip unexported fields
		}

		if i > 0 {
			result += ", "
		}

		jsonTag := field.Tag.Get("json")
		fieldName := field.Name
		if jsonTag != "" {
			fieldName = jsonTag
		}

		switch value.Kind() {
		case reflect.String:
			result += fmt.Sprintf(`"%s": "%v"`, fieldName, value.Interface())
		case reflect.Int, reflect.Int64:
			result += fmt.Sprintf(`"%s": %v`, fieldName, value.Interface())
		default:
			result += fmt.Sprintf(`"%s": "%v"`, fieldName, value.Interface())
		}
	}
	result += "}"
	return result
}

/*
=============================================================================
                              📝 LEARNING NOTES
=============================================================================

🔍 REFLECTION BASICS:
┌─────────────────────────────────────────────────────────────────────────┐
│ // Get type and value                                                   │
│ t := reflect.TypeOf(variable)                                           │
│ v := reflect.ValueOf(variable)                                          │
│                                                                         │
│ // Basic information                                                    │
│ fmt.Println(t.Kind())    // int, string, struct, etc.                  │
│ fmt.Println(t.Name())    // Type name                                   │
│ fmt.Println(v.Type())    // Same as TypeOf                              │
│ fmt.Println(v.Kind())    // Same as TypeOf().Kind()                     │
│                                                                         │
│ // Check capabilities                                                   │
│ v.CanSet()        // Can modify value                                   │
│ v.CanInterface()  // Can get interface{} value                          │
│ v.IsValid()       // Is valid reflect.Value                             │
│ v.IsNil()         // Is nil (for pointers, slices, etc.)                │
└─────────────────────────────────────────────────────────────────────────┘

📊 STRUCT REFLECTION:
┌─────────────────────────────────────────────────────────────────────────┐
│ t := reflect.TypeOf(structInstance)                                     │
│ v := reflect.ValueOf(structInstance)                                    │
│                                                                         │
│ // Field information                                                    │
│ numFields := t.NumField()                                               │
│ field := t.Field(i)              // Get field by index                  │
│ field := t.FieldByName("Name")    // Get field by name                  │
│                                                                         │
│ // Field properties                                                     │
│ field.Name        // Field name                                         │
│ field.Type        // Field type                                         │
│ field.Tag         // Struct tag                                         │
│ field.Tag.Get("json")  // Get specific tag value                        │
│                                                                         │
│ // Field values                                                         │
│ fieldValue := v.Field(i)                    // By index                 │
│ fieldValue := v.FieldByName("Name")         // By name                  │
│ fieldValue.Interface()                      // Get actual value          │
└─────────────────────────────────────────────────────────────────────────┘

🔧 METHOD REFLECTION:
┌─────────────────────────────────────────────────────────────────────────┐
│ t := reflect.TypeOf(instance)                                           │
│ v := reflect.ValueOf(instance)                                          │
│                                                                         │
│ // Method information                                                   │
│ numMethods := t.NumMethod()                                             │
│ method := t.Method(i)                       // By index                 │
│ methodValue := v.MethodByName("MethodName") // By name                  │
│                                                                         │
│ // Call method                                                          │
│ args := []reflect.Value{                                                │
│     reflect.ValueOf("arg1"),                                            │
│     reflect.ValueOf(42),                                                │
│ }                                                                       │
│ results := methodValue.Call(args)                                       │
│                                                                         │
│ // Get result values                                                    │
│ if len(results) > 0 {                                                   │
│     result := results[0].Interface()                                    │
│ }                                                                       │
└─────────────────────────────────────────────────────────────────────────┘

🔄 VALUE MODIFICATION:
┌─────────────────────────────────────────────────────────────────────────┐
│ // Must use pointer to modify                                           │
│ ptr := reflect.ValueOf(&variable)                                       │
│ value := ptr.Elem()  // Dereference pointer                             │
│                                                                         │
│ // Check if settable                                                    │
│ if value.CanSet() {                                                     │
│     // Set different types                                              │
│     value.SetString("new string")                                       │
│     value.SetInt(42)                                                    │
│     value.SetFloat(3.14)                                                │
│     value.SetBool(true)                                                 │
│     value.Set(reflect.ValueOf(newValue))                                │
│ }                                                                       │
│                                                                         │
│ // Modify struct fields                                                 │
│ structValue := reflect.ValueOf(&structInstance).Elem()                  │
│ field := structValue.FieldByName("FieldName")                           │
│ if field.CanSet() {                                                     │
│     field.SetString("new value")                                        │
│ }                                                                       │
└─────────────────────────────────────────────────────────────────────────┘

🎯 COMMON PATTERNS:
┌─────────────────────────────────────────────────────────────────────────┐
│ // Type checking                                                        │
│ if t.Kind() == reflect.Struct {                                         │
│     // Handle struct                                                    │
│ }                                                                       │
│                                                                         │
│ // Interface implementation check                                       │
│ interfaceType := reflect.TypeOf((*MyInterface)(nil)).Elem()             │
│ if someType.Implements(interfaceType) {                                 │
│     // Type implements interface                                        │
│ }                                                                       │
│                                                                         │
│ // Create new instance                                                  │
│ newValue := reflect.New(someType)        // Creates pointer             │
│ instance := newValue.Elem()              // Dereference                 │
│ actualValue := newValue.Interface()      // Get interface{}             │
│                                                                         │
│ // Slice/Array operations                                               │
│ slice := reflect.ValueOf([]int{1, 2, 3})                               │
│ length := slice.Len()                                                   │
│ element := slice.Index(0)                                               │
│                                                                         │
│ // Map operations                                                       │
│ mapValue := reflect.ValueOf(map[string]int{"key": 42})                  │
│ keys := mapValue.MapKeys()                                              │
│ value := mapValue.MapIndex(reflect.ValueOf("key"))                      │
└─────────────────────────────────────────────────────────────────────────┘

⚡ PERFORMANCE CONSIDERATIONS:
• Reflection is slower than direct access
• Use sparingly in performance-critical code
• Cache reflect.Type and reflect.Value when possible
• Consider code generation alternatives
• Profile reflection-heavy code

🚨 COMMON PITFALLS:
❌ Trying to set unexported fields
❌ Not checking CanSet() before setting
❌ Forgetting to dereference pointers
❌ Not handling invalid reflect.Values
❌ Overusing reflection where generics would work

💡 BEST PRACTICES:
• Use reflection judiciously
• Always check IsValid() and CanSet()
• Handle panics from reflection operations
• Document reflection-heavy code well
• Consider alternatives (generics, interfaces)
• Cache reflection results when possible

🎯 REAL-WORLD USES:
• JSON/XML marshaling/unmarshaling
• ORM database mapping
• Dependency injection frameworks
• Testing frameworks
• Configuration parsing
• Generic data processing

=============================================================================
*/