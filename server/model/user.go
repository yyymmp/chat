package model

//用户结构体
type User struct {
	//为了序列化成功，必须保证json标签和redis的 key保持一致
	UserId   int    `json:"userId"`
	UserPwd  string `json:"userPwd"`
	UserName string `json:"userName"`
}
