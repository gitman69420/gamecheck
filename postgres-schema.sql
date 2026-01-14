CREATE TYPE user_games_status_type AS ENUM ('not-started', 'started', 'completed');

CREATE TABLE user_games (
    id SERIAL PRIMARY KEY,
    user_id INTEGER NOT NULL,
    game_id INTEGER NOT NULL,
    is_private BOOLEAN NOT NULL DEFAULT FALSE,
    is_hidden  BOOLEAN NOT NULL DEFAULT FALSE,
    status user_games_status_type DEFAULT 'not-started'
);

CREATE INDEX idx_user_games_user_id ON user_games (user_id);

CREATE INDEX idx_user_games_user_id_hidden_false
    ON user_games (user_id)
    WHERE is_hidden = FALSE;

CREATE INDEX idx_user_games_user_id_hidden_private_false
    ON user_games (user_id)
    WHERE is_hidden = FALSE
    AND is_private = FALSE;

CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    steamid VARCHAR(17) NOT NULL UNIQUE,
    personaname VARCHAR(32) NOT NULL,
    avatarhash CHAR(40) NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

CREATE INDEX idx_users_steamid ON users(steamid);

-- Foreign key reference in `user_games`.`user_id`
ALTER TABLE user_games
ADD CONSTRAINT fk_user_games_userid
FOREIGN KEY (user_id)
REFERENCES users(id);

CREATE TABLE games (
    id INTEGER PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    released DATE,
    background_image TEXT,
    rating DECIMAL(3, 2)
);

-- Foreign key reference in `user_games`.`game_id`
ALTER TABLE user_games
ADD CONSTRAINT fk_user_games_gameid
FOREIGN KEY (game_id)
REFERENCES games(id);
