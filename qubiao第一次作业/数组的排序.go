package main

import "fmt"

func main() {
	/*
	数组的排序
	*/
	//冒泡排序
	arr:=[5]int{15,23,8,10,7}
	fmt.Println(arr)
	//目标：[7,8,10,15,23]
	for i:=1;i<len(arr);i++{
		for j:=0;j<len(arr)-i;j++{
			if arr[j]>arr[j+1]{
				arr[j],arr[j+1] = arr[j+1],arr[j]
			}
		}
	}
	fmt.Println(arr)
}
