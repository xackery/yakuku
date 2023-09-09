package task

import (
	"database/sql"
	"fmt"
	"os"
	"reflect"
	"time"

	"github.com/fatih/structs"
	"github.com/spf13/cobra"
	"gopkg.in/yaml.v3"
)

func Build(cmd *cobra.Command, args []string) error {
	start := time.Now()
	fmt.Printf("Task...")
	var err error
	defer func() {
		fmt.Println(" finished in", time.Since(start).String())
		if err != nil {
			fmt.Println("Error:", err)
			os.Exit(1)
		}
	}()
	err = generate()

	return nil
}

func generate() error {
	r, err := os.Open("task.yaml")
	if err != nil {
		return err
	}
	defer r.Close()
	task := &TaskYaml{}
	dec := yaml.NewDecoder(r)
	dec.KnownFields(true)
	err = dec.Decode(task)
	if err != nil {
		return fmt.Errorf("decode: %w", err)
	}

	err = task.sanitize()
	if err != nil {
		return fmt.Errorf("task sanitize: %w", err)
	}

	err = generateTaskSQL(task)
	if err != nil {
		return fmt.Errorf("generateTaskSQL: %w", err)
	}

	return nil
}

func generateTaskSQL(sp *TaskYaml) error {
	w, err := os.Create("task.sql")
	if err != nil {
		return err
	}
	defer w.Close()

	w.WriteString("REPLACE INTO `tasks` (`taskid`, `activityid`, `req_activity_id`, `step`, `activitytype`, `target_name`, `goalmethod`, `goalcount`, `description_override`, `npc_match_list`, `item_id_list`, `item_list`, `dz_switch_id`, `min_x`, `min_y`, `min_z`, `max_x`, `max_y`, `max_z`, `skill_list`, `spell_list`, `zones`, `zone_version`, `optional`) VALUES\n")

	for taskIndex, task := range sp.Tasks {
		fields := structs.Fields(task)
		w.WriteString("(")
		for fieldIndex, field := range fields {
			if !field.IsExported() {
				continue
			}
			switch field.Kind() {
			case reflect.Slice:
				// assert type
				switch v := field.Value().(type) {
				case []*Activity:
					if len(v) == 0 {
						continue
					}
					//ignore for now
				}
			case reflect.Struct:
				// assert type
				switch v := field.Value().(type) {
				case sql.NullString:
					if v.Valid {
						w.WriteString(fmt.Sprintf("\"%s\"", v.String))
					} else {
						w.WriteString("NULL")
					}
				}
			case reflect.String:
				w.WriteString(fmt.Sprintf("\"%s\"", field.Value()))
			case reflect.Int:
				w.WriteString(fmt.Sprintf("%d", field.Value()))
			case reflect.Float64:
				w.WriteString(fmt.Sprintf("%f", field.Value()))
			case reflect.Bool:
				w.WriteString(fmt.Sprintf("%t", field.Value()))
			default:
				return fmt.Errorf("unknown type %s", field.Kind())
			}
			if fieldIndex == len(fields)-1 {
				w.WriteString("")
			} else {
				w.WriteString(", ")
			}
		}
		if taskIndex == len(sp.Tasks)-1 {
			w.WriteString(");\n")
		} else {
			w.WriteString("),\n")
		}
	}

	return nil
}
