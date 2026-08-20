-- +goose Up
CREATE TABLE callbacks (
	url        TEXT    PRIMARY KEY,
	created_at INTEGER NOT NULL
);

-- +goose Down
DROP TABLE callbacks;
