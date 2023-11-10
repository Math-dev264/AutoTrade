CREATE DATABASE IF NOT EXISTS 'autotrade' USE 'autotrade';

DROP TABLE IF EXISTS users;
CREATE TABLE users
(
	id SERIAL PRIMARY KEY,
	name varchar(60) NOT NULL,
	email varchar(100) UNIQUE NOT NULL,
	password varchar(25) NOT NULL,
	created_at timestamp DEFAULT CURRENT_TIMESTAMP
);

DROP TABLE IF EXISTS authentication;
CREATE TABLE authentication (
	id_user int,
	token int,
	FOREIGN KEY (id_user) REFERENCES users (id)
);
