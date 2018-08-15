package main

import "fmt"

func main() {
	/*
	map,表示映射，用于存储健值对
	*/
	//1.定义map
	var map1 map[int]string
	map2:=map[string]int{"GO":88,"c":89,"java":79,"php":85,"c++":68}
	map3:=make(map[string]int)
	fmt.Println(map1,map2,map3)
	fmt.Println("——————————————————————")
	if map1 ==nil{
		map1=make(map[int]string)
		fmt.Println(map1==nil)
	}
	//向map中添加映射：健值对map名[key]=value
	map1[1]="哈利油"
	map1[2]="烧鸡公"
	map1[3]="蠢得死"
	fmt.Println(map1)
	//获取map中的数据，根据key获取value
	fmt.Println(map1[1])
	fmt.Println(map1[10])
	fmt.Println(map3)
	fmt.Println(map3["N方向有个傻叉！"])
	val1,ok :=map1[1]
	if ok==true{
		fmt.Println("获取到的数据是：",val1)
	}else {
		fmt.Println("key不存在，无法获取。。")
	}
	fmt.Println("---------------")
	//删除map中的键值对
	fmt.Println(map2)
	delete(map2,"java")
	fmt.Println(map2)
}
