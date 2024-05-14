ALTER TABLE "students"
DROP CONSTRAINT unique_phone_st;

ALTER TABLE "students"
DROP CONSTRAINT unique_mail_st;

-- teachers part

ALTER TABLE "teachers"
DROP CONSTRAINT unique_phone_ts;

ALTER TABLE "teachers"
DROP CONSTRAINT unique_mail_ts;