package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	L "./api"

	_ "github.com/lib/pq"
)

var db *sql.DB

const (
	dbhost = "DBHOST"
	dbport = "DBPORT"
	dbuser = "DBUSER"
	dbpass = "DBPASS"
	dbname = "DBNAME"
)

func main() {
	initDb()
	defer db.Close()

	// Rota -> chama uma função
	http.HandleFunc("/api/index", L.IndexHandler)
	http.HandleFunc("/api/repo", L.RepoHandler)

	log.Fatal(http.ListenAndServe("localhost:3131", nil))
}

func initDb() {
	config := dbConfig()
	var err error

	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		config[dbhost], config[dbport], config[dbuser], config[dbpass], config[dbname])

	db, err = sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	fmt.Println("Conectado com sucesso!")
}

func dbConfig() map[string]string {
	conf := make(map[string]string)

	host, ok := os.LookupEnv(dbhost)
	if !ok {
		panic("DBHOST environment variable required but not set")
	}

	port, ok := os.LookupEnv(dbport)
	if !ok {
		panic("DBPORT environment variable required but not set")
	}

	user, ok := os.LookupEnv(dbuser)
	if !ok {
		panic("DBUSER environment variable required but not set")
	}

	password, ok := os.LookupEnv(dbpass)
	if !ok {
		panic("DBPASS environment variable required but not set")
	}

	name, ok := os.LookupEnv(dbname)
	if !ok {
		panic("DBNAME environment variable required but not set")
	}

	conf[dbhost] = host
	conf[dbport] = port
	conf[dbuser] = user
	conf[dbpass] = password
	conf[dbname] = name

	return conf
}
