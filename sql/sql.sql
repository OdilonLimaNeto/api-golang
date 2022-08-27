CREATE DATABASE api;

DROP TABLE IF EXISTS users;

CREATE TABLE users(
id integer NOT NULL PRIMARY KEY,
name text NOT NULL,
nick text NOT NULL,
password text UNIQUE NOT NULL,
email text UNIQUE NOT NULL,
created_at timestamp NOT NULL DEFAULT NOW(),
updated_at timestamp,
deleted_at timestamp DEFAULT NULL
);

SELECT * FROM users;