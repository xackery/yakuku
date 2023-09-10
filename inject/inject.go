package inject

import (
	"fmt"
	"os"

	"github.com/jmoiron/sqlx"
	"github.com/xackery/yakuku/config"
)

// Inject inserts a sql file into the database
func Inject(sqlFile string) error {
	cfg, err := config.Load()
	if err != nil {
		return fmt.Errorf("config load: %w", err)
	}

	db, err := sqlx.Connect("mysql", cfg.InjectDB)
	if err != nil {
		return fmt.Errorf("sqlx connect: %w", err)
	}
	defer db.Close()

	data, err := os.ReadFile(sqlFile)
	if err != nil {
		return fmt.Errorf("read file: %w", err)
	}

	result, err := db.Exec(string(data))
	if err != nil {
		return fmt.Errorf("exec: %w", err)
	}

	rows, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("rows affected: %w", err)
	}

	fmt.Printf("Injected %d rows\n", rows)
	return nil
}
