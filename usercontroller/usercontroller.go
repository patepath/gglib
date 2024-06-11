package usercontroller

import (
	"database/sql"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

type repository interface {
	FindOne(id int64) User
	FindAll() []User
}

func FindOne(c *gin.Context) {
	var id int64

	if id, err = strconv.ParseInt(c.Param("id"), 10, 64); err != nil {
		panic(err)
	}

	var userRepo repository = User{}

	c.JSON(http.StatusOK, userRepo.FindOne(id))
	c.Done()
}

func FindAll(c *gin.Context) {
	var userRepo repository = User{}

	c.JSON(http.StatusOK, userRepo.FindAll())
	c.Done()
}

// --------------------------------------------------------------------------------
var db *sql.DB
var rows *sql.Rows
var err error

type User struct {
	Name     string `json:"name"`
	Position string `json:"position"`
}

func getDSN() string {
	if err := godotenv.Load(".env"); err != nil {
		log.Fatal(err)
	}

	return os.Getenv("dsn")
}

func (u User) FindOne(id int64) User {
	if db, err = sql.Open("mysql", getDSN()); err != nil {
		log.Fatal(err)
	}

	defer db.Close()

	var query = `
		select name, position 
		from user
		where id=?
		order by name;
	`

	if rows, err = db.Query(query, id); err != nil {
		log.Fatal(err)
	}

	defer rows.Close()

	var user User

	if rows.Next() {
		if err = rows.Scan(
			&user.Name,
			&user.Position,
		); err != nil {
			log.Fatal(err)
		}
	}

	return user
}

func (u User) FindAll() []User {
	if db, err = sql.Open("mysql", getDSN()); err != nil {
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

	return users
}
