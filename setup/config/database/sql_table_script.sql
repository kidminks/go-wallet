CREATE TABLE IF NOT EXISTS wallet (
    id bigint PRIMARY KEY AUTO_INCREMENT,
    uuid binary(16) NOT NULL DEFAULT (UUID_TO_BIN(UUID())),
    entity_id bigint,
    entity_uuid varchar(256),
    entity_type varchar(256),
    wallet_type varchar(256),
    minimum_balance double NOT NULL DEFAULT 0,
    maximum_balance double NOT NULL DEFAULT 0,
    status smallint,
    is_active smallint,
    current_balance double NOT NULL DEFAULT 0,
    wallet_metadata json,
    user_metadata json,
    created_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
    created_by varchar(256) NOT NULL,
    updated_by varchar(256) NOT NULL
);

CREATE TABLE IF NOT EXISTS transaction (
  id bigint PRIMARY KEY AUTO_INCREMENT,
  uuid binary(16) NOT NULL DEFAULT (UUID_TO_BIN(UUID())),
  primary_wallet_id bigint NOT NULL,
  secondary_wallet_id bigint NOT NULL,
  primary_wallet_uuid binary(16) NOT NULL,
  secondary_Wallet_uuid binary(16) NOT NULL,
  amount double NOT NULL,
  multiplier tinyint NOT NULL,
  start_date DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
  completion_date DATETIME,
  payment_date DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
  payment_status smallint NOT NULL,
  transaction_metadata json,
  comment text,
  created_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
  updated_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
  created_by varchar(256) NOT NULL,
  updated_by varchar(256) NOT NULL,
  FOREIGN KEY (primary_wallet_id) REFERENCES wallet(id),
  FOREIGN KEY (secondary_wallet_id) REFERENCES wallet(id)
);

CREATE TABLE IF NOT EXISTS transaction_status_history (
    id bigint PRIMARY KEY AUTO_INCREMENT,
    transaction_id bigint NOT NULL,
    transaction_uuid binary(16) NOT NULL,
    status smallint NOT NULL,
    created_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
    created_by varchar(256) NOT NULL,
    updated_by varchar(256) NOT NULL,
    FOREIGN KEY (transaction_id) REFERENCES transaction(id)
);