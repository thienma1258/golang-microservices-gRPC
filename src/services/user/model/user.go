package model

import (
	"github.com/google/uuid"
	"github.com/jinzhu/gorm"
)

// Use pointer value
type User struct {
	gorm.Model
	Name string
	userName string
	passWord string
	profilePicture string

}

func (user *User) BeforeCreate(scope *gorm.Scope) error {
	scope.SetColumn("ID", uuid.New())
	return nil
}