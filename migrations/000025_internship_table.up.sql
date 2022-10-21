CREATE TABLE IF NOT EXISTS internship (
    id bigserial PRIMARY KEY,
    user_id bigint REFERENCES "user" (id) ON DELETE CASCADE,
    title varchar(50) NOT NULL,
    description varchar(255) NOT NULL,
    start_date date NOT NULL,
    end_date date NOT NULL,
    company_name varchar(100) NOT NULL,
    location varchar(100) NOT NULL,
    mentor_name varchar(100) NOT NULL,
    created_at timestamp NOT NULL DEFAULT NOW(),
    updated_at timestamp NOT NULL DEFAULT NOW()
)
