CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE roles
(
    uuid UUID DEFAULT uuid_generate_v4(),
    role VARCHAR(255) NOT NULL UNIQUE,
    PRIMARY KEY (uuid)
);

INSERT INTO roles (role) VALUES ('user'), ('staff'), ('owner');

CREATE TABLE users
(
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

CREATE TABLE staff
(
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

CREATE TABLE regions
(
    uuid   UUID DEFAULT uuid_generate_v4(),
    region VARCHAR(255) NOT NULL UNIQUE,
    PRIMARY KEY (uuid)
);

INSERT INTO regions(region) VALUES
    ('Aqmola'),
    ('Aqtobe'),
    ('Almaty'),
    ('Atyrau'),
    ('East Kazakhstan'),
    ('Zhambyl'),
    ('Qaraghandy'),
    ('Qostanay'),
    ('Qyzylorda'),
    ('Mangystau'),
    ('North Kazakhstan'),
    ('Pavlodar'),
    ('South Kazakhstan'),
    ('West Kazakhstan');

CREATE TABLE cities
(
    uuid UUID DEFAULT uuid_generate_v4(),
    city VARCHAR(255) NOT NULL,
    region_uuid UUID,
    PRIMARY KEY (uuid),
    FOREIGN KEY (region_uuid) REFERENCES regions (uuid) ON DELETE CASCADE
);

-- INSERT INTO cities (city, region_uuid)
--     (SELECT 'Kokshetau', uuid FROM regions WHERE region = 'Aqmola'),
--     SELECT 'Stepnogorsk', uuid FROM regions WHERE region = 'Aqmola',
--     SELECT 'Schuchinsk', uuid FROM regions WHERE region = 'Aqmola',
--     SELECT 'Astana', uuid FROM regions WHERE region = 'Aqmola',
--     SELECT 'Aktobe', uuid FROM regions WHERE region = 'Aqtobe',
--     SELECT 'Almaty', uuid FROM regions WHERE region = 'Almaty',
--     SELECT 'Taldykorgan', uuid FROM regions WHERE region = 'Almaty',
--     SELECT 'Atyrau', uuid FROM regions WHERE region = 'Atyrau',
--     SELECT 'Ridder', uuid FROM regions WHERE region = 'East Kazakhstan',
--     SELECT 'Semey', uuid FROM regions WHERE region = 'East Kazakhstan',
--     SELECT 'Oskemen', uuid FROM regions WHERE region = 'East Kazakhstan',
--     SELECT 'Taraz', uuid FROM regions WHERE region = 'Zhambyl',
--     SELECT 'Balqash', uuid FROM regions WHERE region = 'Qaraghandy',
--     SELECT 'Jezkazgan', uuid FROM regions WHERE region = 'Qaraghandy',
--     SELECT 'Karaganda', uuid FROM regions WHERE region = 'Qaraghandy',
--     SELECT 'Saran', uuid FROM regions WHERE region = 'Qaraghandy',
--     SELECT 'Temirtau', uuid FROM regions WHERE region = 'Qaraghandy',
--     SELECT 'Satpayev', uuid FROM regions WHERE region = 'Qaraghandy',
--     SELECT 'Shakhtinsk', uuid FROM regions WHERE region = 'Qaraghandy',
--     SELECT 'Arkalyk', uuid FROM regions WHERE region = 'Qostanay',
--     SELECT 'Kostanay', uuid FROM regions WHERE region = 'Qostanay',
--     SELECT 'Kyzylorda', uuid FROM regions WHERE region = 'Qyzylorda',
--     SELECT 'Aqtau', uuid FROM regions WHERE region = 'Mangystau',
--     SELECT 'Janaozen', uuid FROM regions WHERE region = 'Mangystau',
--     SELECT 'Petropavl', uuid FROM regions WHERE region = 'North Kazakhstan',
--     SELECT 'Pavlodar', uuid FROM regions WHERE region = 'Pavlodar',
--     SELECT 'Shymkent', uuid FROM regions WHERE region = 'South Kazakhstan',
--     SELECT 'Turkistan', uuid FROM regions WHERE region = 'South Kazakhstan',
--     SELECT 'Oral', uuid FROM regions WHERE region = 'West Kazakhstan';

CREATE TABLE companies
(
    uuid UUID DEFAULT uuid_generate_v4(),
    company VARCHAR(255) NOT NULL UNIQUE,
    PRIMARY KEY (uuid)
);

INSERT INTO companies(company) VALUES ('Toyota'), ('Volkswagen'), ('BMW');

CREATE TABLE models
(
    uuid UUID DEFAULT uuid_generate_v4(),
    model VARCHAR(255) NOT NULL,
    company_uuid UUID,
    PRIMARY KEY (uuid),
    FOREIGN KEY (company_uuid) REFERENCES companies (uuid) ON DELETE CASCADE
);

-- INSERT INTO models (model, company_uuid)
--     SELECT 'GR Supra', uuid FROM companies WHERE company = 'Toyota',
--     SELECT '4Runner', uuid FROM companies WHERE company = 'Toyota',
--     SELECT '86', uuid FROM companies WHERE company = 'Toyota',
--     SELECT 'Camry', uuid FROM companies WHERE company = 'Toyota',
--     SELECT 'Avalon', uuid FROM companies WHERE company = 'Toyota',
--     SELECT 'Tundra', uuid FROM companies WHERE company = 'Toyota',
--     SELECT 'MR2', uuid FROM companies WHERE company = 'Toyota',
--     SELECT 'Mirai', uuid FROM companies WHERE company = 'Toyota',
--     SELECT 'Prius', uuid FROM companies WHERE company = 'Toyota',
--     SELECT 'RAV4', uuid FROM companies WHERE company = 'Toyota',
--     SELECT 'Golf GTI', uuid FROM companies WHERE company = 'Volkswagen',
--     SELECT 'Golf R', uuid FROM companies WHERE company = 'Volkswagen',
--     SELECT 'Taos', uuid FROM companies WHERE company = 'Volkswagen',
--     SELECT 'Jetta', uuid FROM companies WHERE company = 'Volkswagen',
--     SELECT 'Atlas Cross Sport', uuid FROM companies WHERE company = 'Volkswagen',
--     SELECT 'Atlas', uuid FROM companies WHERE company = 'Volkswagen',
--     SELECT 'Arteon', uuid FROM companies WHERE company = 'Volkswagen',
--     SELECT '1-series', uuid FROM companies WHERE company = 'BMW',
--     SELECT '2-series', uuid FROM companies WHERE company = 'BMW',
--     SELECT '3-series', uuid FROM companies WHERE company = 'BMW',
--     SELECT '4-series', uuid FROM companies WHERE company = 'BMW',
--     SELECT '5-series', uuid FROM companies WHERE company = 'BMW',
--     SELECT '6-series', uuid FROM companies WHERE company = 'BMW',
--     SELECT '7-series', uuid FROM companies WHERE company = 'BMW',
--     SELECT '8-series', uuid FROM companies WHERE company = 'BMW',
--     SELECT 'x-series', uuid FROM companies WHERE company = 'BMW',
--     SELECT 'i-series', uuid FROM companies WHERE company = 'BMW';