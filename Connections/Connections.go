package connections

import (
	"database/sql"
	"fmt"
	"os"

	con "mpolitakis.LinkApi/Consts"
	details "mpolitakis.LinkApi/Data/Details"
	ph "mpolitakis.LinkApi/Data/Photo"
	profile "mpolitakis.LinkApi/Data/Profile"
)

// Connections returns a new connection to the Postgres database.
//
// The connection string is defined as a constant within this function.
// If the connection fails, the function will exit with a statprofil code of 1.
func Connections() *sql.DB {
	connStr := fmt.Sprintf("user=%s dbname=%s password=%s host=%s sslmode=disable", con.DbUser, con.DbName, con.DbPassword, con.DbHost)

	conn, err := sql.Open("postgres", connStr)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}

	return conn
}

// BuildSql returns an SQL string for inserting a profiler into the database.
//
// The SQL string is of the form "insert into profilers (profilername, password, email)
// values ('%s', '%s', '%s')". The values are taken from the provided profil.profiler.
func BuildSql(u *profile.Profile) string {
	u.Password = shaHashing(u.Password)
	return fmt.Sprintf("insert into profile (username, password, email)values  ('%s', '%s', '%s')", u.Username, u.Password, u.Email)
}
func BuildSqlDetails(u *details.Details, profileId int) string {
	return fmt.Sprintf("insert into details (profileId, gender, bio, location)values  ('%d', '%s', '%s', '%s')", profileId, u.Gender, u.Bio, u.Location)
}
func BuildSqlPhoto(p *ph.Photo) string {
	return fmt.Sprintf("insert into photo (file, profileId)values  ('%s', '%d')", p.File, p.ProfileId)
}
