package npcgrid

import (
	"context"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/jmoiron/sqlx"
	"github.com/xackery/yakuku/config"
	"github.com/xackery/yakuku/util"
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

	zones := &NpcYaml{}

	if len(filters) > 1 {
		return fmt.Errorf("only one filter is supported")
	}

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
		zones.ZoneID = result.ZoneID
		zones.ZoneShortName = result.ZoneShortName

		zones.Spawn2 = append(zones.Spawn2, result.Spawn2...)
		zones.Grids = append(zones.Grids, result.Grids...)
	}

	w, err := os.Create(path)
	if err != nil {
		return err
	}
	defer w.Close()

	enc := yaml.NewEncoder(w)
	enc.SetIndent(2)
	err = enc.Encode(zones)
	if err != nil {
		return err
	}

	fmt.Println("NPCs from zones", strings.Join(filters, ", "), "exported to", path)
	return nil
}

func importZone(db *sqlx.DB, id int) (*NpcYaml, error) {
	//start := time.Now()
	zoneName := util.ZoneIDToName(id)
	if zoneName == "unknown" {
		return nil, fmt.Errorf("unknown zone id %d", id)
	}
	npc := &NpcYaml{
		ZoneShortName: zoneName,
		ZoneID:        id,
	}
	npc.sanitize()

	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()
	rows, err := db.QueryxContext(ctx, "SELECT id, spawngroupid, zone, x, y, z, heading, pathgrid FROM spawn2 where zone = ?", zoneName)
	if err != nil {
		return nil, fmt.Errorf("db query: %w", err)
	}
	defer rows.Close()

	for rows.Next() {
		r := Spawn2{}
		err = rows.StructScan(&r)
		if err != nil {
			return nil, fmt.Errorf("db struct scan: %w", err)
		}
		r.Zone = ""
		npc.Spawn2 = append(npc.Spawn2, &r)
	}

	rows2, err := db.QueryxContext(ctx, "SELECT * from grid WHERE zoneid = ?", id)
	if err != nil {
		return nil, fmt.Errorf("db query spawn2: %w", err)
	}

	for rows2.Next() {
		r := PathGrid{}
		err = rows2.StructScan(&r)
		if err != nil {
			return nil, fmt.Errorf("db struct scan spawn2: %w", err)
		}
		r.Zoneid = 0
		npc.Grids = append(npc.Grids, &r)
	}

	gridEntryTotal := 0

	rows3, err := db.QueryxContext(ctx, "SELECT * from grid_entries WHERE zoneid = ?", id)
	if err != nil {
		return nil, fmt.Errorf("db query spawn2: %w", err)
	}

	for rows3.Next() {
		r := PathGridEntry{}
		err = rows3.StructScan(&r)
		if err != nil {
			return nil, fmt.Errorf("db struct scan spawn2: %w", err)
		}
		gridEntryTotal++
		for _, grid := range npc.Grids {
			if grid.Id != r.Gridid {
				continue
			}
			r.Gridid = 0
			r.Zoneid = 0

			grid.Entries = append(grid.Entries, &r)
		}
	}

	fmt.Printf("%s (%d) has %d spawn2, %d grids, %d entries\n", zoneName, id, len(npc.Spawn2), len(npc.Grids), gridEntryTotal)
	return npc, nil
}
