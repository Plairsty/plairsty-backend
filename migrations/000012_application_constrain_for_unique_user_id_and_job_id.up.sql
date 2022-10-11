ALTER TABLE application
    ADD CONSTRAINT unique_application UNIQUE (user_id, job_id);