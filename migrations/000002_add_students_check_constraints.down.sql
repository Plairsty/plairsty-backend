ALTER TABLE students
    DROP CONSTRAINT IF EXISTS email_unique;
ALTER TABLE phone
    DROP CONSTRAINT IF EXISTS phone_check;