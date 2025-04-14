package repositories

import (
	"context"

	"github.com/TechmoNoway/golang-ticket-booking-backend/models"
	"gorm.io/gorm"
)

type TicketRepository struct {
	db *gorm.DB
}

func (r *TicketRepository) GetMany(ctx context.Context) ([]*models.Ticket, error) {
	tickets := []*models.Ticket{}

	res := r.db.Model(&models.Ticket{}).Preload("Event").Order("updated_at desc").Find(&tickets)

	if res.Error != nil {
		return nil, res.Error
	}

	return tickets, nil
}

func (r *TicketRepository) GetOne(ctx context.Context, ticketId uint) (*models.Ticket, error) {
	ticket := &models.Ticket{}

	res := r.db.Model(ticket).Preload("Event").Where("id = ?", ticketId).First(ticket)

	if res.Error != nil {
		return nil, res.Error
	}

	return ticket, nil
}

func (r *TicketRepository) CreateOne(ctx context.Context, ticket *models.Ticket) (*models.Ticket, error) {
	res := r.db.Model(ticket).Create(ticket)

	if res.Error != nil {
		return nil, res.Error
	}

	ticket, err := r.GetOne(ctx, ticket.ID)

	if err != nil {
		return nil, err
	}

	return ticket, nil
}

func (r *TicketRepository) UpdateOne(ctx context.Context, ticketId uint, updateData map[string]interface{}) (*models.Ticket, error) {
	ticket := &models.Ticket{}

	updateRes := r.db.Model(ticket).Where("id = ?", ticketId).Updates(updateData)

	if updateRes.Error != nil {
		return nil, updateRes.Error
	}

	getRes, err := r.GetOne(ctx, ticketId)

	if err != nil {
		return nil, err
	}

	return getRes, nil
}

func (r *TicketRepository) DeleteOne(ctx context.Context, ticketId uint) error {

	res := r.db.Delete(&models.Event{}, ticketId)

	return res.Error
}

func NewTicketRepository(db *gorm.DB) models.TicketRepository {
	return &TicketRepository{
		db: db,
	}
}
