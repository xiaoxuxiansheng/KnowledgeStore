package entity

// 账号类型
type AccountType int

const (
	// 账号类型——用户
	AccountTypeCustomer AccountType = 1
	// 账号类型——商家
	AccountTypeMerchant AccountType = 2
)

/**
 * account 充血 entity
 */
type AccountEntity struct {
	// 财富点数
	Points int
	AccountVo
}

/**
 * account valueObject
 */
type AccountVo struct {
	// 账号类型——用户/商家
	Type AccountType
	// 账号
	AccountId string
	// 密码
	Passwd string
}

type TransactionEntity struct {
	// 交易 id
	TxId string
	TransactionVo
}

type TransactionVo struct {
	// 点数发出方
	PayerAccountId string
	// 点数接收方
	PayeeAccountId string
	// 交易点数
	Points int64
}
