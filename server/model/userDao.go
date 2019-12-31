package model

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/garyburd/redigo/redis"
	"gocode/chat/common/message"
)

var (
	MyUserDao *UserDao
)

//使用工厂模式 创建UserDao实例  必然是传入连接  返回dao实例  提供一个函数以供外部获取 注意是函数 不是方法
func NewUserDao(pool *redis.Pool) (userDao *UserDao) {
	return &UserDao{pool: pool}
}

//提供curd方法
type UserDao struct {
	pool *redis.Pool
}

//通过id获取用户
func (this *UserDao) getUserById(conn redis.Conn, id int) (user *User, err error) {
	res, err := redis.String(conn.Do("hget", "users", id))
	fmt.Println("redis取出的数据", res)
	if err != nil {
		if err == redis.ErrNil {
			//不存在此用户
			err = ERROR_USER_NOTEXISTS
		}
		return
	}
	user = &User{}
	//用户存在  反序列化user实例
	err = json.Unmarshal([]byte(res), user)
	user.UserId = id
	fmt.Println("redis反序列化后的数据", user)
	if err != nil {
		fmt.Println("json.Unmarshal err=", err)
		return
	}
	return
}

//对登录验证
func (this *UserDao) Login(userId int, userPwd string) (user *User, err error) {
	conn := this.pool.Get()
	defer conn.Close()
	user, err = this.getUserById(conn, userId)
	if err != nil {
		return
	}
	//检验密码
	if user.UserPwd != userPwd {
		err = ERROR_USER_PWD
		return
	}
	return
}
func (this *UserDao) Register(user *message.User) (err error) {
	conn := this.pool.Get()
	defer conn.Close()
	_, err = this.getUserById(conn, user.UserId)
	if err == nil {
		err = errors.New("用户id已存在")
		return
	}
	//否则将用户存入  系列化数据
	data, err := json.Marshal(user)
	if err != nil {
		return
	}
	_, err = conn.Do("hset", "users", user.UserId, string(data))
	if err != nil {
		fmt.Println("redis 注册用户失败")
		return
	}
	return
}
