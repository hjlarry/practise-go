package main

import "testing"

//go:noinline
func test1() int {
	return 100
}

func test2() int {
	return 100
}

// go test -bench . inline_test.go
// BenchmarkTest/noinline-2         	775735896	         1.46 ns/op
// BenchmarkTest/inline-2           	1000000000	         0.554 ns/op
func BenchmarkTestInline(b *testing.B) {
	b.Run("noinline", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			test1()
		}
	})
	b.Run("inline", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			test2()
		}
	})
}

//go:noinline
func test3() *int {
	x := 100
	return &x
}

func test4() *int {
	x := 100
	return &x
}

// go test -bench . -benchmem inline_test.go
// BenchmarkTestEscape/escape-2           	86273755	        15.5 ns/op	       8 B/op	       1 allocs/op
// BenchmarkTestEscape/noescape-2         	1000000000	         0.408 ns/op	       0 B/op	       0 allocs/op
func BenchmarkTestEscape(b *testing.B) {
	b.Run("escape", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = test3()
		}
	})
	b.Run("noescape", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = test4()
		}
	})
}
