package models

import "time"

type TicketStatus string

const (
	Pending  TicketStatus = "PENDING"
	Accepted TicketStatus = "ACCEPTED"
	Resolved TicketStatus = "RESOLVED"
	Rejected TicketStatus = "REJECTED"
)

type Ticket struct {
	ID                 int       `json:"id"`
	Title              string    `json:"title"`
	Status             string    `json:"status"`
	Description        string    `json:"description"`
	ContactInformation string    `json:"contact_information"`
	CreatedAt          time.Time `json:"created_at"`
	UpdatedAt          time.Time `json:"updated_at"`
}

type CreateTicketRequest struct {
	Title              string `json:"title,omitempty"`
	Description        string `json:"description,omitempty"`
	ContactInformation string `json:"contact_information,omitempty"`
}

type UpdateTicketRequest struct {
	Title              string `json:"title,omitempty"`
	Description        string `json:"description,omitempty"`
	ContactInformation string `json:"contact_information,omitempty"`
}

type UpdateTicketStatusRequest struct {
	Status TicketStatus `json:"status,omitempty"`
}
