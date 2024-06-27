CREATE TABLE "users" (
    id SERIAL PRIMARY KEY,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL,
    first_name VARCHAR(255) NOT NULL,
    last_name VARCHAR(255) NOT NULL,
    email VARCHAR(255) NOT NULL UNIQUE,
    is_admin BOOLEAN NOT NULL DEFAULT FALSE,
    passwd VARCHAR(255) NOT NULL
);

CREATE TABLE accounts (
    id SERIAL PRIMARY KEY,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL,
    name VARCHAR(255) NOT NULL,
    user_id INT NOT NULL,
    entity INT NOT NULL,
    currency INT NOT NULL,
    FOREIGN KEY (user_id) REFERENCES "users"(id)
);

CREATE TABLE transactions (
    id SERIAL PRIMARY KEY,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL,
    concept VARCHAR(255) NOT NULL,
    description TEXT,
    value INT NOT NULL,
    date TIMESTAMP NOT NULL,
    relevance INT NOT NULL,
    account_id INT NOT NULL,
    FOREIGN KEY (account_id) REFERENCES accounts(id)
);

CREATE TABLE credits (
    id SERIAL PRIMARY KEY,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL,
    closing_date TIMESTAMP NOT NULL,
    due_date TIMESTAMP NOT NULL,
    identifier VARCHAR(255) NOT NULL,
    entity INT NOT NULL,
    type INT NOT NULL,
    rate FLOAT NOT NULL,
    total INT NOT NULL,
    user_id INT NOT NULL,
    installments SMALLINT NOT NULL,
    FOREIGN KEY (user_id) REFERENCES "users"(id)
);
