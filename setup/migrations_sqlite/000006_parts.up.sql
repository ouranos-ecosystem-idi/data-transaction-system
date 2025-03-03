CREATE TABLE parts (
    trace_id character varying(256) NOT NULL,
    operator_id character varying(256) NOT NULL,
    plant_id character varying(256) NOT NULL,
    parts_name character varying(256),
    deleted_at timestamp,
    created_at timestamp NOT NULL,
    created_user_id text NOT NULL,
    updated_at timestamp NOT NULL,
    updated_user_id text NOT NULL,
    support_parts_name character varying(256),
    terminated_flag boolean DEFAULT false NOT NULL,
    amount_required double precision,
    amount_required_unit character varying(256),
    parts_label_name character varying(256),
    parts_add_info1 character varying(256),
    parts_add_info2 character varying(256),
    parts_add_info3 character varying(256),
    PRIMARY KEY (trace_id)
    FOREIGN KEY (operator_id, plant_id) REFERENCES plants(operator_id, plant_id) ON UPDATE CASCADE ON DELETE CASCADE
);