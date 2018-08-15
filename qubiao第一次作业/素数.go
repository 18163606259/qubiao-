package main

import (
		"math"
	"fmt"
)

func main() {
	for i := 2; i <= 100; i++ {
		count := 0
		for j := 2; float64(j) < math.Sqrt(float64(i)); j++ {
			if i%j == 0 {
				count++
				break
			}
		}
		if count== 0{
			fmt.Println(i,"是素数")
		}
	}
}