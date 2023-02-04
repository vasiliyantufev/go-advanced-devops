package database

import (
	"database/sql"
	"github.com/vasiliyantufev/go-advanced-devops/internal/config"
	"github.com/vasiliyantufev/go-advanced-devops/internal/storage"

	//_ "github.com/mattn/go-sqlite3"
	log "github.com/sirupsen/logrus"

	_ "github.com/lib/pq"
)

var db DB

type DB struct {
	pool *sql.DB
}

func ConnectDB() error {

	pool, err := sql.Open("postgres",
		config.GetConfigDBServer())

	if err != nil {
		log.Error(err)
		return err
	}
	//defer pool.Close()
	db = DB{pool: pool}

	return nil
}

func Ping() error {
	if err := db.pool.Ping(); err != nil {
		log.Error(err)
		return err
	}
	return nil
}

func CreateTables() {
	var metricsTable = `		
		CREATE TABLE IF NOT EXISTS metrics (
    		id VARCHAR(256),
		    mtype VARCHAR(10),
		    value NUMERIC,
		    delta BIGINT,
			hash  varchar,
		    UNIQUE (id, mtype)
		);

		CREATE UNIQUE INDEX IF NOT EXISTS id_mtype_index
		ON metrics (id, mtype)
 `
	_, err := db.pool.Exec(metricsTable)

	//log.Info(res.LastInsertId())
	if err != nil {
		log.Fatal(err)
		return
	}
	log.Info("CREATE TABLE metrics")
}

func InsertOrUpdateMetrics(agent *storage.MemStorage) error {

	// шаг 1 — объявляем транзакцию
	tx, err := db.pool.Begin()
	if err != nil {
		return err
	}

	// шаг 1.1 — если возникает ошибка, откатываем изменения
	defer tx.Rollback()

	// шаг 2 — готовим инструкцию
	stmt, err := tx.Prepare(`
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
	// шаг 2.1 — не забываем закрыть инструкцию, когда она больше не нужна
	defer stmt.Close()

	for _, val := range agent.GetAllMetrics() {
		// шаг 3 — указываем, что будет добавлено в транзакцию
		if _, err = stmt.Exec(val.ID, val.MType, val.Value, val.Delta, val.Hash); err != nil {
			log.Error(err)
			return err
		}
	}
	// шаг 4 — сохраняем изменения
	return tx.Commit()
	//log.Info(stmt.)

	return nil
}
