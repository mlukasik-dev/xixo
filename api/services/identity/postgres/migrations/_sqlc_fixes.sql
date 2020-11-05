CREATE TABLE roles_with_permissions (
  role_id uuid NOT NULL,
  admin_only bool NOT NULL,
  display_name text NOT NULL,
  description text NOT NULL,
  created_at timestamp NOT NULL,
  updated_at timestamp NOT NULL,

  permissions text[] NOT NULL
);

CREATE TABLE admins_with_roles (
  admin_id uuid NOT NULL,
  first_name text NOT NULL,
  last_name text NOT NULL,
  email text NOT NULL,
  created_at timestamp NOT NULL,
  updated_at timestamp NOT NULL,

  registered bool NOT NULL,
  roles text[] NOT NULL
);

CREATE TABLE users_with_roles (
  user_id uuid NOT NULL,
  account_id uuid NOT NULL,
  first_name text NOT NULL,
  last_name text NOT NULL,
  email text NOT NULL,
  phone_number text,
  created_at timestamp NOT NULL,
  updated_at timestamp NOT NULL,

  registered bool NOT NULL,
  roles text[] NOT NULL
);