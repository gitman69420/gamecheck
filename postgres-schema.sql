CREATE TABLE user_games (
    id SERIAL PRIMARY KEY,
    user_id INTEGER NOT NULL,
    game_id INTEGER NOT NULL,
    is_private BOOLEAN NOT NULL DEFAULT FALSE,
    is_hidden  BOOLEAN NOT NULL DEFAULT FALSE
);

CREATE INDEX idx_user_games_user_id ON user_games (user_id);

CREATE INDEX idx_user_games_user_id_hidden_false
    ON user_games (user_id)
    WHERE is_hidden = FALSE;

CREATE INDEX idx_user_games_user_id_hidden_private_false
    ON user_games (user_id)
    WHERE is_hidden = FALSE
    AND is_private = FALSE;