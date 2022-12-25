-- +migrate Up
-- +migrate StatementBegin

-- USERS TABLE
CREATE SEQUENCE IF NOT EXISTS message_id_pkey_sec
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;

CREATE TABLE IF NOT EXISTS "messages" (
    "message_id" BIGINT DEFAULT nextval('message_id_pkey_sec'::regclass) NOT NULL PRIMARY KEY,
    "text" TEXT,
    "sender" BIGINT NOT NULL,
    "recipient" BIGINT NOT NULL,
    "roomID" VARCHAR(256) NOT NULL,
    "created_at" TIMESTAMP WITHOUT TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    "updated_at" TIMESTAMP WITHOUT TIME ZONE,
    "deleted_at" TIMESTAMP WITHOUT TIME ZONE,
    CONSTRAINT "FK__users_sender" FOREIGN KEY ("sender") REFERENCES "users" ("user_id")
    CONSTRAINT "FK__users_recipient" FOREIGN KEY ("recipient") REFERENCES "users" ("user_id")
);

-- +migrate StatementEnd
-- +migrate Down
