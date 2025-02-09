package dto

import (
	accountentity "github.com/xiaoxuxiansheng/KnowledgeStore/domain/account/entity"
	goodsentity "github.com/xiaoxuxiansheng/KnowledgeStore/domain/goods/entity"
	goodsnorth "github.com/xiaoxuxiansheng/KnowledgeStore/domain/goods/north"
	orderentity "github.com/xiaoxuxiansheng/KnowledgeStore/domain/order/entity"
	ordernorth "github.com/xiaoxuxiansheng/KnowledgeStore/domain/order/north"
	storeentity "github.com/xiaoxuxiansheng/KnowledgeStore/domain/store/entity"
	storenorth "github.com/xiaoxuxiansheng/KnowledgeStore/domain/store/north"
	userentity "github.com/xiaoxuxiansheng/KnowledgeStore/domain/user/entity"
)

/**
 * @brief: 注册账号请求参数
 */
type RegisterCutomerReqDto struct {
	// 用户账号
	AccountId string
	// 商家密码
	Password string
}

func (r *RegisterCutomerReqDto) ToUserEntity() *userentity.UserEntity {
	return &userentity.UserEntity{
		UserId: r.AccountId,
	}
}

func (r *RegisterCutomerReqDto) ToAccountEntity() *accountentity.AccountEntity {
	return &accountentity.AccountEntity{
		Points: 0,
		AccountVo: accountentity.AccountVo{
			Type:      accountentity.AccountTypeCustomer,
			AccountId: r.AccountId,
			Passwd:    r.Password,
		},
	}
}

/**
 * @brief: 登录请求参数
 */
type LoginCutomerReqDto struct {
	// 用户账号
	AccountId string
	// 商家密码
	Password string
}

func (l *LoginCutomerReqDto) ToAccountVo() *accountentity.AccountVo {
	return &accountentity.AccountVo{
		Type:      accountentity.AccountTypeCustomer,
		AccountId: l.AccountId,
		Passwd:    l.Password,
	}
}

/**
 * @brief: 检索店铺请求参数
 */
type ListStoresReqDto struct {
	StoreName string
	Page      int
	Limit     int
}

func (l *ListStoresReqDto) ToQueryOptions() []storenorth.QueryOption {
	opts := []storenorth.QueryOption{
		storenorth.WithPageLimit(l.Page, l.Limit),
	}
	if len(l.StoreName) > 0 {
		opts = append(opts, storenorth.WithStoreName(l.StoreName))
	}
	return opts
}

/**
 * @brief: 检索店铺响应参数
 */
type ListStoresRespDto struct {
	Total int
	List  []*StoreDataDto
}

func NewListStoresRespDto(total int, entities []*storeentity.StoreEntity) *ListStoresRespDto {
	return &ListStoresRespDto{
		Total: total,
		List:  NewStoreDataDtos(entities),
	}
}

func NewStoreDataDtos(entities []*storeentity.StoreEntity) []*StoreDataDto {
	_dtos := make([]*StoreDataDto, 0, len(entities))
	for _, entity := range entities {
		_dtos = append(_dtos, NewStoreDataDto(entity))
	}
	return _dtos
}

func NewStoreDataDto(entity *storeentity.StoreEntity) *StoreDataDto {
	return &StoreDataDto{
		StoreName: entity.StoreName,
	}
}

type StoreDataDto struct {
	StoreName string
}

/**
 * @brief: 检索商品请求参数
 */
type ListGoodsReqDto struct {
	// 商店名称
	StoreName string
	// 商品分类
	Category string
	// 商品名称
	GoodsName string
	// 页码
	Page int
	// 大小
	Limit int
}

func (l *ListGoodsReqDto) ToQueryOptions() []goodsnorth.QueryOption {
	opts := []goodsnorth.QueryOption{
		goodsnorth.WithPageLimit(l.Page, l.Limit),
		goodsnorth.WithStoreName(l.StoreName),
	}
	if len(l.Category) > 0 {
		opts = append(opts, goodsnorth.WithCategory(l.Category))
	}
	if len(l.GoodsName) > 0 {
		opts = append(opts, goodsnorth.WithGoodsName(l.GoodsName))
	}
	return opts
}

/**
 * @brief: 检索商品响应参数
 */
type ListGoodsRespDto struct {
	Total int
	List  []*GoodsDataDto
}

type GoodsDataDto struct {
	GoodsName string
	Category  string
}

func NewListGoodsRespDto(total int, entities []*goodsentity.GoodsEntity) *ListGoodsRespDto {
	return &ListGoodsRespDto{
		Total: total,
		List:  NewGoodsDataDtos(entities),
	}
}

