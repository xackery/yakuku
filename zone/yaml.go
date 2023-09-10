package zone

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
func Yaml(path string, filters []string) error {
	cfg, err := config.Load()
	if err != nil {
		return fmt.Errorf("config load: %w", err)
	}

	db, err := sqlx.Connect("mysql", cfg.OriginalDB)
	if err != nil {
		return fmt.Errorf("db connect: %w", err)
	}
	defer db.Close()

	zones := &ZoneYaml{}

	for _, filter := range filters {
		id, err := strconv.Atoi(filter)
		if err != nil {
			return fmt.Errorf("invalid id %s: %w", filter, err)
		}
		result, err := importZone(db, id)
		if err != nil {
			return fmt.Errorf("import zone %d: %w", id, err)
		}
		if result == nil {
			continue
		}

		zones.Zones = append(zones.Zones, result)
	}

	w, err := os.Create(path)
	if err != nil {
		return err
	}
	defer w.Close()

	enc := yaml.NewEncoder(w)
	err = enc.Encode(zones)
	if err != nil {
		return err
	}

	fmt.Println("Zones from ids", strings.Join(filters, ", "), "exported to", path)
	return nil
}

func importZone(db *sqlx.DB, id int) (*Zone, error) {
	//start := time.Now()
	zone := &Zone{}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	rows, err := db.QueryxContext(ctx, "SELECT * FROM zone where zoneidnumber = ?", id)
	if err != nil {
		return nil, fmt.Errorf("db query: %w", err)
	}
	defer rows.Close()

	zoneShortName := ""
	for rows.Next() {
		err = rows.StructScan(zone)
		if err != nil {
			return nil, fmt.Errorf("db struct scan: %w", err)
		}
		zoneShortName = zone.ShortName.String
	}

	rows2, err := db.QueryxContext(ctx, "SELECT * from zone_points WHERE zone = ?", zoneShortName)
	if err != nil {
		return nil, fmt.Errorf("db query spawn2: %w", err)
	}

	for rows2.Next() {
		r := ZonePoint{}
		err = rows2.StructScan(&r)
		if err != nil {
			return nil, fmt.Errorf("db struct scan zone_points: %w", err)
		}
		zone.Points = append(zone.Points, &r)
	}

	fmt.Printf("%s (%d) has %d points\n", zoneShortName, id, len(zone.Points))
	return zone, nil
}
