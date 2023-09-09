package spell

import (
	"bufio"
	"fmt"
	"os"
	"reflect"
	"strconv"
	"strings"

	"github.com/fatih/structs"
	"github.com/xackery/yakuku/util"
)

type dbReader struct {
	SIDs                 map[string]string
	lastLine             string
	lastSID              int
	changedSpellsUSCount int
}

func (d *dbReader) line(scanner *bufio.Scanner, sid int) string {

	line, ok := d.SIDs[fmt.Sprintf("%d", sid)]
	if ok {
		d.changedSpellsUSCount++
		return line
	}

	var err error
	//fmt.Println("current", sid, category, "| last", d.lastSID, d.lastCategory)
	if d.lastLine == "" || d.lastSID < sid {
		if !scanner.Scan() {
			return ""
		}

		d.lastLine = scanner.Text()
		//fmt.Println("grabbing new line")
		lineBreakdown := strings.Split(d.lastLine, "^")
		d.lastSID, err = strconv.Atoi(lineBreakdown[0])
		if err != nil {
			fmt.Printf("line atoi %s: %s\n", d.lastLine, err)
			os.Exit(1)
		}
	}

	if d.lastSID > sid {
		return ""
	}

	line = d.lastLine

	//fmt.Println("inserting line, grabbing next")
	if !scanner.Scan() {
		return line
	}
	d.lastLine = scanner.Text()
	lineBreakdown := strings.Split(d.lastLine, "^")
	d.lastSID, err = strconv.Atoi(lineBreakdown[0])
	if err != nil {
		fmt.Printf("line atoi %s: %s\n", d.lastLine, err)
		os.Exit(1)
	}
	return line
}

func modifySpellsUS(db *dbReader, sp *SpellYaml) error {
	var err error

	err = util.PrepFile("spells_us", ".txt")
	if err != nil {
		return err
	}

	defer func() {
		os.Remove("spells_us_tmp.txt")
	}()

	r, err := os.Open("spells_us_tmp.txt")
	if err != nil {
		return err
	}
	defer r.Close()
	scanner := bufio.NewScanner(r)

	db.SIDs = map[string]string{}
	for _, spell := range sp.Spells {
		key := fmt.Sprintf("%d", spell.ID)
		fields := structs.Fields(spell)
		line := ""
		for fieldIndex, field := range fields {
			if !field.IsExported() {
				continue
			}
			if field.Tag("db") == "" {
				continue
			}
			switch field.Kind() {
			case reflect.String:
				line += fmt.Sprintf("%s", field.Value())
			case reflect.Int:
				value := field.Value().(int)
				if (field.Name() == "Range" || field.Name() == "Aoerange") && field.Value() == 0 {
				} else {

					line += fmt.Sprintf("%d", value)
				}
			case reflect.Float64:
				line += fmt.Sprintf("%0.0f", field.Value())
			case reflect.Float32:
				line += fmt.Sprintf("%0.0f", field.Value())
			case reflect.Bool:
				line += fmt.Sprintf("%t", field.Value())
			default:
				return fmt.Errorf("unknown type %s", field.Kind())
			}
			if fieldIndex == len(fields)-1 {
				//line += "^"
			} else {
				line += "^"
			}
		}
		db.SIDs[key] = line
	}

	w, err := os.Create("spells_us.txt")
	if err != nil {
		return err
	}
	defer w.Close()

	for sid := 0; sid < 45000; sid++ {
		line := db.line(scanner, sid)
		if line == "" {
			continue
		}

		_, err = w.WriteString(line + "\n")
		if err != nil {
			return fmt.Errorf("line %s writeString: %w", line, err)
		}
	}

	return nil
}
