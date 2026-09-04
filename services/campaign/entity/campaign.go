package entity

import (
	"time"

	"gorm.io/gorm"
)

const (
	CampaignStatusActive = "active"
	CampaignStatusClosed = "closed"
)

type Campaign struct {
	ID              uint           `gorm:"primaryKey;column:id" json:"id"`
	Title           string         `gorm:"column:title;not null" json:"title"`
	Description     string         `gorm:"column:description;not null" json:"description"`
	TargetAmount    float64        `gorm:"column:target_amount;not null" json:"target_amount"`
	CollectedAmount float64        `gorm:"column:collected_amount;not null;default:0" json:"collected_amount"`
	Deadline        time.Time      `gorm:"column:deadline;not null" json:"deadline"`
	Status          string         `gorm:"column:status;not null;default:'active'" json:"status"`
	CreatedByUserID uint           `gorm:"column:created_by_user_id;not null" json:"created_by_user_id"`
	CreatedAt       time.Time      `gorm:"column:created_at" json:"created_at"`
	DeletedAt       gorm.DeletedAt `gorm:"column:deleted_at" json:"-"`

	Categories []Category `gorm:"many2many:campaign_categories;joinForeignKey:campaign_id;joinReferences:category_id" json:"categories,omitempty"`
}

func (Campaign) TableName() string { return "campaigns" }
