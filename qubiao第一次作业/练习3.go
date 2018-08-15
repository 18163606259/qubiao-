package main

import "fmt"

func main (){
	num :=65
	switch num/10  {
	case 0,1,2,3,4,5:
		fmt.Println(num,"不及格")
	case 6:
		fmt.Println(num,"及格")
	case 7:
		fmt.Println(num,"中等")
	case 8:
		fmt.Println(num,"良好")
	case 9,10:
		fmt.Println(num,"优秀")
	}
}
