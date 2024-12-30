CREATE TABLE "users" (
    "id" UUID PRIMARY KEY,
    "username" TEXT UNIQUE NOT NULL,
    "title" VARCHAR(20),
    "status" VARCHAR(50),
    "avatar" TEXT,
    "bio" VARCHAR(255),
    "created_at" TIMESTAMP DEFAULT (now()),
    "updated_at" TIMESTAMP DEFAULT (now())
);