CREATE TABLE IF NOT EXISTS "subjects" (
  "id" UUID PRIMARY KEY,
  "name" VARCHAR(50),
  "type" VARCHAR(50),
  "created_at" TIMESTAMP NOT NULL DEFAULT NOW(),
  "updated_at" TIMESTAMP
);

CREATE TABLE IF NOT EXISTS "time_table" (
  "id" UUID PRIMARY KEY,
  "teacher_id" UUID NOT NULL REFERENCES "teachers" ("id"),
  "student_id" UUID NOT NULL REFERENCES "students" ("id"),
  "subject_id" UUID NOT NULL REFERENCES "subjects" ("id"),
  "from_date" TIMESTAMP NOT NULL,
  "to_date" TIMESTAMP NOT NULL
);