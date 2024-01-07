-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS carriers (
  id uuid NOT NULL,
  name character varying(255) NOT NULL,
  service character varying(255) NOT NULL,
  deadline INT NOT NULL,
  price DOUBLE PRECISION NOT NULL
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS carriers;
-- +goose StatementEnd
