// Package main demonstrates type system mastery: type-safe generics,
// interface segregation, and type assertions/switches.
package main

import "fmt"

// Number constrains Sum to any built-in numeric type, so callers get
// compile-time type safety instead of the runtime risk of interface{}.
type Number interface {
	~int | ~int32 | ~int64 | ~float32 | ~float64
}

func Sum[T Number](values []T) T {
	var total T
	for _, v := range values {
		total += v
	}
	return total
}

// Stack is a generic, type-safe LIFO container: one implementation works
// for any element type, with no runtime type assertions required by callers.
type Stack[T any] struct {
	items []T
}

func (s *Stack[T]) Push(v T) {
	s.items = append(s.items, v)
}

func (s *Stack[T]) Pop() (T, bool) {
	var zero T
	if len(s.items) == 0 {
		return zero, false
	}
	last := s.items[len(s.items)-1]
	s.items = s.items[:len(s.items)-1]
	return last, true
}

// Interface segregation: tiny, single-method interfaces compose cleanly and
// are trivial to mock, unlike one large do-everything interface.
type Named interface {
	Name() string
}

type Aged interface {
	Age() int
}

type Person struct {
	name string
	age  int
}

func (p Person) Name() string { return p.name }
func (p Person) Age() int     { return p.age }

func describe(n Named) string {
	return fmt.Sprintf("this is %s", n.Name())
}

// classify uses a type switch to safely unpack an unpredictable payload
// at runtime, without panicking on unexpected concrete types.
func classify(payload any) string {
	switch v := payload.(type) {
	case int:
		return fmt.Sprintf("int:%d", v)
	case string:
		return fmt.Sprintf("string:%q", v)
	case Named:
		return fmt.Sprintf("named:%s", v.Name())
	default:
		return fmt.Sprintf("unknown type %T", v)
	}
}

func main() {
	fmt.Println("sum ints:", Sum([]int{1, 2, 3, 4}))
	fmt.Println("sum floats:", Sum([]float64{1.5, 2.5, 3.0}))

	var s Stack[string]
	s.Push("a")
	s.Push("b")
	top, ok := s.Pop()
	fmt.Println("stack pop:", top, ok)

	p := Person{name: "Ada", age: 30}
	fmt.Println(describe(p))

	for _, payload := range []any{42, "hello", p, 3.14} {
		fmt.Println("classify:", classify(payload))
	}
}
