package dao

import (
	"database/sql"
	"douban-webend/config"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

var dB *sql.DB

func InitDao() {
	configuration := config.Config
	Initialize(
		configuration.DefaultDbName,
		configuration.DefaultRoot,
		configuration.DefaultPassword,
		configuration.DefaultIpAndPort,
		configuration.DefaultCharset,
	)
}

func Initialize(dbName, root, pwd, ipAndPort, charset string) {
	daraSourceName := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=%s&parseTime=True", root, pwd, ipAndPort, dbName, charset)
	db, err := sql.Open("mysql", daraSourceName)
	if err != nil {
		log.Fatal(err)
	}
	dB = db
}
