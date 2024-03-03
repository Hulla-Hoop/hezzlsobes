package main

import (
	"hezzl/internal/DB/psql"
	"hezzl/internal/DB/rediscash"
	"hezzl/internal/logger"

	"github.com/google/uuid"
	"github.com/joho/godotenv"
)

func main() {

	err := godotenv.Load(".env")
	if err != nil {
		logger.New().L.WithField("godotenv.Load", err).Error(err)
	}
	log := logger.New()
	db, err := psql.InitDb(log)
	if err != nil {
		log.L.WithField("psql.InitDb", err).Error(err)
	}
	r := rediscash.Init(db, log)

	/* reqID := uuid.New().String()
	_, err = r.Create(reqID, "Sakina", 1)
	if err != nil {
		log.L.WithField("rediscash.Create", err).Error(err)
	}
	reqID = uuid.New().String()
	e, err := r.Delete(reqID, 3)
	if err != nil {
		log.L.WithField("rediscash.Delete", err).Error(err)
	}
	log.L.WithField("rediscash.Delete", e).Info() */
	/* reqID := uuid.New().String()
	r.List(reqID, 4, 5) */
	reqID := uuid.New().String()
	r.Reprioritize(reqID, 5, 10)

}
