package models
package models

import (
	"time"
)

type User struct {
	UserID    uint      `gorm:"primaryKey;column:user_id" json:"user_id"`
	Name      string    `gorm:"type:varchar(100);not null" json:"name"`
	Email     string    `gorm:"type:varchar(100);unique;not null" json:"email"`
	Password  string    `gorm:"type:varchar(255);not null" json:"-"`
	Role      string    `gorm:"type:varchar(20);default:'user'" json:"role"`
	CreatedAt time.Time `gorm:"column:created_at" json:"created_at"`
	UpdatedAt time.Time `gorm:"column:updated_at" json:"updated_at"`
}

func (User) TableName() string { return "users" }

type Account struct {
	AccountID uint      `gorm:"primaryKey;column:account_id" json:"account_id"`
	UserID    uint      `gorm:"column:user_id" json:"user_id"`
	Balance   float64   `gorm:"type:decimal(15,2);default:0.00" json:"balance"`
	CreatedAt time.Time `gorm:"column:created_at" json:"created_at"`
	UpdatedAt time.Time `gorm:"column:updated_at" json:"updated_at"`

	User User `gorm:"foreignKey:UserID" json:"user"`
}

func (Account) TableName() string { return "accounts" }

type Campaign struct {
	CampaignID    uint      `gorm:"primaryKey;column:campaign_id" json:"campaign_id"`
	Title         string    `gorm:"type:varchar(150);not null" json:"title"`
	Description   string    `gorm:"type:text;not null" json:"description"`
	TargetAmount  float64   `gorm:"type:decimal(15,2);not null" json:"target_amount"`
	CurrentAmount float64   `gorm:"type:decimal(15,2);default:0.00" json:"current_amount"`
	Status        string    `gorm:"type:varchar(20);default:'active'" json:"status"`
	CreatedBy     uint      `gorm:"column:created_by" json:"created_by"`
	CreatedAt     time.Time `gorm:"column:created_at" json:"created_at"`
	UpdatedAt     time.Time `gorm:"column:updated_at" json:"updated_at"`

	Creator User `gorm:"foreignKey:CreatedBy" json:"creator"`
}

func (Campaign) TableName() string { return "campaigns" }

type PaymentMethod struct {
	PaymentMethodID uint      `gorm:"primaryKey;column:payment_method_id" json:"payment_method_id"`
	Name            string    `gorm:"type:varchar(50);not null" json:"name"`
	IsActive        bool      `gorm:"default:true" json:"is_active"`
	CreatedAt       time.Time `gorm:"column:created_at" json:"created_at"`
}

func (PaymentMethod) TableName() string { return "payment_methods" }

type Donation struct {
	DonationID uint      `gorm:"primaryKey;column:donation_id" json:"donation_id"`
	CampaignID uint      `gorm:"column:campaign_id" json:"campaign_id"`
	UserID     *uint     `gorm:"column:user_id" json:"user_id"` // Pointer agar bisa null jika donatur anonim/tamu
	Amount     float64   `gorm:"type:decimal(15,2);not null" json:"amount"`
	Comment    string    `gorm:"type:text" json:"comment"`
	Anonymous  bool      `gorm:"default:false" json:"anonymous"`
	CreatedAt  time.Time `gorm:"column:created_at" json:"created_at"`

	Campaign Campaign `gorm:"foreignKey:CampaignID" json:"campaign"`
	User     *User    `gorm:"foreignKey:UserID" json:"user"`
}

func (Donation) TableName() string { return "donations" }

type Transaction struct {
	TransactionID   uint      `gorm:"primaryKey;column:transaction_id" json:"transaction_id"`
	DonationID      uint      `gorm:"column:donation_id" json:"donation_id"`
	PaymentMethodID uint      `gorm:"column:payment_method_id" json:"payment_method_id"`
	Status          string    `gorm:"type:varchar(30);default:'pending'" json:"status"`
	ExternalID      string    `gorm:"type:varchar(100)" json:"external_id"`
	CreatedAt       time.Time `gorm:"column:created_at" json:"created_at"`
	UpdatedAt       time.Time `gorm:"column:updated_at" json:"updated_at"`

	Donation      Donation      `gorm:"foreignKey:DonationID" json:"donation"`
	PaymentMethod PaymentMethod `gorm:"foreignKey:PaymentMethodID" json:"payment_method"`
}

func (Transaction) TableName() string { return "transactions" }