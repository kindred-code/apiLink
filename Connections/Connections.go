package connections

import (
	"database/sql"
	"fmt"
	"os"

	us "mpolitakis.LinkApi/User"
)

func Connections() *sql.DB {
	connStr := "user=postgres dbname=postgres password=Agile5-Oxford2-Snooze2-Populate8-Suitably3-Limb4-Occultist1-Throwback4-Crabbing4-Clamshell0 sslmode=disable"

	conn, err := sql.Open("postgres", connStr)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}

	return conn
}

func BuildSql(u *us.User) string {
	return fmt.Sprintf("insert into users (username, password, email)values  ('%s', '%s', '%s')", u.Username, u.Password, u.Email)
}
