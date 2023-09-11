package startingitems

import (
	"context"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/jmoiron/sqlx"
	"github.com/xackery/yakuku/config"
	"gopkg.in/yaml.v3"
)

// Yaml takes database info and dumps to yaml
func Yaml(yamlFile string, filters []string) error {
	cfg, err := config.Load()
	if err != nil {
		return fmt.Errorf("config load: %w", err)
	}

	db, err := sqlx.Connect("mysql", cfg.OriginalDB)
	if err != nil {
		return fmt.Errorf("db connect: %w", err)
	}
	defer db.Close()

	items, err := importStartingItems(db)
	if err != nil {
		return fmt.Errorf("import item: %w", err)
	}

	w, err := os.Create(yamlFile)
	if err != nil {
		return err
	}
	defer w.Close()

	enc := yaml.NewEncoder(w)
	enc.SetIndent(2)
	err = enc.Encode(items)
	if err != nil {
		return err
	}

	fmt.Println("Items", strings.Join(filters, ", "), "exported to", yamlFile)
	return nil
}

func importStartingItems(db *sqlx.DB) (*StartingItemYaml, error) {
	//start := time.Now()
	item := &StartingItemYaml{}

	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()

	var rows *sqlx.Rows
	var err error

	rows, err = db.QueryxContext(ctx, "SELECT starting_items.*, items.name FROM starting_items INNER JOIN items ON items.id = starting_items.itemid")
	if err != nil {
		return nil, fmt.Errorf("db query: %w", err)
	}
	defer rows.Close()

	for rows.Next() {
		r := StartingItem{}
		err = rows.StructScan(&r)
		if err != nil {
			return nil, fmt.Errorf("db struct scan: %w", err)
		}

		item.StartingItems = append(item.StartingItems, &r)
	}

	if len(item.StartingItems) == 0 {
		return nil, fmt.Errorf("no items found")
	}
	err = item.omitEmpty()
	if err != nil {
		return nil, fmt.Errorf("omit empty: %w", err)
	}

	return item, nil
}
