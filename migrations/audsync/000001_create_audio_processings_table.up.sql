CREATE TYPE "processing_status" AS ENUM (
  'pending',
  'in_progress',
  'completed',
  'failed'
);

CREATE TABLE IF NOT EXISTS "audio_processings" (
  "id" serial PRIMARY KEY,
  "audio_id" varchar(255) NOT NULL,
  "act_id" varchar(255) NOT NULL,
  "album_id" varchar(255) NOT NULL,
  "song_id" varchar(255) NOT NULL,
  "file_path" varchar(255) NOT NULL,
  "name" varchar(255) NOT NULL,
  "act_name" varchar(255) NOT NULL,
  "album_name" varchar(255) NOT NULL,
  "cover_art_url" varchar(255),
  "status" processing_status NOT NULL,
  "processing_start_time" timestamp,
  "processing_end_time" timestamp,
  "error_message" text,
  "event_id" varchar(255) UNIQUE NOT NULL,
  "created_at" timestamp NOT NULL
);