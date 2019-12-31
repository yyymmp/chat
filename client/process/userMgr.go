package process

import (
	"fmt"
	"gocode/chat/client/model"
	"gocode/chat/common/message"
)

//客户端维护的在线用户列表map 全局变量

/**
上线一个加上一个  但初始化工作应该在哪？
在客户端用户登录成功后，会获取服务器在线服务列表，所以可以在这个时候初始化这个map
*/
var onlineUsers map[int]*message.User = make(map[int]*message.User, 10)
var currentUser model.CurrtUser

//编写一个函数  处理返回的信息--通知信息类型   参数 服务器返回消息
func updateUserStatus(notifyUsersStatusMes *message.NotifyUsersStatusMes) {
	//可以先判断是否有这个用户
	user, ok := onlineUsers[notifyUsersStatusMes.UserId]
	if !ok {
		//不存在此用户
		user = &message.User{
			UserId: notifyUsersStatusMes.UserId,
		}
	}
	//如果存在 直接更新状态即可
	user.UserStatus = notifyUsersStatusMes.Status
	onlineUsers[notifyUsersStatusMes.UserId] = user
	outputOnlineUser()
}

//展示在线用户
func outputOnlineUser() {
	fmt.Println("当前用户在线列表：")
	for index, _ := range onlineUsers {
		fmt.Println("用户id:\t", index)
	}
}
