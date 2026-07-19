// Package main demonstrates high-performance memory management:
// value vs. pointer receiver semantics, and sync.Pool for allocation reuse.
//
// To see escape analysis in action, run:
//
//	go build -gcflags="-m" ./advanced/memory
//
// and look for "escapes to heap" / "does not escape" diagnostics.
package main

import (
	"bytes"
	"fmt"
	"sync"
)

// Point uses a value receiver: calling Translate never mutates the caller's
// copy, which keeps Point cheap, immutable-by-default, and stack-friendly.
type Point struct{ X, Y int }

func (p Point) Translate(dx, dy int) Point {
	p.X += dx
	p.Y += dy
	return p
}

// Buffer is large enough that copying it on every call would be wasteful,
// so its methods use a pointer receiver to mutate shared state in place.
type Buffer struct {
	data [1024]byte
	len  int
}

func (b *Buffer) Reset() {
	b.len = 0
}

func (b *Buffer) Append(p []byte) {
	n := copy(b.data[b.len:], p)
	b.len += n
}

// bufferPool reuses *bytes.Buffer instances instead of allocating a fresh
// one per request, keeping the memory footprint stable under load.
var bufferPool = sync.Pool{
	New: func() any {
		return new(bytes.Buffer)
	},
}

func render(id int) string {
	buf := bufferPool.Get().(*bytes.Buffer)
	buf.Reset()
	defer bufferPool.Put(buf)

	fmt.Fprintf(buf, "response-%d", id)
	return buf.String()
}

func main() {
	// Value receiver: original is untouched, a translated copy is returned.
	origin := Point{X: 0, Y: 0}
	moved := origin.Translate(3, 4)
	fmt.Println("origin unchanged:", origin, "moved copy:", moved)

	// Pointer receiver: mutates the same underlying array in place.
	var buf Buffer
	buf.Append([]byte("hello"))
	buf.Append([]byte(" world"))
	fmt.Println("buffer bytes:", buf.len)

	// sync.Pool: reused buffers across many simulated requests.
	for i := 0; i < 3; i++ {
		fmt.Println("pooled render:", render(i))
	}
}
