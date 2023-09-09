package aa

import "fmt"

type AAYaml struct {
	Skills []*AASkill `yaml:"skills,omitempty"`
}

type AASkill struct {
	ID              int       `yaml:"id,omitempty" db:"id"`
	Icon            string    `yaml:"icon,omitempty" db:"icon"`
	Name            string    `yaml:"aa_name,omitempty" db:"name"`
	NameSID         int       `yaml:"aa_name_sid,omitempty"`
	Category        int       `yaml:"category,omitempty" db:"category"`
	Classes         int       `yaml:"classes,omitempty" db:"classes"`
	Races           int       `yaml:"races,omitempty" db:"races"` // if 0, default to 65535
	DrakkinHeritage int       `yaml:"drakkin_heritage,omitempty" db:"drakkin_heritage"`
	Deities         int       `yaml:"deities,omitempty" db:"deities"`
	Status          int       `yaml:"status,omitempty" db:"status"`
	Type            int       `yaml:"type,omitempty" db:"type"`
	Charges         int       `yaml:"charges,omitempty" db:"charges"`
	GrantOnly       int       `yaml:"grant_only,omitempty" db:"grant_only"`
	FirstRankID     int       `yaml:"first_rank_id,omitempty" db:"first_rank_id"`
	Enabled         int       `yaml:"enabled,omitempty" db:"enabled"`
	ResetOnDeath    int       `yaml:"reset_on_death,omitempty" db:"reset_on_death"`
	LevelReq        int       `yaml:"level_req,omitempty" db:"level_req"`
	Ranks           []*AARank `yaml:"ranks,omitempty" db:"ranks"`
}

type AARank struct {
	Index          int    `yaml:"index,omitempty" db:"index"`
	ID             int    `yaml:"id,omitempty" db:"id"`
	UpperHotkeySID int    `yaml:"upper_hotkey_sid,omitempty" db:"upper_hotkey_sid"`
	LowerHotkeySID int    `yaml:"lower_hotkey_sid,omitempty" db:"lower_hotkey_sid"`
	Title          string `yaml:"title,omitempty" db:"title"`
	TitleSID       int    `yaml:"title_sid,omitempty" db:"title_sid"`
	Description    string `yaml:"description,omitempty" db:"description"`
	DescriptionSID int    `yaml:"desc_sid,omitempty" db:"desc_sid"`
	Cost           int    `yaml:"cost,omitempty" db:"cost"`
	LevelReq       int    `yaml:"level_req,omitempty" db:"level_req"`
	SpellID        int    `yaml:"spell_id,omitempty" db:"spell"`
	SpellType      int    `yaml:"spell_type,omitempty" db:"spell_type"`
	RecastTime     int    `yaml:"recast_time,omitempty" db:"recast_time"`
	Expansion      int    `yaml:"expansion,omitempty" db:"expansion"`
	PrevID         int    `yaml:"prev_id,omitempty" db:"prev_id"`
	NextID         int    `yaml:"next_id,omitempty" db:"next_id"`
	Slot1          struct {
		EffectID int `yaml:"effect_id,omitempty"`
		Base1    int `yaml:"base1,omitempty"`
		Base2    int `yaml:"base2"`
	} `yaml:"slot1,omitempty"`
	Slot2 struct {
		EffectID int `yaml:"effect_id,omitempty"`
		Base1    int `yaml:"base1,omitempty"`
		Base2    int `yaml:"base2"`
	} `yaml:"slot2,omitempty"`
	Slot3 struct {
		EffectID int `yaml:"effect_id,omitempty"`
		Base1    int `yaml:"base1,omitempty"`
		Base2    int `yaml:"base2"`
	} `yaml:"slot3,omitempty"`
	Slot4 struct {
		EffectID int `yaml:"effect_id,omitempty"`
		Base1    int `yaml:"base1,omitempty"`
		Base2    int `yaml:"base2"`
	} `yaml:"slot4,omitempty"`
	Slot5 struct {
		EffectID int `yaml:"effect_id,omitempty"`
		Base1    int `yaml:"base1,omitempty"`
		Base2    int `yaml:"base2"`
	} `yaml:"slot5,omitempty"`
	Slot6 struct {
		EffectID int `yaml:"effect_id,omitempty"`
		Base1    int `yaml:"base1,omitempty"`
		Base2    int `yaml:"base2"`
	} `yaml:"slot6,omitempty"`
	Slot7 struct {
		EffectID int `yaml:"effect_id,omitempty"`
		Base1    int `yaml:"base1,omitempty"`
		Base2    int `yaml:"base2"`
	} `yaml:"slot7,omitempty"`
}

func (e *AAYaml) sanitize() error {

	abilityNames := make(map[int]string)
	titleNames := make(map[int]string)
	descNames := make(map[int]string)

	uniqueRankIDs := make(map[int]bool)

	for skillIndex, skill := range e.Skills {
		abilityName, ok := abilityNames[skill.ID]
		if !ok {
			abilityNames[skill.ID] = skill.Name
			abilityName = skill.Name
		}
		if abilityName != skill.Name {
			return fmt.Errorf("ability name mismatch for skill id %d between '%s' and '%s'", skill.ID, abilityName, skill.Name)
		}
		titleName, ok := titleNames[skill.NameSID]
		if !ok {
			titleNames[skill.NameSID] = skill.Name
			titleName = skill.Name
		}
		if titleName != skill.Name {
			return fmt.Errorf("title name mismatch for nameSID %d for skill id %d between '%s' and '%s'", skill.NameSID, skillIndex, titleName, skill.Name)
		}
		for rankIndex, rank := range skill.Ranks {
			if rank.ID == 0 {
				return fmt.Errorf("rank id is 0 for skill id %d rank %d", skillIndex, rankIndex)
			}
			_, ok := uniqueRankIDs[rank.ID]
			if ok {
				return fmt.Errorf("duplicate rank id %d for skill id %d rank %d", rank.ID, skillIndex, rankIndex)
			}
			uniqueRankIDs[rank.ID] = true
			if rank.TitleSID != 0 {
				titleName, ok := titleNames[rank.TitleSID]
				if !ok {
					titleNames[rank.TitleSID] = rank.Title
					titleName = rank.Title
				}
				if titleName != rank.Title {
					return fmt.Errorf("title name mismatch for titleSID %d skill id %d rank %d between '%s' and '%s'", rank.TitleSID, skillIndex, rankIndex, titleName, rank.Title)
				}
			}

			descName, ok := descNames[rank.DescriptionSID]
			if !ok {
				descNames[rank.DescriptionSID] = rank.Description
				descName = rank.Description
			}
			if descName != rank.Description {
				return fmt.Errorf("description name mismatch for descriptionSID %d skill id %d rank %d between '%s' and '%s'", rank.DescriptionSID, skillIndex, rankIndex, descName, rank.Description)
			}
		}
	}
	return nil
}

func (e *AAYaml) omitEmpty() error {
	for _, skill := range e.Skills {
		if skill.Category == -1 {
			skill.Category = 0
		}
		for _, rank := range skill.Ranks {
			if rank.SpellID == -1 {
				rank.SpellID = 0
			}
			if rank.PrevID == -1 {
				rank.PrevID = 0
			}
			if rank.NextID == -1 {
				rank.NextID = 0
			}
			if rank.UpperHotkeySID == -1 {
				rank.UpperHotkeySID = 0
			}
			if rank.LowerHotkeySID == -1 {
				rank.LowerHotkeySID = 0
			}
		}
	}
	return nil
}
