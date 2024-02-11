-- Create pet store schema
CREATE SCHEMA IF NOT EXISTS PET_STORE

-- Create the 'users' table
CREATE TABLE IF NOT EXISTS PET_STORE.USERS (
    id UUID PRIMARY KEY NOT NULL,
    first_name VARCHAR(20) NOT NULL,
    last_name VARCHAR(20) NOT NULL,
    username VARCHAR(20) NOT NULL,
    email VARCHAR(255) NOT NULL,
    password VARCHAR(25) NOT NULL,
    phone VARCHAR(20) NOT NULL,
    user_status SMALLINT NOT NULL,
    deleted boolean NOT NULL DEFAULT false,
    created_at bigint NOT NULL,
    modified_at bigint NOT NULL,
    CONSTRAINT username_unique UNIQUE (username)
);  

-- Create an index on the username column
CREATE INDEX IF NOT EXISTS USERS_USERNAME_IDX ON PET_STORE.USERS (username);

-- Grant permissions to the 'bob2' user
GRANT ALL ON PET_STORE.USERS TO bob2;
