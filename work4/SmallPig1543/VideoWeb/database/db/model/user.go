package model

import (
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID        uint `gorm:"primarykey"`
	CreateAt  string
	UserName  string `gorm:"unique"`
	PassWord  string
	AvatarURL string //头像URL
}

func (user *User) SetPassword(password string) error {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 12)
	if err != nil {
		return err
	}
	user.PassWord = string(bytes)
	return nil
}

func (user *User) CheckPassword(password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(user.PassWord), []byte(password))
	return err == nil
}
