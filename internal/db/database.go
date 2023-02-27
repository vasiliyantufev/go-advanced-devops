package database

import (
	"context"
	"database/sql"
	"time"

	"github.com/vasiliyantufev/go-advanced-devops/internal/config"
	"github.com/vasiliyantufev/go-advanced-devops/internal/storage"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	log "github.com/sirupsen/logrus"

	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/lib/pq"
)

//var db DB

type DB struct {
	pool *sql.DB
}

func NewDB(c *config.ConfigServer) (*DB, error) {
	pool, err := sql.Open("postgres", c.DSN)
	if err != nil {
		log.Error(err)
		return nil, err
	}

	ctx, cnl := context.WithTimeout(context.Background(), 10*time.Second)
	defer cnl()
	if err := pool.PingContext(ctx); err != nil {
		return nil, err
	}
	return &DB{pool: pool}, nil
}

func (db DB) Ping() error {
	if err := db.pool.Ping(); err != nil {
		log.Error(err)
		return err
	}
	return nil
}

func (db DB) Close() error {
	return db.pool.Close()
}

func (db DB) CreateTablesMigration() {

	driver, err := postgres.WithInstance(db.pool, &postgres.Config{})
	if err != nil {
		log.Error(err)
	}
	m, err := migrate.NewWithDatabaseInstance(
		"file://./migrations",
		"postgres", driver)
	if err != nil {
		log.Error(err)
	}
	if err = m.Up(); err != nil {
		log.Error(err)
	}
}

func (db DB) InsertOrUpdateMetrics(metrics *storage.MemStorage) error {

	stmt, err := db.pool.Prepare(`
			INSERT INTO metrics 
			VALUES($1, $2, $3, $4, $5)
			ON CONFLICT (id, mtype)
			DO UPDATE SET
				value = $3,
				delta = $4
			`)

	if err != nil {
		log.Error(err)
		return err
	}

	defer stmt.Close()

	for _, metric := range metrics.GetAllMetrics() {
		if _, err = stmt.Exec(metric.ID, metric.MType, metric.Value, metric.Delta, metric.Hash); err != nil {
			log.Error(err)
			return err
		}
	}

	return nil
}
