package main

import (
	"fmt"
	"log"

	"github.com/wincam/go-filter/settings"
)

func main() {
	err := settings.LoadConfig()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(settings.Config.Input.Directory)
}
