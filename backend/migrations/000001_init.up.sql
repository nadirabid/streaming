BEGIN;

-- Video ----------------------------------------------

CREATE TABLE IF NOT EXISTS videos (
  id SERIAL PRIMARY KEY,
  created_at timestamp with time zone,
  updated_at timestamp with time zone,
  deleted_at timestamp with time zone,
);

COMMIT;
