ALTER TABLE IF EXISTS certifications
    ADD COLUMN IF NOT EXISTS allotted_date timestamp;
ALTER TABLE IF EXISTS certifications
    ADD COLUMN IF NOT EXISTS expiry timestamp;
ALTER TABLE IF EXISTS certifications
    ADD COLUMN IF NOT EXISTS type varchar(255);
ALTER TABLE IF EXISTS certifications
    ADD COLUMN IF NOT EXISTS level int
        CONSTRAINT level CHECK (level >= 0 AND level <= 3);
ALTER TABLE IF EXISTS certifications
    ADD COLUMN IF NOT EXISTS category varchar(255);
ALTER TABLE IF EXISTS certifications
    ADD COLUMN IF NOT EXISTS issuer varchar(255);
ALTER TABLE IF EXISTS certifications
    ADD COLUMN IF NOT EXISTS domain varchar(255);
ALTER TABLE IF EXISTS certifications
    ADD COLUMN IF NOT EXISTS user_id int;
ALTER TABLE IF EXISTS certifications
    ADD CONSTRAINT fk_user_id FOREIGN KEY (user_id) REFERENCES "user" (id);
ALTER TABLE IF EXISTS certifications
    DROP COLUMN IF EXISTS provider;
