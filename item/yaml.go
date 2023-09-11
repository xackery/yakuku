package item

import (
	"context"
	"fmt"
	"os"
	"strconv"
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

	items := &ItemYaml{}

	for _, filter := range filters {
		id, err := strconv.Atoi(filter)
		if err != nil {
			return fmt.Errorf("invalid id %s: %w", filter, err)
		}
		result, err := importItem(db, id)
		if err != nil {
			return fmt.Errorf("import item %d: %w", id, err)
		}

		items.Items = append(items.Items, result.Items...)
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

func importItem(db *sqlx.DB, id int) (*ItemYaml, error) {
	//start := time.Now()
	item := &ItemYaml{}

	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()

	var rows *sqlx.Rows
	var err error

	if id == -1 {
		rows, err = db.QueryxContext(ctx, "SELECT * from items")
	} else {
		rows, err = db.QueryxContext(ctx, "SELECT * from items WHERE id = ?", id)
	}
	if err != nil {
		return nil, fmt.Errorf("db query: %w", err)
	}
	defer rows.Close()

	for rows.Next() {
		r := Item{}
		err = rows.StructScan(&r)
		if err != nil {
			return nil, fmt.Errorf("db struct scan: %w", err)
		}

		item.Items = append(item.Items, &r)
	}

	if len(item.Items) == 0 {
		return nil, fmt.Errorf("no items found")
	}
	err = item.omitEmpty()
	if err != nil {
		return nil, fmt.Errorf("omit empty: %w", err)
	}

	return item, nil
}
