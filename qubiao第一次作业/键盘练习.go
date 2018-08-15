package main

import "fmt"

func main(){
	var user,pwd string
	fmt.Println("用户名：")
	fmt.Scanln(&user)
	fmt.Println("密码：")
	fmt.Scanln(&pwd)
	if user =="admin" &&pwd =="123" ||user =="zhangsan" &&pwd =="zhangsan123"{
		fmt.Println("登录成功！")
	}else {
		fmt.Println("登陆失败！")
	}
}
