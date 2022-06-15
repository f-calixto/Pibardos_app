#!/bin/bash
psql -U pibardos -d pibardos_app <<-EOSQL
CREATE TABLE IF NOT EXISTS debts (
    group_id VARCHAR(255),
    lender_id VARCHAR(255),
    borrower_id VARCHAR(255),
    amount INTEGER,
    UNIQUE(group_id, lender_id, borrower_id)
);
CREATE TABLE IF NOT EXISTS debt_requests (
    id VARCHAR(255) PRIMARY KEY UNIQUE,
    group_id VARCHAR(255),
    lender_id VARCHAR(255),
    borrower_id VARCHAR(255),
    date VARCHAR(255),
    description VARCHAR(255),
    amount INTEGER,
    status INTEGER
);
EOSQL
