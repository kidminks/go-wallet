package model

type Wallet struct {
	Id int64 `json:"id"`
	Uuid string `json:"uuid"`
	EntityId int64  `json:"entity_id"`
	EntityUuid string `json:"entity_uuid"`
	EntityType string `json:"entity_type"`
	WalletType string `json:"wallet_type"`
	MinimumBalance float64 `json:"minimum_balance"`
	MaximumBalance float64 `json:"maximum_balance"`
	Status int `json:"status"`
	IsActive bool `json:"is_active"`
	CurrentBalance float64 `json:"current_balance"`
	WalletMetadata StringInterfaceMap `json:"wallet_metadata"`
	UserMetadata StringInterfaceMap `json:"user_metadata"`
	CreatedAt int64 `json:"created_at"`
	CreatedBy string `json:"created_by"`
	UpdatedAt int64 `json:"updated_at"`
	UpdatedBy string `json:"updated_by"`
}
