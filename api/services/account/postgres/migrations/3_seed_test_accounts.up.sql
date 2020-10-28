-- Seeding accounts
INSERT INTO accounts(display_name)
  VALUES ('First client') ON CONFLICT DO NOTHING;