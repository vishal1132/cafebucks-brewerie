package main

import (
	"log"

	"github.com/vishal1132/cafebucks/config"
)

func main() {
	cfg, err := config.LoadEnv()
	if err != nil {
		log.Println(err)
	}
	l := config.DefaultLogger(cfg)
	if err := runserver(cfg, l); err != nil {
		l.Fatal().Err(err).Msg("Failed to run order service server")
	}
}
