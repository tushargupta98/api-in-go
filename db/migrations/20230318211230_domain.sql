-- +goose Up
-- +goose StatementBegin
CREATE TABLE domain (
  id SERIAL PRIMARY KEY,
  label VARCHAR(100) NOT NULL,
  created_at TIMESTAMP NOT NULL DEFAULT NOW(),
  updated_at TIMESTAMP NOT NULL DEFAULT NOW()
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE domain;
-- +goose StatementEnd
