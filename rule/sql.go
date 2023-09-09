package rule

import (
	"fmt"
	"os"
	"time"

	_ "github.com/go-sql-driver/mysql"

	"gopkg.in/yaml.v3"
)

func Sql(srcYaml, dstSql string) error {
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
	err = sqlGenerate(srcYaml, dstSql)
	return nil
}

func sqlGenerate(srcYaml, dstSql string) error {

	r, err := os.Open(srcYaml)
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

	err = generateRuleSQL(rule, dstSql)
	if err != nil {
		return fmt.Errorf("generateRuleSQL: %w", err)
	}
	return nil
}

func generateRuleSQL(sp *RuleYaml, dstSql string) error {
	w, err := os.Create(dstSql)
	if err != nil {
		return err
	}
	defer w.Close()

	for _, rule := range sp.Rules {
		w.WriteString(fmt.Sprintf("UPDATE `rule_values` SET rule_value = '%s' WHERE rule_name = '%s';\n", rule.Value, rule.Name))
	}
	return nil
}
