package main

import "fmt"

func main() {
	var a = [5][2]int{{2, 5}, {5, 6}, {8, 4}, {44, 59}, {49, 29}}
	var i, j int
	for i = 0; i < 5; i++ {
		for j = 0; j < 2; j++ {
			fmt.Printf("a[%d][%d] = %d\n", i, j, a[i][j])
		}

	}
}
