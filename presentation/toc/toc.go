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

// Register 注册用户
// @Summary 注册用户
// @Description 注册用户
// @Tags C端接口
// @Accept application/json
// @Produce application/json
// @Param def body RegisterReqVo true "注册用户请求参数"
// @Success 200 {object} RegisterRespVo
// @Router /api/customer/v1/register [post]
func (ctrl *Controller) RegisterCustomer(c *gin.Context) {
	var req RegisterReqVo
	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, nil)
		return
	}

	_, _ = ctrl.service.RegisterCustomer(c.Request.Context(), nil)
	c.JSON(http.StatusOK, &RegisterRespVo{})
}

func (ctrl *Controller) LoginCustomer(c *gin.Context) {
	var req RegisterReqVo
	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, nil)
		return
	}

	_, _ = ctrl.service.LoginCustomer(c.Request.Context(), nil)
	c.JSON(http.StatusOK, &RegisterRespVo{})
}

func (ctrl *Controller) ListStore(c *gin.Context) {
	var req RegisterReqVo
	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, nil)
		return
	}

	_, _ = ctrl.service.ListStore(c.Request.Context(), nil)
	c.JSON(http.StatusOK, &RegisterRespVo{})
}

func (ctrl *Controller) ListGoods(c *gin.Context) {
	var req RegisterReqVo
	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, nil)
		return
	}

	_, _ = ctrl.service.ListGoods(c.Request.Context(), nil)
	c.JSON(http.StatusOK, &RegisterRespVo{})
}

func (ctrl *Controller) QueryGoods(c *gin.Context) {
	var req RegisterReqVo
	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, nil)
		return
	}

	_, _ = ctrl.service.QueryGoods(c.Request.Context(), nil)
	c.JSON(http.StatusOK, &RegisterRespVo{})
}

func (ctrl *Controller) BuyGoods(c *gin.Context) {
	var req RegisterReqVo
	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, nil)
		return
	}

	_, _ = ctrl.service.BuyGoods(c.Request.Context(), nil)
	c.JSON(http.StatusOK, &RegisterRespVo{})
}

func (ctrl *Controller) PayOrder(c *gin.Context) {
	var req RegisterReqVo
	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, nil)
		return
	}

	_, _ = ctrl.service.PayOrder(c.Request.Context(), nil)
	c.JSON(http.StatusOK, &RegisterRespVo{})
}
