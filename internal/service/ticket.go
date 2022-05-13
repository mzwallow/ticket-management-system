package service

import (
	"net/http"
	"strconv"

	"github.com/mzwallow/ticket-management-system/internal/pkg/models"
	"github.com/mzwallow/ticket-management-system/internal/repo"

	"github.com/gin-gonic/gin"
)

type TicketService struct {
	repo *repo.TicketRepository
}

// NewTicketService creates app service.
func NewTicketService(repo *repo.TicketRepository) *TicketService {
	return &TicketService{repo: repo}
}

// CheckHealth pings server connection.
func (s *TicketService) CheckHealth(ctx *gin.Context) {
	ctx.String(http.StatusOK, "pong")
}

// CreateTicket inserts a new ticket to database.
func (s *TicketService) CreateTicket(ctx *gin.Context) {
	var req models.CreateTicketRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		panic(err)
	}

	if err := s.repo.CreateTicket(req); err != nil {
		panic(err)
	}

	ctx.Status(http.StatusOK)
}

// ListAllTickets lists all tickets and sorting them by `id` by default.
// sort can be either `status` or `updated_at`.
func (s *TicketService) ListAllTickets(ctx *gin.Context) {
	// Get `sort` from query param
	sort := ctx.Query("sort")

	res, err := s.repo.ListAllTickets(sort)
	if err != nil {
		panic(err)
	}

	ctx.JSON(http.StatusOK, res)
}

// ListTicketByID returns a ticket by given `id“.
func (s *TicketService) ListTicketByID(ctx *gin.Context) {
	// Convert id string from param to int
	idString := ctx.Param("id")
	id, err := strconv.ParseInt(idString, 10, 64)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest,
			gin.H{"message": "id must be an int and cannot empty"})
	}

	res, err := s.repo.ListTicketByID(int(id))
	if err != nil {
		panic(err)
	}

	ctx.JSON(http.StatusOK, res)
}

// UpdateTicketByID updates ticket information by given `id`.
func (s *TicketService) UpdateTicketByID(ctx *gin.Context) {
	// Convert id string from param to int
	idString := ctx.Param("id")
	id, err := strconv.ParseInt(idString, 10, 64)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest,
			gin.H{"message": "id must be an int and cannot empty"})
	}

	var req models.UpdateTicketRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		panic(err)
	}

	if err := s.repo.UpdateTicketByID(int(id), req); err != nil {
		panic(err)
	}

	ctx.Status(http.StatusOK)
}

// UpdateTicketStatusByID updates ticket status by given `id`.
func (s *TicketService) UpdateTicketStatusByID(ctx *gin.Context) {
	// Convert id string from param to int
	idString := ctx.Param("id")
	id, err := strconv.ParseInt(idString, 10, 64)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest,
			gin.H{"message": "id must be an int and cannot empty"})
	}

	var req models.UpdateTicketStatusRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		panic(err)
	}

	if err := s.repo.UpdateTicketStatusByID(int(id), req); err != nil {
		panic(err)
	}

	ctx.Status(http.StatusOK)
}
