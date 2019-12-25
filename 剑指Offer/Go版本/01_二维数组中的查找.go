package main

import (
	"fmt"
)

func Find2(target int, array [2][3]int) bool {
	for _, v1 := range array {
		for _, v2 := range v1 {
			if v2 == target {
				return true
			}
		}
	}
	return false
}

func main() {
	var arr = [2][3]int{{1, 4, 3}, {7, 5, 6}}
	result := Find2(5, arr)
	fmt.Println(result)
}
