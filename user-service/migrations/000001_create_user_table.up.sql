CREATE TABLE IF NOT EXISTS users(
    id SERIAL PRIMARY KEY,
    name TEXT,
    email TEXT,
    created_at TIME DEFAULT CURRENT_TIMESTAMP,
    updated_at TIME DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIME 
);