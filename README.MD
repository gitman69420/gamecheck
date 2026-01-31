`Ello there

# Tech Stack
## Go backend (Gin) `WIP`

## Postgres DB `WIP`

## Nuxt frontend `Zero`

## Caching with Valkey `WIP`
Valkey instance runs separate from the REST API

Only implemented for Games search functionality

## OAuth with OpenID from Steam `WIP`
Actual login with Steam not yet implemented

We're using [Steam's public API](https://developer.valvesoftware.com/wiki/Steam_Web_API) to get user info, for which we need a Steam API key.

#### Steam API key
We can get that key [here](https://steamcommunity.com/dev/apikey). Make sure you have a Steam account for this.

## Games API from RAWG `DONE`
It's documentation provided in this [link](https://api.rawg.io/docs).

To skip writing the translation layer and defining the types, I'm used this [package](https://github.com/dimuska139/rawg-sdk-go/).

#### RAWG API key
You'll need to log into RAWG and get an API key [here](https://rawg.io/apidocs). Click on "Get API Key".


# Running in Local

You'll need to install:
- `Make` to run the Makefile targets
- `Docker` to run containers
- `SQLC` to run the command locally (if you want to debug generated code)

Start up the containers by runninng `make dev`

You can bring them down with `make dev-down`

You'd also need both the [Steam API key](#steam-api-key) and [RAWG API key](#rawg-api-key) to get auth and searching games working respectively.