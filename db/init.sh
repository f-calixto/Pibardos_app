#!/bin/bash
psql -U pibardos -d pibardos_app <<-EOSQL
CREATE TABLE IF NOT EXISTS users (
    id VARCHAR(255) PRIMARY KEY,
    username VARCHAR(255),
    email VARCHAR(255),
    pwd VARCHAR(255),
    created_at VARCHAR(255),
    country VARCHAR(2),
    birthday VARCHAR(255),
    status VARCHAR(255),
    avatar VARCHAR(255),
    UNIQUE (id, username, email)
);
CREATE TABLE IF NOT EXISTS groups (
    id VARCHAR(255) PRIMARY KEY,
    name VARCHAR(255),
    size INTEGER,
    country VARCHAR(2),
    admin_id VARCHAR(255),
    access_code VARCHAR(6),
    access_code_issue_time INTEGER,
    avatar VARCHAR(255),
    created_at VARCHAR(255),
    UNIQUE (id, name)
);
CREATE TABLE IF NOT EXISTS users_groups (
    user_id VARCHAR(255) REFERENCES users(id),
    group_id VARCHAR(255) REFERENCES groups(id)
);
CREATE TABLE IF NOT EXISTS group_activities (
    id VARCHAR(255) PRIMARY KEY,
    group_id VARCHAR(255) REFERENCES groups(id),
    title VARCHAR(255),
    date VARCHAR(255),
    guest_list VARCHAR(255)[],
    UNIQUE (id)
);
CREATE TABLE IF NOT EXISTS debts (
    id VARCHAR(255) PRIMARY KEY,
    group_id VARCHAR(255) REFERENCES groups(id),
    lender_id VARCHAR(255),
    borrower_id VARCHAR(255),
    UNIQUE (id)
);
CREATE TABLE IF NOT EXISTS debt_requests (
    id VARCHAR(255) PRIMARY KEY,
    group_id VARCHAR(255) REFERENCES groups(id),
    lender_id VARCHAR(255),
    borrower_id VARCHAR(255),
    created_at VARCHAR(255),
    description VARCHAR(255),
    amount INTEGER,
    status INTEGER,
    UNIQUE (id)
);
EOSQL
