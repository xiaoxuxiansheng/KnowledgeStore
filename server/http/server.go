package http

import (
	"github.com/xiaoxuxiansheng/KnowledgeStore/presentation/tob"
	"github.com/xiaoxuxiansheng/KnowledgeStore/presentation/toc"

	"github.com/gin-gonic/gin"
	"go.uber.org/dig"
)

type Server struct {
	engine  *gin.Engine
	tobCtrl *tob.Controller
	toCCtrl *toc.Controller
}

type ServerParam struct {
	dig.In
	BusinessCtrl *tob.Controller
	CustomerCtrl *toc.Controller
}

func NewServer(serverParam ServerParam) *Server {
	s := Server{
		engine:  gin.Default(),
		tobCtrl: serverParam.BusinessCtrl,
		toCCtrl: serverParam.CustomerCtrl,
	}

	s.registerBusinessRouter()
	s.registerCustomerRouter()
	return &s
}

func (s *Server) Run() error {
	return s.engine.Run(":8080")
}

func (s *Server) registerBusinessRouter() {
	group := s.engine.Group("api/tob/v1")
	// 注册商家账号
	group.POST("/business/register", s.tobCtrl.RegisterBusiness)
	// 登录商家账号
	group.POST("/business/login", s.tobCtrl.LoginBusiness)
	// 创建店铺
	group.POST("/store/create", s.tobCtrl.CreateStore)
	// 发布店铺
	group.POST("/store/online", s.tobCtrl.OnlineStore)
	// 下线店铺
	group.POST("/store/offline", s.tobCtrl.OfflineStore)
	// 创建商品
	group.POST("/goods/create", s.tobCtrl.CreateGoods)
	// 发布商品
	group.POST("/goods/online", s.tobCtrl.OnlineGoods)
	// 下线商品
	group.POST("/goods/offline", s.tobCtrl.OfflineGoods)
}

func (s *Server) registerCustomerRouter() {
	group := s.engine.Group("api/toc/v1")
	// 注册客户账户
	group.POST("/customer/register", s.toCCtrl.RegisterCustomer)
	// 登录客户账户
	group.POST("/customer/login", s.toCCtrl.LoginCustomer)
	// 查询店铺列表
	group.GET("/store/list", s.toCCtrl.ListStore)
	// 查询商品列表
	group.GET("/goods/list", s.toCCtrl.ListGoods)
	// 查询商品详情
	group.GET("/goods/query", s.toCCtrl.QueryGoods)
	// 购买商品
	group.POST("/goods/buy", s.toCCtrl.BuyGoods)
	// 支付订单
	group.POST("/order/pay", s.toCCtrl.PayOrder)
}
