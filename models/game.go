package models

type Game struct {
	GameId   int    `json:"id"`
	GameName string `json:"name"`
}

func ListGames() []Game {
	return []Game{
		{GameId: 123, GameName: "Cyberpunk 2077"},
		{GameId: 234, GameName: "Borderlands 2"},
	}
}
