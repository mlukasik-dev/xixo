INSERT INTO permissions(role_id, method)
VALUES -- Users
	(
		(
			SELECT role_id
			FROM roles
			WHERE super_admin = TRUE
		),
		'/xixo.identity.v1.IdentityService/ListUsers'
	),
	(
		(
			SELECT role_id
			FROM roles
			WHERE super_admin = TRUE
		),
		'/xixo.identity.v1.IdentityService/GetUsersCount'
	),
	(
		(
			SELECT role_id
			FROM roles
			WHERE super_admin = TRUE
		),
		'/xixo.identity.v1.IdentityService/GetUser'
	),
	(
		(
			SELECT role_id
			FROM roles
			WHERE super_admin = TRUE
		),
		'/xixo.identity.v1.IdentityService/CreateUser'
	),
	(
		(
			SELECT role_id
			FROM roles
			WHERE super_admin = TRUE
		),
		'/xixo.identity.v1.IdentityService/UpdateUser'
	),
	(
		(
			SELECT role_id
			FROM roles
			WHERE super_admin = TRUE
		),
		'/xixo.identity.v1.IdentityService/DeleteUser'
	),
	-- Admins
	(
		(
			SELECT role_id
			FROM roles
			WHERE super_admin = TRUE
		),
		'/xixo.identity.v1.IdentityService/ListAdmins'
	),
	(
		(
			SELECT role_id
			FROM roles
			WHERE super_admin = TRUE
		),
		'/xixo.identity.v1.IdentityService/GetAdminsCount'
	),
	(
		(
			SELECT role_id
			FROM roles
			WHERE super_admin = TRUE
		),
		'/xixo.identity.v1.IdentityService/GetAdmin'
	),
	(
		(
			SELECT role_id
			FROM roles
			WHERE super_admin = TRUE
		),
		'/xixo.identity.v1.IdentityService/CreateAdmin'
	),
	(
		(
			SELECT role_id
			FROM roles
			WHERE super_admin = TRUE
		),
		'/xixo.identity.v1.IdentityService/UpdateAdmin'
	),
	(
		(
			SELECT role_id
			FROM roles
			WHERE super_admin = TRUE
		),
		'/xixo.identity.v1.IdentityService/DeleteAdmin'
	),
	-- Roles
	(
		(
			SELECT role_id
			FROM roles
			WHERE super_admin = TRUE
		),
		'/xixo.identity.v1.IdentityService/ListRoles'
	),
	(
		(
			SELECT role_id
			FROM roles
			WHERE super_admin = TRUE
		),
		'/xixo.identity.v1.IdentityService/GetRolesCount'
	),
	(
		(
			SELECT role_id
			FROM roles
			WHERE super_admin = TRUE
		),
		'/xixo.identity.v1.IdentityService/GetRole'
	),
	(
		(
			SELECT role_id
			FROM roles
			WHERE super_admin = TRUE
		),
		'/xixo.identity.v1.IdentityService/CreateRole'
	),
	(
		(
			SELECT role_id
			FROM roles
			WHERE super_admin = TRUE
		),
		'/xixo.identity.v1.IdentityService/UpdateRole'
	),
	(
		(
			SELECT role_id
			FROM roles
			WHERE super_admin = TRUE
		),
		'/xixo.identity.v1.IdentityService/DeleteRole'
	);
INSERT INTO permissions(role_id, method)
VALUES (
		(
			SELECT role_id
			FROM roles
			WHERE account_admin = TRUE
		),
		'/xixo.identity.v1.IdentityService/ListUsers'
	),
	(
		(
			SELECT role_id
			FROM roles
			WHERE account_admin = TRUE
		),
		'/xixo.identity.v1.IdentityService/GetUsersCount'
	),
	(
		(
			SELECT role_id
			FROM roles
			WHERE account_admin = TRUE
		),
		'/xixo.identity.v1.IdentityService/GetUser'
	);