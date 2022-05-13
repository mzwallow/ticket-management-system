package repo

import (
	"context"
	"fmt"

	"github.com/mzwallow/ticket-management-system/internal/database"
	"github.com/mzwallow/ticket-management-system/internal/pkg/models"

	"github.com/pkg/errors"
)

type TicketRepository struct {
	db *database.DB
}

// NewTicketRepository create database related service.
func NewTicketRepository(db *database.DB) *TicketRepository {
	return &TicketRepository{db: db}
}

// CreateTicket inserts a new ticket to database.
func (r *TicketRepository) CreateTicket(ticket models.CreateTicketRequest) error {
	conn, err := r.db.Connect()
	if err != nil {
		return err
	}
	defer conn.Close(context.Background())

	if _, err := conn.Prepare(context.Background(), "create_ticket",
		`
			INSERT INTO tickets (title, description, contact_information)
			VALUES ($1, $2, $3);
		`); err != nil {
		return errors.Wrap(err, "prepare 'create_ticket' statement failed")
	}

	if _, err := conn.Exec(context.Background(), "create_ticket",
		ticket.Title,
		ticket.Description,
		ticket.ContactInformation); err != nil {
		return errors.Wrap(err, "execute 'create_ticket' statement failed")
	}

	return nil
}

// ListAllTickets lists all tickets and sorting them by `id` by default.
// sort can be either `status` or `updated_at`.
func (r *TicketRepository) ListAllTickets(sort string) ([]models.Ticket, error) {
	conn, err := r.db.Connect()
	if err != nil {
		return []models.Ticket{}, err
	}
	defer conn.Close(context.Background())

	// Check given sort method
	var order string
	switch sort {
	case "status":
		order = "status"
	case "updated_at":
		order = "updated_at DESC"
	default:
		order = "id"
	}

	sql := fmt.Sprintf(`
		SELECT id, title, status, description, 
			   contact_information, created_at, updated_at
		FROM tickets ORDER BY %s;`, order)

	if _, err := conn.Prepare(context.Background(), "list_tickets", sql); err != nil {
		return []models.Ticket{}, errors.Wrap(err, "prepare 'list_tickets' statement failed")
	}

	rows, err := conn.Query(context.Background(), "list_tickets")
	if err != nil {
		return []models.Ticket{}, errors.Wrap(err, "execute 'list_tickets' statement failed")
	}
	defer rows.Close()

	var results []models.Ticket
	for rows.Next() {
		var result models.Ticket
		if err := rows.Scan(&result.ID,
			&result.Title,
			&result.Status,
			&result.Description,
			&result.ContactInformation,
			&result.CreatedAt,
			&result.UpdatedAt); err != nil {
			return nil, errors.Wrapf(err, "ListAllTickets: scan row failed")
		}
		results = append(results, result)
	}
	if err := rows.Err(); err != nil {
		return []models.Ticket{}, errors.Wrap(err, "ListAllTickets: reading rows failed")
	}

	return results, nil
}

// ListTicketByID returns a ticket by given `id``.
func (r *TicketRepository) ListTicketByID(id int) (models.Ticket, error) {
	conn, err := r.db.Connect()
	if err != nil {
		return models.Ticket{}, err
	}
	defer conn.Close(context.Background())

	if _, err := conn.Prepare(context.Background(), "list_ticket_by_id",
		`
			SELECT id, title, status, description, 
				   contact_information, created_at, updated_at
			FROM tickets
			WHERE id = $1;
		`); err != nil {
		return models.Ticket{}, errors.Wrap(err, "prepare 'list_ticket_by_id' statement failed")
	}

	var result models.Ticket
	if err := conn.QueryRow(context.Background(), "list_ticket_by_id", id).Scan(&result.ID,
		&result.Title,
		&result.Status,
		&result.Description,
		&result.ContactInformation,
		&result.CreatedAt,
		&result.UpdatedAt); err != nil {
		return models.Ticket{}, errors.Wrap(err, "ListTicketByID: ticket not found")
	}

	return result, nil
}

// UpdateTicketByID updates ticket information by given `id`.
func (r *TicketRepository) UpdateTicketByID(id int, update models.UpdateTicketRequest) error {
	conn, err := r.db.Connect()
	if err != nil {
		return err
	}
	defer conn.Close(context.Background())

	if _, err := conn.Prepare(context.Background(), "update_ticket",
		`
			UPDATE tickets 
			SET title 				= $2, 
				description 		= $3, 
				contact_information = $4,
				updated_at 			= current_timestamp
			WHERE id = $1;
		`); err != nil {
		return errors.Wrap(err, "prepare 'update_ticket' statement failed")
	}

	if _, err := conn.Exec(context.Background(), "update_ticket",
		id,
		update.Title,
		update.Description,
		update.ContactInformation); err != nil {
		return errors.Wrap(err, "execute 'update_ticket' statement failed")
	}

	return nil
}

// UpdateTicketStatusByID updates ticket status by given `id`.
func (r *TicketRepository) UpdateTicketStatusByID(id int, update models.UpdateTicketStatusRequest) error {
	conn, err := r.db.Connect()
	if err != nil {
		return err
	}
	defer conn.Close(context.Background())

	// Check if ticket status is valid
	switch update.Status {
	case models.Pending, models.Accepted, models.Resolved, models.Rejected:
	default:
		return errors.New("invalid status")
	}

	if _, err := conn.Prepare(context.Background(), "update_ticket_status",
		`
			UPDATE tickets 
			SET status 				= $2, 
				updated_at 			= current_timestamp
			WHERE id = $1;
		`); err != nil {
		return errors.Wrap(err, "prepare 'update_ticket_status' statement failed")
	}

	if _, err := conn.Exec(context.Background(), "update_ticket_status",
		id,
		update.Status); err != nil {
		return errors.Wrap(err, "execute 'update_ticket_status' statement failed")
	}

	return nil
}
