package main

import "fmt"

func main(){
	//break 强制结束某个case的执行
	a :=2
	switch a {
	case 1:
		fmt.Println(1)
		fmt.Println(1)
		fmt.Println(1)
	case 2:
		fmt.Println(2)
		fmt.Println(2)
		break
		fmt.Println(2)
	case 3:
		fmt.Println(3)
		fmt.Println(3)
		fmt.Println(3)
	}
}
