package main

import (
	// infrastructure
	"github.com/xiaoxuxiansheng/KnowledgeStore/infrastructure/dao/account"
	"github.com/xiaoxuxiansheng/KnowledgeStore/infrastructure/dao/goods"
	"github.com/xiaoxuxiansheng/KnowledgeStore/infrastructure/dao/merchant"
	"github.com/xiaoxuxiansheng/KnowledgeStore/infrastructure/dao/order"
	"github.com/xiaoxuxiansheng/KnowledgeStore/infrastructure/dao/store"
	"github.com/xiaoxuxiansheng/KnowledgeStore/infrastructure/dao/user"

	// domain
	accountdomain "github.com/xiaoxuxiansheng/KnowledgeStore/domain/account/service"
	goodsdomain "github.com/xiaoxuxiansheng/KnowledgeStore/domain/goods/service"
	merchantdomain "github.com/xiaoxuxiansheng/KnowledgeStore/domain/merchant/service"
	orderdomain "github.com/xiaoxuxiansheng/KnowledgeStore/domain/order/service"
	storedomain "github.com/xiaoxuxiansheng/KnowledgeStore/domain/store/service"
	userdomain "github.com/xiaoxuxiansheng/KnowledgeStore/domain/user/service"

	//application
	businessapp "github.com/xiaoxuxiansheng/KnowledgeStore/application/tob/service"
	customerapp "github.com/xiaoxuxiansheng/KnowledgeStore/application/toc/service"

	// presentation
	businesspresentation "github.com/xiaoxuxiansheng/KnowledgeStore/presentation/tob"
	customerpresentation "github.com/xiaoxuxiansheng/KnowledgeStore/presentation/toc"

	// server
	"github.com/xiaoxuxiansheng/KnowledgeStore/server"

	"go.uber.org/dig"
)

var container *dig.Container

func init() {
	container = dig.New()

	// 基础设施层组件注入
	provideInfra(container)
	// 领域层组件注入
	provideDomain(container)
	// 应用层组件注入
	provideApplication(container)
	// 展示层组件注入
	providePresentation(container)
}

func provideInfra(c *dig.Container) {
	c.Provide(user.NewDao)
	c.Provide(store.NewDao)
	c.Provide(order.NewDao)
	c.Provide(merchant.NewDao)
	c.Provide(goods.NewDao)
	c.Provide(account.NewDao)
}

func provideDomain(c *dig.Container) {
	c.Provide(userdomain.NewService)
	c.Provide(storedomain.NewService)
	c.Provide(orderdomain.NewService)
	c.Provide(merchantdomain.NewService)
	c.Provide(goodsdomain.NewService)
	c.Provide(accountdomain.NewService)
}

func provideApplication(c *dig.Container) {
	c.Provide(businessapp.NewService)
	c.Provide(customerapp.NewService)
}

func providePresentation(c *dig.Container) {
	c.Provide(businesspresentation.NewController)
	c.Provide(customerpresentation.NewController)
}

func GetServer() server.Server {
	var s server.Server
	if err := container.Invoke(func(_server server.Server) {
		s = _server
	}); err != nil {
		panic(err)
	}
	return s
}
