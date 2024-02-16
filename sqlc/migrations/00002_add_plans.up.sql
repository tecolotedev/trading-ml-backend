CREATE TABLE "plans" (
  "id" SERIAL PRIMARY KEY,
  "name" varchar(30),
  "max_historical_bars" INTEGER,
  "max_symbols" INTEGER,
  "max_indicators_per_symbol" INTEGER,
  "max_models" INTEGER,
  "created_at" TIMESTAMP DEFAULT (now()),
  "last_updated" TIMESTAMP DEFAULT (now())
);

ALTER TABLE "users" ADD FOREIGN KEY ("plan_id") REFERENCES "plans" ("id");
