package aa

import (
	"bufio"
	"context"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/goccy/go-yaml"
	"github.com/jmoiron/sqlx"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func Import(cmd *cobra.Command, args []string) error {
	if !viper.IsSet("db_host") {
		return fmt.Errorf("db_host is not set, pass it as an argument --db_host=... or set it in .luaject.yaml")
	}

	db, err := sqlx.Connect("mysql", fmt.Sprintf("%s:%s@tcp(%s)/%s?parseTime=true&multiStatements=true&interpolateParams=true&collation=utf8mb4_unicode_ci&charset=utf8mb4,utf8", viper.GetString("db_user"), viper.GetString("db_pass"), viper.GetString("db_host"), viper.GetString("db_name")))
	if err != nil {
		return fmt.Errorf("db connect: %w", err)
	}
	defer db.Close()

	aa := &AAYaml{}
	err = aa.sanitize()
	if err != nil {
		return err
	}

	// read dbstr_us.txt line by line

	r, err := os.Open("dbstr_us_original.txt")
	if err != nil {
		return err
	}
	defer r.Close()
	scanner := bufio.NewScanner(r)

	titles := make(map[int]string)
	descriptions := make(map[int]string)

	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			continue
		}
		records := strings.Split(line, "^")
		if len(records) < 3 {
			return fmt.Errorf("invalid line: %s", line)
		}
		sid, err := strconv.Atoi(records[0])
		if err != nil {
			return fmt.Errorf("invalid line: %s", line)
		}
		category, err := strconv.Atoi(records[1])
		if err != nil {
			return fmt.Errorf("invalid line: %s", line)
		}
		if category != 1 && category != 4 {
			continue
		}

		if category == 1 {
			titles[sid] = records[2]
		} else {
			descriptions[sid] = records[2]
		}
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	rows, err := db.QueryxContext(ctx, "SELECT * FROM aa_ability ORDER BY id")
	if err != nil {
		return fmt.Errorf("db query: %w", err)
	}
	defer rows.Close()

	for rows.Next() {
		r := &AASkill{}
		err = rows.StructScan(&r)
		if err != nil {
			return fmt.Errorf("db struct scan: %w", err)
		}
		aa.Skills = append(aa.Skills, r)
	}

	err = importRanks(db, aa, titles, descriptions)
	if err != nil {
		return fmt.Errorf("importRanks: %w", err)
	}

	data, err := yaml.MarshalWithOptions(aa, yaml.WithComment(
		yaml.CommentMap{
			"$.skills": []*yaml.Comment{yaml.LineComment("Test!")},
		},
	))
	if err != nil {
		return fmt.Errorf("marshal: %w", err)
	}
	err = os.WriteFile("aa_dump.yaml", data, 0644)
	if err != nil {
		return err
	}
	fmt.Println("Created aa_dump.yaml")

	return nil
}

func importRanks(db *sqlx.DB, aa *AAYaml, titles map[int]string, descriptions map[int]string) error {

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	rows, err := db.QueryxContext(ctx, "SELECT * FROM aa_ranks")
	if err != nil {
		return fmt.Errorf("db query: %w", err)
	}
	defer rows.Close()

	ranks := []*AARank{}

	for rows.Next() {
		r := &AARank{}
		err = rows.StructScan(r)
		if err != nil {
			return fmt.Errorf("db struct scan: %w", err)
		}

		if r.TitleSID > 0 {
			r.Title = titles[r.TitleSID]
		}
		if r.DescriptionSID > 0 {
			r.Description = descriptions[r.DescriptionSID]
		}

		ranks = append(ranks, r)
	}

	for _, r := range ranks {
		for _, skill := range aa.Skills {
			if skill.FirstRankID == r.ID {
				//fmt.Println("Skill", skill.ID, "has first rank", r.ID)
				err = searchRanks(skill, ranks, r.ID)
				if err != nil {
					return fmt.Errorf("searchRanks: %w", err)
				}
				break
			}
		}
	}

	err = importRankEffects(db, aa)
	if err != nil {
		return fmt.Errorf("importRankEffects: %w", err)
	}

	aa.omitEmpty()

	return nil
}

func importRankEffects(db *sqlx.DB, aa *AAYaml) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	rows, err := db.QueryxContext(ctx, "SELECT * FROM aa_rank_effects")
	if err != nil {
		return fmt.Errorf("db query: %w", err)
	}
	defer rows.Close()

	type RankEffect struct {
		RankID   int `db:"rank_id"`
		Slot     int `db:"slot"`
		EffectID int `db:"effect_id"`
		Base1    int `db:"base1"`
		Base2    int `db:"base2"`
	}

	for rows.Next() {
		r := &RankEffect{}
		err = rows.StructScan(r)
		if err != nil {
			return fmt.Errorf("db struct scan: %w", err)
		}

		isFound := false
		for _, skill := range aa.Skills {
			for _, rank := range skill.Ranks {
				if rank.ID == r.RankID {
					switch r.Slot {
					case 1:
						rank.Slot1.EffectID = r.EffectID
						rank.Slot1.Base1 = r.Base1
						rank.Slot1.Base2 = r.Base2
					case 2:
						rank.Slot2.EffectID = r.EffectID
						rank.Slot2.Base1 = r.Base1
						rank.Slot2.Base2 = r.Base2
					case 3:
						rank.Slot3.EffectID = r.EffectID
						rank.Slot3.Base1 = r.Base1
						rank.Slot3.Base2 = r.Base2
					case 4:
						rank.Slot4.EffectID = r.EffectID
						rank.Slot4.Base1 = r.Base1
						rank.Slot4.Base2 = r.Base2
					}
					isFound = true
					break
				}
				if isFound {
					break
				}
			}
			if isFound {
				break
			}
		}
		if !isFound {
			fmt.Println("Rank", r.RankID, "not found")
		}
	}

	return nil
}

func searchRanks(skill *AASkill, ranks []*AARank, id int) error {
	isFirst := true
	index := 1
	for {
		if id < 1 {
			return nil
		}
		next := findRank(ranks, id)
		if next == nil {
			fmt.Println("Skill", skill.ID, skill.Name, "has rank", id, "but it's not found in the ranks table")
			return nil
			//return fmt.Errorf("rank %d not found for skill %d", id, skill.ID)
		}
		if isFirst {
			skill.NameSID = next.TitleSID
			isFirst = false
		}
		next.Index = index
		index++
		skill.Ranks = append(skill.Ranks, next)
		id = next.NextID
	}
}

func findRank(ranks []*AARank, id int) *AARank {
	for _, rankEntry := range ranks {
		if rankEntry.ID == id {
			return rankEntry
		}
	}
	return nil
}
