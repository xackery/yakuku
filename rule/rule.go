package rule

import "fmt"

type RuleYaml struct {
	Rules []*Rule `yaml:"rules"`
}

type Rule struct {
	Name      string `yaml:"name,omitempty" db:"rule_name"`
	Value     string `yaml:"value,omitempty" db:"rule_value"`
	RulesetID int    `yaml:"ruleset_id,omitempty" db:"ruleset_id"`
	Notes     string `yaml:"notes,omitempty" db:"notes"`
}

func (e *RuleYaml) sanitize() error {
	for _, rule := range e.Rules {
		if rule.RulesetID == 0 {
			rule.RulesetID = 1
		}
		if rule.Name == "" {
			return fmt.Errorf("rule name must not be empty")
		}
		if rule.Value == "" {
			return fmt.Errorf("rule value must not be empty")
		}
	}
	return nil
}
