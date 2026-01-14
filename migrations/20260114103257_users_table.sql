-- Create "users" table
CREATE TABLE "public"."users" (
  "id" serial NOT NULL,
  "steamid" character varying(17) NOT NULL,
  "personaname" character varying(32) NOT NULL,
  "avatarhash" character(40) NOT NULL,
  "created_at" timestamptz NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY ("id"),
  CONSTRAINT "users_steamid_key" UNIQUE ("steamid")
);
-- Create index "idx_users_steamid" to table: "users"
CREATE INDEX "idx_users_steamid" ON "public"."users" ("steamid");
