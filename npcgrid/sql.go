package npcgrid

import (
	"fmt"
	"os"
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
	buf := ""

	for _, npc := range sp.Spawn2 {
		fields := structs.Fields(npc)

		buf += "REPLACE INTO `spawn2` SET "
		for _, field := range fields {
			if !field.IsExported() {
				continue
			}

			fieldBuf := util.FieldParse(field)
			if fieldBuf != "" {
				buf += fieldBuf + ", "
				continue
			}

			return fmt.Errorf("unknown type %s %s", field.Tag("db"), field.Kind())
		}
		buf = buf[:len(buf)-2]
		buf += ";\n"
		w.WriteString(buf)
		npcCount++
	}

	gridCount := 0
	for _, grid := range sp.Grids {
		fields := structs.Fields(grid)

		buf += "REPLACE INTO `grid` SET "
		for _, field := range fields {
			if !field.IsExported() {
				continue
			}

			if field.Tag("db") == "zoneid" {
				buf += fmt.Sprintf("`zoneid` = '%d', ", sp.ZoneID)
				continue
			}

			fieldBuf := util.FieldParse(field)
			if fieldBuf != "" {
				buf += fieldBuf + ", "
				continue
			}

			return fmt.Errorf("unknown type %s %s", field.Tag("db"), field.Kind())
		}
		buf = buf[:len(buf)-2]
		buf += ";\n"
		w.WriteString(buf)
		gridCount++
		for _, gridEntry := range grid.Entries {
			if gridEntry.Gridid != grid.Id {
				continue
			}
			fields := structs.Fields(gridEntry)

			buf += "REPLACE INTO `grid_entries` SET "
			for _, field := range fields {
				if !field.IsExported() {
					continue
				}

				if field.Tag("db") == "zoneid" {
					buf += fmt.Sprintf("`zoneid` = '%d', ", sp.ZoneID)
					continue
				}

				fieldBuf := util.FieldParse(field)
				if fieldBuf != "" {
					buf += fieldBuf + ", "
					continue
				}

				return fmt.Errorf("unknown type %s %s", field.Tag("db"), field.Kind())
			}
			buf = buf[:len(buf)-2]
			buf += ";\n"
			w.WriteString(buf)
		}
	}

	fmt.Printf(" %d npcs ", npcCount)

	return nil
}
