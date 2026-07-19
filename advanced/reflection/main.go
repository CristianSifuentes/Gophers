// Package main demonstrates runtime metaprogramming: parsing struct tags
// with reflect, and the intended use of `go generate` and the race detector.
//
// Reflection bypasses compile-time safety and costs more at runtime, so it's
// used sparingly here — just enough to build a generic struct-tag parser.
//
// The code in this package is intentionally free of data races. To see the
// race detector catch a real one, run:
//
//	go run -race ./advanced/reflection
//
// on code that writes a shared variable from multiple goroutines without
// synchronization (see racyExample, left unused on purpose).
//
//go:generate echo "this is where a real generator command would run"
package main

import (
	"fmt"
	"reflect"
)

// User's `tag` struct tags are ordinary metadata until something reads them
// at runtime via reflection — that's what describeFields does below.
type User struct {
	Name  string `tag:"name"`
	Email string `tag:"email"`
	Age   int    `tag:"age"`
}

// describeFields walks any struct's fields with reflect and prints each
// field's name, tag, and value — a generic serializer needs exactly this.
func describeFields(v any) {
	t := reflect.TypeOf(v)
	val := reflect.ValueOf(v)

	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		fmt.Printf("field=%s tag=%s value=%v\n",
			field.Name, field.Tag.Get("tag"), val.Field(i))
	}
}

// racyExample shows (without running) the shape of a data race: two
// goroutines writing counter concurrently with no synchronization.
// `go run -race` on a call site that invokes this would report it.
func racyExample() int {
	counter := 0
	done := make(chan struct{})
	for i := 0; i < 2; i++ {
		go func() {
			counter++ // unsynchronized write — flagged by -race
			done <- struct{}{}
		}()
	}
	<-done
	<-done
	return counter
}

func main() {
	u := User{Name: "Grace", Email: "grace@example.com", Age: 36}
	describeFields(u)

	fmt.Println("type via reflect:", reflect.TypeOf(u))
	fmt.Println("kind via reflect:", reflect.ValueOf(u).Kind())
}
