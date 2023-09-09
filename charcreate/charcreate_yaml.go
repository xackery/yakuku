package charcreate

type CharCreateYaml struct {
	Items       []*CharItem   `yaml:"items,omitempty" db:"items"`
	CharCreates []*CharCreate `yaml:"charCreates,omitempty" db:"charCreates"`
	Allocations []*Allocation `yaml:"allocations,omitempty" db:"allocations"`
}

type CharCreate struct {
	Race    string       `yaml:"race"`
	Items   []*CharItem  `yaml:"items,omitempty" db:"items"`
	Classes []*CharClass `yaml:"classes,omitempty" db:"classes"`
}

type CharClass struct {
	Class   string       `yaml:"class"`
	Items   []*CharItem  `yaml:"items,omitempty" db:"items"`
	Deities []*CharDeity `yaml:"deities,omitempty" db:"deities"`
}

type CharDeity struct {
	Deity string      `yaml:"deity"`
	Items []*CharItem `yaml:"items,omitempty" db:"items"`
	Zones []*CharZone `yaml:"zones,omitempty" db:"zones"`
}

type CharZone struct {
	Zone          string        `yaml:"zone"`
	Items         []*CharItem   `yaml:"items,omitempty" db:"items"`
	ExpansionsReq int           `yaml:"expansions_req"`
	AllocationID  int           `yaml:"allocation_id"`
	Choices       []*CharChoice `yaml:"choices,omitempty" db:"choices"`
}

type CharChoice struct {
	Index        int     `yaml:"index,omitempty"`
	SpawnX       float32 `yaml:"spawn_x,omitempty"`
	SpawnY       float32 `yaml:"spawn_y,omitempty"`
	SpawnZ       float32 `yaml:"spawn_z,omitempty"`
	SpawnHeading float32 `yaml:"spawn_heading,omitempty"`
	BindX        float32 `yaml:"bind_x,omitempty"`
	BindY        float32 `yaml:"bind_y,omitempty"`
	BindZ        float32 `yaml:"bind_z,omitempty"`
	BindID       int     `yaml:"bind_id,omitempty"`
}

type CharItem struct {
	ID          int    `yaml:"id"`
	ItemID      int    `yaml:"item_id,omitempty"`
	Name        string `yaml:"name,omitempty"`
	ItemCharges int    `yaml:"item_charges,omitempty"`
	GM          int    `yaml:"gm,omitempty"`
	Slot        int    `yaml:"slot,omitempty"`
}

type Allocation struct {
	ID       int `yaml:"id"`
	BaseStr  int `yaml:"base_str"`
	BaseSta  int `yaml:"base_sta"`
	BaseDex  int `yaml:"base_dex"`
	BaseAgi  int `yaml:"base_agi"`
	BaseInt  int `yaml:"base_int"`
	BaseWis  int `yaml:"base_wis"`
	BaseCha  int `yaml:"base_cha"`
	AllocStr int `yaml:"alloc_str"`
	AllocSta int `yaml:"alloc_sta"`
	AllocDex int `yaml:"alloc_dex"`
	AllocAgi int `yaml:"alloc_agi"`
	AllocInt int `yaml:"alloc_int"`
	AllocWis int `yaml:"alloc_wis"`
	AllocCha int `yaml:"alloc_cha"`
}

func (e *CharCreateYaml) sanitize() error {
	for _, charCreate := range e.CharCreates {
		err := charCreate.sanitize()
		if err != nil {
			return err
		}
	}
	return nil
}

func (charCreate *CharCreate) sanitize() error {

	return nil
}

func (charCreate *CharCreate) omitEmpty() error {
	return nil
}
