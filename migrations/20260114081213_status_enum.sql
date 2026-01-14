-- Create enum type "user_games_status_type"
CREATE TYPE "public"."user_games_status_type" AS ENUM ('not-started', 'started', 'completed');
-- Modify "user_games" table
ALTER TABLE "public"."user_games" ADD COLUMN "status" "public"."user_games_status_type" NULL DEFAULT 'not-started';
