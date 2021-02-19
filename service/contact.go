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

	if tmp.Id != 0 {
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

}

// 好友列表
func (s *ContactService) SearchFriend(userid int64) []model.User {
	contacts := make([]model.Contact, 0)
	objIds := make([]int64, 0)
	DbEngin.Where("ownerid = ? and cate = ?", userid, model.CONCAT_CATE_USER).Find(&contacts)
	for _, v := range contacts {
		objIds = append(objIds, v.Dstobj)
	}
	coms := make([]model.User, 0)
	if len(objIds) == 0 {
		return coms
	}
	DbEngin.In("id", objIds).Find(&coms)
	return coms
}

// 查找工会
func (s *ContactService) SearchCommunity(userid int64) []model.Community {
	contacts := make([]model.Contact, 0)
	objIds := make([]int64, 0)
	DbEngin.Where("ownerid = ? and cate = ?", userid, model.CONCAT_CATE_COMUNITY).Find(&contacts)
	for _, v := range contacts {
		objIds = append(objIds, v.Dstobj)
	}
	coms := make([]model.Community, 0)
	if len(objIds) == 0 {
		return coms
	}
	DbEngin.In("id", objIds).Find(&coms)
	return coms
}

// 加入工会
func (s *ContactService) JoinCommunity(userid, comId int64) error {
	cot := model.Contact{
		Ownerid: userid,
		Dstobj:  comId,
		Cate:    model.CONCAT_CATE_COMUNITY,
	}
	DbEngin.Get(&cot)
	if cot.Id == 0 {
		cot.Createat = time.Now()
		_, err := DbEngin.InsertOne(cot)
		return err
	} else {
		return nil
	}
}

// 创建公会
func (s *ContactService) CreateCommunity(comm model.Community) (ret model.Community, err error) {
	if len(comm.Name) == 0 {
		err = errors.New("缺少群名称")
		return ret, err
	}
	if comm.Ownerid == 0 {
		err = errors.New("请先登录")
		return ret, err
	}
	com := model.Community{
		Ownerid: comm.Ownerid,
	}
	num, err := DbEngin.Count(&com)
	if num > 5 {
		err = errors.New("一个用户最多只能创建5个群")
		return com, err
	} else {
		comm.Createat = time.Now()
		session := DbEngin.NewSession()
		session.Begin()
		_, err = session.InsertOne(&comm)
		if err != nil {
			session.Rollback()
			return com, err
		}
		_, err = session.InsertOne(
			model.Contact{
				Ownerid:  comm.Ownerid,
				Dstobj:   comm.Id,
				Cate:     model.CONCAT_CATE_COMUNITY,
				Createat: time.Now(),
			})
		if err != nil {
			session.Rollback()
		} else {
			session.Commit()
		}
		return com, err
	}

}
