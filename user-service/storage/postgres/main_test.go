package postgres

import (
	"log"
	"os"
	"projects/user-service/config"
	"projects/user-service/pkg/db"
	"projects/user-service/pkg/logger"
	"testing"
)

var pgRepo *userRepo


func TestMain(m *testing.M) {
	conf := config.Load()

	connDb, err := db.ConnectDB(conf)
	if err != nil {
		log.Fatal("error while connect to db", logger.Error(err))
	}
	pgRepo = NewUserRepo(connDb)

	os.Exit(m.Run())

}
