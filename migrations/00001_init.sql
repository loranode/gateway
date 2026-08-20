-- +goose Up
CREATE TABLE nodes (
	num         INTEGER PRIMARY KEY,
	long_name   TEXT    NOT NULL DEFAULT '',
	short_name  TEXT    NOT NULL DEFAULT '',
	hw_model    TEXT    NOT NULL DEFAULT '',
	role        TEXT    NOT NULL DEFAULT '',
	snr         REAL    NOT NULL DEFAULT 0,
	rssi        INTEGER NOT NULL DEFAULT 0,
	hops_away   INTEGER,
	latitude    REAL,
	longitude   REAL,
	altitude    INTEGER,
	battery     INTEGER,
	voltage     REAL,
	via_mqtt    INTEGER NOT NULL DEFAULT 0,
	is_favorite INTEGER NOT NULL DEFAULT 0,
	last_heard  INTEGER,
	updated_at  INTEGER NOT NULL
);

CREATE TABLE messages (
	id        INTEGER PRIMARY KEY,
	from_num  INTEGER NOT NULL,
	from_id   TEXT    NOT NULL,
	to_num    INTEGER NOT NULL,
	channel   INTEGER NOT NULL,
	text      TEXT    NOT NULL,
	snr       REAL    NOT NULL DEFAULT 0,
	rssi      INTEGER NOT NULL DEFAULT 0,
	hops_away INTEGER,
	rx_time   INTEGER NOT NULL
);

-- +goose Down
DROP TABLE messages;
DROP TABLE nodes;
