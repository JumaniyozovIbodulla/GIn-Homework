ALTER TABLE "students"
ADD CONSTRAINT unique_phone_st UNIQUE (phone);

ALTER TABLE "students"
ADD CONSTRAINT unique_mail_st UNIQUE (mail);

-- teachers part

ALTER TABLE "teachers"
ADD CONSTRAINT unique_phone_ts UNIQUE (phone);

ALTER TABLE "teachers"
ADD CONSTRAINT unique_mail_ts UNIQUE (mail);