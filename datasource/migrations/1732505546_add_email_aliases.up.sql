CREATE TABLE "domain_email_alias" (
    domain_id UUID NOT NULL references domains(id),
    address varchar NOT NULL,
    PRIMARY KEY(domain_id, address),
    destinations text[] NOT NULL,
    created_at timestamp NOT NULL,
    updated_at timestamp NOT NULL
);