func NewGoodsDataDtos(entities []*goodsentity.GoodsEntity) []*GoodsDataDto {
	_dtos := make([]*GoodsDataDto, 0, len(entities))
	for _, entity := range entities {
		_dtos = append(_dtos, NewGoodsDataDto(entity))
	}
	return _dtos
}

func NewGoodsDataDto(entity *goodsentity.GoodsEntity) *GoodsDataDto {
	return &GoodsDataDto{
		GoodsName: entity.GoodsName,
		Category:  entity.Category,
	}
}

/**
 * @brief: 查询商品详情请求参数
 */
type QueryGoodsReqDto struct {
	AccountId string
	StoreName string
	GoodsName string
}

func (q *QueryGoodsReqDto) ToQueryOptions() []goodsnorth.QueryOption {
	return []goodsnorth.QueryOption{
		goodsnorth.WithUserId(q.AccountId),
		goodsnorth.WithStoreName(q.StoreName),
		goodsnorth.WithGoodsName(q.GoodsName),
		goodsnorth.WithPageLimit(1, 1),
	}
}

/**
 * @brief: 查询商品详情响应参数
 */
type QueryGoodsRespDto struct {
	Owned        bool
	GoodsContent string
	Category     string
	GoodsName    string
}

func NewQueryGoodsRespDto(goods *goodsentity.GoodsEntity) *QueryGoodsRespDto {
	return &QueryGoodsRespDto{
		Owned:        goods.Owned,
		GoodsContent: goods.GoodsContent,
		Category:     goods.Category,
		GoodsName:    goods.GoodsName,
	}
}

/**
 * @brief: 购买商品请求参数
 */
type BuyGoodsReqDto struct {
	AccountId string
	StoreName string
	GoodsName string
}

func (b *BuyGoodsReqDto) ToQueryGoodsOption() []goodsnorth.QueryOption {
	return []goodsnorth.QueryOption{
		goodsnorth.WithUserId(b.AccountId),
		goodsnorth.WithStoreName(b.StoreName),
		goodsnorth.WithGoodsName(b.GoodsName),
	}
}

func (b *BuyGoodsReqDto) ToOrderVo(points int) *orderentity.OrderVo {
	return &orderentity.OrderVo{
		UserId:    b.AccountId,
		StoreName: b.StoreName,
		GoodsName: b.GoodsName,
		Points:    points,
	}
}

type ListOrderReqDto struct {
	AccountId string
	Status    string
}

func (l *ListOrderReqDto) ToQueryOptions() []ordernorth.QueryOption {
	return []ordernorth.QueryOption{
		ordernorth.WithUserId(l.AccountId),
		ordernorth.WithStatus(l.Status),
	}
}

type ListOrderRespDto struct {
	Total int
	List  []*OrderDataDto
}

type OrderDataDto struct {
	OrderId   string
	Points    int
	StoreName string
	GoodsName string
}

func NewListOrderRespDto(total int, datas []*orderentity.OrderEntity) *ListOrderRespDto {
	return &ListOrderRespDto{
		Total: total,
	}
}

func NewOrderDataDtos(datas []*orderentity.OrderEntity) []*OrderDataDto {
	_datas := make([]*OrderDataDto, 0, len(datas))
	for _, data := range datas {
		_datas = append(_datas, NewOrderDataDto(data))
	}
	return _datas
}

func NewOrderDataDto(data *orderentity.OrderEntity) *OrderDataDto {
	return &OrderDataDto{
		OrderId:   data.OrderId,
		Points:    data.Points,
		StoreName: data.StoreName,
		GoodsName: data.GoodsName,
	}
}

/**
 * @brief: 支付订单请求参数
 */
type PayOrderReqDto struct {
	AccountId string
	OrderId   string
}

func (p *PayOrderReqDto) ToOrderEntity() *orderentity.OrderEntity {
	return &orderentity.OrderEntity{
		UserId:  p.AccountId,
		OrderId: p.OrderId,
	}
}

func (p *PayOrderReqDto) ToAccountEntities(order *orderentity.OrderEntity) (payer *accountentity.AccountEntity, payee *accountentity.AccountEntity) {
	return &accountentity.AccountEntity{
			Points: order.Points,
			AccountVo: accountentity.AccountVo{
				Type:      accountentity.AccountTypeCustomer,
				AccountId: p.AccountId,
			},
		}, &accountentity.AccountEntity{
			Points: order.Points,
			AccountVo: accountentity.AccountVo{
				Type:      accountentity.AccountTypeMerchant,
				AccountId: order.StoreName,
			},
		}
}
