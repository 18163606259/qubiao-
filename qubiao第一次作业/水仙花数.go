package main

import "fmt"

func main(){
	for i:=100;i<1000;i++{
		a:=i/100
		b:=i/10%10
		c:=i%10
		if a*a*a+b*b*b+c*c*c==i{
			fmt.Println(i)
		}
	}
	//方法二：a,b,c
	for a:=1;a<10;a++{
		for b:=0;b<10;b++{
			for c :=0;c<10;c++{
				n:=a*100+b*10+c
				
				fmt.Println(n)
			}
		}
	}
	
	fmt.Print()
	fmt.Println("——————————————————————————————————————————————————————————-")
	
	for a:=0;a<=20;a++{
		for b:=0;b<=33;b++{
			c:=100-a-b
			if a*5+b*3+c/3==100 && c&3 ==0{
				fmt.Println(a,b,c)
			}
		}
	}
	
}
