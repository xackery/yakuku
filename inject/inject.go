package inject

import (
	"fmt"
	"os"
	"strconv"
	"strings"

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

	inserts := strings.Split(string(data), ";")

	totalRows := 0
	for _, insert := range inserts {
		insert = strings.TrimSpace(insert)
		if insert == "" {
			continue
		}
		result, err := db.Exec(insert)
		if err != nil {
			if strings.Contains(err.Error(), "at row") {
				row := strings.Split(err.Error(), "at row")
				if len(row) < 2 {
					return fmt.Errorf("exec split: %w", err)
				}
				rowNumber := strings.TrimSpace(row[1])

				rowNum, err := strconv.Atoi(rowNumber)
				if err != nil {
					return fmt.Errorf("exec atoi: %w", err)
				}

				records := strings.Split(insert, "\n")
				if len(records) < rowNum {
					return fmt.Errorf("exec len: %w", err)
				}

				fmt.Printf("Query line %d: %s\n", rowNum, records[rowNum-1])
			}
			return fmt.Errorf("exec: %w", err)
		}

		rows, err := result.RowsAffected()
		if err != nil {
			return fmt.Errorf("rows affected: %w", err)
		}
		totalRows += int(rows)
	}

	fmt.Printf("Injected %d rows\n", totalRows)
	return nil
}
