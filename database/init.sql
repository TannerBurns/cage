CREATE DATABASE cage;
\c cage;
CREATE EXTENSION pgcrypto;
CREATE TABLE players (
    id SERIAL PRIMARY KEY,
    first VARCHAR(255) NOT NULL,
    last VARCHAR(255) NOT NULL,
    alias VARCHAR(255),
    email VARCHAR(255) NOT NULL,
    phone VARCHAR(64) NOT NULL,
    birthday DATE NOT NULL,
    joined DATE NOT NULL DEFAULT CURRENT_DATE,
    address VARCHAR(255) NOT NULL,
    city VARCHAR(255) NOT NULL,
    state VARCHAR(64) NOT NULL,
    zip VARCHAR(64) NOT NULL,
    region VARCHAR(255) NOT NULL,
    referral BOOLEAN NOT NULL,
    referredby VARCHAR(255),
    referredtype VARCHAR(255),
    banned BOOLEAN NOT NULL,
    notes TEXT,
    picture TEXT
);
CREATE TABLE employees (
    id SERIAL PRIMARY KEY,
    first VARCHAR(255) NOT NULL,
    last VARCHAR(255) NOT NULL,
    alias VARCHAR(255),
    email VARCHAR(255) NOT NULL,
    phone VARCHAR(64) NOT NULL,
    birthday DATE NOT NULL,
    joined DATE NOT NULL DEFAULT CURRENT_DATE,
    address VARCHAR(255),
    city VARCHAR(255) NOT NULL,
    state VARCHAR(64) NOT NULL,
    zip VARCHAR(64) NOT NULL,
    region VARCHAR(255) NOT NULL,
    notes TEXT,
    picture TEXT
);
CREATE TABLE logins (
    id SERIAL PRIMARY KEY,
    employee_id INTEGER NOT NULL,
    FOREIGN KEY (employee_id) REFERENCES employees (id) ON DELETE CASCADE,
    username TEXT NOT NULL,
    password TEXT NOT NULL
);
CREATE TABLE memberships (
    id SERIAL PRIMARY KEY,
    player_id INTEGER NOT NULL,
    FOREIGN KEY (player_id) REFERENCES players (id) ON DELETE CASCADE,
    employee_id INTEGER NOT NULL,
    FOREIGN KEY (employee_id) REFERENCES employees (id) ON DELETE CASCADE,
    created DATE NOT NULL DEFAULT CURRENT_DATE,
    active BOOLEAN NOT NULL,
    active_date DATE,
    deactive_date DATE,
    amount INTEGER,
    playtime INTEGER NOT NULL,
    notes TEXT
);
CREATE TABLE compensations (
    id SERIAL PRIMARY KEY,
    employee_id INTEGER NOT NULL,
    FOREIGN KEY (employee_id) REFERENCES employees (id),
    player_id INTEGER NOT NULL,
    FOREIGN KEY (player_id) REFERENCES players (id),
    type VARCHAR(255),
    amount NUMERIC(2),
    notes TEXT
);
CREATE TABLE membertransactions (
    id SERIAL PRIMARY KEY,
    membership_id INTEGER NOT NULL,
    player_id INTEGER NOT NULL,
    employee_id INTEGER NOT NULL,
    created DATE NOT NULL DEFAULT CURRENT_DATE,
    type VARCHAR(255) NOT NULL,
    amount INTEGER
);
CREATE TABLE games (
    id SERIAL PRIMARY KEY,
    created DATE NOT NULL DEFAULT CURRENT_DATE,
    game_id INTEGER NOT NULL,
    name VARCHAR(255) NOT NULL,
    category VARCHAR(255),
    max_players INTEGER,
    minimum NUMERIC,
    maximum NUMERIC,
    interval NUMERIC,
    rules TEXT,
    notes TEXT
);
CREATE TABLE playertransactions (
    id SERIAL PRIMARY KEY,
    player_id INTEGER NOT NULL,
    game_id INTEGER NOT NULL,
    created DATE NOT NULL DEFAULT CURRENT_DATE,
    time_played INTEGER NOT NULL
);

CREATE TABLE roles (
    id SERIAL PRIMARY KEY,
    employee_id INTEGER NOT NULL,
    FOREIGN KEY (employee_id) REFERENCES employees (id) ON DELETE CASCADE,
    role VARCHAR(255),
    role_id INTEGER,
    notes TEXT
);
CREATE TABLE permissions (
    id SERIAL PRIMARY KEY,
    role_id INTEGER NOT NULL,
    access_id INTEGER NOT NULL
);
CREATE TABLE house (
    id SERIAL PRIMARY KEY,
    created DATE NOT NULL DEFAULT CURRENT_DATE,
    amount NUMERIC,
    type VARCHAR(255),
    total NUMERIC
);
INSERT INTO permissions (role_id, access_id) VALUES(1, 1);
INSERT INTO permissions (role_id, access_id) VALUES(2, 2);
INSERT INTO permissions (role_id, access_id) VALUES(3, 3);
with alogin as (
    INSERT INTO employees (first, 
                            last, 
                            alias, 
                            email, 
                            phone, 
                            birthday, 
                            address, 
                            city, 
                            state, 
                            zip, 
                            region, 
                            notes, 
                            picture)
    VALUES('Admin', 
            'Admin', 
            '', 
            'admin@pokeraustin.com', 
            '1111111111', 
            '1990-02-01',
            '1111 Somewhere Lane', 
            'Austin', 
            'Texas', 
            78751, 
            'United States', 
            '', 
            '')
    RETURNING id
)
INSERT INTO logins (employee_id, username, password)
VALUES ((select id from alogin), 'admin',  crypt('admin', gen_salt('bf')));
INSERT INTO roles (employee_id, role, role_id, notes) 
VALUES (1, 'Admin', 1, 'Admin account');