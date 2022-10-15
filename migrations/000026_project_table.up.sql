CREATE TABLE IF NOT EXISTS projects
(
    id          bigserial PRIMARY KEY,
    user_id     bigint REFERENCES "user" (id) ON DELETE CASCADE,
    title       varchar(255) NOT NULL,
    description varchar(500) NOT NULL,
    leader      varchar(50)  NOT NULL,
    member_ids  text[]       NOT NULL,
    guide_name  varchar(50)  NOT NULL,
    project_url varchar(255) NOT NULL,
    semester    integer      NOT NULL,
    start_date  date         NOT NULL,
    end_date    date         NOT NULL,
    created_at  timestamp    NOT NULL DEFAULT NOW(),
    updated_at  timestamp    NOT NULL DEFAULT NOW()
)