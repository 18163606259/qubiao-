package main

import "fmt"

func main() {
	/*
	向切片中添加数据
	*/
	s1:=[]int{}
	fmt.Println(s1)
	fmt.Println(len(s1),cap(s1))
	fmt.Printf("%p\t",s1)
	s1=append(s1,1)
	fmt.Println(s1)
	fmt.Println(len(s1),cap(s1))
	fmt.Printf("%p\t",s1)
	s1 = append(s1,2)
	fmt.Println(s1)
	fmt.Println(len(s1),cap(s1))
	fmt.Printf("%p\t",s1)
	s1 =append(s1,3)
	fmt.Println(s1)
	fmt.Println(len(s1),cap(s1))
	fmt.Printf("%p\t",s1)
	s1 = append(s1,4,5,6)
	fmt.Println(s1)
	fmt.Println(len(s1),cap(s1))
	s1 = append(s1,7,8,9)
	fmt.Println(s1)
	fmt.Println(len(s1),cap(s1))
}
