package npc

import (
	"fmt"
	"os"
	"reflect"
	"time"

	"github.com/fatih/structs"
	_ "github.com/go-sql-driver/mysql"
	"github.com/xackery/yakuku/util"

	"gopkg.in/yaml.v3"
)

func Sql(srcYaml, dstSql string) error {
	start := time.Now()
	fmt.Printf("NPC... ")
	var err error
	defer func() {
		fmt.Println("finished in", time.Since(start).String())
		if err != nil {
			fmt.Println("Error:", err)
			os.Exit(1)
		}
	}()
	err = npcGenerate(srcYaml, dstSql)
	return nil
}

func npcGenerate(srcYaml, dstSql string) error {
	r, err := os.Open(srcYaml)
	if err != nil {
		return err
	}
	defer r.Close()
	npc := &NpcYaml{}
	dec := yaml.NewDecoder(r)
	dec.KnownFields(true)
	err = dec.Decode(npc)
	if err != nil {
		return fmt.Errorf("decode: %w", err)
	}

	err = npc.sanitize()
	if err != nil {
		return fmt.Errorf("npc sanitize: %w", err)
	}

	err = generateNpcSQL(npc, dstSql)
	if err != nil {
		return fmt.Errorf("generateNpcSQL: %w", err)
	}
	return nil
}

func generateNpcSQL(sp *NpcYaml, dstSql string) error {
	w, err := os.Create(dstSql)
	if err != nil {
		return err
	}
	defer w.Close()

	npcCount := 0
	spawnBuf := ""
	buf := ""

	spawnEntries := map[int][]*Spawn{}

	for _, npc := range sp.Npcs {
		fields := structs.Fields(npc)

		buf += fmt.Sprintf("REPLACE INTO `npcs` SET ")
		for _, field := range fields {
			if !field.IsExported() {
				continue
			}

			fieldBuf := util.FieldParse(field)
			if fieldBuf != "" {
				buf += fieldBuf + ", "
				continue
			}

			switch field.Kind() {
			case reflect.Slice:
				// assert type
				switch v := field.Value().(type) {
				case []*Spawn:
					if len(v) == 0 {
						continue
					}
					for _, spawn := range v {
						spawnEntries[spawn.SpawngroupID] = append(spawnEntries[spawn.SpawngroupID], spawn)
					}
				case []*Spawn2:
					if len(v) == 0 {
						continue
					}
					spawnEntry, err := generateSpawn2SQL(v)
					if err != nil {
						return fmt.Errorf("generateSpawn2SQL: %w", err)
					}
					spawnBuf += spawnEntry
				}
				continue
			default:
				return fmt.Errorf("unknown type %s %s", field.Tag("db"), field.Kind())
			}
		}
		buf = buf[:len(buf)-2]
		buf += ";\n"
		w.WriteString(buf)
		npcCount++
	}

	spawnGroupBuf := ""

	for _, spawns := range spawnEntries {
		spawnEntry, err := generateSpawnEntrySQL(spawns)
		if err != nil {
			return fmt.Errorf("generateSpawnEntrySQL: %w", err)
		}
		spawnBuf += spawnEntry

		spawnGroup, err := generateSpawnGroupSQL(spawns)
		if err != nil {
			return fmt.Errorf("generateSpawnGroupSQL: %w", err)
		}
		spawnGroupBuf += spawnGroup
	}

	w.WriteString("\n" + spawnGroupBuf)
	w.WriteString("\n" + spawnBuf)

	fmt.Printf(" %d npcs ", npcCount)

	return nil
}

func generateSpawnEntrySQL(spawns []*Spawn) (string, error) {
	buf := ""

	spawnEntryFields := []string{
		"spawngroupID",
		"npcID",
		"chance",
		"condition_value_filter",
		"min_expansion",
		"max_expansion",
		"content_flags",
		"content_flags_disabled",
	}
	for _, spawn := range spawns {
		fields := structs.Fields(spawn)

		buf += fmt.Sprintf("REPLACE INTO `spawnentry` SET ")
		for _, field := range fields {
			if !field.IsExported() {
				continue
			}

			isSpawnEntry := false
			for _, spawnEntryField := range spawnEntryFields {
				if field.Tag("db") != spawnEntryField {
					continue
				}
				isSpawnEntry = true
				break
			}

			if !isSpawnEntry {
				continue
			}

			fieldBuf := util.FieldParse(field)
			if fieldBuf != "" {
				buf += fieldBuf + ", "
				continue
			}

			return "", fmt.Errorf("unknown type %s %s", field.Tag("db"), field.Kind())
		}
		buf = buf[:len(buf)-2]
		buf += ";\n"
	}

	return buf, nil
}

func generateSpawn2SQL(spawns []*Spawn2) (string, error) {
	buf := ""

	spawn2Fields := []string{
		"id",
		"spawngroupID",
		"zone",
		"version",
		"x",
		"y",
		"z",
		"heading",
		"respawntime",
		"variance",
		"pathgrid",
		"path_when_zone_idle",
		"_condition",
		"cond_value",
		"enabled",
		"animation",
		"min_expansion",
		"max_expansion",
		"content_flags",
		"content_flags_disabled",
	}
	for _, spawn := range spawns {
		fields := structs.Fields(spawn)

		buf += "REPLACE INTO `spawn2` SET "
		for _, field := range fields {
			if !field.IsExported() {
				continue
			}

			isSpawn2 := false

			for _, spawn2Field := range spawn2Fields {
				if field.Tag("db") != spawn2Field {
					continue
				}
				isSpawn2 = true
				break
			}

			if !isSpawn2 {
				continue
			}
			fieldBuf := util.FieldParse(field)
			if fieldBuf != "" {
				buf += fieldBuf + ", "
				continue
			}
			return "", fmt.Errorf("unknown type %s %s", field.Tag("db"), field.Kind())
		}
		buf = buf[:len(buf)-2]
		buf += ";\n"
	}

	return buf, nil
}

func generateSpawnGroupSQL(spawns []*Spawn) (string, error) {
	buf := ""

	spawnGroupFields := []string{
		"id",
		"name",
		"spawn_limit",
		"dist",
		"max_x",
		"min_x",
		"max_y",
		"min_y",
		"delay",
		"mindelay",
		"despawn",
		"despawn_timer",
		"wp_spawns",
	}
	for _, spawn := range spawns {
		fields := structs.Fields(spawn)

		buf += "REPLACE INTO `spawngroup` SET "
		for _, field := range fields {
			if !field.IsExported() {
				continue
			}

			isSpawn2 := false

			for _, spawn2Field := range spawnGroupFields {
				if field.Tag("db") != spawn2Field {
					continue
				}
				isSpawn2 = true
				break
			}

			if !isSpawn2 {
				continue
			}
			fieldBuf := util.FieldParse(field)
			if fieldBuf != "" {
				buf += fieldBuf + ", "
				continue
			}
			return "", fmt.Errorf("unknown type %s %s", field.Tag("db"), field.Kind())
		}
		buf = buf[:len(buf)-2]
		buf += ";\n"
	}

	return buf, nil
}
