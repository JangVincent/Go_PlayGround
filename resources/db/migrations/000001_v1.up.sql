CREATE TABLE IF NOT EXISTS Users (
    uuid VARCHAR(512) PRIMARY KEY,
    nickname VARCHAR(512) UNIQUE NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP NOT NULL
);
