package main

import "fmt"

func main (){
	year :=2012
	if year % 4 ==0 && year% 100 !=0 ||year %400 ==0{
		fmt.Println(year,"是闰年")
	}else {
		fmt.Println("不是闰年")
	}
}
