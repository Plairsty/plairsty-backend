CREATE TABLE IF NOT EXISTS certifications
(
    id          bigserial PRIMARY KEY,
    title       varchar(150) NOT NULL,
    description varchar(250) NOT NULL,
    url         varchar(250) NOT NULL,
    provider    varchar(50)  NOT NULL,
    cert_id     varchar(100) NOT NULL DEFAULT '',
    created_at  timestamp    NOT NULL DEFAULT NOW(),
    updated_at  timestamp    NOT NULL DEFAULT NOW()
);