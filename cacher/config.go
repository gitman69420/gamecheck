package cacher

import (
	"context"
	"log"

	"github.com/valkey-io/valkey-go"
)

func NewClient(valkeyUrl string, valkeyPassword string) (valkey.Client, error) {
	client, err := valkey.NewClient(valkey.ClientOption{InitAddress: []string{valkeyUrl}, Password: valkeyPassword})
	if err != nil {
		return nil, err
	}

	ping := client.B().Ping().Message("ping").Build()
	result := client.Do(context.Background(), ping)

	if result.Error() != nil {
		return nil, result.Error()
	}

	res, err := result.ToString()
	if err != nil {
		return nil, err
	}

	log.Default().Println("Pinged Valkey with result: ", res)

	return client, nil
}
