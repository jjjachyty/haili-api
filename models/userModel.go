package models

import "time"
import "haili/common"
import "fmt"

type UserModel struct{}
type User struct {
	ID       uint      `gorm:"column:ID"`
	Name     string    `gorm:"column:Name"`
	PassWord string    `gorm:"column:PassWord"`
	Type     string    `gorm:"column:Type"`
	Status   string    `gorm:"column:Status"`
	ReMark   string    `gorm:"column:ReMark"`
	CreateAt time.Time `gorm:"column:CreateAt"`
	UpdateAt time.Time `gorm:"column:UpdateAt"`
}

func (UserModel) Find(params map[string]interface{}) []User {
	var users []User
	fmt.Println(params)
	common.DB.Where(params).Find(&users)
	return users
}
func (User) TableName() string {
	return "users"
}
