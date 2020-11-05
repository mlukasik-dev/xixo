-- Droping in reverse order of creation
DROP TABLE IF EXISTS users_roles;
DROP TABLE IF EXISTS users;
DROP TABLE IF EXISTS admins_roles;
DROP TABLE IF EXISTS admins;
DROP TABLE IF EXISTS permissions;
DROP TABLE IF EXISTS roles;

DROP DOMAIN IF EXISTS email;
DROP EXTENSION IF EXISTS "uuid-ossp";
