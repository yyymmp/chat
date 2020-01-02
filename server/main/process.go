package main

import (
	"fmt"
	"gocode/chat/common/message"
	process2 "gocode/chat/server/process"
	utils2 "gocode/chat/server/utils"
	"io"
	"net"
)

type Processor struct {
	Conn net.Conn
}

//请求分发
//根据接受到的消息类型返回消息
func (this *Processor)  ServerProcessMes(mess *message.Message) (err error) {
	switch mess.Type {
	//接受的消息类型是登录消息
	case message.LoginMesType:
		up := process2.UserProcess{
			Conn: this.Conn,
		}
		err = up.ServerProcessLogin(mess)
	//接收的消息类型是注册类型
	case message.RegisterResMesType:
		up := process2.UserProcess{
			Conn: this.Conn,
		}
		err = up.ServerProcessRegister(mess)
	//群发消息
	case message.SmsMesType:
		fmt.Println("服务器接收到客户端的群发消息")
		up := &process2.SmsProcess{}
		//转发任务
		up.SendGroupMes(mess)
	default:
		fmt.Println("暂无此类型 ")
	}
	return
}

func (this *Processor) process2() (err error) {
	for {
		transfer := &utils2.Transfer{
			Conn: this.Conn,
		}
		mess, err := transfer.ReadPkg()
		if err != nil {
			if err == io.EOF {
				fmt.Println("客户端退出,服务器正常退出")
				return err
			} else {
				fmt.Println("readPag失败,服务器正常退出")
				return err
			}
		}
		fmt.Println("服务器反序列化得到的消息", mess)
		fmt.Println("进入serverProcessMes")
		//处理 扥发
		err = this.ServerProcessMes(&mess)
		if err != nil {
			return err
		}
	}
}
