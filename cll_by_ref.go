// call by reference

package main
import "fmt"
func main(){
    var a int = 102
    var b int = 345
    fmt.Printf("before swapping %d\n", a)
    fmt.Printf("before swapping %d\n", b)
    swap(&a, &b)
    fmt.Printf("after swapping %d\n", a)
    fmt.Printf("after swapping %d", b)
    
}
func swap(x *int, y *int){
    var temp int
    temp = *x
    *x = *y
    *y = temp
    
}

// call by value



package main
import "fmt"
func main(){
    var a int = 102
    var b int = 345
    fmt.Printf("before swapping %d\n", a)
    fmt.Printf("before swapping %d\n", b)
    swap(a, b)
    fmt.Printf("after swapping %d\n", a)
    fmt.Printf("after swapping %d", b)
    
}
func swap(x int, y int){
    var temp int
    temp = x 
    x = y
    y = temp
    
}