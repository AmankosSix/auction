CREATE TABLE users (
    id serial not null unique,
    name varchar(255) not null,
    password_hash varchar(255) not null,
    email varchar(255) not null,
    phone varchar(255) not null,
    registered_at timestamp not null,
    last_visit_at timestamp not null
);