package main

import "fmt"

func main() {
	var x int = 24
	var ip *int
	ip = &x
	fmt.Printf("address of s a variable %d\n", &x)
	fmt.Printf("address stored in ip variable %d\n", ip)
	fmt.Printf("value of ip variable is %d", *ip)

}
