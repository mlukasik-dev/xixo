CREATE OR REPLACE FUNCTION updated_timestamp_func()
RETURNS TRIGGER AS $$
BEGIN
  NEW.updated_at = NOW();
  RETURN NEW;
END;
$$ LANGUAGE plpgsql;

-- CREATING TRIGGERS
DO $$
DECLARE
  t text;
BEGIN
  FOR t IN
    SELECT table_name FROM information_schema.columns WHERE column_name = 'updated_at'
  LOOP
    EXECUTE format('DROP TRIGGER IF EXISTS trigger_update_timestamp ON %I;
                    CREATE TRIGGER trigger_update_timestamp
                    BEFORE UPDATE ON %I
                    FOR EACH ROW EXECUTE PROCEDURE updated_timestamp_func()', t,t);
  END loop;
END;
$$ language plpgsql;
