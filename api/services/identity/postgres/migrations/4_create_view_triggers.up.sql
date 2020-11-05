/*
 * Roles and permissions related triggers
 */
-- INSTEAD OF INSERT TRIGGER
CREATE OR REPLACE FUNCTION create_role_with_permissions() RETURNS TRIGGER AS $$
DECLARE
	created_role_id admins.admin_id%TYPE;
	created_role_with_permissions roles_with_permissions%ROWTYPE;
	input_permission_method permissions.method%TYPE;
BEGIN
	INSERT INTO
		roles(admin_only, super_admin, account_admin, display_name, description)
	VALUES (
		NEW.admin_only, NEW.super_admin, NEW.account_admin, NEW.display_name,
		COALESCE(NEW.description, '')
	)
		RETURNING role_id INTO created_role_id;

	IF NEW.permissions IS NULL THEN
		SELECT * FROM roles_with_permissions INTO created_role_with_permissions
			WHERE role_id = created_role_id;
		RETURN created_role_with_permissions;
	END IF;

	FOREACH input_permission_method IN ARRAY NEW.permissions
	LOOP
		INSERT INTO permissions(role_id, method)
			VALUES (created_role_id, input_permission_method);
	END LOOP;

	SELECT * FROM roles_with_permissions INTO created_role_with_permissions
		WHERE role_id = created_role_id;
	RETURN created_role_with_permissions;
END;
$$ LANGUAGE PLPGSQL;

-- Deleting trigger for the sake of idempotency of the up migration
DROP TRIGGER IF EXISTS create_role_with_permissions ON roles_with_permissions;

CREATE TRIGGER create_role_with_permissions INSTEAD OF
INSERT ON roles_with_permissions
FOR EACH ROW EXECUTE PROCEDURE create_role_with_permissions();


-- INSTEAD OF UPDATE TRIGGER
CREATE OR REPLACE FUNCTION UPDATE_ROLE_WITH_PERMISSIONS() RETURNS TRIGGER AS $$
DECLARE
	old_permission_method permissions.method%TYPE;
	new_permission_method permissions.method%TYPE;
	updated_role_with_permissions roles_with_permissions%ROWTYPE;
BEGIN
	UPDATE roles SET
		admin_only = NEW.admin_only,
		display_name = NEW.display_name,
		description = NEW.description
	WHERE role_id = NEW.role_id;

	IF NEW.permissions = OLD.permissions THEN
		SELECT * FROM roles_with_permissions INTO updated_role_with_permissions
			WHERE role_id = NEW.role_id;
		RETURN updated_role_with_permissions;
	END IF;

	-- Grant new permissions
	FOREACH new_permission_method IN ARRAY NEW.permissions
	LOOP
		IF NOT new_permission_method = ANY (OLD.permissions) THEN
			INSERT INTO permissions(role_id, method) VALUES (NEW.role_id, new_permission_method);
		END IF;
	END LOOP;

	-- Deny permissions
	FOREACH old_permission_method IN ARRAY OLD.permissions
	LOOP
		IF NOT old_permission_method = ANY (NEW.permissions) THEN
			DELETE FROM permissions WHERE role_id = NEW.role_id AND method = old_permission_method;
		END IF;
	END LOOP;

	SELECT * FROM roles_with_permissions INTO updated_role_with_permissions
		WHERE role_id = NEW.role_id;
	RETURN updated_role_with_permissions;
END;
$$ LANGUAGE PLPGSQL;


-- Deleting trigger for the sake of idempotency of the up migration
DROP TRIGGER IF EXISTS UPDATE_ROLE_WITH_PERMISSIONS ON roles_with_permissions;

CREATE TRIGGER UPDATE_ROLE_WITH_PERMISSIONS INSTEAD OF
UPDATE ON roles_with_permissions
FOR EACH ROW EXECUTE PROCEDURE UPDATE_ROLE_WITH_PERMISSIONS();


/*
 * Admins related triggers
 */
-- INSTEAD OF INSERT TRIGGER
CREATE OR REPLACE FUNCTION create_admin_with_roles() RETURNS TRIGGER AS $$
DECLARE
	created_admin_id admins.admin_id%TYPE;
	created_admin_with_roles admins_with_roles%ROWTYPE;
	input_role_id roles.role_id%TYPE;
BEGIN
	ASSERT NEW.registered IS NULL, 'registered field is output only';
	INSERT INTO admins(first_name, last_name, email)
		VALUES (NEW.first_name, NEW.last_name, NEW.email)
			RETURNING admin_id INTO created_admin_id;

	IF NEW.roles IS NULL THEN
		SELECT * FROM admins_with_roles INTO created_admin_with_roles
			WHERE admin_id = created_admin_id;
		RETURN created_admin_with_roles;
	END IF;

	FOREACH input_role_id IN ARRAY NEW.roles
	LOOP
		INSERT INTO admins_roles(admin_id, role_id)
			VALUES (created_admin_id, input_role_id);
	END LOOP;

	SELECT * FROM admins_with_roles INTO created_admin_with_roles
		WHERE admin_id = created_admin_id;
	RETURN created_admin_with_roles;
END;
$$ LANGUAGE PLPGSQL;

-- Deleting trigger for the sake of idempotency of the up migration
DROP TRIGGER IF EXISTS create_admin_with_roles ON admins_with_roles;

CREATE TRIGGER create_admin_with_roles INSTEAD OF
INSERT ON admins_with_roles
FOR EACH ROW EXECUTE PROCEDURE create_admin_with_roles();


