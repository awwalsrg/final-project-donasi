package domain

import "errors"

var (
	ErrTransactionNotFound = errors.New("transaction not found")
	ErrAlreadyVerified     = errors.New("transaction has already been verified")
	ErrInvalidAmount       = errors.New("amount must be greater than zero")
	ErrCampaignUnavailable = errors.New("campaign is closed or does not exist")
)
