package model

import (
	"gorm.io/datatypes"
	"gorm.io/gorm"
)

type Wallet struct {
	gorm.Model
	Id int64 `gorm:"primaryKey:autoIncrement" json:"id"`
	Uuid string `gorm:"type:varchar(36);default:(UUID());uniqueIndex" json:"uuid"`
	EntityId int64  `gorm:"uniqueIndex:unique_entity;not null" json:"entity_id"`
	EntityUuid string `gorm:"uniqueIndex:unique_entity;type:varchar(36);not null" json:"entity_uuid"`
	EntityType string `gorm:"uniqueIndex:unique_entity;type:varchar(36);not null" json:"entity_type"`
	WalletType string `gorm:"type:varchar(36);not null" json:"wallet_type"`
	MinimumBalance float64 `gorm:"default:0" json:"minimum_balance"`
	MaximumBalance float64 `gorm:"default:0" json:"maximum_balance"`
	Status int `json:"status"`
	IsActive bool `gorm:"default:false" json:"is_active"`
	CurrentBalance float64 `gorm:"default:0" json:"current_balance"`
	WalletMetadata datatypes.JSON `json:"wallet_metadata"`
	UserMetadata datatypes.JSON `json:"user_metadata"`
	CreatedBy string `gorm:"type:varchar(36)" json:"created_by"`
	UpdatedBy string `gorm:"type:varchar(36)" json:"updated_by"`
}
