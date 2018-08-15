package main

import "fmt"

func main (){
	/*
	switch语句的使用
	*/
	a :=5
	switch a {
	case 1:
		fmt.Println("第一季度")
		fmt.Println("第一季度")
	case 2:
		fmt.Println("第二季度")
		fmt.Println("第二季度")
	case 3:
		fmt.Println("第三季度")
		fmt.Println("第三季度")
	case 4:
		fmt.Println("第四季度")
		fmt.Println("第四季度")
	default:
		fmt.Println("默认季度")
	}

}
