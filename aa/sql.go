package aa

import (
	"fmt"
	"os"

	"github.com/xackery/yakuku/util"
)

func generateAASQL(aa *AAYaml) error {
	w, err := os.Create("aa.sql")
	if err != nil {
		return err
	}
	defer w.Close()

	w.WriteString("TRUNCATE aa_ability;\n")
	w.WriteString("INSERT INTO aa_ability (id, `name`, category, classes, races, drakkin_heritage, deities, `status`, `type`, charges, grant_only, first_rank_id, `enabled`, reset_on_death) VALUES\n")
	for skillIndex, skill := range aa.Skills {
		if len(skill.Ranks) == 0 {
			continue
		}
		w.WriteString(fmt.Sprintf("	(%d, '%s', -1, %d, 65535, 127, 131071, 0, %d, 0, %d, %d, 1, 0)", skill.ID, util.EscapeSQL(skill.Name), skill.Classes, skill.Type, skill.GrantOnly, skill.Ranks[0].ID))
		if skillIndex == len(aa.Skills)-1 {
			w.WriteString(";\n\n")
		} else {
			w.WriteString(",\n")
		}
	}

	w.WriteString("TRUNCATE aa_rank_prereqs;\n")

	w.WriteString("TRUNCATE aa_rank_effects;\n")
	w.WriteString("INSERT INTO aa_rank_effects (rank_id, slot, effect_id, base1, base2) VALUES\n")
	levelReq := 1

	for skillIndex, skill := range aa.Skills {
		for rankIndex, rank := range skill.Ranks {
			isDone := false

			if rank.Slot1.EffectID == 0 && rank.Slot2.EffectID == 0 && rank.Slot3.EffectID == 0 && rank.Slot4.EffectID == 0 {
				continue
			}

			if rank.LevelReq > levelReq {
				levelReq = rank.LevelReq
			}

			if rank.LevelReq == 0 {
				rank.LevelReq = levelReq
			}

			w.WriteString(fmt.Sprintf("	(%d, 1, %d, %d, %d)", rank.ID, rank.Slot1.EffectID, rank.Slot1.Base1, rank.Slot1.Base2))
			if rankIndex == len(skill.Ranks)-1 && skillIndex == len(aa.Skills)-1 {
				isDone = true
			} else {
				w.WriteString(",\n")
			}

			if rank.Slot2.EffectID != 0 {
				if isDone {
					w.WriteString(",\n")
				}
				w.WriteString(fmt.Sprintf("	(%d, 2, %d, %d, %d)", rank.ID, rank.Slot2.EffectID, rank.Slot2.Base1, rank.Slot2.Base2))
				if rankIndex == len(skill.Ranks)-1 && skillIndex == len(aa.Skills)-1 {
					isDone = true
				} else {
					w.WriteString(",\n")
				}
			}

			if rank.Slot3.EffectID != 0 {
				if isDone {
					w.WriteString(",\n")
				}
				w.WriteString(fmt.Sprintf("	(%d, 3, %d, %d, %d)", rank.ID, rank.Slot3.EffectID, rank.Slot3.Base1, rank.Slot3.Base2))
				if rankIndex == len(skill.Ranks)-1 && skillIndex == len(aa.Skills)-1 {
					isDone = true
				} else {
					w.WriteString(",\n")
				}
			}

			if rank.Slot4.EffectID != 0 {
				if isDone {
					w.WriteString(",\n")
				}
				w.WriteString(fmt.Sprintf("	(%d, 4, %d, %d, %d)", rank.ID, rank.Slot4.EffectID, rank.Slot4.Base1, rank.Slot4.Base2))
				if rankIndex == len(skill.Ranks)-1 && skillIndex == len(aa.Skills)-1 {
					isDone = true
				} else {
					w.WriteString(",\n")
				}
			}

			if rank.Slot5.EffectID != 0 {
				if isDone {
					w.WriteString(",\n")
				}
				w.WriteString(fmt.Sprintf("	(%d, 5, %d, %d, %d)", rank.ID, rank.Slot5.EffectID, rank.Slot5.Base1, rank.Slot5.Base2))
				if rankIndex == len(skill.Ranks)-1 && skillIndex == len(aa.Skills)-1 {
					isDone = true
				} else {
					w.WriteString(",\n")
				}
			}

			if rank.Slot6.EffectID != 0 {
				if isDone {
					w.WriteString(",\n")
				}
				w.WriteString(fmt.Sprintf("	(%d, 6, %d, %d, %d)", rank.ID, rank.Slot6.EffectID, rank.Slot6.Base1, rank.Slot6.Base2))
				if rankIndex == len(skill.Ranks)-1 && skillIndex == len(aa.Skills)-1 {
					isDone = true
				} else {
					w.WriteString(",\n")
				}
			}

			if isDone {
				w.WriteString(";\n\n")
			}
		}
	}

	w.WriteString("TRUNCATE aa_ranks;\n")
	w.WriteString("INSERT INTO aa_ranks (id, upper_hotkey_sid, lower_hotkey_sid, title_sid, desc_sid, cost, level_req, spell, spell_type, recast_time, expansion, prev_id, next_id) VALUES\n")

	lastSID := 0
	for skillIndex, skill := range aa.Skills {
		if skill.NameSID != 0 {
			lastSID = skill.NameSID
		}
		for rankIndex, rank := range skill.Ranks {
			prevID := -1
			nextID := -1

			if rank.LevelReq == 0 && skill.LevelReq != 0 {
				rank.LevelReq = skill.LevelReq
			}

			if rank.TitleSID != 0 {
				lastSID = rank.TitleSID
			}

			if rankIndex > 0 {
				prevID = skill.Ranks[rankIndex-1].ID
			}
			if rankIndex < len(skill.Ranks)-1 {
				nextID = skill.Ranks[rankIndex+1].ID
			}

			w.WriteString(fmt.Sprintf("	(%d, -1, -1, %d, %d, %d, %d, 0, %d, %d, 0, %d, %d)", rank.ID, lastSID, rank.DescriptionSID, rank.Cost, rank.LevelReq, rank.SpellID, rank.SpellType, prevID, nextID))
			if rankIndex == len(skill.Ranks)-1 && skillIndex == len(aa.Skills)-1 {
				w.WriteString(";\n\n")
			} else {
				w.WriteString(",\n")
			}
		}
	}

	return nil
}
