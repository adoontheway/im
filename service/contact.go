package service

import (
	"errors"
	"im/model"
	"time"
)

type ContactService struct {
}

// 登录函数
func (s *ContactService) AddFriend(userid, destid int64) error {
	if userid == destid {
		return errors.New("不能添加自己为好友")
	}
	tmp := model.Contact{}
	DbEngin.Where("ownerid = ?", userid).
		And("dstid = ?", destid).
		And("cat = ?", model.CONCAT_CATE_USER).
		Get(&tmp)

	if tmp.Id == 0 {
		return errors.New("改用户已经被添加过了")
	}
	session := DbEngin.NewSession()
	session.Begin()
	_, e2 := session.InsertOne(model.Contact{
		Ownerid:  userid,
		Dstobj:   destid,
		Cate:     model.CONCAT_CATE_USER,
		Createat: time.Now(),
	})
	_, e3 := session.InsertOne(model.Contact{
		Ownerid:  destid,
		Dstobj:   userid,
		Cate:     model.CONCAT_CATE_USER,
		Createat: time.Now(),
	})
	if e2 == nil && e3 == nil {
		session.Commit()
		return nil
	} else {
		session.Rollback()
		if e2 != nil {
			return e2
		} else {
			return e3
		}
	}

	return nil
}
