CREATE TABLE "user_domain" (
    user_id UUID NOT NULL references users(id),
    domain_id UUID NOT NULL references domains(id),
    PRIMARY KEY(user_id, domain_id),
    admin bool NOT NULL DEFAULT false,
    email bool NOT NULL DEFAULT false,
    created_at timestamp NOT NULL,
    updated_at timestamp NOT NULL
);