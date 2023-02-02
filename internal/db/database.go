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

func ConnectDB() {
	pool, err := pgxpool.Connect(context.Background(), config.GetConfigDBServer())
	if err != nil {
		log.Error(err)
	}
	db = DB{pool: pool}
}

//func GetPool() *pgxpool.Pool {
//	return db.pool
//}

func Ping() error {
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()

	if err := db.pool.Ping(ctx); err != nil {
		log.Error(err)
		return err
	}

	return nil
}
