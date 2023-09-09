package spell

import (
	"database/sql"
	"fmt"
	"os"
	"reflect"
	"time"

	"github.com/fatih/structs"
	"github.com/spf13/cobra"
	"github.com/xackery/yakuku/util"
	"gopkg.in/yaml.v3"
)

func Build(cmd *cobra.Command, args []string) error {
	start := time.Now()
	db := &dbReader{}
	dbSpell := &dbSpellReader{}
	fmt.Printf("Spell...")
	var err error
	defer func() {
		fmt.Println(" finished in", time.Since(start).String())
		if err != nil {
			fmt.Printf("Error: %s\n", err)
			os.Exit(1)
		}
		fmt.Println("Spell changed", db.changedSpellsUSCount, "spells_us entries")
	}()
	err = generate(db, dbSpell)
	return nil
}

func generate(db *dbReader, dbSpell *dbSpellReader) error {
	r, err := os.Open("spell.yaml")
	if err != nil {
		return err
	}
	defer r.Close()
	spell := &SpellYaml{}
	dec := yaml.NewDecoder(r)
	dec.KnownFields(true)
	err = dec.Decode(spell)
	if err != nil {
		return fmt.Errorf("decode: %w", err)
	}

	err = spell.sanitize()
	if err != nil {
		return fmt.Errorf("spell sanitize: %w", err)
	}

	err = generateSpellSQL(spell)
	if err != nil {
		return fmt.Errorf("generateSpellSQL: %w", err)
	}

	err = modifySpellsUS(db, spell)
	if err != nil {
		return fmt.Errorf("modifySpellsUS: %w", err)
	}

	err = modifyDBStr(dbSpell, spell)
	if err != nil {
		return fmt.Errorf("modifyDBStr: %w", err)
	}

	return nil
}

func generateSpellSQL(sp *SpellYaml) error {

	w, err := os.Create("spell.sql")
	if err != nil {
		return err
	}
	defer w.Close()

	//(3,'Summon Corpse','PLAYER_1','','','','','','',10000,0,0,0,5000,1500,12000,0,0,0,700,70,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,2512,2106,17355,-1,-1,-1,1,1,1,1,-1,-1,-1,-1,100,100,100,100,100,100,100,100,100,100,100,100,0,-1,0,0,91,254,254,254,254,254,254,254,254,254,254,254,6,20,14,-1,0,0,255,255,255,255,35,255,255,255,255,255,35,255,255,255,255,255,43,0,0,4,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,100,0,0,98,0,0,0,0,0,0,0,0,0,3,125,64,0,-1,0,0,0,100,0,0,0,0,0,0,0,0,0,0,0,0,0,5,100,49,52,0,0,0,-1,-1,0,0,50,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,-1,0,0,0,1,0,0,0,1,1,0,-1,0,0,0,1,39,-1,0,1,0,0,1,0,1,0,0,0,0,0,0),
	for _, spell := range sp.Spells {
		w.WriteString("REPLACE INTO `spells_new` SET ")
		fields := structs.Fields(spell)
		for fieldIndex, field := range fields {
			if !field.IsExported() {
				continue
			}
			if field.Tag("db") == "" {
				continue
			}
			switch field.Kind() {
			case reflect.String:
				w.WriteString(fmt.Sprintf("`%s` = '%s'", field.Tag("db"), util.EscapeSQL(field.Value().(string))))
			case reflect.Int:
				w.WriteString(fmt.Sprintf("`%s` = %d", field.Tag("db"), field.Value()))
			case reflect.Float64:
				w.WriteString(fmt.Sprintf("`%s` = %f", field.Tag("db"), field.Value()))
			case reflect.Float32:
				w.WriteString(fmt.Sprintf("`%s` = %f", field.Tag("db"), field.Value()))
			case reflect.Bool:
				w.WriteString(fmt.Sprintf("`%s` = %t", field.Tag("db"), field.Value()))
			case reflect.Struct:
				switch val := field.Value().(type) {
				case time.Time:
					if field.Tag("db") == "updated" {
						w.WriteString(fmt.Sprintf("`%s` = NOW()", field.Tag("db")))
					} else {
						w.WriteString(fmt.Sprintf("`%s` = DATETIME('%s')", field.Tag("db"), val.Format("2006-01-02 15:04:05")))
					}
				case sql.NullString:
					if val.Valid {
						w.WriteString(fmt.Sprintf("`%s` = '%s'", field.Tag("db"), util.EscapeSQL(field.Value().(string))))
					} else {
						w.WriteString(fmt.Sprintf("`%s` = NULL", field.Tag("db")))
					}
				case sql.NullTime:
					if val.Valid {
						w.WriteString(fmt.Sprintf("`%s` = DATETIME('%s')", field.Tag("db"), val.Time.Format("2006-01-02 15:04:05")))
					} else {
						w.WriteString(fmt.Sprintf("`%s` = NULL", field.Tag("db")))
					}
				default:
					return fmt.Errorf("unknown type %s %s", field.Tag("db"), field.Kind())
				}
			default:
				return fmt.Errorf("unknown type %s %s", field.Tag("db"), field.Kind())
			}
			if fieldIndex == len(fields)-1 {
				w.WriteString(";\n")
			} else {
				w.WriteString(", ")
			}
		}
	}

	return nil
}
