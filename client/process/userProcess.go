package process

import (
	"encoding/binary"
	"encoding/json"
	"fmt"
	"gocode/chat/client/utils"
	"gocode/chat/common/message"
	"net"
	"os"
)

type UserProcess struct {
}

//客户端--登录
func (this *UserProcess) Login(userId int, userPwd string) (err error) {
	//连接服务器
	conn, err := net.Dial("tcp", "127.0.0.1:8899")
	if err != nil {
		fmt.Println("客户端主动连接服务器端失败")
		return
	} else {
		//fmt.Println("客户端主动连接服务器套接字成功")
	}
	defer conn.Close()
	//发送消息给服务器
	var mess message.Message
	var loginMes message.LoginMes
	loginMes = message.LoginMes{
		UserId:  userId,
		UserPwd: userPwd,
	}
	serloginMes, err := json.Marshal(loginMes)
	if err != nil {
		fmt.Println("消息序列化失败")
		return
	}
	mess.Type = message.LoginMesType //放入消息类型
	mess.Data = string(serloginMes)  //放入序列化后的消息
	//消息序列化
	data, err := json.Marshal(mess) //得到要发送的数据
	if err != nil {
		fmt.Println("消息体序列化失败")
		return
	}
	//先发送长度  再发送数据
	length := len(data) //将长度转化成一个可以表示长度的切片 数字->字节序列
	//将int转为 切片类型
	var pklen uint32
	pklen = uint32(length)
	var buf [4]byte                             //数组
	binary.BigEndian.PutUint32(buf[0:4], pklen) //大端   包.结构体.方法
	//fmt.Println(buf[0:4])    [0 0 0 81]
	//发送长度
	//fmt.Print(buf[0:4])  //类似于	fmt.Print(buf[0:4])
	n, err := conn.Write(buf[0:4])
	if err != nil || n != 4 {
		fmt.Println("消息长度发生失败")
		return
	}
	//发送数据内容
	n, err = conn.Write(data)
	if err != nil {
		fmt.Println("消息内容发生失败")
		return
	}
	//处理服务器返回消息
	transfer := &utils.Transfer{
		Conn: conn,
	}
	mess, err = transfer.ReadPkg()
	//mess.Data 反序列化
	var resMsg message.LoginResMes
	json.Unmarshal([]byte(mess.Data), &resMsg)
	if resMsg.Code == 200 {
		fmt.Println("登录成功")
		//初始化currtUser保存当前客户端信息和客户端连接
		currentUser.Conn = conn
		currentUser.UserId = userId
		currentUser.UserStatus = message.UserOnline
		//打印当前在线用户id
		for _, v := range resMsg.Usersid {
			if v == userId {
				continue
			}
			//初始化客户端在线用户map
			user := &message.User{
				UserId:     v,
				UserStatus: message.UserOnline, //状态默认上线
			}
			onlineUsers[v] = user
		}
		fmt.Println("\n\n")

		//这里还需要启动一个协程 保持和客户端通信 如果服务器有数据推送，那么就获取数据
		go ServerProcessMes(conn)
		//进入成功界面
		for {
			ShowMenu()
		}
	} else {
		fmt.Println(resMsg.Error)
	}
	return
}

//客户端--注册
func (this *UserProcess) Register(userId int, userPwd string, userName string) (err error) {
	//连接服务器
	conn, err := net.Dial("tcp", "127.0.0.1:8899")
	if err != nil {
		fmt.Println("客户端主动连接服务器端失败")
		return
	} else {
		fmt.Println("客户端主动连接服务器套接字成功")
	}
	defer conn.Close()
	var mess message.Message //总的消息
	var registerMes message.RegisterMes
	registerMes.User.UserId = userId
	registerMes.User.UserPwd = userPwd
	registerMes.User.UserName = userName
	serloginMes, err := json.Marshal(registerMes)
	if err != nil {
		fmt.Println("json.Marshal err=", err)
		return
	}
	mess.Type = message.RegisterResMesType
	mess.Data = string(serloginMes)
	data, err := json.Marshal(mess)
	if err != nil {
		fmt.Println("mess json.Marshal err=", err)
		return
	}
	//传到服务器
	transfer := utils.Transfer{
		Conn: conn,
	}
	err = transfer.WritePkg(data)
	if err != nil {
		fmt.Println("注册写包失败")
		return
	}
	//接受服务器注册返回消息
	mes, err := transfer.ReadPkg() //mes.type 就是RegisterResMes 类型
	if err != nil {
		fmt.Println("transfer.ReadPkg() err = ", err)
		return
	}
	fmt.Println(mes)
	var registerResMes message.RegisterResMes
	json.Unmarshal([]byte(mes.Data), &registerResMes)
	if registerResMes.Code == 200 {
		//注册成功
		fmt.Println("注册成功,请重新登录")
	} else {
		fmt.Println("注册失败 err= ", err)
	}
	os.Exit(0)
	return
}
