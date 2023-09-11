package spell

import (
	"bufio"
	"context"
	"fmt"
	"os"
	"strconv"
	"strings"
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

	if len(args) < 2 {
		fmt.Println("need spell id for import, too big to do everything (-1 if you do anyways)")
		os.Exit(1)
	}

	db, err := sqlx.Connect("mysql", fmt.Sprintf("%s:%s@tcp(%s)/%s?parseTime=true&multiStatements=true&interpolateParams=true&collation=utf8mb4_unicode_ci&charset=utf8mb4,utf8", viper.GetString("db_user"), viper.GetString("db_pass"), viper.GetString("db_host"), viper.GetString("db_name")))
	if err != nil {
		return fmt.Errorf("db connect: %w", err)
	}
	defer db.Close()

	id, err := strconv.Atoi(args[1])
	if err != nil {
		return fmt.Errorf("invalid id: %w", err)
	}

	err = importSpell(db, id)
	if err != nil {
		return fmt.Errorf("import item %d: %w", id, err)
	}

	return nil
}

func importSpell(db *sqlx.DB, id int) error {

	spell := &SpellYaml{}
	err := spell.sanitize()
	if err != nil {
		return err
	}

	descIDs := make(map[int]string)
	err = util.PrepFile("dbstr_us", ".txt")
	if err != nil {
		return err
	}

	defer func() {
		os.Remove("dbstr_us_tmp.txt")
	}()

	r, err := os.Open("dbstr_us_tmp.txt")
	if err != nil {
		return err
	}
	defer r.Close()
	scanner := bufio.NewScanner(r)

	lineNumber := 0
	for scanner.Scan() {
		lineNumber++
		line := scanner.Text()
		records := strings.Split(line, "^")
		id, err := strconv.Atoi(records[0])
		if err != nil {
			return fmt.Errorf("strconv spell id line %d: %w", lineNumber, err)
		}
		category, err := strconv.Atoi(records[1])
		if err != nil {
			return fmt.Errorf("strconv spell category line %d: %w", lineNumber, err)
		}
		switch category {
		case 6:
			descIDs[id] = strings.TrimSpace(records[2])
			if id == 907 {
				fmt.Println(id)
			}
		}
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var rows *sqlx.Rows

	outName := "spell_dump.yaml"
	if id == -1 {
		rows, err = db.QueryxContext(ctx, "SELECT * FROM spells_new")
	} else {
		rows, err = db.QueryxContext(ctx, "SELECT * FROM spells_new WHERE id = ?", id)
		outName = fmt.Sprintf("spell_dump_%d.yaml", id)
	}
	if err != nil {
		return fmt.Errorf("db query: %w", err)
	}
	defer rows.Close()

	for rows.Next() {
		r := Spell{}
		err = rows.StructScan(&r)
		if err != nil {
			return fmt.Errorf("db struct scan: %w", err)
		}

		r.Description = descIDs[r.ID]

		err = r.omitEmpty()
		if err != nil {
			return fmt.Errorf("omit empty: %w", err)
		}
		spell.Spells = append(spell.Spells, &r)
	}

	/*node := yaml.Node{
		Kind:        yaml.ScalarNode,
		Value:       "value",
		LineComment: "Test!",
		Line:        1,
		Column:      1,
	}*/

	if len(spell.Spells) == 0 {
		fmt.Println("No spells found")
		return nil
	}

	w, err := os.Create(outName)
	if err != nil {
		return err
	}
	defer w.Close()

	//err = node.Encode(w)
	enc := yaml.NewEncoder(w)
	enc.SetIndent(2)
	err = enc.Encode(spell)
	if err != nil {
		return err
	}
	fmt.Println("Created spell_dump.yaml")

	return nil
}
