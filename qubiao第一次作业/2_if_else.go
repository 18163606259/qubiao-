package main

import "fmt"

func main(){
	a := 20
	if a >= 18 {
		fmt.Println("成年")
	}else {
		fmt.Println("未成年")
	}
	b :="女"
	if b =="男"{
		fmt.Println("请去男厕所")
	}else {
		fmt.Println("请去女厕所")
	}
	num :=60
	if num %2 == 0 {
		fmt.Println("偶数")
	}else {
		fmt.Println("奇数")
	}
}