CREATE TABLE IF NOT EXISTS api_account
(
    id          SERIAL PRIMARY KEY,
    username    VARCHAR(30) UNIQUE,
    password    VARCHAR(200),
    createAt    INT
);