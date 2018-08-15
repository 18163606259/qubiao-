package main

import (
	"fmt"
	)

func main() {
	/*
	切片
	*/
	//创建数组
	a:=[4]int{1,2,3,4}
	fmt.Println(a)
	fmt.Printf("%T\n",a)
	fmt.Println("数组的长度：",len(a),"数组的容量：",cap(a))
	//切片
	//1.
	s1:=[]int{1,2,3,4}
	fmt.Println(s1)
	fmt.Printf("%T\n",s1)
	fmt.Println("切片的长度：",len(s1),"切片的容量：",cap(s1))
	//2.
	s2:=[]int{}
	fmt.Println(s2)
	fmt.Println("切片的长度：",len(s2),"切片的容量：",cap(s2))
	//3.
	s3:=make([]int,3,8)
	fmt.Println(s3)
	fmt.Printf("长度是：%d,容量是：%d\n",len(s3),cap(s3))
	//修改内容
	s3[0]=1
	s3[1]=2
	s3[2]=3
	fmt.Println(s3)
	//遍历
	for i:=0;i<len(s3);i++{
		fmt.Println(s3[i])
	}
	//4.直接从数组上创建切片
	b:=[10]int{1,2,3,4,5,6,7,8,9,10}
	fmt.Println(b)
	s4:=b[1:6]
	fmt.Println(s4)
	fmt.Printf("长度是：%d,容量是：%d\n",len(s4),cap(s4))
	s5:=b[0:6]
	fmt.Print(s5)
	fmt.Printf("长度是：%d,容量是：%d\n",len(s5),cap(s5))
	s6:=b[5:]
	fmt.Print(s6)
	fmt.Printf("长度是：%d,容量是：%d\n",len(s6),cap(s6))
	
}
