CREATE TABLE users (
    id bigserial PRIMARY KEY,
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
    updated_at timestamptz,
    CONSTRAINT UC_User UNIQUE (user_name,email,id_card)
);

CREATE TABLE user_images (
    id bigserial PRIMARY KEY,
    user_id bigint NOT NULL,
    url text NOT NULL,
    type int NOT NULL
);

CREATE TABLE ROLE (
    id bigserial PRIMARY KEY,
    name text NOT NULL
);

CREATE TABLE user_role (
    id bigserial PRIMARY KEY,
    user_id bigint NOT NULL,
    role_id bigint NOT NULL,
    UNIQUE (user_id)
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
    id bigserial PRIMARY KEY,
    title text NOT  NULL ,
    description text  NOT  NULL,
    code text NOT NULL,
    owner text NOT NULL,
    organization text NOT NULL,
    info text NOT NULL ,
    address text NOT  NULL ,
    register_start_date timestamptz NOT NULL,
    register_end_date timestamptz NOT NULL,
    bid_start_date timestamptz NOT NULL,
    bid_end_date timestamptz NOT NULL,
    start_price int NOT NULL,
    step_price int NOT NULL,
    status int NOT NULL,
    type int NOT NULL,
    updated_at timestamptz,
    created_at timestamptz NOT NULL DEFAULT (now()),
    UNIQUE (code)
);

CREATE TABLE auction_images (
    id bigserial PRIMARY KEY,
    auction_id bigint NOT NULL,
    url text NOT NULL
);

CREATE TABLE register_auction (
    id bigserial PRIMARY KEY,
    auction_id bigint NOT NULL,
    user_id bigint NOT NULL,
    status int NOT NULL,
    updated_at timestamptz,
    created_at timestamptz NOT NULL DEFAULT (now()),
    CONSTRAINT UC_Register_Auction UNIQUE (auction_id,user_id)
);

CREATE TABLE register_auction_images (
    id bigserial PRIMARY KEY,
    register_auction_id bigint NOT NULL,
    url text NOT NULL
);

CREATE TABLE bid (
    id bigserial PRIMARY KEY,
    auction_id bigint NOT NULL,
    user_id bigint NOT NULL,
    price int NOT NULL,
    status int NOT NULL,
    updated_at timestamptz NOT NULL DEFAULT (now()),
    created_at timestamptz NOT NULL DEFAULT (now())
);

