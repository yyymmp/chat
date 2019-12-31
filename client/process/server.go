package process

import (
	"encoding/json"
	"fmt"
	"gocode/chat/client/utils"
	"gocode/chat/common/message"
	"net"
	"os"
)

/**
1 显示登录成功页面
2 保持和服务器联系
*/
func ShowMenu() {
	fmt.Println("------------------恭喜xxx登录成功--------------")
	fmt.Println("\t\t\t1 显示在线用户列表")
	fmt.Println("\t\t\t2 发送消息")
	fmt.Println("\t\t\t3 信息列表")
	fmt.Println("\t\t\t4 退出系统")
	var key int
	var content string
	fmt.Scanf("%d", &key)
	fmt.Println(key)
	switch key {
	case 1:
		outputOnlineUser()
	case 2:
		fmt.Println("请输入要群发的消息")
		fmt.Scanf("%s", &content)
		sms := &SmsProcess{}
		sms.SendGroupMes(content)
	case 3:
		fmt.Println("信息列表")
	case 4:
		fmt.Println("退出系统")
		os.Exit(0)
	default:
		fmt.Println("输入不正确")
	}
}

//不断服务器连接  保持与服务器通信
func ServerProcessMes(conn net.Conn) {
	transfer := &utils.Transfer{
		Conn: conn,
	}
	for {
		fmt.Println("客户端正在读取服务器推送的消息---")
		msg, err := transfer.ReadPkg()
		if err != nil {
			fmt.Println("通信协程服务器通信失败", err)
			return
		}
		//读到消息
		//1 判断接受的数据类型
		switch msg.Type {
		case message.NotifyUsersStatusMesType:
			//处理用户上线消息通知
			//1 取出用户id  和 状态
			var notifyUsersStatusMes message.NotifyUsersStatusMes
			json.Unmarshal([]byte(msg.Data), &notifyUsersStatusMes)
			//2 通知其他用户上线
			updateUserStatus(&notifyUsersStatusMes)
		case message.SmsMesType:
			//得到转发消息
			output(&msg)
		default:
			fmt.Println("暂时不识别该消息类型")
		}
		fmt.Println(msg)

	}

}
