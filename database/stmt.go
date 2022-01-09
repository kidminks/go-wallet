package database

const CreateWalletQueryValue = "(?,?,?,?,?,?,?,?,?,?,?,?,?)"
const CreateWalletQuery = "INSERT INTO wallet (entity_id, entity_uuid, entity_type, wallet_type, minimum_balance, maximum_balance, status, in_active, current_balance, wallet_metadata, user_metadata, created_by, updated_by) VALUES "
const SelectWalletByIdQuery = "SELECT * FROM wallet where id = ? limit 1"
