-- +goose Up
-- +goose StatementBegin
CREATE TABLE chats (
    id SERIAL PRIMARY KEY,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL,
    deleted_at TIMESTAMP,
    message VARCHAR(255),
    user_id INTEGER,
    chat_room_id INTEGER,
    CONSTRAINT chats_user_id_foreign 
        FOREIGN KEY (user_id) 
        REFERENCES users (id) 
        ON UPDATE CASCADE 
        ON DELETE NO ACTION,
    CONSTRAINT chats_chat_room_id_foreign 
        FOREIGN KEY (chat_room_id) 
        REFERENCES chat_rooms (id) 
        ON UPDATE CASCADE 
        ON DELETE NO ACTION
);

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE chats;
-- +goose StatementEnd
