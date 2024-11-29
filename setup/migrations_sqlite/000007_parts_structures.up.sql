CREATE TABLE parts_structures (
    trace_id character varying(256) NOT NULL,
    parent_trace_id character varying(256) NOT NULL,
    deleted_at timestamp,
    created_at timestamp NOT NULL,
    created_user_id text NOT NULL,
    updated_at timestamp NOT NULL,
    updated_user_id text NOT NULL,
    PRIMARY KEY (trace_id, parent_trace_id),
    FOREIGN KEY (trace_id) REFERENCES parts(trace_id) ON UPDATE CASCADE ON DELETE CASCADE
);