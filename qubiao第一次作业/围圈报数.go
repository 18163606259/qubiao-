package main

import "fmt"

func main()  {
	const n=7
		//创建bool类型的数组
		arr:=[n]bool{}
		for i:=0;i<len(arr);i++{
			arr[i]= true
		}
		fmt.Println(arr)
		//定义需要的变量
		count :=0    //计数器
		sum :=0     //需要退出的人数
		index:=0    //数组的下标
		for sum <n-1 {
			if arr[index] == true {
				count++
				if count == 3 {
					count = 0
					sum++
					arr[index] = false
				}
				
			}
			index++
			if index == len(arr) {
				index = 0
			}
		}
		fmt.Println(arr)
		for i:=0;i<len(arr);i++{
			if arr[i] == true{
				fmt.Println(i+1)
			}
		}
}
