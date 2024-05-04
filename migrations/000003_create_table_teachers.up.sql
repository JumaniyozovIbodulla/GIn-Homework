CREATE TABLE IF NOT EXISTS "teachers" (
  "id" uuid PRIMARY KEY,
  "first_name" VARCHAR(50),
  "last_name" VARCHAR(50),
  "subject_id" uuid,
  "start_working" TIMESTAMP,
  "phone" VARCHAR(50),
  "mail" VARCHAR(50),
  "created_at" TIMESTAMP NOT NULL DEFAULT NOW(),
  "updated" TIMESTAMP
);