package main

import "fmt"

func main() {
	/*
	切片的数据类型
	*/
	s1:=make([]int,4,8) //切片
	s2:=[]int{1,2,3,4}
	a1:=[4]int{1,2,3,4}
	fmt.Println(s1,s2,a1)
	fmt.Printf("%T,%T,%T\n",s1,s2,a1)
	fmt.Println("——————————————————————————————————")
	a2:=a1
	fmt.Println(a1,a2)
	a2[0] =100
	fmt.Println(a1,a2)
	s3:=s2
	fmt.Println(s2,s3)
	s3[0]=100
	fmt.Println(s2,s3)
	fmt.Printf("%p,%p\n",s2,s3)//以地址的格式打印
}
