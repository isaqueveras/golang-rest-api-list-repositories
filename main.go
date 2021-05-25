package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"
	_ "github.com/lib/pq"
)

var db *sql.DB

const (
	dbhost = "DBhost"
	dbport = "DBport"
	dbuser = "DBuser"
	dbpass = "DBpass"
	dbname = "DBname"
)

func main() {
	initDB()
	defer db.Close()

	http.HandleFunc("/api/index", indexHandler)
	http.HandleFunc("/api/repo/", repoHandler)
	log.Fatal(http.ListenAndServe("localhost:3131", nil))
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	
}

func repoHandler(w http.ResponseWriter, r *http.Request) {
	
}

func initDB() {
	config := dbConfig()
	var err error
	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s " + "password=%s dbname=%s sslmode=disable", 
		config[dbhost], 
		config[dbport], 
		config[dbuser], 
		config[dbpass], 
		config[dbname]
	)

	db, err = sql.Open("postgres", psqlInfo)
	if (err != nil) {
		panic(err)
	}

	err = db.Ping()
	if (err != nil) {
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

	host, ok := os.LookupEnv(dbport)
	if !ok {
		panic("DBPORT environment variable required but not set")
	}

	host, ok := os.LookupEnv(dbport)
	if !ok {
		panic("DBPORT environment variable required but not set")
	}

	host, ok := os.LookupEnv(dbuser)
	if !ok {
		panic("DBUSER environment variable required but not set")
	}

	host, ok := os.LookupEnv(dbpass)
	if !ok {
		panic("DBPASS environment variable required but not set")
	}

	host, ok := os.LookupEnv(dbname)
	if !ok {
		panic("DBNAME environment variable required but not set")
	}

	conf(dbhost) = host
	conf(dbport) = port
	conf(dbuser) = user
	conf(dbpassword) = password
	conf(dbname) = name

	return conf
}
