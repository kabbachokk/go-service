package main

import (
	"fmt"
	"log"

	"database/sql"

	"github.com/golang-migrate/migrate/v4"

	"github.com/golang-migrate/migrate/v4/database/sqlite3"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/mattn/go-sqlite3"

	"go.service/internal/pkg/util"
)

type Config struct {
	DBDriver         string `mapstructure:"DB_DRIVER"`
	DBSource         string `mapstructure:"DB_SOURCE"`
	DBMigrationsPath string `mapstructure:"DB_MIGRATIONS_PATH"`
}

func main() {
	config := new(Config)
	if err := util.LoadConfig(".", "migrate", config); err != nil {
		log.Fatal("cannot load config:", err)
	}

	conn, err := util.ConnectDb(config.DBDriver, config.DBSource)
	if err != nil {
		log.Fatal("cannot connect to db:", err)
	}

	// Starting migration job.

	if err := dbMigrate(config, conn); err != nil {
		log.Fatal("migrate error:", err)
	}

	fmt.Println("migration complete")
}

// This function executes the migration scripts.
func dbMigrate(config *Config, conn *sql.DB) (err error) {
	d, err := sqlite3.WithInstance(conn, &sqlite3.Config{})
	if err != nil {
		return
	}

	m, err := migrate.NewWithDatabaseInstance(
		config.DBMigrationsPath,
		config.DBSource,
		d,
	)
	if err != nil {
		return
	}

	err = m.Up()
	if err != nil && err != migrate.ErrNoChange {
		return
	}

	return nil
}
