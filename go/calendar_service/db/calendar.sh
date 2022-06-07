#!/bin/bash
psql -U pibardos -d pibardos_app <<-EOSQL
CREATE TABLE IF NOT EXISTS calendar_events (
    id VARCHAR(255) PRIMARY KEY UNIQUE,
    group_id VARCHAR(255),
    title VARCHAR(255),
    start_date VARCHAR(255),
    end_date VARCHAR(255),
    creator_id VARCHAR(255),
    cancelled BOOLEAN,
    guest_list VARCHAR(255)[]
);
EOSQL
