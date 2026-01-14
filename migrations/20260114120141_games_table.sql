-- Create "games" table
CREATE TABLE "public"."games" (
  "id" integer NOT NULL,
  "name" character varying(255) NOT NULL,
  "released" date NULL,
  "background_image" text NULL,
  "rating" numeric(3,2) NULL,
  PRIMARY KEY ("id")
);
-- Modify "user_games" table
ALTER TABLE "public"."user_games" ADD CONSTRAINT "fk_user_games_gameid" FOREIGN KEY ("game_id") REFERENCES "public"."games" ("id") ON UPDATE NO ACTION ON DELETE NO ACTION;
