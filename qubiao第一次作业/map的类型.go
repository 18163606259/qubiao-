package main

import "fmt"

func main() {
	/*
	map的类型
	*/
	//引用类型
	map1:=make(map[string]string)
	map1["name"]="哈利油"
	fmt.Printf("%T\n",map1)
}
