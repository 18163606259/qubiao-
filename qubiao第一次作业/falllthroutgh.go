package main

import (
	"fmt"
	)

func main(){
	/*
	穿透switch
	*/
	a :=2
	switch a {
	case 1:
		fmt.Println("第一季度")
	case 2:
		fmt.Println("第二季度")
		fallthrough//向下穿透
	case 3:
		fmt.Println("第三季度")
		fallthrough
	case 4:
		fmt.Println("第四季度")
	}
	month := 5
	day:= 0
	switch month {
	case 1,3,5,7,8,10,12:
		day = 31
	case 4,6,9,11:
		day = 30
	case 2:
		day = 28
	}
	fmt.Println(month,"月的天数是：",day,"天")
}
