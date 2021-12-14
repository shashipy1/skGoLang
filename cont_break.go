package main
import "fmt"
func main()  {
	var a int = 15
	for a < 25 {
	// 	a++
	// 	fmt.Printf("%d\n", a)
		// if a == 20{
		// 	break
		if a == 20{
			a++
			continue
		}
		fmt.Printf("%d\n", a)
		a++
	}
	
}