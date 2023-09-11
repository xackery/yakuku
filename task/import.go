package task

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/jmoiron/sqlx"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"gopkg.in/yaml.v3"
)

// Import takes database info and dumps to yaml
func Import(cmd *cobra.Command, args []string) error {
	if !viper.IsSet("db_host") {
		return fmt.Errorf("db_host is not set, pass it as an argument --db_host=... or set it in .luaject.yaml")
	}

	db, err := sqlx.Connect("mysql", fmt.Sprintf("%s:%s@tcp(%s)/%s?parseTime=true&multiStatements=true&interpolateParams=true&collation=utf8mb4_unicode_ci&charset=utf8mb4,utf8", viper.GetString("db_user"), viper.GetString("db_pass"), viper.GetString("db_host"), viper.GetString("db_name")))
	if err != nil {
		return fmt.Errorf("db connect: %w", err)
	}
	defer db.Close()

	task := &TaskYaml{}
	err = task.sanitize()
	if err != nil {
		return err
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	rows, err := db.QueryxContext(ctx, "SELECT * FROM tasks")
	if err != nil {
		return fmt.Errorf("db query: %w", err)
	}
	defer rows.Close()

	for rows.Next() {
		r := Task{}
		err = rows.StructScan(&r)
		if err != nil {
			return fmt.Errorf("db struct scan: %w", err)
		}

		activityRows, err := db.QueryxContext(ctx, "SELECT * FROM task_activities WHERE taskid = ?", r.ID)
		if err != nil {
			return fmt.Errorf("db query activities for taskid %d: %w", r.ID, err)
		}

		for activityRows.Next() {
			a := Activity{}
			err = activityRows.StructScan(&a)
			if err != nil {
				return fmt.Errorf("db struct scan activity: %w", err)
			}
			a.TaskID = 0
			r.Activities = append(r.Activities, &a)
		}

		err = r.omitEmpty()
		if err != nil {
			return fmt.Errorf("omit empty: %w", err)
		}
		task.Tasks = append(task.Tasks, &r)
	}

	w, err := os.Create("task_dump.yaml")
	if err != nil {
		return err
	}
	defer w.Close()

	enc := yaml.NewEncoder(w)
	enc.SetIndent(2)
	err = enc.Encode(task)
	if err != nil {
		return err
	}
	fmt.Println("Created task_dump.yaml")
	return nil
}
