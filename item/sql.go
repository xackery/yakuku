package item

import (
	"fmt"
	"os"
	"time"

	"github.com/fatih/structs"
	_ "github.com/go-sql-driver/mysql"
	"github.com/xackery/yakuku/util"

	"gopkg.in/yaml.v3"
)

func Sql(srcYaml, dstSql string) error {
	start := time.Now()
	fmt.Printf("Item... ")
	var err error
	defer func() {
		fmt.Println("finished in", time.Since(start).String())
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
	item := &ItemYaml{}
	dec := yaml.NewDecoder(r)
	dec.KnownFields(true)
	err = dec.Decode(item)
	if err != nil {
		return fmt.Errorf("decode: %w", err)
	}

	err = item.sanitize()
	if err != nil {
		return fmt.Errorf("item sanitize: %w", err)
	}

	err = generateItemSQL(item, dstSql)
	if err != nil {
		return fmt.Errorf("generateItemSQL: %w", err)
	}
	return nil
}

func generateItemSQL(sp *ItemYaml, dstSql string) error {
	w, err := os.Create(dstSql)
	if err != nil {
		return err
	}
	defer w.Close()

	itemCount := 0

	buf := ""
	for _, item := range sp.Items {
		fields := structs.Fields(item)

		buf += "REPLACE INTO `items` SET "
		for _, field := range fields {
			if !field.IsExported() {
				continue
			}
			fieldBuf := util.FieldParse(field)
			if fieldBuf == "" {
				buf += fieldBuf + ", "
				continue
			}
			return fmt.Errorf("unknown field type: %s", field.Kind().String())
		}
		buf = buf[:len(buf)-2]
		buf += ";\n"
		w.WriteString(buf)
		itemCount++
	}
	fmt.Printf(" %d items ", itemCount)

	return nil
}
