package main

import "fmt"

func main() {
	for i:=1; i<=5;i++ {
		for j :=0; j<i;j++{
			fmt.Print("*")
		}
		fmt.Println()
	}
	fmt.Println()
	fmt.Println("*******************************")
	for m:=1;m<5; m++{
		for n:=0;n<5-m;n++{
			fmt.Print("*")
		}
		fmt.Println()

	}
}