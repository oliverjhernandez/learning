CREATE TABLE IF NOT EXISTS users (
  id bigserial PRIMARY KEY,
  created_at TIMESTAMP(0) with time zone NOT NULL DEFAULT NOW(),
  updated_at TIMESTAMP(0) with time zone NOT NULL DEFAULT NOW(),
  first_name text NOT NULL,
  last_name text NOT NULL,
  email citext UNIQUE NOT NULL,
  password_hash bytea NOT NULL,
  version integer NOT NULL DEFAULT 1,
  activated bool NOT NULL
);
