package main

import "fmt"

func main() {
	var x = []int{3, 5, 6, 2, 4}
	var avg float32
	avg = getavg(x, 5)
	fmt.Printf("avg is %f", avg)
}
func getavg(arr []int, size int) float32 {
	var i, sum int
	var avg float32
	for i = 0; i < size; i++ {
		sum = sum + arr[i]
	}
	avg = float32(sum / size)
	return avg
}
