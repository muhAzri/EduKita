CREATE TABLE IF NOT EXISTS question_stats (
    question_id UUID PRIMARY KEY,
    total_attempts INTEGER NOT NULL DEFAULT 0,
    correct_attempts INTEGER NOT NULL DEFAULT 0,
    FOREIGN KEY (question_id) REFERENCES questions(id)
);