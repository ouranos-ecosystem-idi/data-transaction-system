CREATE TABLE cfp_infomation (
    trace_id character varying(256) NOT NULL,
    deleted_at timestamp,
    created_at timestamp NOT NULL,
    created_user_id text NOT NULL,
    updated_at timestamp NOT NULL,
    updated_user_id text NOT NULL,
    cfp_id character varying(256) NOT NULL,
    ghg_emission numeric,
    ghg_declared_unit character varying(20),
    cfp_type character varying(20) NOT NULL,
    dqr_type character varying(256) NOT NULL,
    te_r numeric,
    ge_r numeric,
    ti_r numeric,
    PRIMARY KEY (trace_id, cfp_type),
    FOREIGN KEY (trace_id) REFERENCES parts(trace_id) ON UPDATE CASCADE ON DELETE CASCADE
);