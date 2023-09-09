package aa

import (
	"fmt"
	"os"
	"time"

	"github.com/spf13/cobra"
	"gopkg.in/yaml.v3"
)

func Build(cmd *cobra.Command, args []string) error {
	cmd.SilenceUsage = true
	start := time.Now()
	db := &dbReader{}
	fmt.Printf("AA...")
	var err error
	defer func() {
		fmt.Println(" finished in", time.Since(start).String())
		if err != nil {
			fmt.Println("Error:", err)
			os.Exit(1)
		}
		fmt.Println("AA changed", db.changedDBStrCount, "DBStr entries")
	}()
	err = generate(db)

	return nil
}

func generate(db *dbReader) error {
	r, err := os.Open("aa.yaml")
	if err != nil {
		return err
	}
	defer r.Close()
	aa := &AAYaml{}
	dec := yaml.NewDecoder(r)
	dec.KnownFields(true)
	err = dec.Decode(aa)
	if err != nil {
		return fmt.Errorf("decode: %w", err)
	}

	err = aa.sanitize()
	if err != nil {
		return fmt.Errorf("sanitize: %w", err)
	}

	err = generateAASQL(aa)
	if err != nil {
		return fmt.Errorf("generateAASQL: %w", err)
	}

	err = modifyDBStr(db, aa)
	if err != nil {
		return fmt.Errorf("modifyDBStr: %w", err)
	}

	err = generateWeb(aa)
	if err != nil {
		return fmt.Errorf("generateWeb: %w", err)
	}
	return nil
}
