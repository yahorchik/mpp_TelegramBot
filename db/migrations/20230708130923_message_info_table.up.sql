BEGIN;
CREATE TABLE user_info
(
    user_id text primary key,
    user_nickname text,
    user_firstname text,
    user_lastname text
);
CREATE TABLE message_info
(
    user_id text,
    message_text text,
    message_date timestamptz,
    foreign key (user_id) REFERENCES user_info (user_id)
);
CREATE INDEX idx_message_info on message_info(user_id);
END;