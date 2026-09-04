package repository

import (
	"donation-service/entity"

	"gorm.io/gorm"
)

type TransactionRepository interface {
	FindByCampaignID(campaignID uint) ([]entity.Transaction, error)
	Create(transaction *entity.Transaction) error
	// Tambahan fungsi sesuai Usecase Diagram
	FindByID(id uint) (*entity.Transaction, error)
	FindByUserID(userID uint) ([]entity.Transaction, error)
	FindAll() ([]entity.Transaction, error)
	Update(transaction *entity.Transaction) error
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

// Tambahan: Ambil 1 transaksi spesifik (buat dipakai pas admin mau Accept/Reject)
func (r *transactionRepository) FindByID(id uint) (*entity.Transaction, error) {
	var transaction entity.Transaction
	err := r.db.First(&transaction, id).Error
	if err != nil {
		return nil, err
	}
	return &transaction, nil
}

// Tambahan: Ambil daftar transaksi milik 1 user spesifik (fitur View Own Transactions)
func (r *transactionRepository) FindByUserID(userID uint) ([]entity.Transaction, error) {
	var transactions []entity.Transaction
	err := r.db.Where("user_id = ?", userID).Find(&transactions).Error
	return transactions, err
}

// Tambahan: Ambil semua transaksi (fitur Review Pending Transactions buat Admin)
func (r *transactionRepository) FindAll() ([]entity.Transaction, error) {
	var transactions []entity.Transaction
	err := r.db.Find(&transactions).Error
	return transactions, err
}

// Tambahan: Update transaksi (fitur Accept/Reject ubah status donasi)
func (r *transactionRepository) Update(transaction *entity.Transaction) error {
	return r.db.Save(transaction).Error
}
