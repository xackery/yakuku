package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/xackery/yakuku/charcreate"
	sqlinject "github.com/xackery/yakuku/inject"
	"github.com/xackery/yakuku/item"
	"github.com/xackery/yakuku/npc"
	"github.com/xackery/yakuku/npcgrid"
	"github.com/xackery/yakuku/rule"
	"github.com/xackery/yakuku/startingitems"
	"github.com/xackery/yakuku/zone"
)

var (
	// Version is the current version of yakuku
	Version = "0.0.1"
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

	switch strings.ToLower(os.Args[1]) {
	case "version":
		fmt.Println("Yakuku v" + Version)
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
		fmt.Println("Usage: yakuku yaml [rule|charcreate|item|npc|npcgrid|startingitems|zone] [out_path] [filters]")
		fmt.Println("This command will a yaml dump based on the original database")
		os.Exit(1)
	}

	_, err = os.Stat("dbstr_us_original.txt")
	if err != nil {
		fmt.Println("Please copy dbstr_us.txt into this path, and rename it to dbstr_us_original.txt")
		os.Exit(1)
	}

	_, err = os.Stat("spells_us_original.txt")
	if err != nil {
		fmt.Println("Please copy spells_us.txt into this path, and rename it to spells_us_original.txt")
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
	case "item":
		err = item.Yaml(path, filters)
		if err != nil {
			return fmt.Errorf("item: %w", err)
		}
	case "npc":
		err = npc.Yaml(path, filters)
		if err != nil {
			return fmt.Errorf("npc: %w", err)
		}
	case "npcgrid":
		err = npcgrid.Yaml(path, filters)
		if err != nil {
			return fmt.Errorf("npcgrid: %w", err)
		}
	case "charcreate":
		err = charcreate.Yaml(path, filters)
		if err != nil {
			return fmt.Errorf("charcreate: %w", err)
		}
	case "startingitems":
		err = startingitems.Yaml(path, filters)
		if err != nil {
			return fmt.Errorf("startingitems: %w", err)
		}
	case "zone":
		err = zone.Yaml(path, filters)
		if err != nil {
			return fmt.Errorf("zone: %w", err)
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
	if len(args) < 3 {
		fmt.Println("Usage: yakuku sql [rule|charcreate|item|npc|npcgrid|startingitems|zone] [src.yaml] [dst.sql]")
		fmt.Println("This command will generate sql based on the yaml dump")
		os.Exit(1)
	}

	cmd := strings.ToLower(args[0])
	srcYaml := args[1]
	dstSql := args[2]

	switch cmd {
	case "rule":
		err = rule.Sql(srcYaml, dstSql)
		if err != nil {
			return fmt.Errorf("rule: %w", err)
		}
	case "charcreate":
		err = charcreate.Sql(srcYaml, dstSql)
		if err != nil {
			return fmt.Errorf("charcreate: %w", err)
		}
	case "item":
		err = item.Sql(srcYaml, dstSql)
		if err != nil {
			return fmt.Errorf("item: %w", err)
		}
	case "npc":
		err = npc.Sql(srcYaml, dstSql)
		if err != nil {
			return fmt.Errorf("npc: %w", err)
		}
	case "npcgrid":
		err = npcgrid.Sql(srcYaml, dstSql)
		if err != nil {
			return fmt.Errorf("npcgrid: %w", err)
		}
	case "startingitems":
		err = startingitems.Sql(srcYaml, dstSql)
		if err != nil {
			return fmt.Errorf("startingitems: %w", err)
		}
	case "zone":
		err = zone.Sql(srcYaml, dstSql)
		if err != nil {
			return fmt.Errorf("zone: %w", err)
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
	if len(args) < 1 {
		fmt.Println("Usage: yakuku inject [sql_path]")
		fmt.Println("This command will inject sql into the target database")
		os.Exit(1)
	}

	err = sqlinject.Inject(args[0])
	if err != nil {
		return fmt.Errorf("inject %s: %w", args[0], err)
	}

	return nil
}
