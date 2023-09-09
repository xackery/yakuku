package charcreate

import (
	"context"
	"database/sql"
	"fmt"
	"os"
	"time"

	"github.com/jmoiron/sqlx"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/xackery/yakuku/util"
	"gopkg.in/yaml.v3"
)

func Build(cmd *cobra.Command, args []string) error {
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
	err = generate()
	return nil
}

func generate() error {
	charCreate := &CharCreateYaml{}
	r, err := os.Open("charcreate.yaml")
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

	err = generateCharCreateSQL(charCreate)
	if err != nil {
		return fmt.Errorf("generateCharCreateSQL: %w", err)
	}

	return nil
}

func generateCharCreateSQL(sp *CharCreateYaml) error {
	w, err := os.Create("charcreate.sql")
	if err != nil {
		return err
	}
	defer w.Close()

	w.WriteString("REPLACE INTO `char_create_combinations` (`allocation_id`, `race`, `class`, `deity`, `start_zone`, `expansions_req`) VALUES\n")

	for charCreateIndex, charCreate := range sp.CharCreates {
		for _, class := range charCreate.Classes {
			for _, deity := range class.Deities {
				for _, zone := range deity.Zones {
					w.WriteString(fmt.Sprintf("(%d, %d, %d, %d, %d, %d",
						zone.AllocationID,
						util.RaceNameToID(charCreate.Race),
						util.ClassNameToID(class.Class),
						util.DeityNameToID(deity.Deity),
						util.ZoneNameToID(zone.Zone),
						zone.ExpansionsReq))

					//w.WriteString(fmt.Sprintf("%d, ", util.RaceNameToID(charCreate.Race)))
					if charCreateIndex == len(sp.CharCreates)-1 {
						w.WriteString(");\n")
					} else {
						w.WriteString("),\n")
					}
				}
			}
		}
	}

	w.WriteString("\n\n")

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

	return nil
}

