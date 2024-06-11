CREATE TABLE IF NOT EXISTS questions (
    id UUID PRIMARY KEY,
    learning_topic_id UUID NOT NULL,
    content TEXT NOT NULL,
    answers TEXT[] NOT NULL,
    correct_answer_index INTEGER NOT NULL,
    created_at BIGINT NOT NULL DEFAULT EXTRACT(EPOCH FROM CURRENT_TIMESTAMP) * 1000,
    updated_at BIGINT NOT NULL DEFAULT EXTRACT(EPOCH FROM CURRENT_TIMESTAMP) * 1000,
    FOREIGN KEY (learning_topic_id) REFERENCES learning_topics(id)
);

CREATE INDEX IF NOT EXISTS idx_learning_topic_id ON questions (learning_topic_id);
