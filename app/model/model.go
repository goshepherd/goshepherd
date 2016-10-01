package model

import(
	"gopkg.in/gorp.v1"
	_ "github.com/mattn/go-sqlite3"
	"database/sql"
	"log"
	"os"
)

func InitDb() *gorp.DbMap {
	db, err := sql.Open("sqlite3", "./goshepherd.sqlite")
	checkErr(err, "sql.Open failed")


	dbmap := &gorp.DbMap{Db: db, Dialect: gorp.SqliteDialect{}}
	dbmap.TraceOn("[gorp]", log.New(os.Stdout, "myapp:", log.Lmicroseconds))
	return dbmap
}

func checkErr(err error, msg string) {
	if err != nil {
		log.Fatalln(msg, err)
	}
}
