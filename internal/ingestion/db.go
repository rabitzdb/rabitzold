package ingestion

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"time"
)

var db *sql.DB

func InitDb(user string, pass string, host string, port string, database string) {
	//<username>:<pw>@tcp(<HOST>:<port>)/<dbname>
	connStr := user + ":" + pass + "@tcp(" + host + ":" + port + ")/"+database
	dbMy, err := sql.Open("mysql",connStr)
	if err != nil {
		log.Fatal(err)
	}
	defer dbMy.Close()

	err = dbMy.Ping()
	if err != nil {
		log.Fatal(err.Error())
	}

	dbMy.SetConnMaxLifetime(time.Minute*5);
	dbMy.SetMaxIdleConns(0);
	dbMy.SetMaxOpenConns(20);

	db = dbMy
}