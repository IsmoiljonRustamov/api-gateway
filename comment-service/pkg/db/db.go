package db

import (
	"database/sql"
	"fmt"
	"projects/comment-service/config"

	_ "github.com/lib/pq"
)

func ConnectDB(cfg config.Config) (*sql.DB,error) {
	str := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
	cfg.PostgresHost,cfg.PostgresPort,cfg.PostgresUser,cfg.PostgresPassword,cfg.PostgresDatabase)

	connDB,err := sql.Open("postgres",str)
	if err != nil {
		return nil,err
	}

	return connDB,nil
}