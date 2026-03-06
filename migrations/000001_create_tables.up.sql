CREATE TYPE "languages" AS ENUM('uz','ru','en');
CREATE TYPE "statuses" AS ENUM('pending','confirmed', 'delivered', 'cancelled');

CREATE TABLE IF NOT EXISTS "users" (
    "id"                UUID PRIMARY KEY DEFAULT GEN_RANDOM_UUID(),
    "name"              VARCHAR(255) NOT NULL,
    "email"             TEXT UNIQUE NOT NULL,
    "lang"              languages DEFAULT 'uz',
    "password_hash"     TEXT NOT NULL,
    "created_at"        TIMESTAMPTZ DEFAULT NOW()
);

CREATE TABLE IF NOT EXISTS "roles" (
    "id"            UUID PRIMARY KEY DEFAULT GEN_RANDOM_UUID(),
    "name"          VARCHAR(255) NOT NULL, -- admin, manager
    "description"   TEXT DEFAULT '',
    "created_at"    TIMESTAMPTZ DEFAULT NOW()
);

CREATE TABLE IF NOT EXISTS "permissions" (
    "id"            UUID PRIMARY KEY DEFAULT GEN_RANDOM_UUID(),
    "name"          VARCHAR(255) NOT NULL, -- create_flower, delete_flower
    "description"   TEXT,
    "created_at"    TIMESTAMPTZ DEFAULT NOW()
);

CREATE TABLE IF NOT EXISTS "role_permissions" (
    "role_id"           UUID NOT NULL REFERENCES "roles"("id") ON DELETE CASCADE,
    "permission_id"     UUID NOT NULL REFERENCES "permissions"("id") ON DELETE CASCADE,
    PRIMARY KEY ("role_id", "permission_id")
);

CREATE TABLE IF NOT EXISTS "user_roles" (
    "user_id"   UUID NOT NULL REFERENCES "users"("id") ON DELETE CASCADE,
    "role_id"   UUID NOT NULL REFERENCES "roles"("id") ON DELETE CASCADE,
    PRIMARY KEY ("user_id", "role_id")
);

