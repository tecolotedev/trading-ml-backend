CREATE TABLE "users" (
  "id" SERIAL PRIMARY KEY,
  "username" VARCHAR NOT NULL,
  "password" VARCHAR NOT NULL,
  "email" VARCHAR UNIQUE NOT NULL,
  "verified" BOOLEAN DEFAULT False,
  "created_at" TIMESTAMP DEFAULT (now())
);