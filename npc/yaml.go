package npc

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

		zones.Npcs = append(zones.Npcs, result.Npcs...)
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

	fmt.Println("NPCs from zones", strings.Join(filters, ", "), "exported to", path)
	return nil
}

func importZone(db *sqlx.DB, id int) (*NpcYaml, error) {
	//start := time.Now()
	zoneName := util.ZoneIDToName(id)
	if zoneName == "unknown" {
		return nil, fmt.Errorf("unknown zone id %d", id)
	}
	npc := &NpcYaml{}
	npc.sanitize()

	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()
	rows, err := db.QueryxContext(ctx, "SELECT npc_types.* FROM npc_types where id < ? and id > ?", (id*1000)+1000, (id*1000)-1)
	if err != nil {
		return nil, fmt.Errorf("db query: %w", err)
	}
	defer rows.Close()

	for rows.Next() {
		r := Npc{}
		err = rows.StructScan(&r)
		if err != nil {
			return nil, fmt.Errorf("db struct scan: %w", err)
		}

		npc.Npcs = append(npc.Npcs, &r)
	}

	err = importSpawns(db, id, npc)
	if err != nil {
		return nil, fmt.Errorf("import spawns: %w", err)
	}

	rows2, err := db.QueryxContext(ctx, "SELECT * from spawn2 WHERE zone = ?", zoneName)
	if err != nil {
		return nil, fmt.Errorf("db query spawn2: %w", err)
	}

	for rows2.Next() {
		r := Spawn2{}
		err = rows2.StructScan(&r)
		if err != nil {
			return nil, fmt.Errorf("db struct scan spawn2: %w", err)
		}
		npc.Spawn2 = append(npc.Spawn2, &r)
	}

	if len(npc.Npcs) == 0 {
		fmt.Printf("%s (%d) has 0 npcs ", zoneName, id)
		return nil, nil
	}
	fmt.Printf("%s (%d) has %d npcs\n", zoneName, id, len(npc.Npcs))
	return npc, nil
}

func importSpawns(db *sqlx.DB, zoneID int, npc *NpcYaml) error {

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	rows, err := db.QueryxContext(ctx, "SELECT * FROM spawnentry INNER JOIN spawngroup ON spawngroup.id = spawngroupid where npcID < ? and npcID > ?", (zoneID*1000)+1000, (zoneID*1000)-1)
	if err != nil {
		return fmt.Errorf("db query: %w", err)
	}
	defer rows.Close()

	for rows.Next() {
		r := Spawn{}
		err = rows.StructScan(&r)
		if err != nil {
			return fmt.Errorf("db struct scan: %w", err)
		}

		for _, npc := range npc.Npcs {
			if r.NpcID != npc.ID {
				continue
			}
			if npc.Spawns == nil {
				npc.Spawns = make([]*Spawn, 0)
			}
			r.NpcID = 0
			r.Id = 0
			npc.Spawns = append(npc.Spawns, &r)
		}

	}
	return nil
}
