package models

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

// Task is a struct containing task data
type Task struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

// TaskCollection is a collection of tasks
type TaskCollection struct {
	Tasks []Task `json:"items"`
}

// GetTasks selects all tasks from the DB, puts them into a collection
// and returns them
func GetTasks(db *sql.DB) TaskCollection {
	sql := "SELECT * FROM tasks"
	rows, err := db.Query(sql)
	// exit if the SQL doesn't work for some reason
	if err != nil {
		panic(err)
	}

	// make sure to clean up when the program exits
	defer rows.Close()

	result := TaskCollection{}

	for rows.Next() {
		task := Task{}
		err2 := rows.Scan(&task.ID, &task.Name)
		// exit if error
		if err2 != nil {
			panic(err2)
		}
		result.Tasks = append(result.Tasks, task)
	}

	return result
}

// inserts a new task into DB and returns the new id on success, panics
// on failure
func PutTask(db *sql.DB, name string) (int64, error) {
	sql := "INSERT INTO tasks(name) VALUES(?)"

	// create a prepared SQL statement
	stmt, err := db.Prepare(sql)
	// exit if error
	if err != nil {
		panic(err)
	}

	// cleanup after program exits
	defer stmt.Close()

	// replace the '?' in prepared SQL with 'name'
	result, err2 := stmt.Exec(name)
	// exit if error
	if err2 != nil {
		panic(err2)
	}

	return result.LastInsertId()
}

// the most mysterious function in all of creation
func DeleteTask(db *sql.DB, id int) (int64, error) {
	sql := "DELETE FROM tasks WHERE id = ?"

	// create a prepared SQL statement
	stmt, err := db.Prepare(sql)

	// exit if we get an error
	if err != nil {
		panic(err)
	}

	// replace the '?' in the prepared SQL with 'id'
	result, err2 := stmt.Exec(id)
	if err2 != nil {
		panic(err2)
	}

	return result.RowsAffected()
}
