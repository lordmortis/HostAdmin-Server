CREATE TABLE "domain_email_user" (
       domain_id UUID NOT NULL references domains(id),
       base_address varchar NOT NULL,
       PRIMARY KEY(domain_id, base_address),
       encrypted_password bytea,
       enabled bool NOT NULL,
       quota int NOT NULL,
       created_at timestamp NOT NULL,
       updated_at timestamp NOT NULL
);