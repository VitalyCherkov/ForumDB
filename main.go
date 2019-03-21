package main

import (
	"ForumDB/database"
	"ForumDB/models"
	"ForumDB/router"
	"fmt"
	"net/http"
	"os"
)

func main() {

	env := &models.Env{}
	mainRouter := router.Init(env)

	db, err := database.Init(
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_NAME"),
	)
	if err != nil {
		panicMsg := fmt.Sprintf("Can not connect to postgres: %v\n", err)
		panic(panicMsg)
	}
	env.DB = db
	defer db.Close()

	fmt.Println(http.ListenAndServe(":"+os.Getenv("PORT"), mainRouter))
}
