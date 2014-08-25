-- +goose Up
-- SQL in section 'Up' is executed when this migration is applied
CREATE TABLE goals (
  id int NOT NULL,
  name text,
  tallies int,
  PRIMARY KEY(id)
);

-- +goose Down
-- SQL section 'Down' is executed when this migration is rolled back
DROP TABLE goals;
