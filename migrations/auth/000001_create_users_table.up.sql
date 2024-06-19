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
  "username" varchar UNIQUE NOT NULL,
  "full_name" varchar NOT NULL,
  "birth_date" date NOT NULL,
  "gender" user_gender NOT NULL,
  "email" varchar UNIQUE NOT NULL,
  "phone_number" varchar UNIQUE NOT NULL,
  "password" varchar(60) NOT NULL,
  "status" auth_status NOT NULL,
  "subscription_status" subscription_status NOT NULL,
  "created_at" timestamp NOT NULL,
  "updated_at" timestamp NOT NULL
);

CREATE TABLE "user_preference" (
  "id" serial PRIMARY KEY,
  "user_id" int UNIQUE NOT NULL,
  "email_notifications" bool NOT NULL,
  "sms_notifications" bool NOT NULL
);

CREATE TABLE "social_links" (
  "id" serial PRIMARY KEY,
  "user_id" int NOT NULL,
  "platform" varchar(50) NOT NULL,
  "url" varchar(255) NOT NULL
);

ALTER TABLE "user_preference" ADD FOREIGN KEY ("user_id") REFERENCES "users" ("id");

ALTER TABLE "social_links" ADD FOREIGN KEY ("user_id") REFERENCES "users" ("id");