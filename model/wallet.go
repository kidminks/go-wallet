package model

import (
	"gorm.io/datatypes"
	"gorm.io/gorm"
)

type Wallet struct {
	gorm.Model
	Id int64 `json:"id"`
	Uuid string `gorm:"type:varchar(36);default:(UUID())" json:"uuid"`
	EntityId int64  `json:"entity_id"`
	EntityUuid string `gorm:"type:varchar(36)" json:"entity_uuid"`
	EntityType string `gorm:"type:varchar(36)" json:"entity_type"`
	WalletType string `gorm:"type:varchar(36)" json:"wallet_type"`
	MinimumBalance float64 `json:"minimum_balance"`
	MaximumBalance float64 `json:"maximum_balance"`
	Status int `json:"status"`
	IsActive bool `json:"is_active"`
	CurrentBalance float64 `json:"current_balance"`
	WalletMetadata datatypes.JSON `json:"wallet_metadata"`
	UserMetadata datatypes.JSON `json:"user_metadata"`
	CreatedBy string `gorm:"type:varchar(36)" json:"created_by"`
	UpdatedBy string `gorm:"type:varchar(36)" json:"updated_by"`
}
