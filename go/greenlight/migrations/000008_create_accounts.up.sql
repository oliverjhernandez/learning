CREATE TABLE IF NOT EXISTS accounts (
    id SERIAL PRIMARY KEY,
    created_at TIMESTAMP (0) WITH TIME ZONE NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP (0) WITH TIME ZONE NOT NULL DEFAULT NOW(),
    title VARCHAR(255) NOT NULL,
    user_id INT NOT NULL,
    entity_id INT NOT NULL,
    FOREIGN KEY (user_id) REFERENCES users (id),
    FOREIGN KEY (entity_id) REFERENCES entities (id)
);
