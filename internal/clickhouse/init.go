package clickhouse

import (
	"context"
	"fmt"
	"hezzl/internal/config"
	"hezzl/internal/logger"
	"hezzl/internal/model"
	"time"

	"github.com/ClickHouse/clickhouse-go/v2"
	"github.com/ClickHouse/clickhouse-go/v2/lib/driver"
)

type ClickHouse struct {
	log *logger.Logger
	cl  driver.Conn
}

func Init(log *logger.Logger) *ClickHouse {
	cfg := config.NewClick()

	conn, err := clickhouse.Open(&clickhouse.Options{
		Addr: []string{fmt.Sprintf("%s:%s", cfg.HOST, cfg.PORT)},
		Auth: clickhouse.Auth{
			Username: cfg.USER,
			Password: cfg.PASS,
		},
		Debug:           true,
		DialTimeout:     time.Second,
		MaxOpenConns:    10,
		MaxIdleConns:    5,
		ConnMaxLifetime: time.Hour,
	})
	if err != nil {
		log.L.WithField("ClickHouse.Init", err).Error(err)
		return nil
	}

	if err := conn.Ping(context.Background()); err != nil {
		if exception, ok := err.(*clickhouse.Exception); ok {
			fmt.Printf("[%d] %s \n%s\n", exception.Code, exception.Message, exception.StackTrace)
		} else {
			fmt.Println(err)
		}
	}
	return &ClickHouse{
		log: log,
		cl:  conn,
	}
}

func (c *ClickHouse) Create(Rows map[time.Time]model.LogGoods) error {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	query := "INSERT INTO log (id, name, description, project_id, removed,eventTime,priority)"
	batch, err := c.cl.PrepareBatch(ctx, query)
	if err != nil {
		c.log.L.WithField("ClickHouse.Create", "").Error(err)
		return err
	}
	for _, v := range Rows {
		err := batch.Append(
			v.ID,
			v.Name,
			v.Description,
			v.ProjectID,
			v.Removed,
			v.EventTime,
			v.Priority,
		)
		if err != nil {
			c.log.L.WithField("ClickHouse.Create", "").Error(err)
			return err
		}
	}

	return batch.Send()
}

func (c *ClickHouse) Close() {
	c.cl.Close()
}
