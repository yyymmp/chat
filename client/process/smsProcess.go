package process

import (
	"encoding/json"
	"fmt"
	"gocode/chat/client/utils"
	"gocode/chat/common/message"
)

type SmsProcess struct {
}

//客户端群发消息
func (this *SmsProcess) SendGroupMes(content string) (err error) {

	var mess message.Message
	mess.Type = message.SmsMesType
	var smsMes message.SmsMes
	smsMes.Content = content
	smsMes.UserId = currentUser.UserId
	smsMes.UserStatus = currentUser.UserStatus
	data, err := json.Marshal(smsMes)
	if err != nil {
		fmt.Println("json.Marshal(smsMes) err=", err)
		return
	}
	mess.Data = string(data)
	data, err = json.Marshal(mess)
	if err != nil {
		fmt.Println("json.Marshal(smsMes) err=", err)
		return
	}
	//发送
	transfer := &utils.Transfer{
		Conn: currentUser.Conn,
	}
	err = transfer.WritePkg(data)
	if err != nil {
		fmt.Println("transfer.WritePkg(data) err=", err)
		return
	}
	fmt.Println("客户端消息发送成功")
	return
}
