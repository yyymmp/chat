package process2

import "errors"

//userMgr全局唯一 对在线用户进行操作
var (
	userMgr *UserMgr
)

type UserMgr struct {
	OnlineUsers map[int]*UserProcess
}

func init() {
	userMgr = &UserMgr{
		OnlineUsers: make(map[int]*UserProcess, 1024),
	}
}

//对此map的增删改查  增加/更新
func (this *UserMgr) AddOnlineUsers(userProcess *UserProcess) {
	this.OnlineUsers[userProcess.UserId] = userProcess
}

//删除
func (this *UserMgr) DelOnlineUsers(userId int) {
	delete(this.OnlineUsers, userId)
}

//查询 返回所有在线用户
func (this *UserMgr) GetAllOnlineUsers() map[int]*UserProcess {
	return this.OnlineUsers
}

//查询某个用户 通过id
func (this *UserMgr) GetAllOnlineById(userid int) (up *UserProcess, err error) {
	up, ok := this.OnlineUsers[userid] //返回的是UserProcess结构体指针
	if !ok {
		err = errors.New("用户id不在线" + string(userid))
		return
	}
	return
}
