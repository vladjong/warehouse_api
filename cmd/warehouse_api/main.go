package main

import (
	"context"
	"log"

	"github.com/adough/warehouse_api/internal/config"
)

func main() {
  ctx := context.Background()
  _, err := config.New(ctx)
  if err != nil {
    log.Fatal(err)
  }
}
