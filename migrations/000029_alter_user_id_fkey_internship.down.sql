ALTER TABLE IF EXISTS internship
    DROP CONSTRAINT IF EXISTS fk_user_id;
ALTER TABLE IF EXISTS internship
    ADD CONSTRAINT fk_user_id FOREIGN KEY (user_id) REFERENCES "user" (id);