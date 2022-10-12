CREATE TABLE IF NOT EXISTS "user"
(
    id             bigserial PRIMARY KEY,
    username       varchar(255) NOT NULL,
    hashedPassword varchar(255) NOT NULL
);