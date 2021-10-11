CREATE TABLE users (
    id serial PRIMARY KEY,
    user_name text NOT NULL,
    password text NOT NULL,
    full_name text NOT NULL,
    email text NOT NULL,
    address text NOT NULL,
    phone text NOT NULL,
    birthdate date,
    id_card text NOT NULL,
    id_card_address text NOT NULL,
    id_card_date date NOT NULL,
    bank_id text NOT NULL,
    bank_owner text NOT NULL,
    bank_name text NOT NULL,
    status integer NOT NULL,
    organization_name text,
    organization_id text,
    organization_date date,
    organization_address text,
    position text,
    created_at timestamptz NOT NULL DEFAULT (now()),
    updated_at timestamptz
);

CREATE TABLE user_images (
    id serial PRIMARY KEY,
    user_id int NOT NULL,
    url text NOT NULL,
    type int NOT NULL
);

CREATE TABLE ROLE (
    id serial PRIMARY KEY,
    name text NOT NULL
);

CREATE TABLE user_role (
    id serial PRIMARY KEY,
    user_id int NOT NULL,
    role_id int NOT NULL
);

CREATE OR REPLACE FUNCTION update_modified_column ()
    RETURNS TRIGGER
    AS $$
BEGIN
    NEW.updated_at = now();
    RETURN NEW;
END;
$$
LANGUAGE 'plpgsql';

CREATE TRIGGER update_user_updatetime
    BEFORE UPDATE ON users
    FOR EACH ROW
    EXECUTE PROCEDURE update_modified_column ();

CREATE TABLE auctions (
    id serial PRIMARY KEY,
    code text NOT NULL,
    owner text NOT NULL,
    organization text NOT NULL,
    register_start_date date NOT NULL,
    register_end_date date NOT NULL,
    bid_start_date date NOT NULL,
    bid_end_date date NOT NULL,
    start_price int NOT NULL,
    status int NOT NULL,
    updated_at timestamptz,
    created_at timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE auction_images (
    id serial PRIMARY KEY,
    auction_id int NOT NULL,
    url text NOT NULL
);

CREATE TABLE register_auction (
    id serial PRIMARY KEY,
    auction_id int NOT NULL,
    user_id int NOT NULL,
    status int NOT NULL,
    updated_at timestamptz,
    created_at timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE register_auction_images (
    id serial PRIMARY KEY,
    register_auction_id int NOT NULL,
    url text NOT NULL
);
