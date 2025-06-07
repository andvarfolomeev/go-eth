CREATE TABLE blocks (
    number BIGINT PRIMARY KEY,
    hash BYTEA NOT NULL,
    parent_hash BYTEA NOT NULL,
    state_root BYTEA NOT NULL,
    receipts_root BYTEA NOT NULL,
    transactions_root BYTEA NOT NULL,
    difficulty BIGINT NOT NULL,
    gas_limit BIGINT NOT NULL,
    gas_used BIGINT NOT NULL,
    timestamp BIGINT NOT NULL,
    nonce BYTEA NOT NULL,
    mix_hash BYTEA NOT NULL,
    sha3_uncles BYTEA NOT NULL,
    extra_data BYTEA NOT NULL,
    logs_bloom BYTEA NOT NULL,
    size BIGINT NOT NULL,
    miner BYTEA NOT NULL
);

CREATE TABLE receipts (
    block_hash BYTEA NOT NULL,
    block_number BIGINT NOT NULL REFERENCES blocks (number) ON DELETE CASCADE,
    transaction_hash BYTEA PRIMARY KEY,
    gas_used BIGINT NOT NULL,
    cumulative_gas_used BIGINT NOT NULL,
    effective_gas_price BIGINT NOT NULL,
    contract_address TEXT,
    status SMALLINT NOT NULL
);

CREATE TABLE txs (
    hash BYTEA PRIMARY KEY,
    block_hash BYTEA NOT NULL,
    block_number BIGINT NOT NULL REFERENCES blocks (number) ON DELETE CASCADE,
    "to" BYTEA NOT NULL,
    "from" BYTEA NOT NULL,
    nonce BIGINT NOT NULL,
    gas BIGINT NOT NULL,
    gas_price BIGINT NOT NULL,
    max_fee_per_gas BIGINT NOT NULL,
    max_priority_fee_per_gas BIGINT NOT NULL,
    value TEXT NOT NULL,
    input BYTEA NOT NULL,
    type SMALLINT NOT NULL
);
