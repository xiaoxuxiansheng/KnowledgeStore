package user

import "gorm.io/gorm"

type UserPo struct {
	gorm.Model
	UserId string `json:"user_id" gorm:"column:user_id;NOT NULL"`
	Passwd string `json:"passwd" gorm:"column:passwd;NOT NULL"`
}
