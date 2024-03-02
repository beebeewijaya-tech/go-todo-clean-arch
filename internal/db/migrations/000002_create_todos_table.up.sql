CREATE TABLE IF NOT EXISTS todos (
    id UUID PRIMARY KEY,
    title VARCHAR(300),
    body TEXT,
    user_id uuid,
    created_at TIMESTAMPTZ default now(),

    CONSTRAINT fk_user_id FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE
);

CREATE INDEX t_title_index ON todos(title);
