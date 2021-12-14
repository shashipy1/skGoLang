package main

import "fmt"

func main() {
	var x int = 34
	var ptr *int
	var pptr **int
	ptr = &x
	pptr = &ptr
	fmt.Printf("value of x is %d\n", x)
	fmt.Printf("value of x is %d\n", *ptr)
	fmt.Printf("value of x is %d\n", **pptr)
	fmt.Printf("add of var is %d\n", &x)
	fmt.Printf("add of var is %d\n", ptr)
	fmt.Printf("add of var is %d\n", pptr)

}
