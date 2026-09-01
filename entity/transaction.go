package entity

import "time"

type Transaction struct {
	ID               uint       `gorm:"primaryKey;column:id" json:"id"`
	UserID           uint       `gorm:"column:user_id;not null" json:"user_id"`
	CampaignID       uint       `gorm:"column:campaign_id;not null" json:"campaign_id"`
	Amount           float64    `gorm:"type:numeric(15,2);not null" json:"amount"`
	Status           string     `gorm:"type:varchar(50);not null;default:'pending'" json:"status"`
	VerifiedByUserID *uint      `gorm:"column:verified_by_user_id" json:"verified_by_user_id"`
	VerifiedAt       *time.Time `gorm:"column:verified_at" json:"verified_at"`
	CreatedAt        time.Time  `gorm:"column:created_at" json:"created_at"`
}

func (Transaction) TableName() string { return "transactions" }
