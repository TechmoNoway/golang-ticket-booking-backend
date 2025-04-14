package models

import (
	"context"
	"time"
)

type Ticket struct {
	ID        uint      `json:"id" gorm:"primarykey"`
	EventID   uint      `json:"eventId"`
	Event     Event     `json:"event" gorm:"foreignkey:EventID;contraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Entered   bool      `json:"Entered" default:"false"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

type TicketRepository interface {
	GetMany(ctx context.Context) ([]*Ticket, error)
	GetOne(ctx context.Context, ticketId uint) (*Ticket, error)
	CreateOne(ctx context.Context, ticket *Ticket) (*Ticket, error)
	UpdateOne(ctx context.Context, ticketId uint, updateData map[string]interface{}) (*Ticket, error)
	DeleteOne(ctx context.Context, ticketId uint) error
}

type ValidateTicket struct {
	TicketId uint `json:"ticketId"`
}
