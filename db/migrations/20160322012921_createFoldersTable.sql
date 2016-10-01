
-- +goose Up
-- SQL in section 'Up' is executed when this migration is applied
CREATE TABLE folders (
  folder_id integer NOT NULL PRIMARY KEY AUTOINCREMENT,
  folder_name integer NOT NULL,
  created_at timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  updated_at timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP
);

CREATE INDEX folder_name ON folders(folder_name);
CREATE INDEX folder_updated ON folders(updated_at);

-- +goose Down
-- SQL section 'Down' is executed when this migration is rolled back
DROP TABLE folders;