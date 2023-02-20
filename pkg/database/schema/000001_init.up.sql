CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE roles (
    uuid UUID DEFAULT uuid_generate_v4(),
    role VARCHAR(255) NOT NULL UNIQUE,
    PRIMARY KEY (uuid)
);

INSERT INTO roles (role) VALUES ('user'), ('admin');

CREATE TABLE users (
    uuid UUID DEFAULT uuid_generate_v4(),
    name VARCHAR(255) NOT NULL,
    password_hash VARCHAR(255) NOT NULL,
    email VARCHAR(255) NOT NULL UNIQUE,
    phone VARCHAR(255) NOT NULL,
    registered_at TIMESTAMP NOT NULL,
    last_visit_at TIMESTAMP NOT NULL,
    role_uuid UUID,
    PRIMARY KEY (uuid),
    FOREIGN KEY (role_uuid) REFERENCES roles (uuid) ON DELETE CASCADE
);