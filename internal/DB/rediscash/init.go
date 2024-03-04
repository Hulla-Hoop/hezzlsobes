package rediscash

import (
	"context"
	"hezzl/internal/DB/psql"
	"hezzl/internal/config"
	"hezzl/internal/logger"
	"strconv"

	"github.com/redis/go-redis/v9"
)

type Cash struct {
	r   *redis.Client
	db  *psql.SqlPostgres
	log *logger.Logger
}

func Init(db *psql.SqlPostgres, l *logger.Logger) *Cash {
	c := config.NewCfgRedis()
	l.L.WithField("redis.Init", "").Info(c)
	cDB, err := strconv.Atoi(c.DB)
	if err != nil {
		l.L.WithField("Redis.Init", err).Info("Невозможно конвертировать конфиг")
	}

	d := redis.NewClient(&redis.Options{
		Addr:     c.HOST + ":" + c.PORT,
		Password: c.PASSWORD,
		DB:       cDB,
	})

	s, err := d.Ping(context.Background()).Result()
	if err != nil {
		l.L.WithField("redis.Init", err).Error()
	}

	l.L.WithField("redis.Init", s).Info()

	return &Cash{
		r:   d,
		db:  db,
		log: l,
	}
}

func (c *Cash) DelVal(id int) error {
	key := strconv.Itoa(id)
	ctx := context.Background()
	err := c.r.Del(ctx, key).Err()
	if err != nil {
		return err
	}
	return nil
}

func (c *Cash) Close() {
	c.db.Close()
	c.r.Close()
}
