-- +goose Up
CREATE TABLE channels (
	idx        INTEGER PRIMARY KEY,
	name       TEXT    NOT NULL DEFAULT '',
	role       TEXT    NOT NULL DEFAULT '',
	updated_at INTEGER NOT NULL
);

-- +goose Down
DROP TABLE channels;
