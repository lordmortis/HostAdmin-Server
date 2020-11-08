CREATE TABLE "user_domain" (
    user_id UUID NOT NULL references users(id),
    domain_id UUID NOT NULL references domains(id),
    PRIMARY KEY(user_id, domain_id),
    admin bool not null default false,
    email bool not null default false,
    created_at timestamp,
    updated_at timestamp
);