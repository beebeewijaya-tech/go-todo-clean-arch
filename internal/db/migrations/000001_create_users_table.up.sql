CREATE TABLE IF NOT EXISTS users (
    id UUID PRIMARY KEY ,
    email VARCHAR(300) UNIQUE,
    password VARCHAR(500),
    created_at TIMESTAMPTZ default now()
);

CREATE INDEX u_email_index ON users(email);