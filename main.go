package main

import (
	"context"
	"fmt"
	"os"

	dapr "github.com/dapr/go-sdk/client"
)

func main() {
	// start the Dapr client
	client, err := dapr.NewClient()
	if err != nil {
		panic(err)
	}
	defer client.Close()

	// read the token from env var
	token := os.Getenv("DAPR_TOKEN")

	// set the token on the client
	client.WithAuthToken(token)

	// call Dapr
	m, err := client.GetMetadata(context.Background())
	if err != nil {
		panic(err)
	}

	fmt.Println("Received metadata id: " + m.ID)
}
