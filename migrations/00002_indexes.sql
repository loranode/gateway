-- +goose Up
CREATE INDEX idx_messages_to_rx ON messages (to_num, rx_time);

-- +goose Down
DROP INDEX idx_messages_to_rx;
