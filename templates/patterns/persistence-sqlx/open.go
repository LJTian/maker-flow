package persist

import (
	"fmt"
	"os"
	"strings"

	"github.com/jmoiron/sqlx"
)

// Config selects a database/sql driver name and DSN.
// Product repos blank-import the concrete driver (see README).
type Config struct {
	Driver string // sqlite | postgres | mysql
	DSN    string
}

// FromEnv reads DB_DRIVER and DATABASE_URL.
// For sqlite, SQLITE_PATH may be used when DATABASE_URL is empty
// (normalized to a file: DSN).
func FromEnv() (Config, error) {
	driver := strings.ToLower(strings.TrimSpace(os.Getenv("DB_DRIVER")))
	if driver == "" {
		driver = "sqlite"
	}
	dsn := strings.TrimSpace(os.Getenv("DATABASE_URL"))
	if dsn == "" && driver == "sqlite" {
		path := strings.TrimSpace(os.Getenv("SQLITE_PATH"))
		if path == "" {
			path = "notes.db"
		}
		dsn = path
	}
	if dsn == "" {
		return Config{}, fmt.Errorf("persist: set DATABASE_URL (or SQLITE_PATH for sqlite)")
	}
	return Config{Driver: driver, DSN: dsn}, nil
}

// Open opens a sqlx.DB. The matching driver must already be registered
// via blank import in the product main package (or test).
func Open(cfg Config) (*sqlx.DB, error) {
	name, err := sqlDriverName(cfg.Driver)
	if err != nil {
		return nil, err
	}
	db, err := sqlx.Open(name, cfg.DSN)
	if err != nil {
		return nil, fmt.Errorf("persist: open %s: %w", name, err)
	}
	if err := db.Ping(); err != nil {
		_ = db.Close()
		return nil, fmt.Errorf("persist: ping %s: %w", name, err)
	}
	return db, nil
}

func sqlDriverName(driver string) (string, error) {
	switch strings.ToLower(strings.TrimSpace(driver)) {
	case "sqlite":
		// Prefer pure-Go modernc ("sqlite"). Product may use mattn ("sqlite3") —
		// call Open with Driver "sqlite3" or blank-import modernc.
		return "sqlite", nil
	case "sqlite3":
		return "sqlite3", nil
	case "postgres", "postgresql":
		return "postgres", nil
	case "mysql":
		return "mysql", nil
	default:
		return "", fmt.Errorf("persist: unsupported DB_DRIVER %q (want sqlite|postgres|mysql)", driver)
	}
}
