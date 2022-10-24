package entity

import (
	"gorm.io/gorm"
)

type Admin struct {
	gorm.Model
	Name     string
	Email    string `gorm:"uniqueIndex"`
	Password string `json:"-"`
	Donators []Donator `grom:"foreignKey:AdminID"`
}

type TypeFund struct {
	gorm.Model
	TypeFund string
	Donators []Donator `grom:"foreignKey:TypeFundID"`
}

type Organization struct {
	gorm.Model
	Organization string
	Donators []Donator `grom:"foreignKey:OrganizationID"`
}

type Donator struct {
	gorm.Model

	UserName string
	DateTime string
	UserInfo string
	UserNote string
	Amount   int
	NameFund string

	AdminID *uint
	Admin   Admin
	TypeFundID *uint
	TypeFund   TypeFund
	OrganizationID *uint
	Organization   Organization
}
