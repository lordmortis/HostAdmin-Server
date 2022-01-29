CREATE TABLE "domains" (
    id UUID NOT NULL PRIMARY KEY,
    name varchar UNIQUE NOT NULL,
    created_at timestamp NOT NULL,
    updated_at timestamp NOT NULL
);