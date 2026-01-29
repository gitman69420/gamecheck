package auth

import (
	"errors"
	"gamecheck-backend/config"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var jwtSecret string

func LoadJwtSecret() string {
	if jwtSecret != "" {
		return jwtSecret
	}

	jwtSecret = os.Getenv(config.JWT_SECRET)
	return jwtSecret
}

func GenerateJWT(userDetails UserDetails, duration time.Duration) (string, time.Time, error) {
	expiryTime := time.Now().Add(duration)

	claims := Claims{
		UserId:      userDetails.UserId,
		SteamId:     userDetails.SteamId,
		PersonaName: userDetails.PersonaName,
		AvatarHash:  userDetails.AvatarHash,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expiryTime),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			NotBefore: jwt.NewNumericDate(time.Now()),
		},
	}

	secret := LoadJwtSecret()

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err := token.SignedString([]byte(secret))

	return signedToken, expiryTime, err

}

func VerifyJWT(tokenString string) (*Claims, error) {
	secret := LoadJwtSecret()

	token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(t *jwt.Token) (any, error) {
		return []byte(secret), nil
	})

	if err != nil {
		return nil, err
	}

	if !token.Valid {
		return nil, errors.New("Token is invalid or is expired")
	}

	claims, ok := token.Claims.(*Claims)

	if !ok {
		return nil, errors.New("Invalid claims from token")
	}

	return claims, nil
}
