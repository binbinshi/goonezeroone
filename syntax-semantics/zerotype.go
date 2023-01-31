package main

import "unsafe"

type A [0][256]int
type S struct {
	x A
	y [1<<30]A
	z [1<<30]struct{}
}

type T [1<<30]S

func main() {
	var a A
	var s S
	var t T
	println(unsafe.Sizeof(a))
	println(unsafe.Sizeof(s))
	println(unsafe.Sizeof(t))
}
