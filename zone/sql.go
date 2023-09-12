package zone

import (
	"fmt"
	"os"
	"reflect"
	"time"

	"github.com/fatih/structs"
	_ "github.com/go-sql-driver/mysql"
	"github.com/xackery/yakuku/util"

	"gopkg.in/yaml.v3"
)

func Sql(srcYaml, dstSql string) error {
	start := time.Now()
	fmt.Printf("Zone... ")
	var err error
	defer func() {
		fmt.Println("finished in", time.Since(start).String())
		if err != nil {
			fmt.Println("Error:", err)
			os.Exit(1)
		}
	}()
	err = zoneGenerate(srcYaml, dstSql)
	return nil
}

func zoneGenerate(srcYaml, dstSql string) error {
	r, err := os.Open(srcYaml)
	if err != nil {
		return err
	}
	defer r.Close()
	zone := &ZoneYaml{}
	dec := yaml.NewDecoder(r)
	dec.KnownFields(true)
	err = dec.Decode(zone)
	if err != nil {
		return fmt.Errorf("decode: %w", err)
	}

	err = zone.sanitize()
	if err != nil {
		return fmt.Errorf("npc sanitize: %w", err)
	}

	err = generateZoneSQL(zone, dstSql)
	if err != nil {
		return fmt.Errorf("generateZoneSQL: %w", err)
	}
	return nil
}

func generateZoneSQL(sp *ZoneYaml, dstSql string) error {
	w, err := os.Create(dstSql)
	if err != nil {
		return err
	}
	defer w.Close()

	zoneCount := 0
	zonePointBuf := ""
	buf := ""

	for _, zone := range sp.Zones {
		fields := structs.Fields(zone)

		buf += "REPLACE INTO `zone` SET "
		for _, field := range fields {
			if !field.IsExported() {
				continue
			}

			fieldBuf := util.FieldParse(field)
			if fieldBuf != "" {
				buf += fieldBuf + ", "
				continue
			}

			switch field.Kind() {
			case reflect.Slice:
				// assert type
				switch v := field.Value().(type) {
				case []*ZonePoint:
					if len(v) == 0 {
						continue
					}
					zonePoint, err := generateZonePointSQL(v, zone.ShortName.String)
					if err != nil {
						return fmt.Errorf("generateZonePointSQL: %w", err)
					}
					zonePointBuf += zonePoint
				}
				continue
			default:
				return fmt.Errorf("unknown type %s %s", field.Tag("db"), field.Kind())
			}
		}
		buf = buf[:len(buf)-2]
		buf += ";\n"
		zoneCount++
	}
	w.WriteString(buf)
	w.WriteString("\n" + zonePointBuf)

	fmt.Printf(" %d zones ", zoneCount)
	return nil
}

func generateZonePointSQL(points []*ZonePoint, zoneName string) (string, error) {
	buf := ""

	for _, point := range points {
		fields := structs.Fields(point)

		buf += "REPLACE INTO `zone_points` SET "
		for _, field := range fields {
			if !field.IsExported() {
				continue
			}

			if field.Tag("db") == "zone" {
				buf += fmt.Sprintf("`zone` = '%s', ", zoneName)
				continue
			}

			fieldBuf := util.FieldParse(field)
			if fieldBuf != "" {
				buf += fieldBuf + ", "
				continue
			}

			return "", fmt.Errorf("unknown type %s %s", field.Tag("db"), field.Kind())
		}
		buf = buf[:len(buf)-2]
		buf += ";\n"
	}

	return buf, nil
}
