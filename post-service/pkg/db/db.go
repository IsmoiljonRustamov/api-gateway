package db

import (
	"database/sql"
	"fmt"
	"projects/post-service/config"

	_ "github.com/lib/pq"
)

func ConnectDB(cfg config.Config) (*sql.DB, error) {
	str := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		cfg.PostgresHost, cfg.PostgresPort, cfg.PostgresUser, cfg.PostgresPassword, cfg.PostgresDatabase)

	connDB, err := sql.Open("postgres", str)
	if err != nil {
		return nil, err
	}

	return connDB, nil
}

func ConnectDBForSuite(cfg config.Config) (*sql.DB, func()) {
	psqlString := fmt.Sprintf("host=%s port=%s user=%s password=%s database=%s sslmode=disable",
		cfg.PostgresHost,
		cfg.PostgresPort,
		cfg.PostgresUser,
		cfg.PostgresPassword,
		cfg.PostgresDatabase,
	)

	connDB,err := sql.Open("postgres",psqlString)
	if err != nil {
		return nil,func() {}
	}

	cleaUpfunc := func ()  {
		connDB.Close()
	}

	return connDB,cleaUpfunc

}
