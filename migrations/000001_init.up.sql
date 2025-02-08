CREATE TABLE IF NOT EXISTS tickets (
   id CHAR(8) PRIMARY KEY,
   chat_id BIGINT NOT NULL,
   topic VARCHAR(64) NOT NULL,
   closed BOOLEAN DEFAULT FALSE,
   created_at TIMESTAMP WITH TIME ZONE
);
