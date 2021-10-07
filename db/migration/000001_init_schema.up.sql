CREATE TABLE users (
    id  serial PRIMARY KEY,
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
    organization_name text ,
    organization_id text ,
    organization_date date ,
    organization_address text ,
    position text,
    created_at  timestamptz NOT NULL DEFAULT (now()),
    updated_at timestamptz
);
CREATE TABLE user_images (
    id  serial PRIMARY KEY,
    user_id  int  NOT NULL,
    url text NOT NULL,
    type  int  NOT NULL
);

CREATE TABLE role (
   id  serial PRIMARY KEY,
   name text NOT NULL
);

CREATE TABLE user_role (
  id  serial PRIMARY KEY,
  user_id int NOT NULL,
  role_id int NOT NULL
);