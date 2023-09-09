package config

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"strings"
	"sync"
)

var (
	mu  sync.Mutex
	cfg *Config
)

// Config represents a txt config
type Config struct {
	OriginalDB string
	InjectDB   string
}

// Load loads the config
func Load() (*Config, error) {
	mu.Lock()
	defer mu.Unlock()

	if cfg != nil {
		return cfg, nil
	}

	cfg = &Config{}

	data, err := os.ReadFile("yakuku.txt")
	if err != nil {

		if strings.Contains(err.Error(), "no such file or directory") {
			err = os.WriteFile("yakuku.txt", []byte(`# format: dbuser:dbpassword@tcp(dbhost:dbport)/dbname
# change this to a db for 'yaml'
ORIGINAL_DB=ro:ro@tcp(content-cdn.projecteq.net:16033)/peq_content
# change this to your target DB for 'inject' command (optional)
INJECT_DB=root:root@tcp(localhost:3306)/peq
`), 0644)
			if err != nil {
				return nil, err
			}

			fmt.Println("A new yakuku.txt file was generated. Please edit it and re-run.")
			os.Exit(1)

			return cfg, nil
		}
		return nil, err
	}

	reader := bufio.NewReader(bytes.NewBuffer(data))

	for {
		line, err := reader.ReadString('\n')
		if err != nil {
			break
		}
		if !strings.Contains(line, "=") {
			continue
		}
		firstEquals := strings.Index(line, "=")
		key := strings.TrimSpace(line[:firstEquals])
		value := strings.TrimSpace(line[firstEquals+1:])
		switch strings.ToLower(key) {
		case "original_db":
			cfg.OriginalDB = value + "?parseTime=true&multiStatements=true&interpolateParams=true&collation=utf8mb4_unicode_ci&charset=utf8mb4,utf8"
		case "inject_db":
			cfg.InjectDB = value + "?parseTime=true&multiStatements=true&interpolateParams=true&collation=utf8mb4_unicode_ci&charset=utf8mb4,utf8"
		}
	}

	return cfg, nil
}
