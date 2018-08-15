package main

import "fmt"

func main() {
	/*
	copy()
	*/
	m:=[]int{1,2,3,4,5}
	n:=[]int{7,8,9}
	fmt.Println(m)
	fmt.Println(n)
	
	//拷贝数据
	copy(m,n)//将N的数据拷贝到M里
	fmt.Println(m)
	fmt.Println(n)
	fmt.Println("————————————————————")
	copy(n,m[1:4])//将m中的下标1到3的数据拷贝到n里
	fmt.Println(m)
	fmt.Println(n)
	fmt.Println("————————————————————")
	copy(n[1:],m[3:])
	fmt.Println(m)
	fmt.Println(n)
	fmt.Println("————————————————————")
	s3 := [] int {1,2,3,4}
	s4 := s3
	s5 :=make([]int,4,4)
	copy(s5,s3)
	fmt.Println(s4,s5)
}

