// sclising and appends

package main
import "fmt"
func main(){
    arr:=[5]int{5,24,56,2,6}
    slc:=arr[:]
    slc = append(slc,49)
    fmt.Println(slc,len(slc), arr, len(arr))
}