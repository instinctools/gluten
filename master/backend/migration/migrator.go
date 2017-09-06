package migration

import (
	"bitbucket.org/instinctools/gluten/shared/logging"
	"github.com/mattes/migrate"
	_ "github.com/mattes/migrate/database/postgres"
	_ "github.com/mattes/migrate/source/file"
)

func ApplyMigrations(folder string, connectionString string) {
	m, err := migrate.New(folder, connectionString)
	if err != nil {
		logging.WithFields(logging.Fields{
			"error": err,
		}).Error("Creating migration error")
	}
	m.Up()
}
