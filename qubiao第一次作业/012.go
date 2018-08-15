package main

import "fmt"

func main(){
	res :=1
	for i:=1; i<=5;i++  {
		res *= i
	}
	fmt.Println(res)
	for i:=58;i>=23;i-- {
		fmt.Println(i)
	}
	for i :=65;i<=90;i++{
		fmt.Printf("%q\t",i)
	}
	fmt.Println()
	fmt.Println("_______________________________________________")
	for i:=1;i<=100;i++{
		fmt.Print(i,"\t")
		if i%10 ==0{
			fmt.Println()
		}
	}
}
