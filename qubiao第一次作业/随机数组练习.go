package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	/*
	给定一个整数数组，长度为10，数字取自随机数
	*/
	arr:=[10]int{}
	fmt.Println(arr)
	rand.Seed(time.Now().UnixNano())
	for i:=0;i<len(arr);i++{
		x:=0
		//判断没有在数组中储存过
		for {
			x= rand.Intn(10)+1
			flag :=true
			for j:=0;j<i;j++{
				if arr[j]==x{
					flag = false
					break
				}
			}
			if flag{
				break
			}
		}
		arr[i]=x
	}
	fmt.Println(arr)
}