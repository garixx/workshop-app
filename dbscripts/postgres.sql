CREATE TABLE users
(
    id            SERIAL PRIMARY KEY,
    login         character varying(20)    NOT NULL UNIQUE,
    username      character varying(20)    NOT NULL,
    password_hash character varying(255)   NOT NULL,
    created_at    timestamp with time zone NOT NULL DEFAULT now()
);

CREATE TABLE tokens
(
    id         SERIAL PRIMARY KEY,
    login      character varying(20)    NOT NULL UNIQUE,
    token      character varying(20)    NOT NULL UNIQUE,
    created_at timestamp with time zone NOT NULL DEFAULT now(),
    expired_in integer                  NOT NULL DEFAULT 43200
);