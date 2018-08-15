package main

import (
	"math/rand"
	"fmt"
	"time"
)

func main() {
	/*
	随机数
	*/
	t1:=time.Now()
	//fmt.Println(t1)
	timeStamp1:=t1.Unix()
	rand.Seed(timeStamp1)
	num1:=rand.Intn(10)
	fmt.Println(num1)
	num2:=rand.Intn(10)
	fmt.Println(num2)
	num3:=rand.Intn(10)
	fmt.Println(num3)
	num4:=rand.Intn(10)
	fmt.Println(num4)
	fmt.Print()
	fmt.Println("————————————————————")
	num5:=rand.Intn(100)+1
	fmt.Println(num5)
}
