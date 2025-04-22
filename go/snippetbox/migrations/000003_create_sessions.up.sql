CREATE TABLE sessions (
    token CHAR(43) PRIMARY KEY,
    session_data BLOB NOT NULL,
    expiry TIMESTAMP (6) NOT NULL
);
