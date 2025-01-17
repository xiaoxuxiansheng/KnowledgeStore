package entity

type OrderEntity struct {
	OrderId    string
	UserId     string
	MerchantId string
	GoodsIds   []string
	Points     int64
}

type OrderVo struct {
	UserId     string
	MerchantId string
	GoodsIds   []string
	Points     int64
}
