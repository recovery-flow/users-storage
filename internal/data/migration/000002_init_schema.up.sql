CREATE TABLE "users" (
    "id" UUID PRIMARY KEY,
    "username" VARCHAR(20) UNIQUE NOT NULL,
    "title" VARCHAR(30),
    "status" VARCHAR(50),
    "avatar" TEXT,
    "bio" VARCHAR(255),
    "city" UUID,
    "created_at" TIMESTAMP DEFAULT (now()),
    "updated_at" TIMESTAMP DEFAULT (now())
);