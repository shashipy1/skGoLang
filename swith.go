package main
import "fmt"
func main()  {
	var s string = "Rahul"
	switch {
	case s == "S":
		fmt.Println("S IS EQUAL TO R")	
	
	case s == "":
		fmt.Println("OK")	
	
	case s == "Rahul":
		fmt.Println("thik hai bahi")
default:
	fmt.Println("Nothing")
}

}