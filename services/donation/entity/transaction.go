package entity

import "time"

const (
	TransactionStatusPending  = "pending"
	TransactionStatusAccepted = "accepted"
	TransactionStatusRejected = "rejected"
)

type Transaction struct {
	ID               uint       `gorm:"primaryKey;column:id" json:"id"`
	UserID           uint       `gorm:"column:user_id;not null" json:"user_id"`
	CampaignID       uint       `gorm:"column:campaign_id;not null" json:"campaign_id"`
	Name             string     `gorm:"column:name;not null" json:"name"`
	Amount           float64    `gorm:"column:amount;not null" json:"amount"`
	Message          string     `gorm:"column:message" json:"message,omitempty"`
	IsAnonymous      bool       `gorm:"column:is_anonymous;not null;default:false" json:"is_anonymous"`
	Status           string     `gorm:"column:status;not null;default:'pending'" json:"status"`
	VerifiedByUserID *uint      `gorm:"column:verified_by_user_id" json:"verified_by_user_id,omitempty"`
	VerifiedAt       *time.Time `gorm:"column:verified_at" json:"verified_at,omitempty"`
	CreatedAt        time.Time  `gorm:"column:created_at" json:"created_at"`
}

func (Transaction) TableName() string { return "transactions" }
