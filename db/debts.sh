#!/bin/bash
psql -U pibardos -d pibardos_app <<-EOSQL
CREATE TABLE IF NOT EXISTS debts (
    id VARCHAR(255) PRIMARY KEY UNIQUE,
    group_id VARCHAR(255) REFERENCES groups(id),
    lender_id VARCHAR(255),
    borrower_id VARCHAR(255)
);
CREATE TABLE IF NOT EXISTS debt_requests (
    id VARCHAR(255) PRIMARY KEY UNIQUE,
    group_id VARCHAR(255) REFERENCES groups(id),
    lender_id VARCHAR(255),
    borrower_id VARCHAR(255),
    created_at VARCHAR(255),
    description VARCHAR(255),
    amount INTEGER,
    status INTEGER
);
EOSQL
