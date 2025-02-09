package entity

type OrderEntity struct {
	OrderId   string
	UserId    string
	StoreName string
	GoodsName string
	Points    int
}

type OrderVo struct {
	UserId    string
	StoreName string
	GoodsName string
	Points    int
}
