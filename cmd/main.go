package main //entry point of the application

import (
	"database/sql"
	"log"

	"github.com/go-sql-driver/mysql"
	"github.com/hilalerdogan/mobile-application-backend/cmd/api"
	"github.com/hilalerdogan/mobile-application-backend/config"
	"github.com/hilalerdogan/mobile-application-backend/db"
)

func main() {
	db, err := db.NewMySQLStorage(mysql.Config{
		User:                 config.Envs.DBUser,
		Passwd:               config.Envs.DBPasswd,
		Addr:                 config.Envs.DBAddr,
		DBName:               config.Envs.DBName,
		Net:                  "tcp",
		AllowNativePasswords: true,
		ParseTime:            true,
	})
	if err != nil {
		log.Fatal(err)
	}
	initStotage(db)
	server := api.NewAPIServer(":8080", db)
	if err := server.Run(); err != nil {
		log.Fatal(err)
	}
}

func initStotage(db *sql.DB) {
	err := db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Database is connected")

}
