package item

import (
	"database/sql"
	"fmt"
	"os"
	"reflect"
	"time"

	"github.com/fatih/structs"
	_ "github.com/go-sql-driver/mysql"

	"github.com/spf13/cobra"
	"gopkg.in/yaml.v3"
)

func Build(cmd *cobra.Command, args []string) error {
	start := time.Now()
	fmt.Printf("Item...")
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

	r, err := os.Open("item.yaml")
	if err != nil {
		return err
	}
	defer r.Close()
	item := &ItemYaml{}
	dec := yaml.NewDecoder(r)
	dec.KnownFields(true)
	err = dec.Decode(item)
	if err != nil {
		return fmt.Errorf("decode: %w", err)
	}

	err = item.sanitize()
	if err != nil {
		return fmt.Errorf("item sanitize: %w", err)
	}

	err = generateItemSQL(item)
	if err != nil {
		return fmt.Errorf("generateItemSQL: %w", err)
	}
	return nil
}

func generateItemSQL(sp *ItemYaml) error {
	w, err := os.Create("item.sql")
	if err != nil {
		return err
	}
	defer w.Close()

	for _, item := range sp.Items {
		fields := structs.Fields(item)

		w.WriteString("REPLACE INTO `items` SET ")
		for fieldIndex, field := range fields {
			if !field.IsExported() {
				continue
			}
			switch field.Kind() {
			case reflect.String:
				w.WriteString(fmt.Sprintf("`%s` = '%s'", field.Tag("db"), field.Value()))
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
						w.WriteString(fmt.Sprintf("`%s` = '%s'", field.Tag("db"), field.Value()))
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
		//w.WriteString(fmt.Sprintf(" WHERE id = %d;\n", item.ID))
	}

	return nil
}
