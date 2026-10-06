package main

import (
	"context"
	"fmt"

	"dagger.io/dagger"
)

func main() {
	ctx := context.Background()

	if err := Build(ctx); err != nil {
		fmt.Println("Error: ", err)
		panic(err)
	}
}

func Build(ctx context.Context) error {
	client, err := dagger.Connect(ctx)
	if err != nil {
		return err
	}
	defer client.Close()

	python := client.Container().From("python:3.12.2-bookworm").
		WithDirectory("python", client.Host().Directory("py-files")).
		WithExec([]string{"python", "--version"})

	python = python.WithExec([]string{"python", "python/world.py"})

	_, err = python.Directory("output").Export(ctx, "output")

	if err != nil {
		return err
	}

	return nil
}