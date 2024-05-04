CREATE TABLE IF NOT EXISTS "teachers" (
  "id" UUID PRIMARY KEY,
  "first_name" VARCHAR(50),
  "last_name" VARCHAR(50),
  "subject_id" UUID,
  "start_working" TIMESTAMP,
  "phone" VARCHAR(50),
  "mail" VARCHAR(50),
  "created_at" TIMESTAMP NOT NULL DEFAULT NOW(),
  "updated_at" TIMESTAMP
);