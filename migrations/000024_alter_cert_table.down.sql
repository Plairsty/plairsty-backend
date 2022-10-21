-- DROP ALTER TABLE IF EXISTS certifications ADD COLUMN IF NOT EXISTS allotted_date timestamp;
-- ALTER TABLE IF EXISTS certifications ADD COLUMN IF NOT EXISTS expiry timestamp;
-- ALTER TABLE IF EXISTS certifications ADD COLUMN IF NOT EXISTS type varchar(255);
-- ALTER TABLE IF EXISTS certifications ADD COLUMN IF NOT EXISTS level varchar(255);
-- ALTER TABLE IF EXISTS certifications ADD COLUMN IF NOT EXISTS category varchar(255);
-- ALTER TABLE IF EXISTS certifications ADD COLUMN IF NOT EXISTS issuer varchar(255);
-- ALTER TABLE IF EXISTS certifications ADD COLUMN IF NOT EXISTS domain varchar(255);

ALTER TABLE IF EXISTS certifications
    DROP COLUMN IF EXISTS allotted_date;
ALTER TABLE IF EXISTS certifications
    DROP COLUMN IF EXISTS expiry;
ALTER TABLE IF EXISTS certifications
    DROP COLUMN IF EXISTS type;
ALTER TABLE IF EXISTS certifications
    DROP COLUMN IF EXISTS level CASCADE;
ALTER TABLE IF EXISTS certifications
    DROP COLUMN IF EXISTS category;
ALTER TABLE IF EXISTS certifications
    DROP COLUMN IF EXISTS issuer;
ALTER TABLE IF EXISTS certifications
    DROP COLUMN IF EXISTS domain;
ALTER TABLE IF EXISTS certifications
    ADD COLUMN IF NOT EXISTS provider varchar(255);

ALTER TABLE IF EXISTS certifications
    DROP COLUMN IF EXISTS user_id;

ALTER TABLE IF EXISTS certifications
    DROP CONSTRAINT IF EXISTS fk_user_id;