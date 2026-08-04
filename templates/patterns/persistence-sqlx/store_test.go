package persist

import (
	"testing"

	_ "modernc.org/sqlite"
)

func TestOpenAndMigrate(t *testing.T) {
	db, err := Open(Config{Driver: "sqlite", DSN: "file::memory:?cache=shared"})
	if err != nil {
		t.Fatalf("Open: %v", err)
	}
	t.Cleanup(func() { _ = db.Close() })

	if err := Migrate(db, "sqlite"); err != nil {
		t.Fatalf("Migrate: %v", err)
	}
}

func TestFromEnvSQLiteDefaults(t *testing.T) {
	t.Setenv("DB_DRIVER", "")
	t.Setenv("DATABASE_URL", "")
	t.Setenv("SQLITE_PATH", "")

	cfg, err := FromEnv()
	if err != nil {
		t.Fatalf("FromEnv: %v", err)
	}
	if cfg.Driver != "sqlite" {
		t.Fatalf("Driver=%q want sqlite", cfg.Driver)
	}
	if cfg.DSN != "notes.db" {
		t.Fatalf("DSN=%q want notes.db", cfg.DSN)
	}
}

func TestFromEnvRequiresDatabaseURL(t *testing.T) {
	t.Setenv("DB_DRIVER", "postgres")
	t.Setenv("DATABASE_URL", "")

	_, err := FromEnv()
	if err == nil {
		t.Fatal("expected error when DATABASE_URL missing for postgres")
	}
}

func TestMigrateUnknownDriver(t *testing.T) {
	db, err := Open(Config{Driver: "sqlite", DSN: "file::memory:?cache=shared"})
	if err != nil {
		t.Fatalf("Open: %v", err)
	}
	t.Cleanup(func() { _ = db.Close() })

	if err := Migrate(db, "unknown"); err == nil {
		t.Fatal("expected error for unknown driver")
	}
}

func TestOpenRequiresDriver(t *testing.T) {
	if _, err := Open(Config{DSN: "file::memory:?cache=shared"}); err == nil {
		t.Fatal("expected error for empty driver")
	}
}

func TestFromEnvDatabaseURLOverridesSQLitePath(t *testing.T) {
	t.Setenv("DB_DRIVER", "sqlite")
	t.Setenv("DATABASE_URL", "file::memory:?cache=shared")
	t.Setenv("SQLITE_PATH", "/ignored.db")

	cfg, err := FromEnv()
	if err != nil {
		t.Fatalf("FromEnv: %v", err)
	}
	if cfg.DSN != "file::memory:?cache=shared" {
		t.Fatalf("DSN=%q want file::memory:?cache=shared", cfg.DSN)
	}
}
