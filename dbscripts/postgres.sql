CREATE TABLE users (
                       id SERIAL PRIMARY KEY,
                       login character varying(255) NOT NULL UNIQUE,
                       username character varying(255) NOT NULL,
                       password_hash character varying(255) NOT NULL,
                       created_at timestamp with time zone NOT NULL DEFAULT now()
);

--INSERT INTO users (id, login, username, password_hash)VALUES (1, 'testuser', 'me', 'a12345');