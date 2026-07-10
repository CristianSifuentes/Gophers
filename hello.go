package main

import (
	"container/list"
	"fmt"
	"reflect"
)

func main() {
	println("Hi, this is the firt print in console using go")

	// This is a comment
	/*
		multi comments
	*/

	// variables
	var name string = "variable"
	println(name)

	name = "other vriable"
	println(name)

	// name = 6
	// cannot use 6 (untyped int constant) as string value in assignmentcompilerIncompatibleAssign

	//Strong typed
	// infer data type

	var otherVariable2 = ""
	println(otherVariable2)

	var myIny int32 = 7
	myIny = myIny * 10
	println(myIny)

	myIny = 10

	// println(myIny + name)
	// invalid operation: myIny + name (mismatched types int32 and string)compilerMismatchedTypes
	println(name + string(myIny))

	println("%s", name, myIny, "he")

	fmt.Println("reflect int32", reflect.TypeOf(myIny))

	var myFloat = 6.5
	println(myFloat)
	fmt.Println("reflect float64", reflect.TypeOf(myFloat))
	println(float64(myIny) + myFloat)

	var myBool bool = true
	println(myBool)
	fmt.Println(ternary(myBool, "The variable is true", "The variable is false"))

	// variable declared and initialized abreviaded way
	myString := "Variable"
	fmt.Println(myString)

	//Constants
	const myConstant = "Constant"
	fmt.Println(myConstant)

	//Control Flow
	if true && myBool {
		println("True")
	} else if myString == "X" {
		println("myString = X")
	} else {
		println("else")
	}

	// Data Structre
	// Array

	var myArray [3]int
	println(len(myArray))
	println("Array[2]", myArray[2])

	myArray[2] = 3
	// myArray[4] = 3
	// invalid argument: index 4 out of bounds [0:3]compilerInvalidIndex
	// is a smart lenguaje

	println("Before Array[2]", myArray[2])

	// Map
	myMap := make(map[string]int)
	myMap["Cris"] = 35
	myMap["Artias"] = 20
	myMap["John"] = 29
	println(myMap)

	myMap2 := map[string]int{"Cris": 36, "Artias": 30}
	myMap2["Cris"] = 36
	println(myMap2)

	//List
	myList := list.New()
	myList.PushBack(1)
	myList.PushBack(2)
	myList.PushBack(3)
	myList.PushBack(4)
	//In Go, a pointer is a variable that stores the memory address of another value instead of the value itself.
	println(myList.Back().Value)

	// Bucles
	for i := 0; i < myList.Len(); i++ {
		println(i)

	}

	for i := 0; i < len(myMap2); i++ {
		println(i)

	}

	for index, value := range myArray {
		fmt.Println(index, value)
	}

}

// ternary mimics the ?: operator Go doesn't have: a function call is an
// expression, so its result can be passed directly into another call.
func ternary(cond bool, whenTrue, whenFalse string) string {
	if cond {
		return whenTrue
	}
	return whenFalse
}
