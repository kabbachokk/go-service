package util

import (
	"database/sql"
	"fmt"

	"github.com/cenkalti/backoff/v4"

	_ "github.com/mattn/go-sqlite3"
)

func ConnectDb(dbDriver string, dbSource string) (conn *sql.DB, err error) {
	conn, err = sql.Open(dbDriver, dbSource)
	if err != nil {
		return
	}

	b := backoff.NewExponentialBackOff()
	b.MaxElapsedTime = 0

	err = backoff.Retry(func() error {
		fmt.Println("Connecting to a database ...")
		if err := conn.Ping(); err != nil {
			return fmt.Errorf("ping failed %v", err)
		}
		return nil
	}, b)

	return
}
