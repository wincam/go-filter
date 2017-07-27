package main

import (
	"log"

	"github.com/wincam/go-filter/filter"
	"github.com/wincam/go-filter/settings"
)

func main() {
	err := settings.LoadConfig()
	if err != nil {
		log.Fatal(err)
	}

	for _, filterConfig := range settings.Config.FilterConfigs {
		err = filter.RunFilter(filterConfig)
		if err != nil {
			log.Print(err)
		}
	}
}
