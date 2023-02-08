package main

import (
	"context"
	"fmt"
	"os"

	"dagger.io/dagger"
// 	"github.com/digitalocean/godo"
)

func main() {
	ctx := context.Background()

	// initialize Dagger client
	client, err := dagger.Connect(ctx, dagger.WithLogOutput(os.Stdout))
	if err != nil {
		panic(err)
	}
	defer client.Close()

	contextDir := client.Host().Directory(".")

	ref, err := client.Container().Build(contextDir).WithEnvVariable("DIGITALOCEAN_TOKEN", os.Getenv(("DIGITALOCEAN_TOKEN"))).
		Publish(ctx, "registry.digitalocean.com/hivenetes/buzz:latest")

	if err != nil {
		panic(err)
	}
	fmt.Printf("Published image to: %s\n", ref)
}
