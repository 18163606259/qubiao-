package main

import "fmt"

func main() {
	//获取数组中的元素的下标和元素的数值
	arr:=[5]int{19,32,56,78,61}
	for i:=0;i<len(arr);i++{
		fmt.Println(i,arr[i])
	}
	fmt.Println("————————————————————————————————")
	for i,v:=range arr{
		fmt.Println("下标：",i,"数值：",v)
	}
	fmt.Println("————————————————————————————")
	sum :=0
	for  _,v:=range arr{
		sum +=v
	}
	fmt.Println(sum)
}
