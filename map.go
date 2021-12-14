 // maps, there is a key and set value 

package main
import "fmt"
func main(){
    var myMap map[string]int
    myMap = map[string]int{}
    myMap["A"]=12
    myMap["b"]=39
    myMap["C"]=48
    myMap["d"]=35
    myMap["e"]=64
    myMap["f"]=7
    
    // delete maps
    
    delete(myMap, "b")
    delete(myMap, "e")
   
    fmt.Println(myMap)
}