CREATE TABLE accounts (
    id SERIAL PRIMARY KEY,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL,
    title VARCHAR(255) NOT NULL,
    user_id INT NOT NULL,
    entity INT NOT NULL,
    currency INT NOT NULL,
    FOREIGN KEY (user_id) REFERENCES users (id)
);
