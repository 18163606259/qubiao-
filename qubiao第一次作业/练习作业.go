package main

import "fmt"

func main (){
	for i:=100;i<=200;i++{
		if i % 3==1 && i % 4 ==2 && i % 5 ==3{
			fmt.Println(i)
			break
		}
	}
	fmt.Print()
	fmt.Println("————————————————————————————————————————")
	for x:=0;x<150;x++{
		y:=150-x
		if x/y==3 && x % y ==10{
			fmt.Println(x,y)
		}
		
	}
	fmt.Print()
	fmt.Println("————————————————————————————————————————")
	for i:=1000;i<10000;i++{
		x:=i%1000
		if x*3==i{
			fmt.Println(i)
		}
	}
	fmt.Print()
	fmt.Println("————————————————————————————————————————")
	count:=0
	for i:=1;i<=500;i++{
		a:=i/100
		b:=i/10%10
		c:=i%10
		if a!=4 && b!=4 && c!=4{
			fmt.Println(i)
			count++
		}
	}
	fmt.Println("共有：",count)
}
