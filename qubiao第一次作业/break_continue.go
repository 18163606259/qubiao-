package main

import "fmt"

func main(){
	//从1到100内能被3整除，不能被5整除的数，只输出前5个
	count :=0
	for i:=1;i<=100;i++{
		if i % 3 ==0 && i % 5 !=0{
			fmt.Println(i,"\t")
			count++
			if count ==5{
				break
			}
		}
		
	}
}
