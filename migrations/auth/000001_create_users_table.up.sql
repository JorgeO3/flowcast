CREATE TYPE "user_gender" AS ENUM (
  'male',
  'female',
  'non_binary',
  'other'
);

CREATE TYPE "subscription_status" AS ENUM (
  'active',
  'inactive',
  'suspended'
);

CREATE TYPE "auth_status" AS ENUM (
  'active',
  'inactive',
  'locked'
);

CREATE TABLE "users" (
  "id" serial PRIMARY KEY,
  "username" varchar,
  "full_name" varchar,
  "birth_date" date,
  "gender" user_gender,
  "email" varchar,
  "phone_number" varchar,
  "password" varchar(60),
  "status" auth_status,
  "subscription_status" subscription_status,
  "created_at" timestamp,
  "updated_at" timestamp
);

CREATE TABLE "user_preference" (
  "id" serial PRIMARY KEY,
  "user_id" int UNIQUE,
  "email_notifications" bool,
  "sms_notifications" bool
);

CREATE TABLE "social_links" (
  "id" serial PRIMARY KEY,
  "user_id" int,
  "platform" varchar(50),
  "url" varchar(255)
);

ALTER TABLE "user_preference" ADD FOREIGN KEY ("user_id") REFERENCES "users" ("id");

ALTER TABLE "social_links" ADD FOREIGN KEY ("user_id") REFERENCES "users" ("id");