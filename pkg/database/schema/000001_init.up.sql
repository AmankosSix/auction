CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE users (
    uuid uuid DEFAULT uuid_generate_v4(),
    name varchar(255) not null,
    password_hash varchar(255) not null,
    email varchar(255) not null,
    phone varchar(255) not null,
    registered_at timestamp not null,
    last_visit_at timestamp not null
);