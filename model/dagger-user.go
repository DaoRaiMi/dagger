package model

import (
	"time"

	"github.com/daoraimi/dagger/box/orm"
	"github.com/pkg/errors"
)

type DaggerUser struct {
	Base
	Username      string     `gorm:"column:username"`
	Password      string     `gorm:"column:password"`
	RoleId        uint64     `gorm:"column:role_id"`
	Nickname      string     `gorm:"column:nickname"`
	Phone         string     `gorm:"column:phone"`
	Email         string     `gorm:"column:email"`
	Status        uint32     `gorm:"column:status"`
	LastLoginTime *time.Time `gorm:"column:last_login_time"`
}

func (DaggerUser) TableName() string {
	return "dagger_user"
}

func (DaggerUser) FindByUserID(ID uint64) (*DaggerUser, error) {
	var user DaggerUser
	if err := orm.R().Where("id = ?", ID).
		First(&user).Error; err != nil {
		if orm.IsRecordNotFoundError(err) {
			return nil, nil
		}
		return nil, errors.WithStack(err)
	}
	return &user, nil
}

func (DaggerUser) FindByUsername(username string) (*DaggerUser, error) {
	var row DaggerUser
	if err := orm.R().Where("`username` = ?", username).First(&row).Error; err != nil {
		if orm.IsRecordNotFoundError(err) {
			return nil, nil
		}
		return nil, errors.WithStack(err)
	}
	return &row, nil
}

func (DaggerUser) FindByPhone(phone string) (*DaggerUser, error) {
	var row DaggerUser
	if err := orm.R().Where("`phone` = ?", phone).First(&row).Error; err != nil {
		if orm.IsRecordNotFoundError(err) {
			return nil, nil
		}
		return nil, errors.WithStack(err)
	}
	return &row, nil
}

func (DaggerUser) FindByEmail(email string) (*DaggerUser, error) {
	var row DaggerUser
	if err := orm.R().Where("`email` = ?", email).First(&row).Error; err != nil {
		if orm.IsRecordNotFoundError(err) {
			return nil, nil
		}
		return nil, errors.WithStack(err)
	}
	return &row, nil
}

func (DaggerUser) TouchLoginTime(userID uint64) error {
	err := orm.R().Model(DaggerUser{}).Where("`id` = ?", userID).
		Update(map[string]interface{}{"last_login_time": time.Now()}).
		Error
	if err != nil {
		return errors.WithStack(err)
	}
	return nil
}

type DaggerUserList []*DaggerUser

func (l DaggerUserList) IDSetList() []uint64 {
	var IDList []uint64
	for _, user := range l {
		IDList = append(IDList, user.ID)
	}
	return IDList
}
