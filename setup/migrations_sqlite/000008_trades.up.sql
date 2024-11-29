CREATE TABLE trades (
    trade_id character varying(256) NOT NULL,
    downstream_operator_id character varying(256) NOT NULL,
    upstream_operator_id character varying(256) NOT NULL,
    downstream_trace_id character varying(256) NOT NULL,
    upstream_trace_id character varying(256),
    trade_date timestamp,
    deleted_at timestamp,
    created_at timestamp NOT NULL,
    created_user_id text NOT NULL,
    updated_at timestamp NOT NULL,
    updated_user_id text NOT NULL,
    PRIMARY KEY (trade_id),
    FOREIGN KEY (downstream_operator_id) REFERENCES operators(operator_id),
    FOREIGN KEY (downstream_trace_id) REFERENCES parts(trace_id),
    FOREIGN KEY (upstream_operator_id) REFERENCES operators(operator_id)
);