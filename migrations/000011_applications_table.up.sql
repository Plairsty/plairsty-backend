CREATE TABLE IF NOT EXISTS application (
    id bigserial PRIMARY KEY,
    -- Foreign key to the user table
    user_id bigint REFERENCES students(id),
    -- Foreign key to the job table
    job_id bigint REFERENCES jobs(id)
)