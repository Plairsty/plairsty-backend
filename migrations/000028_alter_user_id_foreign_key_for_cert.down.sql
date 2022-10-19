ALTER TABLE IF EXISTS certifications
    DROP CONSTRAINT IF EXISTS fk_user_id;
ALTER TABLE IF EXISTS certifications
    ADD CONSTRAINT fk_user_id FOREIGN KEY (user_id) REFERENCES "user" (id);