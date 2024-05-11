CREATE TABLE IF NOT EXISTS user_elo_history (
    id UUID PRIMARY KEY,
    user_id UUID NOT NULL,
    previous_elo INTEGER,
    new_elo INTEGER,
    created_at BIGINT NOT NULL DEFAULT EXTRACT(EPOCH FROM CURRENT_TIMESTAMP) * 1000,
    updated_at BIGINT NOT NULL DEFAULT EXTRACT(EPOCH FROM CURRENT_TIMESTAMP) * 1000,
    FOREIGN KEY(user_id) REFERENCES users(id)
);

CREATE INDEX IF NOT EXISTS idx_user_id on user_elo_history(user_id)
