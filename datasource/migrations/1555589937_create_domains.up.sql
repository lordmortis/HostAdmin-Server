CREATE TABLE "domains" (
    id UUID NOT NULL PRIMARY KEY,
    name varchar UNIQUE NOT NULL,
    created_at timestamp,
    updated_at timestamp
);