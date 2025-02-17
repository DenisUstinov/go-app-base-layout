package main

import (
	"fmt"

	"github.com/DenisUstinov/go-app-base-layout/pkg/config"
	"github.com/rs/zerolog/log"
)

func main() {
	c, err := config.New()
	if err != nil {
		log.Fatal().Err(err).Msg("config.New")
	}
	fmt.Println(c)
}
