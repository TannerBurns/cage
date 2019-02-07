package models

import (
	"database/sql"
	"fmt"
	"os"
	"strings"

	_ "github.com/lib/pq"
)

type Connection struct {
	Config *ConfigParser
}

func NewSession() *Connection {
	connection := &Connection{}
	connection.Config = &ConfigParser{Path: "./config/config.dev.ini"}
	_ = connection.Config.Parse()
	return connection
}

// Connect to database
func (conn *Connection) Connect() (db *sql.DB, err error) {
	if strings.ToLower(os.Args[1]) == "dev" {
		db, err = sql.Open("postgres",
			fmt.Sprintf("host=%s dbname=%s user=%s password=%s sslmode=disable",
				"localhost",
				"cage",
				"postgres",
				"postgres"))
	} else if conn.Config.Parsed["postgres"]["password"] == "" {
		db, err = sql.Open("postgres",
			fmt.Sprintf("host=%s dbname=%s user=%s sslmode=disable",
				conn.Config.Parsed["postgres"]["host"],
				conn.Config.Parsed["postgres"]["database"],
				conn.Config.Parsed["postgres"]["user"]))
	} else {
		db, err = sql.Open("postgres",
			fmt.Sprintf("host=%s dbname=%s user=%s password=%s sslmode=disable",
				conn.Config.Parsed["postgres"]["host"],
				conn.Config.Parsed["postgres"]["database"],
				conn.Config.Parsed["postgres"]["user"],
				conn.Config.Parsed["postgres"]["password"]))
	}
	return
}
