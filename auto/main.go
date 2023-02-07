package main

import (
	"context"
	"path/filepath"

	"github.com/pulumi/pulumi/sdk/v3/go/auto"
)

func main() {
	ctx := context.Background()
	stack, err := auto.UpsertStackLocalSource(ctx, "test", filepath.Join("..", "stack"))
	if err != nil {
		panic(err)
	}
	stack.Up(ctx)
}
