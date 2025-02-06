package server

import (
	"github.com/xiaoxuxiansheng/KnowledgeStore/server/http"

	"go.uber.org/dig"
)

type Server struct {
	dig.In
	HttpServer *http.Server
}

func (s Server) Run() error {
	return s.HttpServer.Run()
}
