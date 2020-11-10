-- FUNCTION
CREATE OR REPLACE FUNCTION update_timestamp() RETURNS TRIGGER AS $$ BEGIN ASSERT NEW.updated_at = OLD.updated_at
  AND NEW.created_at = OLD.created_at,
  'created_at and updated_at columns are output only';
NEW.updated_at = NOW();
RETURN NEW;
END;
$$ LANGUAGE plpgsql;
-- CREATING TRIGGERS
DO $$
DECLARE t text;
BEGIN FOR t IN
SELECT table_name
FROM information_schema.columns
WHERE column_name = 'updated_at' LOOP EXECUTE format(
    'DROP TRIGGER IF EXISTS update_timestamp ON %I;

    CREATE TRIGGER update_timestamp
      BEFORE UPDATE ON %I
    FOR EACH ROW EXECUTE FUNCTION update_timestamp()',
    t,
    t
  );
END loop;
END;
$$ LANGUAGE plpgsql;