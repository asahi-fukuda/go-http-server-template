
-- +migrate Up
CREATE TABLE organizations (
  id                 VARCHAR(255) NOT NULL,
  `name`             VARCHAR(255) NOT NULL,
  created_at         DATETIME(6)  NOT NULL DEFAULT current_timestamp(6),
  updated_at         DATETIME(6)  NOT NULL DEFAULT current_timestamp(6) ON UPDATE current_timestamp(6),
  PRIMARY KEY (id)
);

-- +migrate Down
DROP TABLE organizations;