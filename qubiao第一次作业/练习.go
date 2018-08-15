package main

import "fmt"

func main() {
	a := -1
	if a > 0 {
		fmt.Println("整数")
	} else if a <0 {
		fmt.Println("负数")
	}else {
		fmt.Println("零")
	}
	b:=4
	if b ==1 {
		fmt.Println("熊大")
	}else if b ==2{
		fmt.Println("熊二")
	}else if b ==3{
		fmt.Println("光头强")
	}else {
		fmt.Println("不知道")
	}
 	num :=50
	if num < 60{
		fmt.Println("不及格")
	}else if num < 70{
		fmt.Println("及格")
	}else if num <80{
		fmt.Println("中等")
	}else if num <90{
		fmt.Println("良好")
	}else if num <=100 {
		fmt.Println("优秀")
	}

}