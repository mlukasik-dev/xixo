INSERT INTO roles(
    admin_only,
    super_admin,
    account_admin,
    display_name,
    description
  )
VALUES (
    TRUE,
    TRUE,
    NULL,
    'Super Admin',
    'Has all the permission to do everything'
  ),
  (
    FALSE,
    NULL,
    TRUE,
    'Account Admin',
    'Has all the permissions to manage the xixo-app'
  ) ON CONFLICT DO NOTHING;
-- Seed admins
INSERT INTO admins(first_name, last_name, email)
VALUES ('Martin', 'Lukasik', 'martilukas7@gmail.com') ON CONFLICT DO NOTHING;
-- Add Super Admin role to the first admin
INSERT INTO admins_roles(admin_id, role_id)
VALUES (
    (
      SELECT admin_id
      FROM admins
      WHERE email = 'martilukas7@gmail.com'
    ),
    (
      SELECT role_id
      FROM roles
      WHERE super_admin = TRUE
    )
  ) ON CONFLICT DO NOTHING;