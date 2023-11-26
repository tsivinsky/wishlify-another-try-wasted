CREATE TABLE IF NOT EXISTS users (
  id bigserial primary key,
  email text,
  login text NOT NULL,
  created_at timestamp NOT NULL DEFAULT NOW(),
  updated_at timestamp NOT NULL DEFAULT NOW()
);

CREATE OR REPLACE FUNCTION users_changed() RETURNS TRIGGER
    LANGUAGE plpgsql
    AS $$
BEGIN
  NEW.updated_at := NOW();
  RETURN NEW;
END;
$$;

CREATE TRIGGER trigger_users_changed
  BEFORE UPDATE ON users
  FOR EACH ROW
  EXECUTE PROCEDURE users_changed();