// Import takes database info and dumps to yaml
func Import(cmd *cobra.Command, args []string) error {
	if !viper.IsSet("db_host") {
		return fmt.Errorf("db_host is not set, pass it as an argument --db_host=... or set it in .luaject.yaml")
	}

	db, err := sqlx.Connect("mysql", fmt.Sprintf("%s:%s@tcp(%s)/%s?parseTime=true&multiStatements=true&interpolateParams=true&collation=utf8mb4_unicode_ci&charset=utf8mb4,utf8", viper.GetString("db_user"), viper.GetString("db_pass"), viper.GetString("db_host"), viper.GetString("db_name")))
	if err != nil {
		return fmt.Errorf("db connect: %w", err)
	}
	defer db.Close()

	charCreate := &CharCreateYaml{}
	err = charCreate.sanitize()
	if err != nil {
		return err
	}

	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()

	rows1, err := db.QueryxContext(ctx, "SELECT * FROM char_create_point_allocations")
	if err != nil {
		return fmt.Errorf("db query allocations: %w", err)
	}
	defer rows1.Close()

	type charCreatePointAllocationDB struct {
		ID       int `db:"id"`
		BaseStr  int `db:"base_str"`
		BaseSta  int `db:"base_sta"`
		BaseDex  int `db:"base_dex"`
		BaseAgi  int `db:"base_agi"`
		BaseInt  int `db:"base_int"`
		BaseWis  int `db:"base_wis"`
		BaseCha  int `db:"base_cha"`
		AllocStr int `db:"alloc_str"`
		AllocSta int `db:"alloc_sta"`
		AllocDex int `db:"alloc_dex"`
		AllocAgi int `db:"alloc_agi"`
		AllocInt int `db:"alloc_int"`
		AllocWis int `db:"alloc_wis"`
		AllocCha int `db:"alloc_cha"`
	}

	for rows1.Next() {
		r := charCreatePointAllocationDB{}
		err = rows1.StructScan(&r)
		if err != nil {
			return fmt.Errorf("db allocation struct scan: %w", err)
		}

		charCreate.Allocations = append(charCreate.Allocations, &Allocation{
			ID:       r.ID,
			BaseStr:  r.BaseStr,
			BaseSta:  r.BaseSta,
			BaseDex:  r.BaseDex,
			BaseAgi:  r.BaseAgi,
			BaseInt:  r.BaseInt,
			BaseWis:  r.BaseWis,
			BaseCha:  r.BaseCha,
			AllocStr: r.AllocStr,
			AllocSta: r.AllocSta,
			AllocDex: r.AllocDex,
			AllocAgi: r.AllocAgi,
			AllocInt: r.AllocInt,
			AllocWis: r.AllocWis,
			AllocCha: r.AllocCha,
		})
	}

	rows, err := db.QueryxContext(ctx, "SELECT * FROM char_create_combinations")
	if err != nil {
		return fmt.Errorf("db query combinations: %w", err)
	}
	defer rows.Close()

	type charCreateDB struct {
		AllocationID  int `db:"allocation_id"`
		Race          int `db:"race"`
		Class         int `db:"class"`
		Deity         int `db:"deity"`
		StartZone     int `db:"start_zone"`
		ExpansionsReq int `db:"expansions_req"`
	}

	for rows.Next() {
		r := charCreateDB{}
		err = rows.StructScan(&r)
		if err != nil {
			return fmt.Errorf("db combinations struct scan: %w", err)
		}

		isFound := false
		var raceIndex int
		var c *CharCreate
		for raceIndex, c = range charCreate.CharCreates {
			if util.RaceNameToID(c.Race) == r.Race {
				isFound = true
				break
			}
		}
		if !isFound {
			charCreate.CharCreates = append(charCreate.CharCreates, &CharCreate{Race: util.RaceIDToName(r.Race)})
			raceIndex = len(charCreate.CharCreates) - 1
		}

		var classIndex int
		isFound = false
		var cl *CharClass
		for classIndex, cl = range charCreate.CharCreates[raceIndex].Classes {
			if util.ClassNameToID(cl.Class) == r.Class {
				isFound = true
				break
			}
		}

		if !isFound {
			charCreate.CharCreates[raceIndex].Classes = append(charCreate.CharCreates[raceIndex].Classes, &CharClass{Class: util.ClassIDToName(r.Class)})
			classIndex = len(charCreate.CharCreates[raceIndex].Classes) - 1
		}

		var deityIndex int
		isFound = false
		var dl *CharDeity
		for deityIndex, dl = range charCreate.CharCreates[raceIndex].Classes[classIndex].Deities {
			if util.DeityNameToID(dl.Deity) == r.Deity {
				isFound = true
				break
			}
		}

		if !isFound {
			charCreate.CharCreates[raceIndex].Classes[classIndex].Deities = append(charCreate.CharCreates[raceIndex].Classes[classIndex].Deities, &CharDeity{Deity: util.DeityIDToName(r.Deity)})
			deityIndex = len(charCreate.CharCreates[raceIndex].Classes[classIndex].Deities) - 1
		}

		var zoneIndex int
		isFound = false
		var zl *CharZone
		for zoneIndex, zl = range charCreate.CharCreates[raceIndex].Classes[classIndex].Deities[deityIndex].Zones {
			if util.ZoneNameToID(zl.Zone) == r.StartZone {
				isFound = true
				break
			}
		}

		if !isFound {
			charCreate.CharCreates[raceIndex].Classes[classIndex].Deities[deityIndex].Zones = append(charCreate.CharCreates[raceIndex].Classes[classIndex].Deities[deityIndex].Zones, &CharZone{Zone: util.ZoneIDToName(r.StartZone)})
			zoneIndex = len(charCreate.CharCreates[raceIndex].Classes[classIndex].Deities[deityIndex].Zones) - 1
		}

		focus := charCreate.CharCreates[raceIndex].Classes[classIndex].Deities[deityIndex].Zones[zoneIndex]
		focus.Zone = util.ZoneIDToName(r.StartZone)
		focus.ExpansionsReq = r.ExpansionsReq
		focus.AllocationID = r.AllocationID
	}

	items, err := startingItemsQuery(ctx, db, 0, 0, 0, 0)
	if err != nil {
		return fmt.Errorf("starting items query: %w", err)
	}
	charCreate.Items = items
	for _, char := range charCreate.CharCreates {
		items, err := startingItemsQuery(ctx, db, util.RaceNameToID(char.Race), 0, 0, 0)
		if err != nil {
			return fmt.Errorf("starting items query race %s: %w", char.Race, err)
		}
		char.Items = items

		for _, class := range char.Classes {
			items, err := startingItemsQuery(ctx, db, util.RaceNameToID(char.Race), util.ClassNameToID(class.Class), 0, 0)
			if err != nil {
				return fmt.Errorf("starting items query race %s class %s: %w", char.Race, class.Class, err)
			}
			class.Items = items

			for _, deity := range class.Deities {

				items, err := startingItemsQuery(ctx, db, util.RaceNameToID(char.Race), util.ClassNameToID(class.Class), util.DeityNameToID(deity.Deity), 0)
				if err != nil {
					return fmt.Errorf("starting items query race %s class %s deity %s: %w", char.Race, class.Class, deity.Deity, err)
				}
				deity.Items = items

				for _, zone := range deity.Zones {
					choices, err := choiceQuery(ctx, db, util.RaceNameToID(char.Race), util.ClassNameToID(class.Class), util.DeityNameToID(deity.Deity), util.ZoneNameToID(zone.Zone))
					if err != nil {
						return fmt.Errorf("choice query class %s deity %s zone %s: %w", class.Class, deity.Deity, zone.Zone, err)
					}

					items, err := startingItemsQuery(ctx, db, util.RaceNameToID(char.Race), util.ClassNameToID(class.Class), util.DeityNameToID(deity.Deity), util.ZoneNameToID(zone.Zone))
					if err != nil {
						return fmt.Errorf("starting items query class %s deity %s zone %s: %w", class.Class, deity.Deity, zone.Zone, err)
					}

					zone.Choices = choices

					if len(class.Items) == 0 {
						class.Items = items
					} else {
						if len(items) > 0 && class.Items[0].ItemID != items[0].ItemID {
							zone.Items = items
						}
					}
				}
			}
		}
	}

	w, err := os.Create("charcreate_dump.yaml")
	if err != nil {
		return err
	}
	defer w.Close()

	enc := yaml.NewEncoder(w)
	err = enc.Encode(charCreate)
	if err != nil {
		return err
	}

	fmt.Println("Created charcreate_dump.yaml")
	return nil
}

