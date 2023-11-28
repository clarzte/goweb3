package server

import (
	"fmt"

	m "cryptom.app/merchants"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type Server struct {
	Port            int
	MongoDbUri      string
	RpcUrl          string
	AllowedOrigins  []string
	ServerReady     chan bool
	StorageHostname string
}

func (s *Server) Start() {
	e := echo.New()

	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: s.AllowedOrigins,
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept},
	}))

	client := m.NewMerchantsEventHandler(s.RpcUrl)
	repo := m.NewMerchantsRepository(s.MongoDbUri)
	storage := m.NewMerchantsStorage(s.StorageHostname)

	service := m.NewMerchantsService(repo, storage, client.ArweaveIDs)
	_ = m.NewMerchantsRoutes(e, service)

	go func() {
		if err := e.Start(fmt.Sprintf(":%d", s.Port)); err != nil {
			e.Logger.Fatal(err)
		}
	}()

	if s.ServerReady != nil {
		s.ServerReady <- true
	}
}
