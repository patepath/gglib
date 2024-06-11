package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
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

func main() {
	fmt.Println("start...")

	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal(err)
	}

	urldb = fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", os.Getenv("username"), os.Getenv("password"), os.Getenv("dbhost"), os.Getenv("port"), os.Getenv("database"))

	r := gin.Default()

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	var v1 = r.Group("/v1")
	{
		v1.GET("/user", getUser)
	}

	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
