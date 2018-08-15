package main

import "fmt"

func main() {
	arr1 :=[4]int{1,2,3,4}
	fmt.Println(arr1)
	arr2 :=arr1
	fmt.Println(arr2)
	arr2[0]=100
	fmt.Println(arr1,arr2)
	num4:=1
	num5:=1
	fmt.Println(num4==num5)
	arr3:=[4]int{6,7,8,9}
	fmt.Println(arr1==arr3)
}
