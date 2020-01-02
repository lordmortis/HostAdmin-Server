CREATE TABLE "users" (
    id UUID NOT NULL PRIMARY KEY,
    username varchar(40) UNIQUE NOT NULL,
    email varchar(40) NOT NULL,
    encrypted_password bytea,
    super_admin bool,
    created_at timestamp,
    updated_at timestamp
);