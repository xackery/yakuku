package charcreate

import (
	"fmt"
	"os"
	"time"

	"github.com/xackery/yakuku/util"
	"gopkg.in/yaml.v3"
)

func Sql(srcYaml, dstSql string) error {
	start := time.Now()
	fmt.Printf("CharCreate...")
	var err error
	defer func() {
		fmt.Println(" finished in", time.Since(start).String())
		if err != nil {
			fmt.Printf("Error: %s\n", err)
			os.Exit(1)
		}
	}()
	err = sqlGenerate(srcYaml, dstSql)
	return nil
}

func sqlGenerate(srcYaml, dstSql string) error {
	charCreate := &CharCreateYaml{}
	r, err := os.Open(srcYaml)
	if err != nil {
		return err
	}
	defer r.Close()
	dec := yaml.NewDecoder(r)
	dec.KnownFields(true)
	err = dec.Decode(charCreate)
	if err != nil {
		return fmt.Errorf("decode: %w", err)
	}

	err = charCreate.sanitize()
	if err != nil {
		return fmt.Errorf("charCreate sanitize: %w", err)
	}

	err = generateCharCreateSQL(charCreate, dstSql)
	if err != nil {
		return fmt.Errorf("generateCharCreateSQL: %w", err)
	}

	return nil
}

func generateCharCreateSQL(sp *CharCreateYaml, dstSql string) error {
	w, err := os.Create(dstSql)
	if err != nil {
		return err
	}
	defer w.Close()

	w.WriteString("REPLACE INTO `char_create_combinations` (`allocation_id`, `race`, `class`, `deity`, `start_zone`, `expansions_req`) VALUES\n")

	buf := ""
	for _, charCreate := range sp.CharCreates {
		for _, class := range charCreate.Classes {
			for _, deity := range class.Deities {
				for _, zone := range deity.Zones {
					buf += fmt.Sprintf("(%d, %d, %d, %d, %d, %d",
						zone.AllocationID,
						util.RaceNameToID(charCreate.Race),
						util.ClassNameToID(class.Class),
						util.DeityNameToID(deity.Deity),
						util.ZoneNameToID(zone.Zone),
						zone.ExpansionsReq)

					buf += "),\n"
				}
			}
		}
	}
	buf = buf[:len(buf)-2]
	buf += ";\n"
	w.WriteString(buf)

	w.WriteString("\n\n")
	fmt.Printf("%d charCreate combinations ", len(sp.CharCreates))

	w.WriteString("REPLACE INTO `char_create_point_allocations` (`id`, `base_str`, `base_sta`, `base_dex`, `base_agi`, `base_int`, `base_wis`, `base_cha`, `alloc_str`, `alloc_sta`, `alloc_dex`, `alloc_agi`, `alloc_int`, `alloc_wis`, `alloc_cha`) VALUES\n")
	for allocIndex, allocation := range sp.Allocations {
		w.WriteString(fmt.Sprintf("(%d, %d, %d, %d, %d, %d, %d, %d, %d, %d, %d, %d, %d, %d, %d",
			allocation.ID,
			allocation.BaseStr,
			allocation.BaseSta,
			allocation.BaseDex,
			allocation.BaseAgi,
			allocation.BaseInt,
			allocation.BaseWis,
			allocation.BaseCha,
			allocation.AllocStr,
			allocation.AllocSta,
			allocation.AllocDex,
			allocation.AllocAgi,
			allocation.AllocInt,
			allocation.AllocWis,
			allocation.AllocCha))

		if allocIndex == len(sp.Allocations)-1 {
			w.WriteString(");\n")
		} else {
			w.WriteString("),\n")
		}
	}
	fmt.Printf("%d allocations ", len(sp.Allocations))

	return nil
}
