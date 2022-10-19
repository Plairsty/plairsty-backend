-- ALter constraint named fk_user_id to reference student table

ALTER TABLE IF EXISTS certifications DROP CONSTRAINT IF EXISTS fk_user_id;
ALTER TABLE IF EXISTS certifications ADD CONSTRAINT fk_user_id FOREIGN KEY (user_id) REFERENCES students (id) ON DELETE CASCADE;