ALTER TABLE application
    ADD COLUMN IF NOT EXISTS status bool default false;