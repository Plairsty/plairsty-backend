ALTER TABLE resume DROP COLUMN IF EXISTS resume_url;

ALTER TABLE resume ADD COLUMN file_key VARCHAR(255) NOT NULL default '';