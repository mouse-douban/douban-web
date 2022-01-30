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

func OpenTransaction() (*sql.Tx, error) {
	return dB.Begin()
}

func RollBackTransaction(tx *sql.Tx) {
	if tx == nil {
		return
	}
	err := tx.Rollback()
	if err != nil {
		log.Println("事务回滚失败！原因：", err)
	}
}

func CommitTransaction(tx *sql.Tx) {
	if tx == nil {
		return
	}
	err := tx.Commit()
	if err != nil {
		log.Println("事务提交失败！原因：", err)
	}
}
