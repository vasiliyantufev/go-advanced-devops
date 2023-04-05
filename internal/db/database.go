// Module database
package database

import (
	"context"
	"database/sql"
	"time"

	"github.com/vasiliyantufev/go-advanced-devops/internal/config/configserver"
	"github.com/vasiliyantufev/go-advanced-devops/internal/storage"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	log "github.com/sirupsen/logrus"

	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/lib/pq"
)

type DB struct {
	pool *sql.DB
}

// NewDB - creates a new database instance
func NewDB(c *configserver.ConfigServer) (*DB, error) {
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

// Ping - checks the database connection
func (db DB) Ping() error {
	if err := db.pool.Ping(); err != nil {
		log.Error(err)
		return err
	}
	return nil
}

// Close - closes the database connection
func (db DB) Close() error {
	return db.pool.Close()
}

// CreateTablesMigration - creates database tables using migrations
func (db DB) CreateTablesMigration(cfg *configserver.ConfigServer) {

	driver, err := postgres.WithInstance(db.pool, &postgres.Config{})
	if err != nil {
		log.Error(err)
	}
	m, err := migrate.NewWithDatabaseInstance(
		cfg.RootPath,
		"postgres", driver)
	if err != nil {
		log.Error(err)
	}
	if err = m.Up(); err != nil {
		log.Error(err)
	}
}

// InsertOrUpdateMetrics - adds new metrics to the database or updates if the entry is already present
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
