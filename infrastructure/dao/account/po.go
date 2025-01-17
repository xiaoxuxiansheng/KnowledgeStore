package account

import "gorm.io/gorm"

type AccountPo struct {
	gorm.Model
	Type int `json:"type" gorm:"column:type;NOT NULL"`
	// 账号
	AccountId string `json:"account_id" gorm:"column:account_id;NOT NULL"`
	// 密码
	Passwd string `json:"passwd" gorm:"column:passwd;NOT NULL"`
	// 账号归属身份唯一 id
	UniqueId string `json:"unique_id" gorm:"column:account_id;NOT NULL"`
	// 财富点数
	Points int64 `json:"points" gorm:"column:points;NOT NULL"`
}
