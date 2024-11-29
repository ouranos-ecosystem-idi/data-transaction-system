CREATE TABLE cfp_certificates (
    cfp_id character varying(256) NOT NULL,
    id integer NOT NULL,
    cfp_certificate text NOT NULL,
    deleted_at timestamp,
    created_at timestamp NOT NULL,
    created_user_id text NOT NULL,
    updated_at timestamp NOT NULL,
    updated_user_id text NOT NULL,
    PRIMARY KEY (cfp_id, id)
);