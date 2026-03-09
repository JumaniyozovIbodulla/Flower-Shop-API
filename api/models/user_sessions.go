package models

import "github.com/google/uuid"


// Redis DB
type UserSession struct {
	ID           uuid.UUID `json:"id" example:"a1810161-00fc-434e-8452-364fd3683f56"`
	UserID       uuid.UUID `json:"user_id" example:"9c1242b2-b211-4a91-ba6c-a58e903327fd"`
	AccessToken  string    `json:"access_token" example:"access_token"`
	RefreshToken string    `json:"refresh_token" example:"refresh_token"`
	UserAgent    string    `json:"user_agent" example:"user_agent"`
	IPAddress    string    `json:"ip_address" example:"0.0.0.0"`
	CreatedAt    int64     `json:"created_at"`
	LastActive   int64     `json:"last_active"`
	ExpiresAt    int64     `json:"expired_at"`
}

