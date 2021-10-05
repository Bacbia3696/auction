CREATE TABLE users (
    Id  serial PRIMARY KEY,
    UserName text NOT NULL,
    Password text NOT NULL,
    FullName text NOT NULL,
    Email text NOT NULL,
    Address text NOT NULL,
    Phone text NOT NULL,
    BirthDate date,
    IdCard text NOT NULL,
    IdCardAddress text NOT NULL,
    IdCardDate date NOT NULL,
    BankId text NOT NULL,
    BankOwner text NOT NULL,
    BankName text NOT NULL,
    Status integer NOT NULL,
    CreatedAt  timestamptz NOT NULL DEFAULT (now()),
    UpdatedAt timestamptz
);
CREATE TABLE user_images (
    Id  serial PRIMARY KEY,
    UserId  int  NOT NULL,
    Url text NOT NULL
);

CREATE TABLE role (
   Id  serial PRIMARY KEY,
   Name text NOT NULL
);

CREATE TABLE user_role (
  Id  serial PRIMARY KEY,
  UserId int NOT NULL,
  RoleId int NOT NULL
);