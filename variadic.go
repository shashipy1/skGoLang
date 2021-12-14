	// variadic function
// pass unlimited number of parameters and not required to defined 
package main
import "fmt"
func variadic(params...int){
    fmt.Printf("%T\n %v\n",params,params)
}

func main(){
    variadic(3,45,5,3,5,5,6,2,3,7)
}