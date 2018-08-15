package main

import "fmt"

func main() {
	//数组：arr:=[10]int{1,2,3,4,5,6,7,8,9,10} 中所有数据的和
	arr:=[10]int{1,2,3,4,5,6,7,8,9,10}
	sum :=0
	for i:=0;i<len(arr);i++{
		sum += arr[i]
	}
	fmt.Println(sum)
}
