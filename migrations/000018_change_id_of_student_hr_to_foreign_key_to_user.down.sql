ALTER TABLE hr
    ADD CONSTRAINT hr_user_fk FOREIGN KEY (id) REFERENCES "user" (id);
ALTER TABLE students
    ADD CONSTRAINT students_user_fk FOREIGN KEY (id) REFERENCES "user" (id);