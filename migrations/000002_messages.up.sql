BEGIN;

CREATE TABLE messages (
    id CHAR(18) PRIMARY KEY,
    sender VARCHAR(32),
    from_support BOOLEAN NOT NULL DEFAULT FALSE,
    ticket_id CHAR(18) REFERENCES tickets (id),
    text TEXT,
    created_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP
);

COMMIT;
