CREATE TABLE "todos" (
  "id" serial PRIMARY KEY,
  "title" varchar NOT NULL,
  "completed" boolean NOT NULL,
  "created_at_UTC" timestamptz NOT NULL DEFAULT (now()),
   "updated_at_UTC" timestamptz NOT NULL DEFAULT (now())
);