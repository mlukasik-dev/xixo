-- https://dba.stackexchange.com/questions/68266/what-is-the-best-way-to-store-an-email-address-in-postgresql
CREATE DOMAIN email AS text CHECK (
  value ~ '^[a-zA-Z0-9.!#$%&''*+/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$'
);

-- https://en.wikipedia.org/wiki/E.164
-- https://stackoverflow.com/questions/6478875/regular-expression-matching-e-164-formatted-phone-numbers
CREATE DOMAIN phone_number AS text CHECK (
  value ~ '^\+?[1-9]\d{1,14}$'
);

CREATE TABLE IF NOT EXISTS roles (
  PRIMARY KEY (role_id),
  role_id uuid DEFAULT gen_random_uuid(),
  -- if true role can be assigned only to the admin
  admin_only boolean NOT NULL DEFAULT FALSE,
  super_admin boolean,
  /* value can be either TRUE or NULL, see constraints */
  account_admin boolean,
  /* value can be either TRUE or NULL, see constraints */
  display_name varchar(50) UNIQUE NOT NULL,
  description text NOT NULL DEFAULT '',
  created_at timestamp NOT NULL DEFAULT NOW(),
  updated_at timestamp NOT NULL DEFAULT NOW(),
  -- Ensures that fields can be only TRUE or NULL
  CONSTRAINT super_admin_true_or_null CHECK (super_admin),
  CONSTRAINT account_admin_true_or_null CHECK (account_admin),
  -- Ensures that only one row in the table can have the field with true value
  CONSTRAINT super_admin_only_one_true UNIQUE (super_admin),
  CONSTRAINT account_admin_only_one_true UNIQUE (account_admin)
);
CREATE TABLE IF NOT EXISTS permissions (
  PRIMARY KEY (permission_id),
  permission_id uuid DEFAULT gen_random_uuid(),
  -- When role is deleted all associated permissions will be deleted
  role_id uuid NOT NULL REFERENCES roles ON DELETE CASCADE,
  method text NOT NULL,
  /* gRPC method */
  created_at timestamp NOT NULL DEFAULT NOW(),
  updated_at timestamp NOT NULL DEFAULT NOW(),
  -- No rows with same role_id and method
  CONSTRAINT permissions_unique_role_id_and_method UNIQUE (role_id, method)
);
CREATE TABLE IF NOT EXISTS admins (
  PRIMARY KEY (admin_id),
  admin_id uuid DEFAULT gen_random_uuid(),
  first_name varchar(35) NOT NULL,
  last_name varchar(35) NOT NULL,
  email text UNIQUE NOT NULL,
  password text,
  /* if NULL means admin was created, but not registered yet */
  created_at timestamp NOT NULL DEFAULT NOW(),
  updated_at timestamp NOT NULL DEFAULT NOW()
);
-- Bridge table
CREATE TABLE IF NOT EXISTS admins_roles (
  PRIMARY KEY (admin_id, role_id),
  admin_id uuid NOT NULL REFERENCES admins ON DELETE CASCADE,
  role_id uuid NOT NULL REFERENCES roles ON DELETE CASCADE
);
CREATE TABLE IF NOT EXISTS users (
  PRIMARY KEY (user_id),
  user_id uuid DEFAULT gen_random_uuid(),
  account_id uuid NOT NULL,
  /* reference to the accounts table from account-service */
  first_name varchar(35) NOT NULL,
  last_name varchar(35) NOT NULL,
  email text NOT NULL,
  password varchar(255),
  /* if NULL means user was created, but not registered yet */
  phone_number phone_number,
  /* max digits is 15 all spaces must be removed */
  created_at timestamp NOT NULL DEFAULT NOW(),
  updated_at timestamp NOT NULL DEFAULT NOW(),
  -- Ensures that emails are unique inside of an account
  CONSTRAINT users_unique_email_and_account_id UNIQUE (email, account_id)
);
-- Bridge table
CREATE TABLE IF NOT EXISTS users_roles (
  PRIMARY KEY (user_id, role_id),
  user_id uuid NOT NULL REFERENCES users ON DELETE CASCADE,
  role_id uuid NOT NULL REFERENCES roles ON DELETE CASCADE
);