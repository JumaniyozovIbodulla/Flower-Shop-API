CREATE TYPE "languages" AS ENUM('uz','ru','en');
CREATE TYPE "order_status" AS ENUM('pending','confirmed', 'delivered', 'cancelled');
CREATE TYPE "media_type" AS ENUM ('image', 'video');

CREATE SEQUENCE "order_number_seq" START 10000;


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
    "name"          VARCHAR(255) NOT NULL UNIQUE, -- admin, manager
    "description"   TEXT DEFAULT '',
    "created_at"    TIMESTAMPTZ DEFAULT NOW()
);

CREATE TABLE IF NOT EXISTS "permissions" (
    "id"            UUID PRIMARY KEY DEFAULT GEN_RANDOM_UUID(),
    "name"          VARCHAR(255) NOT NULL UNIQUE, -- create_flower, delete_flower
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


CREATE TABLE IF NOT EXISTS "flowers" (
    "id"            UUID PRIMARY KEY DEFAULT GEN_RANDOM_UUID(),
    "title"         VARCHAR(255) NOT NULL,
    "description"   TEXT DEFAULT '',
    "price"         BIGINT NOT NULL CHECK ("price" >= 0),
    "stock"         INT NOT NULL CHECK ("stock" >= 0) DEFAULT 0,
    "is_active"     BOOLEAN NOT NULL DEFAULT TRUE,
    "created_at"    TIMESTAMPTZ DEFAULT NOW()
);

CREATE TABLE IF NOT EXISTS "flower_media" (
    "id"                    UUID PRIMARY KEY DEFAULT GEN_RANDOM_UUID(),
    "flower_id"             UUID NOT NULL REFERENCES "flowers"("id") ON DELETE CASCADE,
    "content_type"          media_type NOT NULL,
    "object_name"           TEXT NOT NULL,
    "created_at"            TIMESTAMPTZ DEFAULT NOW()
);

CREATE TABLE IF NOT EXISTS "carts" (
    "id"            UUID PRIMARY KEY DEFAULT GEN_RANDOM_UUID(),
    "customer_id"   UUID NOT NULL REFERENCES "users"("id"),
    "created_at"    TIMESTAMPTZ DEFAULT NOW()
);


CREATE TABLE IF NOT EXISTS "cart_items" (
    "id"        UUID PRIMARY KEY DEFAULT GEN_RANDOM_UUID(),
    "cart_id"   UUID NOT NULL REFERENCES "carts"("id") ON DELETE CASCADE,
    "flower_id" UUID NOT NULL REFERENCES "flowers"("id"),
    "quantity"  INT NOT NULL CHECK ("quantity" > 0),
    "price"     BIGINT NOT NULL CHECK("price" > 0)
);

CREATE TABLE IF NOT EXISTS "orders" (
    "id"            UUID PRIMARY KEY DEFAULT GEN_RANDOM_UUID(),
    "order_number"  BIGINT UNIQUE DEFAULT (nextval('order_number_seq') * 7919),
    "customer_id"   UUID NOT NULL REFERENCES "users"("id"),
    "total_price"   BIGINT NOT NULL CHECK("total_price" > 0),
    "status"        order_status NOT NULL DEFAULT 'pending',
    "created_at"    TIMESTAMPTZ DEFAULT NOW(),
    "updated_at"    TIMESTAMPTZ
);


CREATE TABLE IF NOT EXISTS "order_items" (
    "id"        UUID PRIMARY KEY DEFAULT GEN_RANDOM_UUID(),
    "order_id"  UUID NOT NULL REFERENCES "orders"("id") ON DELETE CASCADE,
    "flower_id" UUID NOT NULL REFERENCES "flowers"("id"),
    "quantity"  INT NOT NULL CHECK ("quantity" > 0),
    "price"     BIGINT NOT NULL CHECK("price" > 0)
);

