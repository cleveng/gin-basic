/**
 * @Author: cleveng | cleveng@gmail.com
 * @Description:
 * @File:  User
 * @Version: 1.0.0
 * @Date: 2020-09-25 10:11
 */

package models

import (
	"encoding/gob"
	"gorm.io/gorm"
	uuid "github.com/satori/go.uuid"
	"time"
)

type User struct {
	Id        uint       `gorm:"primary_key" json:"-"`
	CreatedAt time.Time  `json:"-"`
	UpdatedAt time.Time  `json:"-"`
	DeletedAt *time.Time `json:"-"`

	// 邮箱修改时间
	// 手机修改时间
	TelModifiedAt   *time.Time `json:"tel_modified_at"`
	EmailModifiedAt *time.Time `json:"email_modified_at"`

	Uuid            string `json:"-"`
	Name            string `json:"-"`
	DisplayName     string `json:"display_name"`
	Email           string `json:"email"`
	EmailIsVerified bool   `json:"email_is_verified"` // 有是否验证
	Tel             string `json:"tel"`
	Avatar          string `json:"avatar"`
	Password        string `json:"-"`
	Status          bool   `json:"status"`

	Token string `json:"token" gorm:"-"` // 2020/08/15

}

func init() {
	gob.Register(User{})
}

func (user *User) TableName() string {
	return "test_users"
}

func (user *User) BeforeCreate(scope *gorm.DB) error {
	scope.Statement.SetColumn("created_at", time.Now())
	scope.Statement.SetColumn("updated_at", time.Now())
	url := uuid.Must(uuid.NewV4(), nil).String()
	scope.Statement.SetColumn("uuid", url)
	return nil
}

func (user *User) BeforeUpdate(scope *gorm.DB) error {
	scope.Statement.SetColumn("updated_at", time.Now())
	return nil
}
