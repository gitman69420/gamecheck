`Ello there

# Tech Stack
### Go backend (Gin) `WIP`
### Postgres DB `WIP`
### Nuxt frontend `Zero`
### Caching with Valkey `Zero`
### OAuth with OpenID from Steam `Zero`
[Link](https://partner.steamgames.com/doc/webapi_overview/oauth)
### Games API from RAWG `WIP`
[Link](https://api.rawg.io/docs)


# Running in Local

Have the following deps installed:
<br>(I will later dockerize them)
- GO (1.25.5+)
- Postgres (18)
- Make

### DB essentials
Create a database called `app`, create a user role called `app_user` with password `app_password` and grant it access to the `app` database.

Switch to the `app` db, and run the queries in `./postgres-schema.sql`

Run `make dev`