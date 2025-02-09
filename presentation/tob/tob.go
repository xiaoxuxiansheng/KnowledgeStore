package tob

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/xiaoxuxiansheng/KnowledgeStore/presentation/iservice"
)

type Controller struct {
	service iservice.ToBService
}

func NewController(service iservice.ToBService) *Controller {
	return &Controller{
		service: service,
	}
}

func (ctrl *Controller) EchoRet(c *gin.Context, err error, data interface{}) {
	if err != nil {
		c.JSON(http.StatusOK, map[string]interface{}{
			"errno":  ErrnoUnknown,
			"errmsg": err.Error(),
			"data":   data,
		})
		return
	}
	c.JSON(http.StatusOK, map[string]interface{}{
		"errno":  ErrnoSuccess,
		"errmsg": "ok",
		"data":   data,
	})
}

// Register 注册商家账号
// @Summary 注册商家账号
// @Description 注册商家账号
// @Tags B端接口
// @Accept application/json
// @Produce application/json
// @Param def body RegisterBusinessReqVo true "注册商家账号请求参数"
// @Router /api/tob/v1/business/register [post]
func (ctrl *Controller) RegisterBusiness(c *gin.Context) {
	var req RegisterBusinessReqVo
	if err := c.BindJSON(&req); err != nil {
		ctrl.EchoRet(c, err, nil)
		return
	}

	err := ctrl.service.RegisterBusiness(c.Request.Context(), nil)
	ctrl.EchoRet(c, err, nil)
}

// Register 登录商家账号
// @Summary 登录商家账号
// @Description 登录商家账号
// @Tags B端接口
// @Accept application/json
// @Produce application/json
// @Param def body LoginBusinessReqVo true "登录商家账号请求参数"
// @Success 200 {object} Response
// @Router /api/tob/v1/business/login [post]
func (ctrl *Controller) LoginBusiness(c *gin.Context) {
	var req LoginBusinessReqVo
	if err := c.BindJSON(&req); err != nil {
		ctrl.EchoRet(c, err, nil)
		return
	}

	err := ctrl.service.LoginBusiness(c.Request.Context(), req.ToDto())
	ctrl.EchoRet(c, err, nil)
}

// Register 创建店铺
// @Summary 创建店铺
// @Description 创建店铺
// @Tags B端接口
// @Accept application/json
// @Produce application/json
// @Param def body CreateStoreReqVo true "创建店铺请求参数"
// @Success 200 {object} Response
// @Router /api/tob/v1/store/create [post]
func (ctrl *Controller) CreateStore(c *gin.Context) {
	var req CreateStoreReqVo
	if err := c.BindJSON(&req); err != nil {
		ctrl.EchoRet(c, err, nil)
		return
	}

	err := ctrl.service.CreateStore(c.Request.Context(), req.ToDto())
	ctrl.EchoRet(c, err, nil)
}

// Register 发布店铺
// @Summary 发布店铺
// @Description 发布店铺
// @Tags B端接口
// @Accept application/json
// @Produce application/json
// @Param def body OnlineStoreReqVo true "发布店铺请求参数"
// @Success 200 {object} Response
// @Router /api/tob/v1/store/online [post]
func (ctrl *Controller) OnlineStore(c *gin.Context) {
	var req OnlineStoreReqVo
	if err := c.BindJSON(&req); err != nil {
		ctrl.EchoRet(c, err, nil)
		return
	}

	err := ctrl.service.OnlineStore(c.Request.Context(), req.ToDto())
	ctrl.EchoRet(c, err, nil)
}

// Register 下线店铺
// @Summary 下线店铺
// @Description 下线店铺
// @Tags B端接口
// @Accept application/json
// @Produce application/json
// @Param def body OfflineStoreReqVo true "下线店铺请求参数"
// @Success 200 {object} Response
// @Router /api/tob/v1/store/offline [post]
func (ctrl *Controller) OfflineStore(c *gin.Context) {
	var req OfflineStoreReqVo
	if err := c.BindJSON(&req); err != nil {
		ctrl.EchoRet(c, err, nil)
		return
	}

	err := ctrl.service.OfflineStore(c.Request.Context(), req.ToDto())
	ctrl.EchoRet(c, err, nil)
}

// Register 创建商品
// @Summary 创建商品
// @Description 创建商品
// @Tags B端接口
// @Accept application/json
// @Produce application/json
// @Param def body CreateGoodsReqVo true "创建商品请求参数"
// @Success 200 {object} Response
// @Router /api/tob/v1/goods/create [post]
func (ctrl *Controller) CreateGoods(c *gin.Context) {
	var req CreateGoodsReqVo
	if err := c.BindJSON(&req); err != nil {
		ctrl.EchoRet(c, err, nil)
		return
	}

	err := ctrl.service.CreateGoods(c.Request.Context(), req.ToDto())
	ctrl.EchoRet(c, err, nil)
}

// Register 发布商品
// @Summary 发布商品
// @Description 发布商品
// @Tags B端接口
// @Accept application/json
// @Produce application/json
// @Param def body OnlineGoodsReqVo true "发布商品请求参数"
// @Success 200 {object} Response
// @Router /api/tob/v1/goods/online [post]
func (ctrl *Controller) OnlineGoods(c *gin.Context) {
	var req OnlineGoodsReqVo
	if err := c.BindJSON(&req); err != nil {
		ctrl.EchoRet(c, err, nil)
		return
	}

	err := ctrl.service.OnlineGoods(c.Request.Context(), req.ToDto())
	ctrl.EchoRet(c, err, nil)
}

// Register 下线商品
// @Summary 下线商品
// @Description 下线商品
// @Tags B端接口
// @Accept application/json
// @Produce application/json
// @Param def body OfflineGoodsReqVo true "下线商品请求参数"
// @Success 200 {object} Response
// @Router /api/tob/v1/goods/offline [post]
func (ctrl *Controller) OfflineGoods(c *gin.Context) {
	var req OfflineGoodsReqVo
	if err := c.BindJSON(&req); err != nil {
		ctrl.EchoRet(c, err, nil)
		return
	}

	err := ctrl.service.OfflineGoods(c.Request.Context(), req.ToDto())
	ctrl.EchoRet(c, err, nil)
}
