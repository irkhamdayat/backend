package bootstrap

import (
	"database/sql"
	"errors"

	_ "github.com/lib/pq" // postgres driver
	"github.com/pressly/goose/v3"

	"github.com/Halalins/backend/config"
	"github.com/Halalins/backend/internal/common/util"
)

func StartMigration(actionType string, name string, version *int64) {
	migrationDir := "db/migration"

	db, err := sql.Open("postgres", config.Env.Postgres.DSN)
	util.ContinueOrFatal(err)
	err = goose.SetDialect("postgres")
	util.ContinueOrFatal(err)

	switch actionType {
	case "create":
		err = goose.Create(db, migrationDir, name, "sql")
	case "up":
		err = goose.Up(db, migrationDir, goose.WithAllowMissing())
	case "up-by-one":
		err = goose.UpByOne(db, migrationDir, goose.WithAllowMissing())
	case "up-to":
		err = goose.UpTo(db, migrationDir, *version, goose.WithAllowMissing())
	case "down":
		err = goose.Down(db, migrationDir, goose.WithAllowMissing())
	case "down-to":
		err = goose.DownTo(db, migrationDir, *version, goose.WithAllowMissing())
	case "status":
		err = goose.Status(db, migrationDir)
	case "reset":
		err = goose.Reset(db, migrationDir, goose.WithAllowMissing())
		if err != nil {
			break
		}
		err = goose.Up(db, migrationDir, goose.WithAllowMissing())
	default:
		err = errors.New("invalid command")
	}

	util.ContinueOrFatal(err)
}
