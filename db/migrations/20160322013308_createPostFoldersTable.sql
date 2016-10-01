
-- +goose Up
-- SQL in section 'Up' is executed when this migration is applied
CREATE TABLE post_folders (
  post_id integer NOT NULL PRIMARY KEY AUTOINCREMENT,
  folder_id  integer NOT NULL,
  FOREIGN KEY(post_id) REFERENCES posts(post_id),
  FOREIGN KEY(folder_id) REFERENCES posts(folder_id)
);

-- +goose Down
-- SQL section 'Down' is executed when this migration is rolled back
DROP TABLE post_folders;