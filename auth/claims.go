package auth

import (
	"errors"

	"github.com/golang-jwt/jwt/v5"
)

type SteamUserDetails struct {
	SteamId     string `json:"steamid"`
	PersonaName string `json:"personaname"`
	AvatarHash  string `json:"avatarhash"`
}

type UserDetails struct {
	UserId      int    `json:"user_id"`
	SteamId     string `json:"steamid"`
	PersonaName string `json:"personaname"`
	AvatarHash  string `json:"avatarhash"`
}

type Claims struct {
	UserId      int    `json:"user_id"`
	SteamId     string `json:"steamid"`
	PersonaName string `json:"personaname"`
	AvatarHash  string `json:"avatarhash"`
	jwt.RegisteredClaims
}

func (c Claims) Validate() error {
	if c.SteamId == "" {
		return errors.New("steamid cannot be empty")
	}

	if c.PersonaName == "" {
		return errors.New("personaname cannot be empty")
	}

	return nil
}
