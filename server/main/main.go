package main

import (
	"fmt"
	"gocode/chat/server/model"
	"net"
	"time"
)

func process(conn net.Conn) {
	fmt.Println("一个客户端连接成功")
	defer conn.Close()
	process := &Processor{Conn: conn}
	err := process.process2()
	if err != nil {
		return
	}
}

//全局值需要一个dao，在redis初始化后 进行初始化
func initDao() {
	model.MyUserDao = model.NewUserDao(pool)
}
func init() {
	//服务开始时初始化redis
	initPool("127.0.0.1:6379", 16, 0, 300*time.Second)
	//初始化redis连接池后，理解创建出dao
	initDao()
}

//服务器监听
func main() {
	listen, err := net.Listen("tcp", "0.0.0.0:8899")
	defer listen.Close()
	if err != nil {
		fmt.Println("服务器监听失败")
		return
	}
	for {
		fmt.Println("正在等到客户端连接服务器.....")
		conn, err := listen.Accept()
		if err != nil {
			fmt.Println("一个客户端连接失败，请重连")
		}
		//启动一个协程处理套接字
		go process(conn)
	}
}
