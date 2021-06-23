package model

const ModelUserFile = `package models

type User struct {
	ID   uint   ` + "`" + `gorm:"column:id;primarykey" json:"-"` + "`" + `
	UUID string ` + "`" + `gorm:"column:uuid;size:37;uniqueIndex" json:"uuid"` + "`" + `
	Name string ` + "`" + `gorm:"column:name;size:50;not null;index" json:"name"` + "`" + `
}

func (User) TableName() string {
	return "user"
}

var DevelopUser = &User{UUID: "c66f19b6-569d-4c85-94ae-5f5b45de18f0", Name: "develop"}
`
