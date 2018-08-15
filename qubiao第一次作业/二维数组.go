package main

import "fmt"

func main() {
	/*
	二维数组
	*/
		b:=[3][4]int{
			{1,2,3,4},
			{5,6,7,8},
			{9,10,11,12}}
		fmt.Println(b)
		fmt.Println("二维数组的长度：",len(b))
		fmt.Println("第一个数组：",b[0])
		fmt.Println("第一个一维数组的长度：",len(b[0]))
		fmt.Println("第一个一维数组的第一个数据：",b[0][0])
		fmt.Println("第二个一维数组的第三个数据：",b[1][2])
		fmt.Println("第三个一维数组的第四个数据：",b[2][3])
		//打印输出二维中的每一个数
		for i:=0;i<len(b);i++{
			for j:=0;j<len(b[i]);j++{
				fmt.Print(b[i][j],"\t")
			}
			fmt.Println()
		}
		fmt.Println("——————————————————————————————————")
		//使用range遍历访问
		for _,arr:=range b{
			for _,v:=range arr{
				fmt.Print(v,"\t")
			}
			fmt.Println()
		}
	}

