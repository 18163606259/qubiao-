package main

import (
	"fmt"
	"math"
)

func main() {
	/*
	数据是四位数，在传递过程中加密，每位数字都加5的和除以10
	再将第一位和第四位交互，第二位和第三位交互
	
	3961
	8 14 11 6
	8 4 1 6
	6148
	*/
		num1:=3961
		arr:=[4]int{}
		for i:=0;i<len(arr);i++{
			arr[i] = num1 % int(math.Pow10(4-i))/int(math.Pow10(3-i))
		}
		fmt.Println(arr)
		for i:=0;i<len(arr);i++{
			arr[i]+=5
			arr[i]%=10
		}
		fmt.Println(arr)
		for i:=len(arr)-1;i>=0;i--{
			fmt.Print(arr[i],"\t")
		}
	}
