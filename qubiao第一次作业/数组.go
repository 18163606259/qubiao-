package main

import (
	"fmt"
	)

func main() {
	/*
	数组：array
	*/
	var arr1[5]int
	 fmt.Println(arr1)
		fmt.Println(len(arr1))//获取数组
		arr1[0]=1
		arr1[1]=3
		arr1[2]=5
		arr1[3]=7
		arr1[4]=9
		fmt.Println(arr1)
		fmt.Println(arr1[2])
		//遍历数组：一次访问数组中的每一个数
		for i:=0;i<len(arr1);i++{
		fmt.Println(arr1[i])
	}
		//数组的其他创建写法
		var a[4]int
		fmt.Println(a)
		var b = [5]int{6,7,8,9,10}
		fmt.Println(b)
		var c = [5]int{1,2,3}
		fmt.Println(c)
		var d=[5]int{4:100}//为index为4的位置存储数值100
		fmt.Println(d)
		f:=[4]string{"Rose","王二狗","哈利油","烧鸡公"}
		fmt.Println(f)
		g:=[...]int{1,2,3,4,5} //表示可以自动来推断数组的长度，根据给定的数值的数量
		fmt.Println(g)
		fmt.Println(len(g))
		h:=[...]int{1:100,3:200,9:300}
		fmt.Println(h)
		fmt.Println(len(h))
		
		
}
