-- Delete foreing keys of independent tables
ALTER TABLE "social_links" DROP CONSTRAINT "social_links_user_id_fkey";
ALTER TABLE "user_preference" DROP CONSTRAINT "user_preference_user_id_fkey";

-- Delete user tables
DROP TABLE "social_links";
DROP TABLE "user_preference";
DROP TABLE "users";

-- Delete user types
DROP TYPE "auth_status";
DROP TYPE "subscription_status";
DROP TYPE "user_gender";