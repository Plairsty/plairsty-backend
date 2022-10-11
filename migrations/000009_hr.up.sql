CREATE TABLE IF NOT EXISTS hr
(
    id      bigserial PRIMARY KEY,
    name    TEXT NOT NULL,
    phone   TEXT NOT NULL,
    email   TEXT NOT NULL,
    company TEXT NOT NULL
);

CREATE TABLE IF NOT EXISTS jobs
(
    id             bigserial PRIMARY KEY,
    role           TEXT   NOT NULL,
    department     TEXT   NOT NULL,
    skills         TEXT   NOT NULL,
    experience     TEXT   NOT NULL,
    required_cgpa  TEXT   NOT NULL,
    description    TEXT   NOT NULL,
    location       TEXT   NOT NULL,
    certifications TEXT   NOT NULL,
    title          TEXT   NOT NULL,
    company        TEXT   NOT NULL,
    hr_id          BIGINT NOT NULL,
    FOREIGN KEY (hr_id) REFERENCES hr (id)
);