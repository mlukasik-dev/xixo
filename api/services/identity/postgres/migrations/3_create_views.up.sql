-- View for accessing roles together their with permissions
CREATE OR REPLACE VIEW roles_with_permissions AS 
SELECT
	r.role_id, r.admin_only, r.super_admin, r.account_admin, r.display_name, r.description, r.created_at, r.updated_at,
	ARRAY_REMOVE(ARRAY_AGG(permissions.method), NULL) AS permissions FROM roles r
LEFT JOIN permissions USING(role_id)
		GROUP BY (role_id);

-- View for accessing admins together with their roles
CREATE OR REPLACE VIEW admins_with_roles AS
SELECT a.admin_id, a.first_name, a.last_name, a.email, a.created_at, a.updated_at,
	a.password IS NOT NULL AS registered,
	ARRAY_REMOVE(ARRAY_AGG(ar.role_id), NULL) AS roles
FROM admins a
	LEFT JOIN admins_roles ar USING (admin_id) GROUP BY (a.admin_id);

-- View for accessing users together with their roles
CREATE OR REPLACE VIEW users_with_roles AS
SELECT
	u.user_id, u.account_id, u.first_name, u.last_name, u.email, u.phone_number, u.created_at, u.updated_at,
	u.password IS NOT NULL AS registered,
	ARRAY_REMOVE(ARRAY_AGG(ur.role_id), NULL) AS roles
FROM users u
	LEFT JOIN users_roles ur USING (user_id)
		GROUP BY (u.user_id);