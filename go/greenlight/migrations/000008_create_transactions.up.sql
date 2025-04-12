CREATE TABLE transactions (
    id SERIAL PRIMARY KEY,
    created_at TIMESTAMP (0) WITH TIME ZONE NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP (0) WITH TIME ZONE NOT NULL DEFAULT NOW(),
    concept VARCHAR(255) NOT NULL,
    description TEXT,
    amount INT NOT NULL,
    relevance INT NOT NULL,
    account_id INT NOT NULL,
    FOREIGN KEY (account_id) REFERENCES accounts (id)
);
