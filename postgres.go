package main

mport (
	"database/sql"
	"time"

	_ "github.com/lib/pq"
)

var counts int64

func Connect(dsn string, maxOpenConnections, maxIdleConnections, maxIdleTime int) (*sql.DB, error) {
	for {
		connection, err := openDB(dsn, maxOpenConnections, maxIdleConnections, maxIdleTime)
		if err != nil {
			counts++
		} else {
			return connection, nil
		}
		if counts > 10 {
			return nil, err
		}
		time.Sleep(2 * time.Second)
	}
}

func openDB(dsn string, maxOpenConnections, maxIdleConnections, maxIdleTime int) (*sql.DB, error) {
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		return nil, err
	}
	db.SetMaxOpenConns(maxOpenConnections)
	db.SetMaxIdleConns(maxIdleConnections)
	db.SetConnMaxIdleTime(time.Duration(maxIdleTime) * time.Minute)
	if err := db.Ping(); err != nil {
		return nil, err
	}
	return db, nil
}