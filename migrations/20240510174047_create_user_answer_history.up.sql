CREATE TABLE IF NOT EXISTS history_answers (
    id UUID PRIMARY KEY,
    user_id UUID NOT NULL,
    question_id UUID NOT NULL,
    answer TEXT NOT NULL,
    is_correct BOOLEAN,
    created_at BIGINT NOT NULL DEFAULT EXTRACT(EPOCH FROM CURRENT_TIMESTAMP) * 1000,
    updated_at BIGINT NOT NULL DEFAULT EXTRACT(EPOCH FROM CURRENT_TIMESTAMP) * 1000,
    FOREIGN KEY(user_id) REFERENCES users(id),
    FOREIGN KEY(question_id) REFERENCES questions(id)
);

CREATE INDEX IF NOT EXISTS idx_user_id on history_answers(user_id)

CREATE INDEX IF NOT EXISTS idx_history_answers_user_id_question_id_is_correct ON history_answers (user_id, question_id, is_correct);

CREATE INDEX IF NOT EXISTS idx_questions_learning_topic_id ON questions (learning_topic_id);

CREATE INDEX IF NOT EXISTS idx_questions_id ON questions (id);
