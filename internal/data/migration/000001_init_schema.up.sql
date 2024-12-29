CREATE TABLE "users" (
    "id" UUID PRIMARY KEY,
    "username" TEXT UNIQUE NOT NULL,
    "avatar" TEXT,
    "bio" TEXT,
    "created_at" TIMESTAMP DEFAULT (now()),
    "updated_at" TIMESTAMP DEFAULT (now())
);