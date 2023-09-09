package rule

import (
	"fmt"
	"os"
	"time"

	_ "github.com/go-sql-driver/mysql"

	"gopkg.in/yaml.v3"
)

func Sql(path string) error {
	start := time.Now()
	fmt.Printf("Rule...")
	var err error
	defer func() {
		fmt.Println(" finished in", time.Since(start).String())
		if err != nil {
			fmt.Println("Error:", err)
			os.Exit(1)
		}
	}()
	err = sqlGenerate(path)
	return nil
}

func sqlGenerate(path string) error {

	r, err := os.Open("rule.yaml")
	if err != nil {
		return err
	}
	defer r.Close()
	rule := &RuleYaml{}
	dec := yaml.NewDecoder(r)
	dec.KnownFields(true)
	err = dec.Decode(rule)
	if err != nil {
		return fmt.Errorf("decode: %w", err)
	}

	err = rule.sanitize()
	if err != nil {
		return fmt.Errorf("rule sanitize: %w", err)
	}

	err = generateRuleSQL(rule)
	if err != nil {
		return fmt.Errorf("generateRuleSQL: %w", err)
	}
	return nil
}

func generateRuleSQL(sp *RuleYaml) error {
	w, err := os.Create("rule.sql")
	if err != nil {
		return err
	}
	defer w.Close()

	for _, rule := range sp.Rules {
		w.WriteString(fmt.Sprintf("UPDATE `rule_values` SET rule_value = '%s' WHERE rule_name = '%s';\n", rule.Value, rule.Name))
	}

	return nil
}
