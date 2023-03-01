CREATE TABLE users (
    id bigserial not null primary key,
    first_name varchar(255) not null,
    last_name varchar(255) not null,
    email varchar(255) not null unique,
    encrypted_password varchar(255) not null
);