-- INSTEAD OF UPDATE TRIGGER
CREATE OR REPLACE FUNCTION update_admin_with_roles() RETURNS TRIGGER AS $$
DECLARE
	old_role_id roles.role_id%TYPE;
	new_role_id roles.role_id%TYPE;
	updated_admin_with_roles admins_with_roles%ROWTYPE;
BEGIN
	ASSERT NEW.registered = OLD.registered, 'registered field is output only';
	UPDATE admins SET
		first_name = NEW.first_name,
		last_name = NEW.last_name,
		email = NEW.email
	WHERE admin_id = NEW.admin_id;
	
	IF NEW.roles = OLD.roles THEN
		SELECT * FROM admins_with_roles INTO updated_admin_with_roles
			WHERE admin_id = NEW.admin_id;
		RETURN updated_admin_with_roles;
	END IF;
	
	-- Grant new roles 	
	FOREACH new_role_id IN ARRAY NEW.roles
	LOOP
		IF NOT new_role_id = ANY (OLD.roles) THEN
			INSERT INTO admins_roles(admin_id, role_id) VALUES (NEW.admin_id, new_role_id);
		END IF;
	END LOOP;
	
	-- Deny roles
	FOREACH old_role_id IN ARRAY OLD.roles
	LOOP
		IF NOT old_role_id = ANY (NEW.roles) THEN
			DELETE FROM admins_roles WHERE admin_id = NEW.admin_id AND role_id = old_role_id;
		END IF;
	END LOOP;
	
	SELECT * FROM admins_with_roles INTO updated_admin_with_roles
		WHERE admin_id = NEW.admin_id;
	RETURN updated_admin_with_roles;
END;
$$ LANGUAGE PLPGSQL;

-- Deleting trigger for the sake of idempotency of the up migration
DROP TRIGGER IF EXISTS update_admin_with_roles ON admins_with_roles;

CREATE TRIGGER update_admin_with_roles INSTEAD OF
UPDATE ON admins_with_roles
FOR EACH ROW EXECUTE PROCEDURE update_admin_with_roles();


/*
 * Users related triggers
 */
 -- INSTEAD OF INSERT TRIGGER
CREATE OR REPLACE FUNCTION create_user_with_roles() RETURNS TRIGGER AS $$
DECLARE
	created_user_id users.user_id%TYPE;
	created_user_with_roles users_with_roles%ROWTYPE;
	input_role_id roles.role_id%TYPE;
BEGIN
	ASSERT NEW.registered IS NULL, 'registered field is output only';
	INSERT INTO users(account_id, first_name, last_name, email, phone_number)
		VALUES (NEW.account_id, NEW.first_name, NEW.last_name, NEW.email, NEW.phone_number)
			RETURNING user_id INTO created_user_id;

	IF NEW.roles IS NULL THEN
		SELECT * FROM users_with_roles INTO created_user_with_roles
			WHERE user_id = created_user_id;
		RETURN created_user_with_roles;
	END IF;

	FOREACH input_role_id IN ARRAY NEW.roles
	LOOP
		INSERT INTO users_roles(user_id, role_id)
			VALUES (created_user_id, input_role_id);
	END LOOP;

	SELECT * FROM users_with_roles INTO created_user_with_roles
		WHERE user_id = created_user_id;
	RETURN created_user_with_roles;
END;
$$ LANGUAGE PLPGSQL;

-- Deleting trigger for the sake of idempotency of the up migration
DROP TRIGGER IF EXISTS create_user_with_roles ON users_with_roles;

CREATE TRIGGER create_user_with_roles INSTEAD OF
INSERT ON users_with_roles
FOR EACH ROW EXECUTE PROCEDURE create_user_with_roles();


-- INSTEAD OF UPDATE TRIGGER
CREATE OR REPLACE FUNCTION update_user_with_roles() RETURNS TRIGGER AS $$
DECLARE
	old_role_id roles.role_id%TYPE;
	new_role_id roles.role_id%TYPE;
	updated_user_with_roles users_with_roles%ROWTYPE;
BEGIN
	ASSERT NEW.registered = OLD.registered, 'registered field is output only';
	UPDATE users SET
		first_name = NEW.first_name,
		last_name = NEW.last_name,
		email = NEW.email,
		phone_number = NEW.phone_number
	WHERE user_id = NEW.user_id;
	
	IF NEW.roles = OLD.roles THEN
		SELECT * FROM users_with_roles INTO updated_user_with_roles
			WHERE user_id = NEW.user_id;
		RETURN updated_user_with_roles;
	END IF;
	
	-- Grant new roles 	
	FOREACH new_role_id IN ARRAY NEW.roles
	LOOP
		IF NOT new_role_id = ANY (OLD.roles) THEN
			INSERT INTO users_roles(user_id, role_id) VALUES (NEW.user_id, new_role_id);
		END IF;
	END LOOP;
	
	-- Deny roles
	FOREACH old_role_id IN ARRAY OLD.roles
	LOOP
		IF NOT old_role_id = ANY (NEW.roles) THEN
			DELETE FROM users_roles WHERE user_id = NEW.user_id AND role_id = old_role_id;
		END IF;
	END LOOP;
	
	SELECT * FROM users_with_roles INTO updated_user_with_roles
		WHERE user_id = NEW.user_id;
	RETURN updated_user_with_roles;
END;
$$ LANGUAGE PLPGSQL;

-- Deleting trigger for the sake of idempotency of the up migration
DROP TRIGGER IF EXISTS update_user_with_roles ON users_with_roles;

CREATE TRIGGER update_user_with_roles INSTEAD OF
UPDATE ON users_with_roles
FOR EACH ROW EXECUTE PROCEDURE update_user_with_roles();