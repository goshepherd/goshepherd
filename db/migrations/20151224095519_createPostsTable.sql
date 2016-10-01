
-- +goose Up
-- SQL in section 'Up' is executed when this migration is applied
CREATE TABLE posts (
  post_id integer NOT NULL PRIMARY KEY AUTOINCREMENT,
  title varchar(200) NOT NULL,
  content text NOT NULL,
  created_at timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  updated_at timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP
);

CREATE INDEX post_search ON posts(title, content);
CREATE INDEX post_updated ON posts(updated_at);

-- +goose Down
-- SQL section 'Down' is executed when this migration is rolled back
DROP TABLE posts;
