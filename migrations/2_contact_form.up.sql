CREATE TABLE IF NOT EXISTS contact_form
(
    id       SERIAL PRIMARY KEY,
    username VARCHAR(30),
    email    VARCHAR(50),
    message  VARCHAR(1000),
    clientIp VARCHAR(15),
    sendAt   INT
);