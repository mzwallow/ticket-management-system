package main

import (
	"github.com/mzwallow/ticket-management-system/configs"
	"github.com/mzwallow/ticket-management-system/internal/database"
	"github.com/mzwallow/ticket-management-system/internal/repo"
	"github.com/mzwallow/ticket-management-system/internal/server"
	"github.com/mzwallow/ticket-management-system/internal/service"
)

func main() {
	cfg := configs.GetConfig()
	db := database.InitDB(cfg.DBUrl)
	repo := repo.NewTicketRepository(db)
	service := service.NewTicketService(repo)
	server := server.NewServer(cfg.ListenAddr, service)
	server.Run()
}
