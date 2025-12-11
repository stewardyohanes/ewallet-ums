package models

import "time"

type Outbox struct {
	ID        int		`gorm:"primaryKey"`
	EventType string	`gorm:"size:50"`
	Payload   string	`gorm:"type:json"`
	Status    string	`gorm:"size:20;type:ENUM('pending', 'sent', 'failed')"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (Outbox) TableName() string { return "outbox" }