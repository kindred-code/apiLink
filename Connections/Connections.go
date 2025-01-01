package connections

import (
	"database/sql"
	"fmt"
	"os"

	ph "mpolitakis.LinkApi/Data/Photo"
	us "mpolitakis.LinkApi/Data/User"
)

// Connections returns a new connection to the Postgres database.
//
// The connection string is defined as a constant within this function.
// If the connection fails, the function will exit with a status code of 1.
func Connections() *sql.DB {
	connStr := "user=postgres dbname=postgres password=Agile5-Oxford2-Snooze2-Populate8-Suitably3-Limb4-Occultist1-Throwback4-Crabbing4-Clamshell0 sslmode=disable"

	conn, err := sql.Open("postgres", connStr)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}

	return conn
}

// BuildSql returns an SQL string for inserting a user into the database.
//
// The SQL string is of the form "insert into users (username, password, email)
// values ('%s', '%s', '%s')". The values are taken from the provided us.User.
func BuildSql(u *us.User) string {
	return fmt.Sprintf("insert into users (username, password, email)values  ('%s', '%s', '%s')", u.Username, u.Password, u.Email)
}

func BuildSqlPhoto(p *ph.Photo) string {
	return fmt.Sprintf("insert into photos (file, userid)values  ('%s', '%d')", p.File, p.UserId)
}
