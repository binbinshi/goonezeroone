package main

var g *[0]int
var a,b [0]int
func f() *[0]int{
	return new([0]int)
}

func main() {
	var x,y,z,w [0]int
	g = &z; g =&w
	println(&b == &a)
	println(&x == &y)
	println(&z == &w)
	println(&z == f())
	// go 1.18.1
	/**
	false
	false
	true
	false
	*/

	// go 1.19
	/**
	false
	false
	true
	true
	*/
}
