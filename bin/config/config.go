package config

import (
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/joho/godotenv"
)

// Env struct
type Env struct {
	RootApp    string
	HTTPPort   uint16
	PostgreSQL struct {
		Host     string
		User     string
		Password string
		DBName   string
		Port     uint16
		SSLMode  string
	}
}

// GlobalEnv global environment
var GlobalEnv Env

func init() {
	err := godotenv.Load()

	if err != nil {
		log.Println(err)
	}

	path, err := os.Getwd()
	if err != nil {
		log.Println(err)
	}

	rootApp := strings.TrimSuffix(path, "/bin/config")
	os.Setenv("APP_PATH", rootApp)
	GlobalEnv.RootApp = rootApp

	loadGeneral()
	loadPostgreSQL()
}

func loadGeneral() {
	port, _ := strconv.Atoi(os.Getenv("HTTP_PORT"))
	GlobalEnv.HTTPPort = uint16(port)
}

func loadPostgreSQL() {
	GlobalEnv.PostgreSQL.Host = os.Getenv("POSTGRE_HOST")
	GlobalEnv.PostgreSQL.User = os.Getenv("POSTGRE_USER")
	GlobalEnv.PostgreSQL.Password = os.Getenv("POSTGRE_PASSWORD")
	GlobalEnv.PostgreSQL.DBName = os.Getenv("POSTGRE_DBNAME")
	Portpostgre, _ := strconv.Atoi(os.Getenv("POSTGRE_PORT"))
	GlobalEnv.PostgreSQL.Port = uint16(Portpostgre)
	GlobalEnv.PostgreSQL.SSLMode = os.Getenv("POSTGRE_SSLMODE")
}
