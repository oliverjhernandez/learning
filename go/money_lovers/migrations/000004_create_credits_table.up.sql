CREATE TABLE credits (
    id SERIAL PRIMARY KEY,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL,
    closing_day INT NOT NULL,
    due_day INT NOT NULL,
    identifier VARCHAR(255) NOT NULL,
    entity INT NOT NULL,
    type INT NOT NULL,
    rate FLOAT NOT NULL,
    total INT NOT NULL,
    user_id INT NOT NULL,
    installments SMALLINT NOT NULL,
    FOREIGN KEY (user_id) REFERENCES "users"(id)
);
