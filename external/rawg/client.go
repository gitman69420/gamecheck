package rawg

import (
	"net/http"

	rawgSdk "github.com/dimuska139/rawg-sdk-go/v3"
)

func NewClient(secret string) *rawgSdk.Client {
	config := rawgSdk.Config{
		ApiKey: secret,
		Rps:    3,
	}

	return rawgSdk.NewClient(http.DefaultClient, &config)
}
