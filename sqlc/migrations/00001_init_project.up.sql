CREATE TABLE "users" (
  "id" SERIAL PRIMARY KEY,
  "name" VARCHAR NOT NULL,
  "last_name" VARCHAR NOT NULL,
  "username" VARCHAR(50) UNIQUE NOT NULL,
  "password" VARCHAR NOT NULL,
  "email" VARCHAR UNIQUE NOT NULL,
  "verified" BOOLEAN DEFAULT false,
  "created_at" TIMESTAMP DEFAULT (now()),
  "last_updated" TIMESTAMP DEFAULT (now()),
  "plan_id" INTEGER
);
