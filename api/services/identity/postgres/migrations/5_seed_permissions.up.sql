INSERT INTO permissions(role_id, method) VALUES
	-- Users
	((SELECT role_id FROM roles WHERE super_admin = TRUE), '/xixo.identity.v1.Users/ListUsers'),
	((SELECT role_id FROM roles WHERE super_admin = TRUE), '/xixo.identity.v1.Users/GetUsersCount'),
	((SELECT role_id FROM roles WHERE super_admin = TRUE), '/xixo.identity.v1.Users/GetUser'),
	((SELECT role_id FROM roles WHERE super_admin = TRUE), '/xixo.identity.v1.Users/CreateUser'),
	((SELECT role_id FROM roles WHERE super_admin = TRUE), '/xixo.identity.v1.Users/UpdateUser'),
	((SELECT role_id FROM roles WHERE super_admin = TRUE), '/xixo.identity.v1.Users/DeleteUser'),
	-- Admins
	((SELECT role_id FROM roles WHERE super_admin = TRUE), '/xixo.identity.v1.Admins/ListAdmins'),
	((SELECT role_id FROM roles WHERE super_admin = TRUE), '/xixo.identity.v1.Admins/GetAdminsCount'),
	((SELECT role_id FROM roles WHERE super_admin = TRUE), '/xixo.identity.v1.Admins/GetAdmin'),
	((SELECT role_id FROM roles WHERE super_admin = TRUE), '/xixo.identity.v1.Admins/CreateAdmin'),
	((SELECT role_id FROM roles WHERE super_admin = TRUE), '/xixo.identity.v1.Admins/UpdateAdmin'),
	((SELECT role_id FROM roles WHERE super_admin = TRUE), '/xixo.identity.v1.Admins/DeleteAdmin'),
	-- Roles
	((SELECT role_id FROM roles WHERE super_admin = TRUE), '/xixo.identity.v1.Roles/ListRoles'),
	((SELECT role_id FROM roles WHERE super_admin = TRUE), '/xixo.identity.v1.Roles/GetRolesCount'),
	((SELECT role_id FROM roles WHERE super_admin = TRUE), '/xixo.identity.v1.Roles/GetRole'),
	((SELECT role_id FROM roles WHERE super_admin = TRUE), '/xixo.identity.v1.Roles/CreateRole'),
	((SELECT role_id FROM roles WHERE super_admin = TRUE), '/xixo.identity.v1.Roles/UpdateRole'),
	((SELECT role_id FROM roles WHERE super_admin = TRUE), '/xixo.identity.v1.Roles/DeleteRole');
	

INSERT INTO permissions(role_id, method) VALUES
	((SELECT role_id FROM roles WHERE account_admin = TRUE), '/xixo.identity.v1.Users/ListUsers'),
	((SELECT role_id FROM roles WHERE account_admin = TRUE), '/xixo.identity.v1.Users/GetUsersCount'),
	((SELECT role_id FROM roles WHERE account_admin = TRUE), '/xixo.identity.v1.Users/GetUser');