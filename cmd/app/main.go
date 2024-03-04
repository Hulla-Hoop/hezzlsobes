package main

import (
	"hezzl/internal/logger"
	"hezzl/pkg/firstapp"
	"hezzl/pkg/secondapp"

	"github.com/joho/godotenv"
)

func main() {
	log := logger.New()
	err := godotenv.Load(".env")
	if err != nil {
		log.L.WithField("godotenv.Load", err).Error(err)
	}
	first := firstapp.New(log)

	second := secondapp.New(log)
	go func() {
		second.C.Read()
	}()

	first.Start()

}
