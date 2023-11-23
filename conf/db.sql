-- Table Createion
CREATE TABLE IF NOT EXISTS permissions (
  id BIGINT NOT NULL PRIMARY KEY,
  name VARCHAR(50) NOT NULL UNIQUE,
  description VARCHAR(100),
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP(),
  updated_at DATETIME DEFAULT CURRENT_TIMESTAMP() ON UPDATE CURRENT_TIMESTAMP()
);

CREATE TABLE IF NOT EXISTS roles (
  id BIGINT NOT NULL PRIMARY KEY,
  name VARCHAR(50) NOT NULL UNIQUE,
  description VARCHAR(100),
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  updated_at DATETIME DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  deleted_at DATETIME DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS users (
  id BIGINT NOT NULL PRIMARY KEY,
  name VARCHAR(80) NOT NULL,
  email VARCHAR(80) NOT NULL UNIQUE,
  password VARCHAR(80) NOT NULL,
  created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  updated_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  deleted_at DATETIME
);
--Table Association
CREATE TABLE IF NOT EXISTS role_permission (
  permission_id BIGINT NOT NULL REFERENCES permissions (id) ON UPDATE CASCADE ON DELETE CASCADE,
  role_id BIGINT NOT NULL REFERENCES roles (id) ON UPDATE CASCADE ON DELETE CASCADE
);

CREATE TABLE IF NOT EXISTS role_user (
  role_id BIGINT NOT NULL REFERENCES roles (id) ON UPDATE CASCADE ON DELETE CASCADE,
  user_id BIGINT NOT NULL REFERENCES users (id) ON UPDATE CASCADE ON DELETE CASCADE
);

--Data insertion
INSERT INTO users (id, name, email, password) VALUES
(1,	'Admin',	'admin@admin.com',	'$2a$06$AKTSnF57WrOB39S8KaoUbeYdLNiGVM1Rfl7nU05cEC1u12MMM2nsq');
INSERT INTO users (id, name, email, password) VALUES
(2,	'Dinesh',	'dinesh@admin.com',	'$2a$06$CNzEcDlJyHqKcQnWlNxgxOURwCPrDTGR3hfEcMAzDimSxJCwWbbQW');
INSERT INTO users (id, name, email, password) VALUES
(3,	'Shami',	'shami@admin.com',	'$2a$06$CdMBOO11ODnTYnQjtnxloOKdsSD.05elMaD2hce/nxDBmJ4t.IHlS');


INSERT INTO permissions (id, name, description) VALUES
(1,	'view user',	'Can view a user'),
(2,	'create user',	'Can create user'),
(3,	'edit user',	'Can edit user'),
(4,	'delete user',	'Can delete user'),
(5,	'view role',	'Can view role'),
(6,	'create role',	'Can create role'),
(7,	'edit role',	'Can edit role'),
(8,	'delete role',	'Can delete role'),
(9,	'view permission',	'Can view permission'),
(10,	'create permission',	'Can create permission'),
(11,	'edit permission',	'Can edit permission'),
(12,	'delete permission',	'Can delete permission'),
(13,	'give permission to role',	'Can give permission to role'),
(14,	'remove permission to role',	'Can remove permission to role'),
(15,	'sync permission to role',	'Can sync permission to role'),
(17,	'get permissions by role id',	'Can get all permissions by role ID'),
(18,	'get permissions by role name',	'Can get all permissions by role name'),
(19,	'assign role by user id',	'Can assign role to a user'),
(20,	'sync role by user id',	'Can sync user role'),
(21,	'get role by user id',	'Can get user role');

INSERT INTO roles (id, name, description) VALUES (1,	'Normal',	'User have this role, can view only');
INSERT INTO roles (id, name, description) VALUES (2,	'RW',	'User have this role, can perform Read & Write operation');
INSERT INTO roles (id, name, description) VALUES (3,	'RWD',	'User have this role, can perform Read, Write and Delete operation');
INSERT INTO roles (id, name, description) VALUES (4,	'Admin',	'Admin of the system');

INSERT INTO role_user (role_id, user_id) VALUES (1,	1);
INSERT INTO role_user (role_id, user_id) VALUES (2,	2);
INSERT INTO role_permission ( permission_id,role_id) VALUES (1,	1);
INSERT INTO role_permission (permission_id,role_id) VALUES (2,	1);
INSERT INTO role_permission (permission_id,role_id) VALUES (1,	2);
INSERT INTO role_permission (permission_id,role_id) VALUES (2,	2);
-- Workaround to fix primary key out of sync
--SELECT setval('users_id_seq', (SELECT MAX(id) FROM users)+1);
--SELECT setval('roles_id_seq', (SELECT MAX(id) FROM roles)+1);
--SELECT setval('permissions_id_seq', (SELECT MAX(id) FROM permissions)+1);