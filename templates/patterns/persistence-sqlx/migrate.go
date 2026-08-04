package persist

import (
	"embed"
	"fmt"
	"strings"

	"github.com/jmoiron/sqlx"
)

//go:embed dialect/*.sql
var dialectFS embed.FS

// Migrate applies the dialect migration for driver (sqlite|postgres|mysql).
func Migrate(db *sqlx.DB, driver string) error {
	name, err := dialectFile(driver)
	if err != nil {
		return err
	}
	sqlBytes, err := dialectFS.ReadFile(name)
	if err != nil {
		return fmt.Errorf("persist: read %s: %w", name, err)
	}
	if _, err := db.Exec(string(sqlBytes)); err != nil {
		return fmt.Errorf("persist: migrate %s: %w", name, err)
	}
	return nil
}

func dialectFile(driver string) (string, error) {
	switch strings.ToLower(strings.TrimSpace(driver)) {
	case "sqlite", "sqlite3":
		return "dialect/sqlite.sql", nil
	case "postgres", "postgresql":
		return "dialect/postgres.sql", nil
	case "mysql":
		return "dialect/mysql.sql", nil
	default:
		return "", fmt.Errorf("persist: unsupported driver %q for migrate", driver)
	}
}
