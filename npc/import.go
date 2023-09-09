package npc

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/jmoiron/sqlx"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/xackery/yakuku/util"
	"gopkg.in/yaml.v3"
)

// Import takes database info and dumps to yaml
func Import(cmd *cobra.Command, args []string) error {
	if !viper.IsSet("db_host") {
		return fmt.Errorf("db_host is not set, pass it as an argument --db_host=... or set it in .luaject.yaml")
	}

	err := os.MkdirAll("npc", 0755)
	if err != nil {
		return err
	}

	db, err := sqlx.Connect("mysql", fmt.Sprintf("%s:%s@tcp(%s)/%s?parseTime=true&multiStatements=true&interpolateParams=true&collation=utf8mb4_unicode_ci&charset=utf8mb4,utf8", viper.GetString("db_user"), viper.GetString("db_pass"), viper.GetString("db_host"), viper.GetString("db_name")))
	if err != nil {
		return fmt.Errorf("db connect: %w", err)
	}
	defer db.Close()

	for id := 0; id < 1000; id++ {
		err = importZone(db, id)
		if err != nil {
			return fmt.Errorf("import zone %d: %w", id, err)
		}
	}

	return nil
}

func importZone(db *sqlx.DB, id int) error {
	//start := time.Now()
	zoneName := util.ZoneIDToName(id)
	if zoneName == "unknown" {
		if id > 0 {
			return nil
		}
		zoneName = "global"
	}
	npc := &NpcYaml{}
	npc.sanitize()

	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()
	rows, err := db.QueryxContext(ctx, "SELECT npc_types.* FROM npc_types where id < ? and id > ?", (id*1000)+1000, (id*1000)-1)
	if err != nil {
		return fmt.Errorf("db query: %w", err)
	}
	defer rows.Close()

	for rows.Next() {
		r := Npc{}
		err = rows.StructScan(&r)
		if err != nil {
			return fmt.Errorf("db struct scan: %w", err)
		}

		npc.Npcs = append(npc.Npcs, &r)
	}

	err = importSpawns(db, id, npc)
	if err != nil {
		return fmt.Errorf("import spawns: %w", err)
	}

	rows2, err := db.QueryxContext(ctx, "SELECT * from spawn2 WHERE zone = ?", zoneName)
	if err != nil {
		return fmt.Errorf("db query spawn2: %w", err)
	}

	for rows2.Next() {
		r := Spawn2{}
		err = rows2.StructScan(&r)
		if err != nil {
			return fmt.Errorf("db struct scan spawn2: %w", err)
		}
		npc.Spawn2 = append(npc.Spawn2, &r)
	}

	if len(npc.Npcs) == 0 {
		fmt.Println("No npcs found in", zoneName)
		return nil
	}
	w, err := os.Create(fmt.Sprintf("npc/%s.yaml", zoneName))
	if err != nil {
		return err
	}
	defer w.Close()

	enc := yaml.NewEncoder(w)
	err = enc.Encode(npc)
	if err != nil {
		return err
	}
	//fmt.Printf("Created npc/%s.yaml in %0.2f seconds\n", zoneName, time.Since(start).Seconds())
	return nil
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
			npc.Spawns = append(npc.Spawns, &r)
		}

	}
	return nil
}
