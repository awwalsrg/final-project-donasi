package domain

import "errors"

var (
	ErrCampaignNotFound = errors.New("campaign not found")
	ErrCampaignClosed   = errors.New("campaign is not active")
	ErrDeadlinePassed   = errors.New("campaign deadline has passed")
	ErrInvalidCategory  = errors.New("one or more categories do not exist")
	ErrInvalidAmount    = errors.New("amount must be greater than zero")
)
