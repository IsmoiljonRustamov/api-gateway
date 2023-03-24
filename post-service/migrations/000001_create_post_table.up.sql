CREATE TABLE IF NOT EXISTS posts(
    id SERIAL PRIMARY KEY,
    title TEXT,
    description TEXT,
    user_id INTEGER,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP 
);