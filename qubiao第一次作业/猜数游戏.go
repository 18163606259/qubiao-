package main

import (
	"math/rand"
	"time"
	"fmt"
)

func main() {
	/*
	 猜数游戏
	 */
	rand.Seed(time.Now().UnixNano())
	num := rand.Intn(100) + 1
	
	a := 0
	for i:=1;;i++{
		fmt.Printf("请输入第%d次的猜测数值：",i)
		fmt.Println()
		fmt.Scanln(&a)
		if a > num {
			fmt.Println("猜大了，下次小点")
		} else if a < num {
			fmt.Println("猜小了，下次猜大点'")
		} else {
			fmt.Printf("恭喜你猜对了！共猜%d次",i)
			break
		}
	}
}