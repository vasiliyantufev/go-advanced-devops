package database

import (
	"context"
	"github.com/jackc/pgx/v4/pgxpool"
	log "github.com/sirupsen/logrus"
	"github.com/vasiliyantufev/go-advanced-devops/internal/config"
)

var db DB

type DB struct {
	pool *pgxpool.Pool
}

// Подключение к базе данных
func ConnectDB() {
	pool, err := pgxpool.Connect(context.Background(), config.GetConfigDBServer())
	if err != nil {
		log.Error(err)
	}
	db = DB{pool: pool}
}

func GetPool() *pgxpool.Pool {
	return db.pool
}
