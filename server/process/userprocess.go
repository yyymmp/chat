package process2

import (
	"encoding/json"
	"fmt"
	"gocode/chat/client/utils"
	"gocode/chat/common/message"
	"gocode/chat/server/model"
	utils2 "gocode/chat/server/utils"
	"net"
)

//处理用户相关
type UserProcess struct {
	Conn   net.Conn
	UserId int //表示这个结构体当前是哪个用户的连接

}

//处理服务器登录返回消息
func (this *UserProcess) ServerProcessLogin(mess *message.Message) (err error) {
	//将消息内容反序列化
	var loginMes message.LoginMes
	err = json.Unmarshal([]byte(mess.Data), &loginMes)
	if err != nil {
		fmt.Println("mess.Data 序列化失败")
		return
	}
	fmt.Println(loginMes)
	//判断登录
	//组织登录返回消息类型
	var loginMessRes message.LoginResMes
	//redis 数据库完成验证
	fmt.Println("传入验证的id和密码:", loginMes.UserId, loginMes.UserPwd)
	user, err := model.MyUserDao.Login(loginMes.UserId, loginMes.UserPwd)
	fmt.Println(user)
	if err != nil {
		if err == model.ERROR_USER_NOTEXISTS {
			loginMessRes.Code = 500
			loginMessRes.Error = err.Error()
		} else if err == model.ERROR_USER_PWD {
			loginMessRes.Code = 403
			loginMessRes.Error = err.Error()
		} else {
			loginMessRes.Code = 505
			loginMessRes.Error = "服务器内部服务错误"
		}
	} else {
		loginMessRes.Code = 200
		//登录成功后 将用户加入在线用户的map
		//放入
		this.UserId = loginMes.UserId
		userMgr.AddOnlineUsers(this) //this 即当前的用户处理结构体 里面含有当前用户的连接
		//通知其他在线用户 我上线了
		this.NotifyOthersOnlineUser(loginMes.UserId)
		//返回信息添加当前所有在线用户的id切片
		for index, _ := range userMgr.OnlineUsers {
			loginMessRes.Usersid = append(loginMessRes.Usersid, index) //返回时将用户id也新增
		}
		fmt.Println(loginMessRes.Usersid)
		fmt.Println(user.UserId, "号用户登录成功")
	}
	var resMes message.Message
	resMes.Type = message.LoginResMesType
	r, err := json.Marshal(loginMessRes)
	if err != nil {
		fmt.Println("json.Marshal fail")
		return
	}
	resMes.Data = string(r)
	//将返回信息序列化
	data, err := json.Marshal(resMes)
	if err != nil {
		fmt.Println("返回消息json.Marshal fail")
		return
	}
	//发消息   先发送数据长度  再发送消息内容 与读包一个 封装在一个函数中
	transfer := &utils2.Transfer{
		Conn: this.Conn,
	}
	err = transfer.WritePkg(data)
	return
}

//处理注册
func (this *UserProcess) ServerProcessRegister(mess *message.Message) (err error) {
	var Register message.RegisterMes
	err = json.Unmarshal([]byte(mess.Data), &Register)
	if err != nil {
		fmt.Println("mess.Data 序列化失败")
		return
	}
	var messRes message.Message               //总的回复消息
	messRes.Type = message.RegisterResMesType //消息类型
	var registerResMes message.RegisterResMes //响应信息
	user := Register.User
	err = model.MyUserDao.Register(&user)
	if err != nil {
		//注册成功
		registerResMes.Code = 505
		registerResMes.Error = err.Error()
		fmt.Println("Register err=", err)
	} else {
		//注册失败
		registerResMes.Code = 200
	}
	r, err := json.Marshal(registerResMes)
	if err != nil {
		fmt.Println("json.Marshal(registerResMes) err ", err)
		return
	}
	messRes.Data = string(r)
	data, err := json.Marshal(messRes) //总消息序列化
	//发消息   先发送数据长度  再发送消息内容 与读包一个 封装在一个函数中
	transfer := &utils2.Transfer{
		Conn: this.Conn,
	}
	err = transfer.WritePkg(data)
	return
}

//通知其他人
/**
userId  刚刚上线的用户  通知给其他在线用户
*/
func (this *UserProcess) NotifyOthersOnlineUser(userId int) {
	//遍历在线用户列表
	fmt.Println("即将遍历")
	for id, up := range userMgr.OnlineUsers {
		//不需要通知自己
		if id == userId {
			continue
		}
		//通知其他用户  传入要 发送用户的id  遍历即广播
		fmt.Println("开始通知")
		up.NotifyMeOnlineUser(userId) //up:已在线用户的连接  传入刚刚上线的人的id
	}
}
func (this *UserProcess) NotifyMeOnlineUser(userId int) {
	//发送通知
	fmt.Println("准备发送通知")
	var mess message.Message
	mess.Type = message.NotifyUsersStatusMesType
	var notifyUsersStatusMes message.NotifyUsersStatusMes
	notifyUsersStatusMes.UserId = userId //谁上线了？
	notifyUsersStatusMes.Status = message.UserOnline
	data, err := json.Marshal(notifyUsersStatusMes)
	if err != nil {
		fmt.Println("json.Marshal(notifyUsersStatusMes) err= ", err)
		return
	}
	mess.Data = string(data)
	//将总数据序列化发送
	data, err = json.Marshal(mess)
	if err != nil {
		fmt.Println("json.Marshal(mess) err= ", err)
		return
	}
	transfer := &utils.Transfer{
		Conn: this.Conn,
	}
	err = transfer.WritePkg(data)
	if err != nil {
		fmt.Println("通知广播出错", err)
		return
	}
}
