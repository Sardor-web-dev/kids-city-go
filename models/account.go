package models

type Account struct {
	ID                         string  `gorm:"primaryKey;column:id" json:"id"`
	UserID                     uint    `gorm:"column:userId" json:"userId"`
	Type                       string  `gorm:"column:type" json:"type"`
	Provider                   string  `gorm:"column:provider" json:"provider"`
	ProviderAccountID          string  `gorm:"column:providerAccountId" json:"providerAccountId"`
	RefreshToken               *string `gorm:"column:refresh_token" json:"refresh_token"`
	AccessToken                *string `gorm:"column:access_token" json:"access_token"`
	ExpiresAt                  *int    `gorm:"column:expires_at" json:"expires_at"`
	TokenType                  *string `gorm:"column:token_type" json:"token_type"`
	Scope                      *string `gorm:"column:scope" json:"scope"`
	IDToken                    *string `gorm:"column:id_token" json:"id_token"`
	SessionState               *string `gorm:"column:session_state" json:"session_state"`
	RefreshTokenExpiresIn      *int    `gorm:"column:refresh_token_expires_in" json:"refresh_token_expires_in"`
}

func (Account) TableName() string {
	return "Account"
}
