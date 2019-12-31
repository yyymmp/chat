package process2

import (
	"encoding/json"
	"fmt"
	"gocode/chat/common/message"
	"gocode/chat/server/utils"
	"net"
)

//处理消息相关

//消息转发
type SmsProcess struct {
}

func (this *SmsProcess) SendGroupMes(mess *message.Message) {
	//群发消息
	var data message.SmsMes
	err := json.Unmarshal([]byte(mess.Data), &data) //在这里给mess反序列化了
	if err != nil {
		fmt.Println("json.Unmarshal err=", err)
		return
	}
	d, err := json.Marshal(mess)
	for _, up := range userMgr.OnlineUsers {
		//给在线的每个用户群发
		if id == data.UserId { //不群发给自己
			continue
		}

		this.sendOtherOnline(d, up.Conn)
	}
}
func (this *SmsProcess) sendOtherOnline(data []byte, conn net.Conn) {
	tf := &utils.Transfer{
		Conn: conn,
	}
	err := tf.WritePkg(data)
	if err != nil {
		fmt.Println("转发消息失败：", err)
	}
}
