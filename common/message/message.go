package message

//定义消息类型常量
const (
	LoginMesType             = "LoginMes"
	LoginResMesType          = "LoginResMes"
	RegisterMesType          = "RegisterMes"
	RegisterResMesType       = "RegisterResMes"
	NotifyUsersStatusMesType = "NotifyUsersStatusMes"
	SmsMesType               = "SmsMes"
)

//定义用户状态常量
const (
	UserOnline = iota //在线
)

//通信消息
type Message struct {
	Type string `json:"type"` //消息类型  自定义常量
	Data string `json:"data"` //消息数据
}

//定义登录消息
type LoginMes struct {
	UserId   int    `json:"useId"`
	UserPwd  string `json:"userPwd"`
	UserName string `json:"userName"`
}

//定义服务登录返回消息
type LoginResMes struct {
	Code  int    //状态码  200成功  500 失败
	Error string `json:"error"` //返回错误信息
	//添加在线用户id切片
	Usersid []int
}

//定义注册消息类型
type RegisterMes struct {
	User User `json:"user"`
}

//定义注册返回消息类型
type RegisterResMes struct {
	Code  int    //状态码  200 成功   400 用户已占用
	Error string `json:"error"` //返回错误信息
}

//定义推送新用户上线消息类型
type NotifyUsersStatusMes struct {
	UserId int `json:"userId"` //上线的用户id
	Status int `json:"status"` //用户状态  在线/忙碌
}

//群发消息
type SmsMes struct {
	User    `json:"user"` //发送消息的人  直接是User结构体 匿名结构体
	Content string        `json:"content"` //消息内容
}
