package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/xackery/yakuku/rule"
)

func main() {
	err := run()
	if err != nil {
		fmt.Println("Run failed:", err)
		os.Exit(1)
	}

}

func run() error {
	if len(os.Args) < 2 {
		fmt.Println("Usage: yakuku [yaml|sql|inject]")
		os.Exit(1)
	}

	_, err := os.Stat("dbstr_us_original.txt")
	if err != nil {
		fmt.Println("Please copy dbstr_us.txt into this path, and rename it to dbstr_us_original.txt")
		os.Exit(1)
	}

	_, err = os.Stat("spells_us_original.txt")
	if err != nil {
		fmt.Println("Please copy spells_us.txt into this path, and rename it to spells_us_original.txt")
		os.Exit(1)
	}

	switch strings.ToLower(os.Args[1]) {
	case "yaml":
		return yaml(os.Args[2:])
	case "sql":
		return sql(os.Args[2:])
	case "inject":
		return inject(os.Args[2:])
	default:
		fmt.Println("Unknown command:", os.Args[1])
		fmt.Println("Usage: yakuku [yaml|sql|inject]")
		os.Exit(1)
	}

	return nil
}

func yaml(args []string) error {
	var err error
	if len(args) < 2 {
		fmt.Println("Usage: yakuku yaml [rule|spell|aa|task|charcreate] [out_path] [filters]")
		fmt.Println("This command will a yaml dump based on the original database")
		os.Exit(1)
	}

	cmd := strings.ToLower(args[0])
	path := args[1]

	filters := []string{}
	if len(args) > 2 {
		filters = args[2:]
	}

	switch cmd {
	case "rule":
		err = rule.Yaml(path, filters)
		if err != nil {
			return fmt.Errorf("rule: %w", err)
		}
	default:
		fmt.Println("Unknown command:", cmd)
		fmt.Println("Usage: yakuku yaml [rule|spell|aa|task|charcreate]")
		os.Exit(1)
	}
	return nil
}

func sql(args []string) error {
	var err error
	if len(args) < 2 {
		fmt.Println("Usage: yakuku sql [rule|spell|aa|task|charcreate|all] [yml_path] [out_path]")
		fmt.Println("This command will generate sql based on the yaml dump")
		os.Exit(1)
	}

	cmd := strings.ToLower(args[0])
	path := args[1]

	switch cmd {
	case "rule":
		err = rule.Sql(path)
		if err != nil {
			return fmt.Errorf("rule: %w", err)
		}
	default:
		fmt.Println("Unknown command:", cmd)
		fmt.Println("Usage: yakuku sql [rule|spell|aa|task|charcreate]")
		os.Exit(1)
	}
	return nil
}

func inject(args []string) error {
	var err error
	if len(args) < 2 {
		fmt.Println("Usage: yakuku inject [sql_path]")
		fmt.Println("This command will inject sql into the target database")
		os.Exit(1)
	}

	cmd := strings.ToLower(args[0])
	path := args[1]

	switch cmd {
	case "rule":
		err = rule.Inject(path)
		if err != nil {
			return fmt.Errorf("rule: %w", err)
		}
	default:
		fmt.Println("Unknown command:", cmd)
		fmt.Println("Usage: yakuku inject [rule|spell|aa|task|charcreate]")
		os.Exit(1)
	}
	return nil
}
