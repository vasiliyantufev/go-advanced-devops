CREATE TABLE IF NOT EXISTS metrics (
    id VARCHAR(256),
    mtype VARCHAR(10),
    value NUMERIC,
    delta BIGINT,
    hash  varchar,
    UNIQUE (id, mtype)
    );

CREATE UNIQUE INDEX IF NOT EXISTS id_mtype_index
    ON metrics (id, mtype)