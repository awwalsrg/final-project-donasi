package repositories

import (
	"donasi-app/entity"

	"gorm.io/gorm"
)

type TransactionRepository interface {
	FindByCampaignID(campaignID uint) ([]entity.Transaction, error)
	Create(transaction *entity.Transaction) error
}

type transactionRepository struct {
	db *gorm.DB
}

func NewTransactionRepository(db *gorm.DB) TransactionRepository {
	return &transactionRepository{db: db}
}

func (r *transactionRepository) FindByCampaignID(campaignID uint) ([]entity.Transaction, error) {
	var transactions []entity.Transaction
	err := r.db.Where("campaign_id = ?", campaignID).Find(&transactions).Error
	return transactions, err
}

func (r *transactionRepository) Create(transaction *entity.Transaction) error {
	return r.db.Create(transaction).Error
}
