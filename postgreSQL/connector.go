package postgreSQL

import (
	"database/sql"
	_ "github.com/lib/pq"
	"os"
)

func GetKey() (string, error) {
	dbURL := os.Getenv("DATABASE_URL")
	db, err := sql.Open("postgres", dbURL)
	if err != nil {
		return "", err
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	var key struct {
		Id  string `json:"id"`
		Key string `json:"key"`
	}

	err = db.QueryRow("SELECT id, key FROM key_table WHERE id = 'k';").Scan(&key)
	if err != nil {
		return "", err
	}

	return key.Key, nil
}
