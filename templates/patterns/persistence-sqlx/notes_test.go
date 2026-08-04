package persist

import (
	"context"
	"path/filepath"
	"testing"

	_ "modernc.org/sqlite"
)

func openTestDB(t *testing.T) *Store {
	t.Helper()
	dir := t.TempDir()
	dsn := filepath.Join(dir, "test.db")
	db, err := Open(Config{Driver: "sqlite", DSN: dsn})
	if err != nil {
		t.Fatalf("open: %v", err)
	}
	t.Cleanup(func() { _ = db.Close() })
	if err := Migrate(db, "sqlite"); err != nil {
		t.Fatalf("migrate: %v", err)
	}
	return NewStore(db)
}

func TestNotesCRUD(t *testing.T) {
	store := openTestDB(t)
	ctx := context.Background()

	n, err := store.Create(ctx, "hello", "world")
	if err != nil {
		t.Fatalf("create: %v", err)
	}
	if n.ID == 0 || n.Title != "hello" {
		t.Fatalf("create result: %+v", n)
	}

	got, err := store.Get(ctx, n.ID)
	if err != nil || got.Body != "world" {
		t.Fatalf("get: %+v err=%v", got, err)
	}

	updated, err := store.Update(ctx, n.ID, "hi", "there")
	if err != nil || updated.Title != "hi" || updated.Body != "there" {
		t.Fatalf("update: %+v err=%v", updated, err)
	}

	list, err := store.List(ctx, 10)
	if err != nil || len(list) != 1 {
		t.Fatalf("list: %v len=%d", err, len(list))
	}

	if err := store.Delete(ctx, n.ID); err != nil {
		t.Fatalf("delete: %v", err)
	}
	if _, err := store.Get(ctx, n.ID); err == nil {
		t.Fatal("expected missing after delete")
	}
}

func TestNotesCRUDMemoryDSN(t *testing.T) {
	db, err := Open(Config{Driver: "sqlite", DSN: "file::memory:?cache=shared"})
	if err != nil {
		t.Fatalf("open: %v", err)
	}
	t.Cleanup(func() { _ = db.Close() })
	if err := Migrate(db, "sqlite"); err != nil {
		t.Fatalf("migrate: %v", err)
	}
	store := NewStore(db)
	ctx := context.Background()

	n, err := store.Create(ctx, "mem", "dsn")
	if err != nil {
		t.Fatalf("create: %v", err)
	}
	if n.ID == 0 {
		t.Fatal("expected non-zero id")
	}
}

func TestFromEnvSQLitePath(t *testing.T) {
	t.Setenv("DB_DRIVER", "sqlite")
	t.Setenv("DATABASE_URL", "")
	t.Setenv("SQLITE_PATH", "/tmp/mf-test.db")
	cfg, err := FromEnv()
	if err != nil {
		t.Fatal(err)
	}
	if cfg.Driver != "sqlite" || cfg.DSN != "/tmp/mf-test.db" {
		t.Fatalf("cfg: %+v", cfg)
	}
}

func TestMigrateDialectFiles(t *testing.T) {
	for _, d := range []string{"sqlite", "postgres", "mysql"} {
		if _, err := dialectFile(d); err != nil {
			t.Fatalf("dialect %s: %v", d, err)
		}
	}
}
