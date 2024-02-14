CREATE TABLE "users" (
  "id" SERIAL PRIMARY KEY,
  "username" VARCHAR NOT NULL,
  "password" VARCHAR NOT NULL,
  "email" VARCHAR UNIQUE NOT NULL,
  "verfied" BOOLEAN DEFAULT false
  "created_at" TIMESTAMP DEFAULT (now())
);

CREATE TABLE "accounts" (
  "id" SERIAL PRIMARY KEY,
  "balance" FLOAT DEFAULT 0,
  "currency" VARCHAR NOT NULL,
  "created_at" TIMESTAMP DEFAULT (now()),
  "user_id" INTEGER
);

CREATE TABLE "transfers" (
  "id" SERIAL PRIMARY KEY,
  "amount" FLOAT NOT NULL,
  "reason" VARCHAR,
  "account_id" INTEGER,
  "created_at" TIMESTAMP NOT NULL
);

ALTER TABLE "accounts" ADD FOREIGN KEY ("user_id") REFERENCES "users" ("id");

ALTER TABLE "transfers" ADD FOREIGN KEY ("account_id") REFERENCES "accounts" ("id");
