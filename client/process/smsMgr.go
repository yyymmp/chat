package process

import (
	"encoding/json"
	"fmt"
	"gocode/chat/common/message"
)

//输出群发消息
func output(mess *message.Message) {
	var sms message.SmsMes
	err := json.Unmarshal([]byte(mess.Data), &sms)
	if err != nil {
		fmt.Println("群发客户端反序列化失败")
		return
	}
	//显示
	info, err := fmt.Scanf("用户id：%d说：%s", sms.UserId, sms.Content)
	fmt.Println(info)
	fmt.Println() //换行
}
