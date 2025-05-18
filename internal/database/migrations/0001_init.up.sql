CREATE TABLE roles (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    created_at TIMESTAMPTZ
);

CREATE TABLE permissions (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    value VARCHAR(255) NOT NULL,
    created_at TIMESTAMPTZ
);

CREATE TABLE role_permissions (
    role_id INT NOT NULL REFERENCES roles(id) ON DELETE CASCADE,
    permission_id INT NOT NULL REFERENCES permissions(id) ON DELETE CASCADE,
    PRIMARY KEY (role_id, permission_id)
);

CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    username VARCHAR(255) NOT NULL UNIQUE,
    email VARCHAR(255) NOT NULL UNIQUE,
    password_hash TEXT NOT NULL,
    last_login TIMESTAMPTZ,
    is_active BOOLEAN NOT NULL DEFAULT TRUE,
    role_id INT NOT NULL REFERENCES roles(id),
    updated_at TIMESTAMPTZ,
    created_at TIMESTAMPTZ
);

INSERT INTO roles (name, description, created_at) VALUES
('Admin', 'Administrator role with full access', NOW()),
('User', 'Regular user role with limited access', NOW());

INSERT INTO permissions (name, description, value, created_at) VALUES
('View users', 'Permission to view users', 'users.view', NOW()),
('Edit users', 'Permission to edit users', 'users.edit', NOW()),
('Create users', 'Permission to create users', 'users.create', NOW()),

INSERT INTO role_permissions (role_id, permission_id) VALUES
(1, 1),
(1, 2),
(1, 3),
(2, 1);
