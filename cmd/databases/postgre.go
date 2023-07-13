package database

import (
	"fmt"
	"log"
	"os"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

func PostgresSQL() *sqlx.DB {
	psql := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=%s", os.Getenv("DB_USER"), os.Getenv("DB_PASS"), os.Getenv("DB_NAME"), os.Getenv("DB_SSLMODE"))
	db, err := sqlx.Connect(os.Getenv("DB_DRIVER"), psql)
	if err != nil {
		panic(err)
	}

	err = db.Ping()
	if err != nil {
		panic(err)
	}
	log.Println("Connected to database")
	return db
}

func MigrateDB(db *sqlx.DB) error {
	log.Println("Running migration database...")
	_, err := db.Exec(promoTable)
	if err != nil {
		log.Println("error 1")
		return err
	}
	_, err = db.Exec(orderTable)
	if err != nil {
		log.Println("error 2")
		return err
	}
	_, err = db.Exec(orderPromoTable)
	if err != nil {
		log.Println("error 3")
		return err
	}
	log.Println("Migration database finished")
	return nil
}

func NewDatabases() *sqlx.DB {
	db := PostgresSQL()
	err := MigrateDB(db)
	if err != nil {
		panic(err)
	}
	return db
}
