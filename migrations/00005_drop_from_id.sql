-- +goose Up
ALTER TABLE messages DROP COLUMN from_id;

-- +goose Down
ALTER TABLE messages ADD COLUMN from_id TEXT NOT NULL DEFAULT '';
