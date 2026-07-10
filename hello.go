package main

import (
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

	var myBool = true
	println(myBool)
	if myBool {
		fmt.Println("The variable is true")
	}
}
