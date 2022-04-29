#!/bin/bash
psql --username pibardos --dbname pibardos_app <<-EOSQL
CREATE TABLE IF NOT EXISTS files (
    id VARCHAR(255) PRIMARY KEY,
    name VARCHAR(255),
);
EOSQL
