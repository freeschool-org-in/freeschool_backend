package main

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo"
)

// Module is a slice of lessons with some meta data
type Module struct {
	ID          int    `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	CourseID    int    `json:"courseID"`
}

func createModule(c echo.Context) error {
	m := &Module{}
	if err := c.Bind(m); err != nil {
		return err
	}

	insertCourseSQL := "INSERT INTO module(title, course_id) VALUES(?,?)"

	stmt, err := db.Prepare(insertCourseSQL)
	if err != nil {
		return err
	}

	_, err = stmt.Exec(m.Title, m.CourseID)

	return c.JSON(http.StatusCreated, m)
}

func getModules(c echo.Context) error {
	var mod []Module = make([]Module, 0)
	row, err := db.Query("SELECT id, title, course_id FROM module")
	if err != nil {
		return err
	}
	defer row.Close()

	for row.Next() {
		m := Module{}
		row.Scan(&m.ID, &m.Title, &m.CourseID)
		mod = append(mod, m)
	}
	return c.JSON(http.StatusOK, mod)
}

func getModuleByID(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	row := db.QueryRow("SELECT id, title FROM module WHERE id=?", id)
	m := Module{}
	row.Scan(&m.ID, &m.Title)

	return c.JSON(http.StatusOK, m)
}

func updateModule(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	m := &Module{}

	if err := c.Bind(&m); err != nil {
		return err
	}

	stmt, err := db.Prepare("UPDATE module SET title=? WHERE id=?")
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(m.Title, id)

	return c.JSON(http.StatusOK, m.Title)
}

func deleteModule(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	stmt, err := db.Prepare("DELETE FROM module where id=?")
	if err != nil {
		return err
	}
	defer stmt.Close()

	stmt.Exec(id)

	return c.NoContent(http.StatusOK)

}