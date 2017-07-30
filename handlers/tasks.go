package handlers

import (
	"database/sql"
	"net/http"
	"strconv"

	"todo/models"

	"github.com/labstack/echo"
)

// return arbitrary JSONs in response
type H map[string]interface{}

// GetTasks endpoint
func GetTasks(db *sql.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		// fetch tasks using the model
		return c.JSON(http.StatusOK, models.GetTasks(db))
	}
}

// PutTask endpoint
func PutTasks(db *sql.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		// Instantiate a new task
		var task models.Task
		// map incoming JSON body to the new Task
		c.Bind(&task)
		// add task using our new model
		id, err := models.PutTask(db, task.Name)
		// return a JSON response if successful
		if err == nil {
			return c.JSON(http.StatusCreated, H{
				"created": id})
		} else {
			return err
		}
	}
}

// DeleteTask endpoint
func DeleteTasks(db *sql.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		id, _ := strconv.Atoi(c.Param("id"))

		// use our new model to delete a task
		_, err := models.DeleteTask(db, id)

		// return a JSON response on success
		if err == nil {
			return c.JSON(http.StatusOK, H{
				"deleted": id,
			})
		} else {
			return err
		}
	}
}
