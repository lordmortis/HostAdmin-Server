CREATE TABLE "users" (
    id UUID NOT NULL PRIMARY KEY,
    username varchar(40) UNIQUE NOT NULL,
    email varchar(40) NOT NULL,
    encrypted_password bytea NOT NULL,
    super_admin bool NOT NULL,
    created_at timestamp NOT NULL,
    updated_at timestamp NOT NULL
);