ALTER TABLE IF EXISTS students
    ADD CONSTRAINT email_unique UNIQUE (email);

-- Phone number, check if it is 10 digits
ALTER TABLE IF EXISTS phone
    ADD CONSTRAINT phone_check CHECK (phone ~ '^[0-9]{10}$');