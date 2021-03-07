BEGIN;

-- Generic Template =) ----------------------------------------------

/*
CREATE TABLE IF NOT EXISTS table_name (
  id SERIAL PRIMARY KEY,
  created_at timestamp with time zone,
  updated_at timestamp with time zone,
  deleted_at timestamp with time zone,
);
*/

-- Content ----------------------------------------------

CREATE TABLE IF NOT EXISTS content (
  id SERIAL PRIMARY KEY,
  created_at timestamp with time zone,
  updated_at timestamp with time zone,
  deleted_at timestamp with time zone,
  name text,
  description text,
  asset_path text
);

-- MiniSeries ----------------------------------------------

CREATE TABLE IF NOT EXISTS mini_series (
  id SERIAL PRIMARY KEY,
  created_at timestamp with time zone,
  updated_at timestamp with time zone,
  deleted_at timestamp with time zone,
  name text,
  description text,
  asset_path text,
  episode_number integer,
  content_id integer REFERENCES content(id)
);

COMMIT;
