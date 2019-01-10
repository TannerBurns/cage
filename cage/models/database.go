package models

import (
	"database/sql"
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	_ "github.com/lib/pq"
)

type PostgresConnection struct {
	Host     string
	Database string
	User     string
	Password string
}

// Connect to database
func (conn *PostgresConnection) Connect() (db *sql.DB, err error) {
	dat, err := ioutil.ReadFile(os.Args[1])
	if err != nil {
		return
	}
	sections := strings.Split(string(dat), "\n\r\n")
	for i := range sections {
		lines := strings.Split(sections[i], "\n")
		header := lines[0]
		conf := lines[1:]
		if strings.Contains(header, "postgres") {
			for i2 := range conf {
				if strings.Contains(conf[i2], "host") {
					conn.Host = strings.Split(conf[i2], "=")[1]
				}
				if strings.Contains(conf[i2], "database") {
					conn.Database = strings.Split(conf[i2], "=")[1]
				}
				if strings.Contains(conf[i2], "user") {
					conn.User = strings.Split(conf[i2], "=")[1]
				}
				if strings.Contains(conf[i2], "password") {
					conn.Password = strings.Split(conf[i2], "=")[1]
				}
			}
			break
		}
	}
	if conn.Password == "" {
		db, err = sql.Open("postgres",
			fmt.Sprintf("host=%s dbname=%s user=%s sslmode=disable",
				conn.Host,
				conn.Database,
				conn.User))
	} else {
		db, err = sql.Open("postgres",
			fmt.Sprintf("host=%s dbname=%s user=%s password=%s sslmode=disable",
				conn.Host,
				conn.Database,
				conn.User,
				conn.Password))
	}
	return
}
