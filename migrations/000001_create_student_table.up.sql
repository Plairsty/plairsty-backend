CREATE TABLE IF NOT EXISTS students
(
    ID          int                         NOT NULL,
    FIRST_NAME  varchar(50)                 NOT NULL,
    MIDDLE_NAME varchar(50)                 NOT NULL,
    LAST_NAME   varchar(50)                 NOT NULL,
    EMAIL       varchar(50)                 NOT NULL,
    IMAGE_URL   varchar(50)                 NOT NULL,
    CREATED_AT  timestamp(0) with time zone NOT NULL DEFAULT now()
);

CREATE TABLE IF NOT EXISTS phone
(
    ID bigserial PRIMARY KEY,
    phone int NOT NULL,
    TYPE  VARCHAR(50) NOT NULL,
    FOREIGN KEY (ID) REFERENCES students (ID)
)