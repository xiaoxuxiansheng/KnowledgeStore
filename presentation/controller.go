package presentation

import (
	"github.com/xiaoxuxiansheng/KnowledgeStore/presentation/iservice"
	"go.uber.org/dig"
)

type Controller struct {
	businessService iservice.BusinessService
	customerService iservice.CustomerService
}

type ControllerParam struct {
	dig.In
	BusinessService iservice.BusinessService
	CustomerService iservice.CustomerService
}

func NewController(param ControllerParam) *Controller {
	return &Controller{
		businessService: param.BusinessService,
		customerService: param.CustomerService,
	}
}
