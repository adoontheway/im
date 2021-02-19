package service

import (
	"errors"
	"fmt"
	"im/util"
	"math/rand"
	"time"

	"im/model"
)

type UserService struct {
}

// 注册函数
func (s *UserService) Register(mobile, plainpwd, nickname, avatar, sex string) (user model.User, err error) {
	// Check mobile no is exists
	tmp := model.User{}
	_, err = DbEngin.Where("mobile=?", mobile).Get(&tmp)
	if err != nil {
		return tmp, err
	}

	// if yes return already exists
	if tmp.Id > 0 {
		return tmp, errors.New("手机号已经注册")
	}
	//else do register process
	tmp.Mobile = mobile
	tmp.Avatar = avatar
	tmp.Nickname = nickname
	tmp.Salt = fmt.Sprintf("%6d", rand.Int31n(10000))
	tmp.Sex = sex
	tmp.Passwd = util.MakePasswd(plainpwd, tmp.Salt)
	tmp.Createat = time.Now()
	// 可以是一个随机数
	tmp.Token = fmt.Sprintf("%08d", rand.Int31n(10000000))
	// InsertOne
	_, err = DbEngin.InsertOne(&tmp)
	// 前端恶意插入特殊字符
	// fmt.Println("Affect rows:", rows)
	// if rows == 0 {
	// 	return tmp, errors.New("插入失败")
	// }
	// 数据库操作失败
	if err != nil {
		return tmp, err
	}
	return tmp, err
}

// 登录函数
func (s *UserService) Login(mobile, plainpwd string) (user model.User, err error) {
	tmp := model.User{}
	// Find user by mobile
	DbEngin.Where("mobile=?", mobile).Get(&tmp)
	if tmp.Id == 0 {
		return tmp, errors.New("用户不存在")
	}
	// Compare password
	if !util.ValidatePasswd(plainpwd, tmp.Salt, tmp.Passwd) {
		return tmp, errors.New("密码不正确")
	}
	// Refresh token
	str := fmt.Sprintf("%d", time.Now().Unix())
	tmp.Token = util.MD5Encode(str)
	DbEngin.ID(tmp.Id).Cols("token").Update(&tmp)
	return tmp, nil
}

// 查询某个用户
func (s *UserService) Find(userId int64) (user model.User) {
	tmp := model.User{}
	DbEngin.ID(userId).Get(&tmp)
	return tmp
}
