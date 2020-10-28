CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE IF NOT EXISTS permissions (
  PRIMARY KEY (permission_id),
  permission_id uuid DEFAULT uuid_generate_v4(),
  role_id uuid NOT NULL, /* references role from identity-service */
  method varchar(255) NOT NULL, /* gRPC method */
  created_at timestamp NOT NULL DEFAULT NOW(),
  updated_at timestamp NOT NULL DEFAULT NOW(),
  CONSTRAINT permissions_unique_role_id_and_method UNIQUE(role_id, method)
);

CREATE TABLE IF NOT EXISTS accounts (
  PRIMARY KEY (account_id),
  account_id uuid DEFAULT uuid_generate_v4(),
  display_name varchar(50) UNIQUE NOT NULL, /* agency name or other unique identifier */
  created_at timestamp NOT NULL DEFAULT NOW(),
  updated_at timestamp NOT NULL DEFAULT NOW()
);
