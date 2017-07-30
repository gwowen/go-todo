package main

import (
	"database/sql"
	"todo/handlers"

	"github.com/labstack/echo"
	_ "github.com/mattn/go-sqlite3"
)

func main() {

	db := initDB("storage.db")
	migrate(db)

	e := echo.New()

	e.File("/", "public/index.html")
	e.GET("/tasks", handlers.GetTasks(db))
	e.PUT("/tasks", handlers.PutTasks(db))
	e.DELETE("/tasks/:id", handlers.DeleteTasks(db))
	e.Logger.Fatal(e.Start(":8000"))
}

func initDB(filepath string) *sql.DB {
	db, err := sql.Open("sqlite3", filepath)

	// check for db errors, then exit
	if err != nil {
		panic(err)
	}

	// if there's no errors, but no connection,
	// exit
	if db == nil {
		panic("db nil")
	}

	return db
}

func migrate(db *sql.DB) {
	sql := `
		CREATE TABLE IF NOT EXISTS tasks(
			id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
			name VARCHAR NOT NULL
		);
	`

	_, err := db.Exec(sql)

	// exit if something goes wrong with the
	// SQL statement
	if err != nil {
		panic(err)
	}
}
