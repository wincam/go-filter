package main

import (
	"log"

	"github.com/wincam/go-filter/settings"
)

func main() {
	err := settings.LoadConfig()
	if err != nil {
		log.Fatal(err)
	}
}
