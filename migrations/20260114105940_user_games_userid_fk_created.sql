-- Modify "user_games" table
ALTER TABLE "public"."user_games" ADD CONSTRAINT "fk_user_games_userid" FOREIGN KEY ("user_id") REFERENCES "public"."users" ("id") ON UPDATE NO ACTION ON DELETE NO ACTION;
