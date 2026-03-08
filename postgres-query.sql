-- name: GetUserGames :many
SELECT * FROM user_games
WHERE user_id = $1;

-- name: GetUserPublicGames :many
SELECT * FROM user_games
WHERE user_id = $1 AND is_hidden = FALSE AND is_private = FALSE;

-- name: GetUserInfo :one
SELECT * FROM users WHERE id = $1;

-- name: GetUserPrivateGames :many
SELECT * FROM user_games
WHERE user_id = $1 AND is_hidden = FALSE;

-- name: CreateUserPrivateGame :exec
INSERT INTO user_games (user_id, game_id, is_private)
VALUES ($1, $2, $3);

-- name: PotentiallyCreateUserAndGetUserID :one
INSERT INTO users (steamid, personaname, avatarhash)
VALUES ($1, $2, $3)
ON CONFLICT (steamid)
DO UPDATE
SET
    avatarhash = CASE
        WHEN users.avatarhash IS DISTINCT FROM EXCLUDED.avatarhash
        THEN EXCLUDED.avatarhash
        ELSE users.avatarhash
    END,
    personaname = CASE
        WHEN users.personaname IS DISTINCT FROM EXCLUDED.personaname
        THEN EXCLUDED.personaname
        ELSE users.personaname
    END
RETURNING id;

-- name: UpsertGames :batchexec
INSERT INTO games (
    id,
    name,
    released,
    background_image,
    rating
)
VALUES (
    $1, $2, $3, $4, $5
)
ON CONFLICT (id)
DO UPDATE
SET
    name             = EXCLUDED.name,
    released         = EXCLUDED.released,
    background_image = EXCLUDED.background_image,
    rating           = EXCLUDED.rating,
    updated_at       = CURRENT_TIMESTAMP
WHERE
    games.updated_at < CURRENT_TIMESTAMP - INTERVAL '7 days';