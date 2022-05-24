#!/bin/bash
psql -U pibardos -d pibardos_app <<-EOSQL
CREATE TABLE IF NOT EXISTS group_activities (
    id VARCHAR(255) PRIMARY KEY UNIQUE,
    group_id VARCHAR(255) REFERENCES groups(id),
    title VARCHAR(255),
    date VARCHAR(255),
    guest_list VARCHAR(255)[]
);
EOSQL
