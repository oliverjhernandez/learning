CREATE INDEX IF NOT EXISTS transactions_description_idx ON transactions USING GIN (to_tsvector('simple', description));
