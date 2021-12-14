// dynamic and mixed type 


package main
import "fmt"
func main()  {
	var x float64 = 32.244
	y := 345
	var a, b, c, z = 34, 55, "SHASHI", 34.34
	fmt.Println(x)
	fmt.Println(y) 
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(z)
	fmt.Printf("c is of type %T\n", c)
	// fmt.Printf("%T", x)
	fmt.Printf("y is type of %T \n", y)
	fmt.Printf("z is type of %T", z)

	
}