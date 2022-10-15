CREATE TABLE IF NOT EXISTS resume
(
    id         bigserial PRIMARY KEY,
    resume_url varchar(255) NOT NULL, --- Since resume will be directly uploaded to the server, we will store the url of the resume in the database
    student_id int          NOT NULL REFERENCES students (id),
    created_at timestamp    NOT NULL DEFAULT NOW(),
    updated_at timestamp    NOT NULL DEFAULT NOW()
);