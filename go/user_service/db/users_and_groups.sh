#!/bin/bash
psql -U pibardos -d pibardos_app <<-EOSQL
CREATE TABLE IF NOT EXISTS users (
    id VARCHAR(255) PRIMARY KEY UNIQUE,
    username VARCHAR(255) UNIQUE,
    email VARCHAR(255) UNIQUE,
    created_at VARCHAR(255),
    country VARCHAR(2),
    birthdate VARCHAR(255),
    status VARCHAR(255),
    avatar VARCHAR(255)
);
CREATE TABLE IF NOT EXISTS groups (
    id VARCHAR(255) PRIMARY KEY UNIQUE,
    name VARCHAR(255) UNIQUE,
    size INTEGER,
    country VARCHAR(2),
    admin_id VARCHAR(255),
    access_code VARCHAR(6),
    access_code_issue_time INTEGER,
    avatar VARCHAR(255),
    created_at VARCHAR(255)
);
CREATE TABLE IF NOT EXISTS users_groups (
    user_id VARCHAR(255) REFERENCES users(id),
    group_id VARCHAR(255) REFERENCES groups(id)
);
EOSQL
