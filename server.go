package main

import (
	"github.com/xiaoxuxiansheng/KnowledgeStore/presentation"
	"go.uber.org/dig"
)

type Server struct {
	dig.In
	Ctrl *presentation.Controller
}

func (s Server) Run() error {
	return nil
}
