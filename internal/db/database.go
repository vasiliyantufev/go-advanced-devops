package database

import (
	"context"
	"github.com/jackc/pgx/v4/pgxpool"
	log "github.com/sirupsen/logrus"
	"github.com/vasiliyantufev/go-advanced-devops/internal/config"
	"time"
)

var db DB

type DB struct {
	pool *pgxpool.Pool
}

func ConnectDB() error {
	pool, err := pgxpool.Connect(context.Background(), config.GetConfigDBServer())
	if err != nil {
		log.Error(err)
		return err
	}
	db = DB{pool: pool}
	return nil
}

func Ping() error {
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()

	if err := db.pool.Ping(ctx); err != nil {
		log.Error(err)
		return err
	}

	return nil
}

func CreateTables() {

	var metricsTable = `
        CREATE TABLE IF NOT EXISTS metrics(
			id    varchar(32) PRIMARY KEY,
			mtype varchar(32),
			delta int,
			value double precision,
			hash  varchar(32)
        )
  `
	result, err := db.pool.Exec(context.Background(), metricsTable)
	if err != nil {
		log.Error(err)
		return
	}
	log.Info(result.String() + " metrics")
}
