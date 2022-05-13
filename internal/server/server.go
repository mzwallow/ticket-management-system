package server

import (
	"github.com/mzwallow/ticket-management-system/internal/service"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

type Server struct {
	addr    string
	service *service.TicketService
}

// NewServer creates new REST server.
func NewServer(addr string, service *service.TicketService) *Server {
	return &Server{addr: addr, service: service}
}

func (s *Server) Run() {
	router := gin.Default()
	router.Use(cors.Default())

	v1 := router.Group("/api/v1")
	{
		v1.GET("/ping", s.service.CheckHealth)
		v1.POST("/tickets", s.service.CreateTicket)
		v1.GET("/tickets", s.service.ListAllTickets)
		v1.GET("/tickets/:id", s.service.ListTicketByID)
		v1.PATCH("/tickets/:id", s.service.UpdateTicketByID)
		v1.PATCH("/tickets/:id/status", s.service.UpdateTicketStatusByID)
	}

	router.Run(s.addr)
}
