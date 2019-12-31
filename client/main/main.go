package main

import (
	"fmt"
	"gocode/chat/client/process"
)

var userId int
var userPwd string
var userName string

func init() {

}
func main() {
	loop := true
	var key int
	for loop {
		fmt.Println("---------------------欢迎使用------------------")
		fmt.Println("\t\t\t1 登录系统")
		fmt.Println("\t\t\t2 注册系统")
		fmt.Println("\t\t\t3 退出系统")
		fmt.Print("请输入要执行的操作序号\n")
		fmt.Scanf("%d\n", &key)
		switch key {
		case 1:
			fmt.Println("\t\t\t 登录系统")
			fmt.Println("请输入登录id")
			fmt.Scanf("%d\n", &userId)
			fmt.Println("请输入登录密码")
			fmt.Scanf("%s\n", &userPwd)
			//将登录函数写到另外一个文件
			us := &process.UserProcess{}
			err := us.Login(userId, userPwd)
			if err != nil {
				fmt.Println(err)
			}
			//loop=false //停止循环，进入下一次下一子菜单
		case 2:
			fmt.Println("\t\t\t 注册系统")
			fmt.Println("请输入登录id")
			fmt.Scanf("%d\n", &userId)
			fmt.Println("请输入登录密码")
			fmt.Scanf("%s\n", &userPwd)
			fmt.Println("请输入登录名称")
			fmt.Scanf("%s\n", &userName)
			us := &process.UserProcess{}
			err := us.Register(userId, userPwd, userPwd)
			if err != nil {
				fmt.Println(err)
			}
		case 3:
			fmt.Println("\t\t\t3 退出系统")
			loop = false
		default:
			fmt.Println("输入有误，请重新操作")
		}
	}
	//根据key进入二级菜单
	if key == 1 {
		//登录
		fmt.Println("请输入登录id")
		fmt.Scanf("%d\n", &userId)
		fmt.Println("请输入登录name")
		fmt.Scanf("%s\n", &userPwd)
		//将登录函数写到另外一个文件
		us := process.UserProcess{}
		us.Login(userId, userPwd)
	} else {

	}

}
