package main

import "fmt"

func main (){
	/*
	键盘输入
	*/
	a := 10
	fmt.Println(a)
	fmt.Println("变量a的内存地址是：",&a)

	var num1 int
	var s1 string
	fmt.Println("请输入数据：")
	//fmt.Scanln(&num1,&s1)
	fmt.Scanf("%d,%s",&num1,&s1)
	fmt.Println("num1",num1)
	fmt.Println("s1",s1)
}