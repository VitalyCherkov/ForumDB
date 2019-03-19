package database

import (
	"fmt"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

func makeMigrations(db *sqlx.DB, databaseName string) error {
	driver, err := postgres.WithInstance(db.DB, &postgres.Config{})
	if err != nil {
		fmt.Printf("%v\n", err)
		return err
	}
	fmt.Printf("created driver instance\n")

	m, err := migrate.NewWithDatabaseInstance("file://migrations", databaseName, driver)
	if err != nil {
		fmt.Printf("[ERROR] %v\n", err)
		return err
	}
	fmt.Printf("created migration\n")

	err = m.Steps(1)
	if err != nil {
		fmt.Printf("[ERROR] %v\n", err)
		return err
	}
	fmt.Printf("passed migration\n")

	return nil
}

func Init(user, password, port, dbName string) (db *sqlx.DB, err error) {

	connectionString := fmt.Sprintf(
		"postgresql://%v:%v@localhost:%v/%v?sslmode=disable",
		user,
		password,
		port,
		dbName,
	)

	db, err = sqlx.Open("postgres", connectionString)

	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	fmt.Printf("successfully connected to database\n")

	err = makeMigrations(db, dbName)
	if err != nil {
		fmt.Printf("[ERROR] falid to process migrations: %v\n", err)
		db.Close()
		return nil, err
	}

	return db, nil
}