func choiceQuery(ctx context.Context, db *sqlx.DB, raceID int, classID int, deityID int, zoneID int) ([]*CharChoice, error) {
	choices := []*CharChoice{}

	rows, err := db.QueryxContext(ctx, "SELECT * FROM start_zones WHERE player_class = ? AND player_deity = ? AND player_race = ? AND start_zone = ?", classID, deityID, raceID, zoneID)
	if err != nil {
		return nil, fmt.Errorf("db query choices: %w", err)
	}
	defer rows.Close()

	type charChoiceDB struct {
		X                    float32        `db:"x"`                      // float NOT NULL DEFAULT 0,
		Y                    float32        `db:"y"`                      // float NOT NULL DEFAULT 0,
		Z                    float32        `db:"z"`                      // float NOT NULL DEFAULT 0,
		Heading              float32        `db:"heading"`                // float NOT NULL DEFAULT 0,
		ZoneId               int            `db:"zone_id"`                // int(4) NOT NULL DEFAULT 0,
		BindId               int            `db:"bind_id"`                // int(4) NOT NULL DEFAULT 0,
		PlayerChoice         int            `db:"player_choice"`          // int(2) NOT NULL DEFAULT 0,
		PlayerClass          int            `db:"player_class"`           // int(2) NOT NULL DEFAULT 0,
		PlayerDeity          int            `db:"player_deity"`           // int(4) NOT NULL DEFAULT 0,
		PlayerRace           int            `db:"player_race"`            // int(4) NOT NULL DEFAULT 0,
		StartZone            int            `db:"start_zone"`             // int(4) NOT NULL DEFAULT 0,
		BindX                float32        `db:"bind_x"`                 // float NOT NULL DEFAULT 0,
		BindY                float32        `db:"bind_y"`                 // float NOT NULL DEFAULT 0,
		BindZ                float32        `db:"bind_z"`                 // float NOT NULL DEFAULT 0,
		SelectRank           int            `db:"select_rank"`            // tinyint(3) unsigned NOT NULL DEFAULT 50,
		MinExpansion         int            `db:"min_expansion"`          // tinyint(4) NOT NULL DEFAULT -1,
		MaxExpansion         int            `db:"max_expansion"`          // tinyint(4) NOT NULL DEFAULT -1,
		ContentFlags         sql.NullString `db:"content_flags"`          // varchar(100) DEFAULT NULL,
		ContentFlagsDisabled sql.NullString `db:"content_flags_disabled"` // varchar(100) DEFAULT NULL,
	}

	for rows.Next() {
		r := charChoiceDB{}
		err = rows.StructScan(&r)
		if err != nil {
			return nil, fmt.Errorf("db choice struct scan: %w", err)
		}
		choices = append(choices, &CharChoice{
			Index:        r.PlayerChoice,
			SpawnX:       r.X,
			SpawnY:       r.Y,
			SpawnZ:       r.Z,
			SpawnHeading: r.Heading,
			BindID:       r.BindId,
			BindX:        r.BindX,
			BindY:        r.BindY,
			BindZ:        r.BindZ,
		})
	}

	return choices, nil
}

