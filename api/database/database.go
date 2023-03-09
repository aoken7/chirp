package database

import (
	"database/sql"
	"os"
	"time"

	"github.com/go-sql-driver/mysql"
	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/mysqldialect"
	"github.com/uptrace/bun/extra/bundebug"
)

var DB *bun.DB

func Connect() {
	config := mysql.Config{
		DBName:    os.Getenv("MYSQL_DATABASE"),
		User:      os.Getenv("MYSQL_USER"),
		Passwd:    os.Getenv("MYSQL_PASSWORD"),
		Addr:      "mysql:3306",
		Net:       "tcp",
		ParseTime: true,
		Collation: "utf8mb4_unicode_ci",
		Loc:       time.Local,
	}

	mysql, err := sql.Open("mysql", config.FormatDSN())
	if err != nil {
		panic(err)
	}

	DB := bun.NewDB(mysql, mysqldialect.New())
	DB.AddQueryHook(bundebug.NewQueryHook(
		bundebug.WithVerbose(true),
		bundebug.FromEnv("BUNDEBUG"),
	))
}
