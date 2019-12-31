package utils

import (
	"encoding/binary"
	"encoding/json"
	"fmt"
	"gocode/chat/common/message"
	"net"
)

type Transfer struct {
	Conn net.Conn
	Buff [8096]byte
}

//读取数据 先读消息长度 再读消息内容
func (this *Transfer) ReadPkg() (mess message.Message, err error) {
	//读取数据长度
	_, err = this.Conn.Read(this.Buff[:4]) //从conn套接字中读取 放入buff中   read:当客户段连接关闭时,此方法将不会再阻塞
	if err != nil {
		fmt.Println("read fail")
		return
	}
	var pklen uint32
	//字符切片转化为长度
	pklen = binary.BigEndian.Uint32(this.Buff[0:4])
	//继续读取消息内容  从conn中读取消息内容放入buff
	n, err := this.Conn.Read(this.Buff[:pklen])
	if n != int(pklen) || err != nil {
		fmt.Println("conn read fail,可能出现丢包现象")
		return
	}
	//反序列化得到数据
	err = json.Unmarshal(this.Buff[:pklen], &mess) //这里一定要加上&
	if err != nil {
		fmt.Println("消息内容反序列化失败")
		return
	}
	return
}

//给对方发送数据   先发送长度 再发送消息内容
func (this *Transfer) WritePkg(data []byte) (err error) {
	var pklen uint32
	length := len(data)
	pklen = uint32(length)
	binary.BigEndian.PutUint32(this.Buff[0:4], pklen)
	//写回数据长度
	n, err := this.Conn.Write(this.Buff[0:4])
	if err != nil {
		fmt.Println("数据长度发送失败")
		return
	}
	//写回数据内容
	n, err = this.Conn.Write(data)
	if n != int(pklen) || err != nil {
		fmt.Println("消息内容返回失败")
		return
	}
	return
}