func startingItemsQuery(ctx context.Context, db *sqlx.DB, raceID int, classID int, deityID int, zoneID int) ([]*CharItem, error) {
	items := []*CharItem{}

	rows, err := db.QueryxContext(ctx, "SELECT starting_items.*, items.name, items.lore FROM starting_items INNER JOIN items ON items.id = starting_items.itemid WHERE `class` = ? AND deityid = ? AND race = ? AND zoneid = ?", classID, deityID, raceID, zoneID)
	if err != nil {
		return nil, fmt.Errorf("db query items: %w", err)
	}
	defer rows.Close()

	type charItemDB struct {
		ID                   int            `db:"id"`                     // int(11) unsigned NOT NULL AUTO_INCREMENT,
		Race                 int            `db:"race"`                   // int(11) NOT NULL DEFAULT 0,
		Class                int            `db:"class"`                  // int(11) NOT NULL DEFAULT 0,
		DeityID              int            `db:"deityid"`                // int(11) NOT NULL DEFAULT 0,
		ZoneID               int            `db:"zoneid"`                 // int(11) NOT NULL DEFAULT 0,
		ItemID               int            `db:"itemid"`                 // int(11) NOT NULL DEFAULT 0,
		ItemName             string         `db:"name"`                   // varchar(64) NOT NULL DEFAULT '',
		ItemLore             string         `db:"lore"`                   // varchar(80) NOT NULL DEFAULT '',
		ItemCharges          int            `db:"item_charges"`           // tinyint(3) unsigned NOT NULL DEFAULT 1,
		Gm                   int            `db:"gm"`                     // tinyint(1) NOT NULL DEFAULT 0,
		Slot                 int            `db:"slot"`                   // mediumint(9) NOT NULL DEFAULT -1,
		MinExpansion         int            `db:"min_expansion"`          // tinyint(4) NOT NULL DEFAULT -1,
		MaxExpansion         int            `db:"max_expansion"`          // tinyint(4) NOT NULL DEFAULT -1,
		ContentFlags         sql.NullString `db:"content_flags"`          // varchar(100) DEFAULT NULL,
		ContentFlagsDisabled sql.NullString `db:"content_flags_disabled"` // varchar(100) DEFAULT NULL,
	}

	for rows.Next() {
		r := charItemDB{}
		err = rows.StructScan(&r)
		if err != nil {
			return nil, fmt.Errorf("db item struct scan: %w", err)
		}

		name := r.ItemName
		if r.ItemLore != "" {
			name = fmt.Sprintf("%s (Lore: %s)", r.ItemName, r.ItemLore)
		}
		items = append(items, &CharItem{
			ID:          r.ID,
			ItemID:      r.ItemID,
			Name:        name,
			ItemCharges: r.ItemCharges,
			GM:          r.Gm,
			Slot:        r.Slot,
		})
	}

	return items, nil
}
