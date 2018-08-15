package main

import "fmt"

func main (){

	var a ,b int
	fmt.Println("请输入两个整数：")
	fmt.Scanln(&a,&b)
	fmt.Println("请输入一个运算操作：+，-，*，/")
	var oper string
	fmt.Scanln(&oper)
	fmt.Println(a,b,oper)
	switch oper {
	case "+":
		fmt.Printf("%d %s %d = %d",a,oper,b(a+b))
	case "-":
		fmt.Printf("%d %s %d = %d",a,oper,b(a-b))
	case "*":
		fmt.Printf("%d %s %d = %d",a,oper,b(a*b))
	case "/":
		fmt.Printf("%d %s %d = %d",a,oper,b(a/b))
	default:
		fmt.Println("无法计算")


	}
}