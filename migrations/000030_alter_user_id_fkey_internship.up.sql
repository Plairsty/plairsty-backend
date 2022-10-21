ALTER TABLE IF EXISTS internship DROP CONSTRAINT IF EXISTS fk_user_id;
ALTER TABLE IF EXISTS internship DROP CONSTRAINT IF EXISTS internship_user_id_fkey;

ALTER TABLE IF EXISTS internship ADD CONSTRAINT internship_user_id_fkey FOREIGN KEY (user_id) REFERENCES students (id) ON DELETE CASCADE;