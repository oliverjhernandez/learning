CREATE TABLE transactions (
    id SERIAL PRIMARY KEY,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL,
    concept VARCHAR(255) NOT NULL,
    description TEXT,
    amount INT NOT NULL,
    relevance INT NOT NULL,
    account_id INT NOT NULL,
    FOREIGN KEY (account_id) REFERENCES accounts (id)
);
