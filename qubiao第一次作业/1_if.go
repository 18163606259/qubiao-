package main

import "fmt"

func main(){
	/*
	if语句：做选择结构体
	*/
	a := 23
	if a >= 22{
		fmt.Println(a,"成年。。")
	}
	b := 5
	if b < 0{
		b = -b
	}
	fmt.Println(b)

  	c := 65
  	if c <= 60{
  		fmt.Println(c,"不及格")
	}else {
		fmt.Println(c,"及格")
	}
}
