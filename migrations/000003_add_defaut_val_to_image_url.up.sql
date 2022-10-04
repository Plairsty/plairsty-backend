CREATE TABLE IF NOT EXISTS students
(
    ID          int                         NOT NULL PRIMARY KEY,
    FIRST_NAME  varchar(50)                 NOT NULL,
    MIDDLE_NAME varchar(50)                 NOT NULL,
    LAST_NAME   varchar(50)                 NOT NULL,
    EMAIL       varchar(50)                 NOT NULL,
    IMAGE_URL   varchar(50)                 NOT NULL DEFAULT 'https://www.w3schools.com/howto/img_avatar.png',
    CREATED_AT  timestamp(0) with time zone NOT NULL DEFAULT now()
);

CREATE TABLE IF NOT EXISTS phone
(
    id         bigserial   NOT NULL PRIMARY KEY,
    student_id int         NOT NULL REFERENCES students (id),
    phone      varchar(50) NOT NULL
)