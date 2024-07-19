package main

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/PanosVasilopoulos92/ecom/cmd/api"
	"github.com/PanosVasilopoulos92/ecom/config"
	"github.com/PanosVasilopoulos92/ecom/db"
	"github.com/go-sql-driver/mysql"
)

func main() {
	db, err := db.NewMySQLStorage(mysql.Config{
		User:                 config.Envs.DBUser,
		Passwd:               config.Envs.DBPasswd,
		Addr:                 config.Envs.DBAddress,
		DBName:               config.Envs.DBName,
		Net:                  "tcp",
		AllowNativePasswords: true,
		ParseTime:            true,
	})

	if err != nil {
		log.Fatalln(err)
	}

	initStorage(db)

	server := api.NewAPIServer(":8080", db)

	if err := server.Run(); err != nil {
		log.Fatalln(err)
	}

}

func initStorage(db *sql.DB) {
	err := db.Ping()
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println("Successful connection with database!")
}
