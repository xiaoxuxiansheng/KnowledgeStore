package toc

import "github.com/xiaoxuxiansheng/KnowledgeStore/application/toc/dto"

const (
	ErrnoInternel = -1
	ErrnoSuccess  = 0
)

/**
 * @brief: 注册用户账号请求参数
 */
type RegisterCustomerReqVo struct {
	// 用户账号
	AccountId string `json:"account_id" binding:"required"`
	// 商家密码
	Password string `json:"password" binding:"required"`
}

func (r *RegisterCustomerReqVo) ToDto() *dto.RegisterCutomerReqDto {
	return &dto.RegisterCutomerReqDto{
		AccountId: r.AccountId,
		Password:  r.Password,
	}
}

/**
 * @brief: 登录用户账号请求参数
 */
type LoginCustomerReqVo struct {
	// 用户账号
	AccountId string `json:"account_id" binding:"required"`
	// 商家密码
	Password string `json:"password" binding:"required"`
}

func (l *LoginCustomerReqVo) ToDto() *dto.LoginCutomerReqDto {
	return &dto.LoginCutomerReqDto{
		AccountId: l.AccountId,
		Password:  l.Password,
	}
}

/**
 * @brief: 店铺列表请求参数
 */
type ListStoreReqVo struct {
	// 商店名称
	StoreName string `form:"store_name"`
	// 页码
	Page int `form:"page"`
	// 大小
	Limit int `form:"limit"`
}

func (l *ListStoreReqVo) ToDto() *dto.ListStoresReqDto {
	page, limit := l.Page, l.Limit
	if page <= 0 {
		page = 1
	}
	if limit <= 10 {
		limit = 10
	}

	return &dto.ListStoresReqDto{
		StoreName: l.StoreName,
		Page:      page,
		Limit:     limit,
	}
}

/**
 * @brief: 店铺列表响应参数
 */
type ListStoreRespVo struct {
	Total int            `json:"total"`
	List  []*StoreDataVo `json:"list"`
}

type StoreDataVo struct {
	// 店铺名称
	StoreName string `json:"store_name"`
}

func NewListStoreRespVo(storesDto *dto.ListStoresRespDto) *ListStoreRespVo {
	return &ListStoreRespVo{
		Total: storesDto.Total,
		List:  NewStoreDatas(storesDto.List),
	}
}

func NewStoreDatas(datas []*dto.StoreDataDto) []*StoreDataVo {
	_data := make([]*StoreDataVo, 0, len(datas))
	for _, data := range datas {
		_data = append(_data, NewStoreData(data))
	}
	return _data
}

func NewStoreData(data *dto.StoreDataDto) *StoreDataVo {
	return &StoreDataVo{
		StoreName: data.StoreName,
	}
}

/**
 * @brief: 商品列表请求参数
 */
type ListGoodsReqVo struct {
	// 商店名称
	StoreName string `form:"store_name" binding:"required"`
	// 商品分类
	Category string `form:"category"`
	// 商品名称
	GoodsName string `form:"goods_name"`
	// 页码
	Page int `form:"page"`
	// 大小
	Limit int `form:"limit"`
}

func (l *ListGoodsReqVo) ToDto() *dto.ListGoodsReqDto {
	page, limit := l.Page, l.Limit
	if page <= 0 {
		page = 1
	}
	if limit <= 10 {
		limit = 10
	}

	return &dto.ListGoodsReqDto{
		StoreName: l.StoreName,
		Category:  l.Category,
		GoodsName: l.GoodsName,
		Page:      page,
		Limit:     limit,
	}
}

type ListGoodsRespVo struct {
	Total int            `json:"total"`
	List  []*GoodsDataVo `json:"data"`
}

type GoodsDataVo struct {
	GoodsName string `json:"goods_name"`
	Category  string `json:"category"`
}

func NewListGoodsRespVo(respDto *dto.ListGoodsRespDto) *ListGoodsRespVo {
	return &ListGoodsRespVo{
		Total: respDto.Total,
	}
}

func NewGoodsDatas(datas []*dto.GoodsDataDto) []*GoodsDataVo {
	_datas := make([]*GoodsDataVo, 0, len(datas))
	for _, data := range datas {
		_datas = append(_datas, NewGoodsData(data))
	}
	return _datas
}

func NewGoodsData(data *dto.GoodsDataDto) *GoodsDataVo {
	return &GoodsDataVo{
		GoodsName: data.GoodsName,
		Category:  data.Category,
	}
}

type QueryGoodsReqVo struct {
	AccountId string `form:"store_name" binding:"required"`
	StoreName string `form:"store_name" binding:"required"`
	GoodsName string `form:"goods_name" binding:"required"`
}

func (q *QueryGoodsReqVo) ToDto() *dto.QueryGoodsReqDto {
	return &dto.QueryGoodsReqDto{
		AccountId: q.AccountId,
		StoreName: q.StoreName,
		GoodsName: q.GoodsName,
	}
}

type QueryGoodsRespVo struct {
	Owned        bool   `json:"owned"`
	GoodsContent string `json:"goods_content"`
	Category     string `json:"category"`
	GoodsName    string `json:"goods_name"`
}

func NewQueryGoodsRespVo(resp *dto.QueryGoodsRespDto) *QueryGoodsRespVo {
	return &QueryGoodsRespVo{
		Owned:        resp.Owned,
		GoodsContent: resp.GoodsContent,
		Category:     resp.Category,
		GoodsName:    resp.GoodsName,
	}
}

type BuyGoodsReqVo struct {
	AccountId string `json:"account_id" binding:"required"`
	StoreName string `json:"store_name" binding:"required"`
	GoodsName string `json:"goods_name" binding:"required"`
}

func (b *BuyGoodsReqVo) ToDto() *dto.BuyGoodsReqDto {
	return &dto.BuyGoodsReqDto{
		AccountId: b.AccountId,
		StoreName: b.StoreName,
		GoodsName: b.GoodsName,
	}
}

type ListOrderReqVo struct {
	AccountId string `form:"account_id" binding:"required"`
	Status    string `form:"status"  binding:"required"`
}

func (l *ListOrderReqVo) ToDto() *dto.ListOrderReqDto {
	return &dto.ListOrderReqDto{
		AccountId: l.AccountId,
		Status:    l.Status,
	}
}

type ListOrderRespVo struct {
	Status string         `json:"status"`
	Total  int            `json:"total"`
	List   []*OrderDataVo `json:"list"`
}

type OrderDataVo struct {
	OrderId string `json:"order_id"`
}

func NewListOrderRespVo(listDto *dto.ListOrderRespDto) *ListOrderRespVo {
	return &ListOrderRespVo{
		Total: listDto.Total,
		List:  NewOrderDatas(listDto.List),
	}
}

func NewOrderDatas(datas []*dto.OrderDataDto) []*OrderDataVo {
	_data := make([]*OrderDataVo, 0, len(datas))
	for _, data := range datas {
		_data = append(_data, NewOrderData(data))
	}
	return _data
}

func NewOrderData(data *dto.OrderDataDto) *OrderDataVo {
	return &OrderDataVo{
		OrderId: data.OrderId,
	}
}

type PayOrderReqVo struct {
	AccountId string
	OrderId   string
}

func (p *PayOrderReqVo) ToDto() *dto.PayOrderReqDto {
	return &dto.PayOrderReqDto{
		AccountId: p.AccountId,
		OrderId:   p.OrderId,
	}
}
