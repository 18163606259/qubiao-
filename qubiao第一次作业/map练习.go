package main

import (
	"fmt"
	)

func main() {
	//创建map
	map1:=make(map[string]string)
	map1["name"]="王二狗"
	map1["age"]="30"
	map1["sex"]="男性"
	map1["address"]="长沙市时代阳光大道789号"
	fmt.Println(map1["name"])
	fmt.Println(map1["age"])
	fmt.Println(map1["sex"])
	fmt.Println(map1["address"])
	//修改
	map2:=make(map[string]string)
	map2["name"]="哈利油"
	map2["age"]="30"
	map2["sex"]="男性"
	map2["address"]="长沙市时代阳光大道789号"
	fmt.Println(map2["name"])
	fmt.Println(map2["age"])
	fmt.Println(map2["sex"])
	fmt.Println(map2["address"])
	
	s1:=make([]map[string]string,0,10)
	s1=append(s1,map1)
	s1=append(s1,map2)
	for i:=0;i<len(s1);i++{
		fmt.Printf("第%d个人的信息是：\n",i+1)
		for k,v:=range s1[i] {
			fmt.Println("\t",k,",",
			v)
			
		}
	}
}
