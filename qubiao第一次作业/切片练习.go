package main

import "fmt"

func main() {
	a:=[4]string{"哈哈","哈利油","烧鸡公","逗比"}
	fmt.Println(a)
	fmt.Printf("%T\n",a)
	b:=[]int{1,2,3,4,5,6,7}
	fmt.Println(b)
	fmt.Println("长度是：",len(b),"容量是：",cap(b))
	c:=b[3:5]
	fmt.Println(c)
	fmt.Println("长度是：",len(c),"容量是：",cap(c))
}
