package usercontroller

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

var urldb string
var db *sql.DB
var rows *sql.Rows
var err error

type User struct {
	Name     string `json:"name"`
	Position string `json:"position"`
}

func getUser(c *gin.Context) {
	if db, err = sql.Open("mysql", urldb); err != nil {
		log.Fatal(err)
	}

	defer db.Close()

	var query = `
		select name, position 
		from user
		order by name;
	`

	if rows, err = db.Query(query); err != nil {
		log.Fatal(err)
	}

	defer rows.Close()

	var users []User
	var user User

	for rows.Next() {
		if err = rows.Scan(
			&user.Name,
			&user.Position,
		); err != nil {
			log.Fatal(err)
		}

		users = append(users, user)
	}

	c.JSON(http.StatusOK, users)
	c.Done()
}
