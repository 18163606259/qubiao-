package main

import "fmt"

func main(){
	for i:=1;i<=100;i++{
		if i % 2 == 0{
			fmt.Println(i)
		}
	}
	fmt.Println()
	fmt.Println("***********************************")
	sum :=0
	for i:=1;i<=100;i+=2{
		sum+=i
	}
	fmt.Println(sum)
	fmt.Println()
	fmt.Println("***********************************")
	count :=0
	for i:=1;i<=100;i++{
		if i %3 == 0 && i % 5 !=0{
			fmt.Print(i,"\t")
			count++
			if count % 5 ==0{
				fmt.Println()
			}
		}

	}
}
