-- Modify "games" table
ALTER TABLE "public"."games" ADD COLUMN "updated_at" timestamptz NOT NULL DEFAULT CURRENT_TIMESTAMP;
