package main

import (
	"fmt"
	"sort"
)

func main() {
	/*
	遍历map集合
	*/
	//无序的遍历map
	map1:=map[int]string{1:"英雄联盟",2:"绝地求生",3:"王者荣耀",4:"csOnline"}
	//使用for range
	for k,v:=range map1{
		fmt.Println(k,v)
	}
	//有序的遍历map
	//1.获取map中的key
	s1:=make([]int,0,len(map1))
	for k,_:=range map1{
		s1=append(s1,k)
	}
	fmt.Println(s1)
	//冒泡排序
	sort.Ints(s1)//使用sort包下的方法Ints()给int类型的切片排序
	//根据排序后的s1获取map中对应的value
	for _,k:=range s1{
		fmt.Println(k,map1[k])
	}
	
}
