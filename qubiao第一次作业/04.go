package main

import (
	"fmt"
	)

func main(){
	/*
	switch的其他写法
	*/
	var c string
	c = "a"
	switch c {
	case "i", "o", "u", "e","a":
		fmt.Println("c是一个元音")
	default:
		fmt.Println("c是辅音")
	}
	//多一条初始化语句
	switch lan:="golang";lan {
	case "java":
		fmt.Println("java语言")
	case "c":
		fmt.Println("C语言")
	case "php":
		fmt.Println("php语言")
	case "golang":
		fmt.Println("GO语言")
	default:
		fmt.Println("神经语言")
		
	}
}