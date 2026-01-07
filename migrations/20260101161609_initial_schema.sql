-- Create "user_games" table
CREATE TABLE "public"."user_games" (
  "id" serial NOT NULL,
  "user_id" integer NOT NULL,
  "game_id" integer NOT NULL,
  "is_private" boolean NOT NULL DEFAULT false,
  "is_hidden" boolean NOT NULL DEFAULT false,
  PRIMARY KEY ("id")
);
-- Create index "idx_user_games_user_id" to table: "user_games"
CREATE INDEX "idx_user_games_user_id" ON "public"."user_games" ("user_id");
-- Create index "idx_user_games_user_id_hidden_false" to table: "user_games"
CREATE INDEX "idx_user_games_user_id_hidden_false" ON "public"."user_games" ("user_id") WHERE (is_hidden = false);
-- Create index "idx_user_games_user_id_hidden_private_false" to table: "user_games"
CREATE INDEX "idx_user_games_user_id_hidden_private_false" ON "public"."user_games" ("user_id") WHERE ((is_hidden = false) AND (is_private = false));
