package persistence

import (
	"bitbucket.org/instinctools/gluten/shared/logging"
	"github.com/mattes/migrate"
	_ "github.com/mattes/migrate/database/postgres"
	_ "github.com/mattes/migrate/source/file"
)

func ApplyMigrations() {
	m, err := migrate.New("file://migrations", "postgresql://postgres:1@localhost:5432/gluten?sslmode=disable")
	if err != nil {
		logging.WithFields(logging.Fields{
			"error": err,
		}).Error("Creating migration error")
	}
	m.Up()
}
