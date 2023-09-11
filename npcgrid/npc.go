package npcgrid

type NpcYaml struct {
	ZoneID        int         `yaml:"zone_id,omitempty" db:"zoneid"`
	ZoneShortName string      `yaml:"zone_short_name,omitempty" db:"short_name"`
	Spawn2        []*Spawn2   `yaml:"spawn2,omitempty"`
	Grids         []*PathGrid `yaml:"grids,omitempty"`
}

type Spawn2 struct {
	ID           int     `yaml:"spawn2_id" db:"id"`               // int(11) NOT NULL AUTO_INCREMENT,
	SpawngroupID int     `yaml:"spawngroup_id" db:"spawngroupid"` // int(11) NOT NULL DEFAULT 0,
	Zone         string  `yaml:"zone,omitempty" db:"zone"`        // varchar(32) NOT NULL DEFAULT '',
	X            float32 `yaml:"spawn_x" db:"x"`                  // float(14,6) NOT NULL DEFAULT 0.000000,
	Y            float32 `yaml:"spawn_y" db:"y"`                  // float(14,6) NOT NULL DEFAULT 0.000000,
	Z            float32 `yaml:"spawn_z" db:"z"`                  // float(14,6) NOT NULL DEFAULT 0.000000,
	Heading      float32 `yaml:"heading" db:"heading"`            // float(14,6) NOT NULL DEFAULT 0.000000,
	Pathgrid     int     `yaml:"pathgrid" db:"pathgrid"`          // int(11) NOT NULL DEFAULT 0,
}

type PathGrid struct {
	Id      int              `yaml:"grid_id" db:"id"`              // int(10) NOT NULL DEFAULT 0,
	Zoneid  int              `yaml:"zoneid,omitempty" db:"zoneid"` // int(10) NOT NULL DEFAULT 0,
	Type    int              `yaml:"type" db:"type"`               // int(10) NOT NULL DEFAULT 0,
	Type2   int              `yaml:"type2" db:"type2"`             // int(10) NOT NULL DEFAULT 0,
	Entries []*PathGridEntry `yaml:"entries,omitempty" db:"entries"`
}

type PathGridEntry struct {
	Gridid      int     `yaml:"gridid,omitempty" db:"gridid"` // int(10) NOT NULL DEFAULT 0,
	Zoneid      int     `yaml:"zoneid,omitempty" db:"zoneid"` // int(10) NOT NULL DEFAULT 0,
	Number      int     `yaml:"number" db:"number"`           // int(10) NOT NULL DEFAULT 0,
	X           float32 `yaml:"grid_x" db:"x"`                // float NOT NULL DEFAULT 0,
	Y           float32 `yaml:"grid_y" db:"y"`                // float NOT NULL DEFAULT 0,
	Z           float32 `yaml:"grid_z" db:"z"`                // float NOT NULL DEFAULT 0,
	Heading     float32 `yaml:"grid_heading" db:"heading"`    // float NOT NULL DEFAULT 0,
	Pause       int     `yaml:"pause" db:"pause"`             // int(10) NOT NULL DEFAULT 0,
	Centerpoint int     `yaml:"centerpoint" db:"centerpoint"` // tinyint(4) NOT NULL DEFAULT 0,
}

func (e *NpcYaml) sanitize() error {
	return nil
}
