package model

import (
	"gocode/chat/common/message"
	"net"
)

//设置为全局变量

//维护客户端自己 当前用户自己的信息   连接 注意：一个客户端只有一个自己的连接 不是服务端，维护所有客户端到服务器的连接
//并在登录成功后完成对CurrtUser的初始化比较合适
type CurrtUser struct {
	Conn         net.Conn `json:"conn"`
	message.User `json:"user"`
}
