package main

import (
	"fmt"
	"url-shortener/internal/config"
)

func main() {
	cfg := config.MustLoad()

	fmt.Printf("%+v\n", cfg)
}
