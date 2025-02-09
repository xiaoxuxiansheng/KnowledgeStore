package toc

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/xiaoxuxiansheng/KnowledgeStore/presentation/iservice"
)

type Controller struct {
	service iservice.ToCService
}

func NewController(service iservice.ToCService) *Controller {
	return &Controller{
		service: service,
	}
}

func (ctrl *Controller) EchoRet(c *gin.Context, err error, data interface{}) {
	if err != nil {
		c.JSON(http.StatusOK, map[string]interface{}{
			"errno":  ErrnoInternel,
			"errmsg": err.Error(),
			"data":   nil,
		})
		return
	}
	c.JSON(http.StatusOK, map[string]interface{}{
		"errno":  ErrnoSuccess,
		"errmsg": "ok",
		"data":   data,
	})
}

// Register 注册用户
// @Summary 注册用户
// @Description 注册用户
// @Tags C端接口
// @Accept application/json
// @Produce application/json
// @Param def body RegisterCustomerReqVo true "注册用户请求参数"
// @Router /api/customer/v1/register [post]
func (ctrl *Controller) RegisterCustomer(c *gin.Context) {
	var req RegisterCustomerReqVo
	if err := c.BindJSON(&req); err != nil {
		ctrl.EchoRet(c, err, nil)
		return
	}

	err := ctrl.service.RegisterCustomer(c.Request.Context(), nil)
	ctrl.EchoRet(c, err, nil)
}

// Register 登录用户
// @Summary 登录用户
// @Description 登录用户
// @Tags C端接口
// @Accept application/json
// @Produce application/json
// @Param def body LoginCustomerReqVo true "登录用户请求参数"
// @Router /api/customer/v1/login [post]
func (ctrl *Controller) LoginCustomer(c *gin.Context) {
	var req LoginCustomerReqVo
	if err := c.BindJSON(&req); err != nil {
		ctrl.EchoRet(c, err, nil)
		return
	}

	err := ctrl.service.LoginCustomer(c.Request.Context(), req.ToDto())
	ctrl.EchoRet(c, err, nil)
}

// Register 查询店铺列表
// @Summary 查询店铺列表
// @Description 查询店铺列表
// @Tags C端接口
// @Accept application/json
// @Produce application/json
// @Param def body ListStoreReqVo true "查询店铺列表请求参数"
// @Router /api/store/v1/list [get]
func (ctrl *Controller) ListStore(c *gin.Context) {
	var req ListStoreReqVo
	if err := c.Bind(&req); err != nil {
		ctrl.EchoRet(c, err, nil)
		return
	}

	resp, err := ctrl.service.ListStore(c.Request.Context(), req.ToDto())
	if err != nil {
		ctrl.EchoRet(c, err, nil)
		return
	}
	ctrl.EchoRet(c, err, NewListStoreRespVo(resp))
}

// Register 查询商品列表
// @Summary 查询商品列表
// @Description 查询商品列表
// @Tags C端接口
// @Accept application/json
// @Produce application/json
// @Param def body ListStoreReqVo true "查询商品列表请求参数"
// @Router /api/goods/v1/list [get]
func (ctrl *Controller) ListGoods(c *gin.Context) {
	var req ListGoodsReqVo
	if err := c.Bind(&req); err != nil {
		ctrl.EchoRet(c, err, nil)
		return
	}

	resp, err := ctrl.service.ListGoods(c.Request.Context(), req.ToDto())
	if err != nil {
		ctrl.EchoRet(c, err, nil)
		return
	}
	ctrl.EchoRet(c, err, NewListGoodsRespVo(resp))
}

// Register 查询商品详情
// @Summary 查询商品详情
// @Description 查询商品详情
// @Tags C端接口
// @Accept application/json
// @Produce application/json
// @Param def body QueryGoodsReqVo true "查询商品详情请求参数"
// @Router /api/goods/v1/query [get]
func (ctrl *Controller) QueryGoods(c *gin.Context) {
	var req QueryGoodsReqVo
	if err := c.Bind(&req); err != nil {
		ctrl.EchoRet(c, err, nil)
		return
	}

	resp, err := ctrl.service.QueryGoods(c.Request.Context(), req.ToDto())
	if err != nil {
		ctrl.EchoRet(c, err, nil)
		return
	}
	ctrl.EchoRet(c, err, NewQueryGoodsRespVo(resp))
}

// Register 购买商品
// @Summary 购买商品
// @Description 购买商品
// @Tags C端接口
// @Accept application/json
// @Produce application/json
// @Param def body BuyGoodsReqVo true "购买商品请求参数"
// @Router /api/goods/v1/buy [get]
func (ctrl *Controller) BuyGoods(c *gin.Context) {
	var req BuyGoodsReqVo
	if err := c.BindJSON(&req); err != nil {
		ctrl.EchoRet(c, err, nil)
		return
	}

	err := ctrl.service.BuyGoods(c.Request.Context(), req.ToDto())
	ctrl.EchoRet(c, err, nil)
}

// Register 查看我的订单
// @Summary 查看我的订单
// @Description 查看我的订单
// @Tags C端接口
// @Accept application/json
// @Produce application/json
// @Param def body ListOrderReqVo true "查看我的订单请求参数"
// @Router /api/order/v1/list [get]
func (ctrl *Controller) ListOrder(c *gin.Context) {
	var req ListOrderReqVo
	if err := c.BindJSON(&req); err != nil {
		ctrl.EchoRet(c, err, nil)
		return
	}

	resp, err := ctrl.service.ListOrder(c.Request.Context(), req.ToDto())
	ctrl.EchoRet(c, err, NewListOrderRespVo(resp))
}

// Register 支付订单
// @Summary 支付订单
// @Description 支付订单
// @Tags C端接口
// @Accept application/json
// @Produce application/json
// @Param def body PayOrderReq true "支付订单请求参数"
// @Router /api/order/v1/pay [get]
func (ctrl *Controller) PayOrder(c *gin.Context) {
	var req PayOrderReqVo
	if err := c.BindJSON(&req); err != nil {
		ctrl.EchoRet(c, err, nil)
		return
	}

	err := ctrl.service.PayOrder(c.Request.Context(), req.ToDto())
	ctrl.EchoRet(c, err, nil)
}
