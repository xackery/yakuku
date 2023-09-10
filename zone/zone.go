package zone

import (
	"database/sql"
	"fmt"
)

type ZoneYaml struct {
	Zones []*Zone `yaml:"zones,omitempty"`
}

type Zone struct {
	Zoneidnumber            int            `yaml:"zoneidnumber" db:"zoneidnumber"`                           // int(4) NOT NULL DEFAULT 0,
	ShortName               sql.NullString `yaml:"short_name" db:"short_name"`                               // varchar(32) DEFAULT NULL,
	Id                      int            `yaml:"id" db:"id"`                                               // int(10) NOT NULL AUTO_INCREMENT,
	LongName                string         `yaml:"long_name" db:"long_name"`                                 // text NOT NULL,
	SafeX                   float32        `yaml:"safe_x" db:"safe_x"`                                       // float NOT NULL DEFAULT 0,
	SafeY                   float32        `yaml:"safe_y" db:"safe_y"`                                       // float NOT NULL DEFAULT 0,
	SafeZ                   float32        `yaml:"safe_z" db:"safe_z"`                                       // float NOT NULL DEFAULT 0,
	SafeHeading             float32        `yaml:"safe_heading" db:"safe_heading"`                           // float NOT NULL DEFAULT 0,
	GraveyardId             float32        `yaml:"graveyard_id" db:"graveyard_id"`                           // float NOT NULL DEFAULT 0,
	MinLevel                int            `yaml:"min_level" db:"min_level"`                                 // tinyint(3) unsigned NOT NULL DEFAULT 0,
	MaxLevel                int            `yaml:"max_level" db:"max_level"`                                 // tinyint(3) unsigned NOT NULL DEFAULT 255,
	MinStatus               int            `yaml:"min_status" db:"min_status"`                               // tinyint(3) unsigned NOT NULL DEFAULT 0,
	Version                 int            `yaml:"version" db:"version"`                                     // tinyint(3) unsigned NOT NULL DEFAULT 0,
	Timezone                int            `yaml:"timezone" db:"timezone"`                                   // int(5) NOT NULL DEFAULT 0,
	Maxclients              int            `yaml:"maxclients" db:"maxclients"`                               // int(5) NOT NULL DEFAULT 0,
	Ruleset                 int            `yaml:"ruleset" db:"ruleset"`                                     // int(10) unsigned NOT NULL DEFAULT 0,
	Note                    sql.NullString `yaml:"note" db:"note"`                                           // varchar(80) DEFAULT NULL,
	FileName                sql.NullString `yaml:"file_name" db:"file_name"`                                 // varchar(16) DEFAULT NULL,
	Underworld              float32        `yaml:"underworld" db:"underworld"`                               // float NOT NULL DEFAULT 0,
	Minclip                 float32        `yaml:"minclip" db:"minclip"`                                     // float NOT NULL DEFAULT 450,
	Maxclip                 float32        `yaml:"maxclip" db:"maxclip"`                                     // float NOT NULL DEFAULT 450,
	FogMinclip              float32        `yaml:"fog_minclip" db:"fog_minclip"`                             // float NOT NULL DEFAULT 450,
	FogMaxclip              float32        `yaml:"fog_maxclip" db:"fog_maxclip"`                             // float NOT NULL DEFAULT 450,
	FogBlue                 int            `yaml:"fog_blue" db:"fog_blue"`                                   // tinyint(3) unsigned NOT NULL DEFAULT 0,
	FogRed                  int            `yaml:"fog_red" db:"fog_red"`                                     // tinyint(3) unsigned NOT NULL DEFAULT 0,
	FogGreen                int            `yaml:"fog_green" db:"fog_green"`                                 // tinyint(3) unsigned NOT NULL DEFAULT 0,
	Sky                     int            `yaml:"sky" db:"sky"`                                             // tinyint(3) unsigned NOT NULL DEFAULT 1,
	Ztype                   int            `yaml:"ztype" db:"ztype"`                                         // tinyint(3) unsigned NOT NULL DEFAULT 1,
	ZoneExpMultiplier       string         `yaml:"zone_exp_multiplier" db:"zone_exp_multiplier"`             // decimal(6,2) NOT NULL DEFAULT 0.00,
	Walkspeed               float32        `yaml:"walkspeed" db:"walkspeed"`                                 // float NOT NULL DEFAULT 0.4,
	TimeType                int            `yaml:"time_type" db:"time_type"`                                 // tinyint(3) unsigned NOT NULL DEFAULT 2,
	FogRed1                 int            `yaml:"fog_red1" db:"fog_red1"`                                   // tinyint(3) unsigned NOT NULL DEFAULT 0,
	FogGreen1               int            `yaml:"fog_green1" db:"fog_green1"`                               // tinyint(3) unsigned NOT NULL DEFAULT 0,
	FogBlue1                int            `yaml:"fog_blue1" db:"fog_blue1"`                                 // tinyint(3) unsigned NOT NULL DEFAULT 0,
	FogMinclip1             float32        `yaml:"fog_minclip1" db:"fog_minclip1"`                           // float NOT NULL DEFAULT 450,
	FogMaxclip1             float32        `yaml:"fog_maxclip1" db:"fog_maxclip1"`                           // float NOT NULL DEFAULT 450,
	FogRed2                 int            `yaml:"fog_red2" db:"fog_red2"`                                   // tinyint(3) unsigned NOT NULL DEFAULT 0,
	FogGreen2               int            `yaml:"fog_green2" db:"fog_green2"`                               // tinyint(3) unsigned NOT NULL DEFAULT 0,
	FogBlue2                int            `yaml:"fog_blue2" db:"fog_blue2"`                                 // tinyint(3) unsigned NOT NULL DEFAULT 0,
	FogMinclip2             float32        `yaml:"fog_minclip2" db:"fog_minclip2"`                           // float NOT NULL DEFAULT 450,
	FogMaxclip2             float32        `yaml:"fog_maxclip2" db:"fog_maxclip2"`                           // float NOT NULL DEFAULT 450,
	FogRed3                 int            `yaml:"fog_red3" db:"fog_red3"`                                   // tinyint(3) unsigned NOT NULL DEFAULT 0,
	FogGreen3               int            `yaml:"fog_green3" db:"fog_green3"`                               // tinyint(3) unsigned NOT NULL DEFAULT 0,
	FogBlue3                int            `yaml:"fog_blue3" db:"fog_blue3"`                                 // tinyint(3) unsigned NOT NULL DEFAULT 0,
	FogMinclip3             float32        `yaml:"fog_minclip3" db:"fog_minclip3"`                           // float NOT NULL DEFAULT 450,
	FogMaxclip3             float32        `yaml:"fog_maxclip3" db:"fog_maxclip3"`                           // float NOT NULL DEFAULT 450,
	FogRed4                 int            `yaml:"fog_red4" db:"fog_red4"`                                   // tinyint(3) unsigned NOT NULL DEFAULT 0,
	FogGreen4               int            `yaml:"fog_green4" db:"fog_green4"`                               // tinyint(3) unsigned NOT NULL DEFAULT 0,
	FogBlue4                int            `yaml:"fog_blue4" db:"fog_blue4"`                                 // tinyint(3) unsigned NOT NULL DEFAULT 0,
	FogMinclip4             float32        `yaml:"fog_minclip4" db:"fog_minclip4"`                           // float NOT NULL DEFAULT 450,
	FogMaxclip4             float32        `yaml:"fog_maxclip4" db:"fog_maxclip4"`                           // float NOT NULL DEFAULT 450,
	FogDensity              float32        `yaml:"fog_density" db:"fog_density"`                             // float NOT NULL DEFAULT 0,
	FlagNeeded              string         `yaml:"flag_needed" db:"flag_needed"`                             // varchar(128) NOT NULL DEFAULT '',
	Canbind                 int            `yaml:"canbind" db:"canbind"`                                     // tinyint(4) NOT NULL DEFAULT 1,
	Cancombat               int            `yaml:"cancombat" db:"cancombat"`                                 // tinyint(4) NOT NULL DEFAULT 1,
	Canlevitate             int            `yaml:"canlevitate" db:"canlevitate"`                             // tinyint(4) NOT NULL DEFAULT 1,
	Castoutdoor             int            `yaml:"castoutdoor" db:"castoutdoor"`                             // tinyint(4) NOT NULL DEFAULT 1,
	Hotzone                 int            `yaml:"hotzone" db:"hotzone"`                                     // tinyint(3) unsigned NOT NULL DEFAULT 0,
	Insttype                int            `yaml:"insttype" db:"insttype"`                                   // tinyint(1) unsigned zerofill NOT NULL DEFAULT 0,
	Shutdowndelay           string         `yaml:"shutdowndelay" db:"shutdowndelay"`                         // bigint(16) unsigned NOT NULL DEFAULT 5000,
	Peqzone                 int            `yaml:"peqzone" db:"peqzone"`                                     // tinyint(4) NOT NULL DEFAULT 1,
	Expansion               int            `yaml:"expansion" db:"expansion"`                                 // tinyint(3) NOT NULL DEFAULT 0,
	BypassExpansionCheck    int            `yaml:"bypass_expansion_check" db:"bypass_expansion_check"`       // tinyint(3) NOT NULL DEFAULT 0,
	Suspendbuffs            int            `yaml:"suspendbuffs" db:"suspendbuffs"`                           // tinyint(1) unsigned NOT NULL DEFAULT 0,
	RainChance1             int            `yaml:"rain_chance1" db:"rain_chance1"`                           // int(4) NOT NULL DEFAULT 0,
	RainChance2             int            `yaml:"rain_chance2" db:"rain_chance2"`                           // int(4) NOT NULL DEFAULT 0,
	RainChance3             int            `yaml:"rain_chance3" db:"rain_chance3"`                           // int(4) NOT NULL DEFAULT 0,
	RainChance4             int            `yaml:"rain_chance4" db:"rain_chance4"`                           // int(4) NOT NULL DEFAULT 0,
	RainDuration1           int            `yaml:"rain_duration1" db:"rain_duration1"`                       // int(4) NOT NULL DEFAULT 0,
	RainDuration2           int            `yaml:"rain_duration2" db:"rain_duration2"`                       // int(4) NOT NULL DEFAULT 0,
	RainDuration3           int            `yaml:"rain_duration3" db:"rain_duration3"`                       // int(4) NOT NULL DEFAULT 0,
	RainDuration4           int            `yaml:"rain_duration4" db:"rain_duration4"`                       // int(4) NOT NULL DEFAULT 0,
	SnowChance1             int            `yaml:"snow_chance1" db:"snow_chance1"`                           // int(4) NOT NULL DEFAULT 0,
	SnowChance2             int            `yaml:"snow_chance2" db:"snow_chance2"`                           // int(4) NOT NULL DEFAULT 0,
	SnowChance3             int            `yaml:"snow_chance3" db:"snow_chance3"`                           // int(4) NOT NULL DEFAULT 0,
	SnowChance4             int            `yaml:"snow_chance4" db:"snow_chance4"`                           // int(4) NOT NULL DEFAULT 0,
	SnowDuration1           int            `yaml:"snow_duration1" db:"snow_duration1"`                       // int(4) NOT NULL DEFAULT 0,
	SnowDuration2           int            `yaml:"snow_duration2" db:"snow_duration2"`                       // int(4) NOT NULL DEFAULT 0,
	SnowDuration3           int            `yaml:"snow_duration3" db:"snow_duration3"`                       // int(4) NOT NULL DEFAULT 0,
	SnowDuration4           int            `yaml:"snow_duration4" db:"snow_duration4"`                       // int(4) NOT NULL DEFAULT 0,
	Gravity                 float32        `yaml:"gravity" db:"gravity"`                                     // float NOT NULL DEFAULT 0.4,
	Type                    int            `yaml:"type" db:"type"`                                           // int(3) NOT NULL DEFAULT 0,
	Skylock                 int            `yaml:"skylock" db:"skylock"`                                     // tinyint(4) NOT NULL DEFAULT 0,
	FastRegenHp             int            `yaml:"fast_regen_hp" db:"fast_regen_hp"`                         // int(11) NOT NULL DEFAULT 180,
	FastRegenMana           int            `yaml:"fast_regen_mana" db:"fast_regen_mana"`                     // int(11) NOT NULL DEFAULT 180,
	FastRegenEndurance      int            `yaml:"fast_regen_endurance" db:"fast_regen_endurance"`           // int(11) NOT NULL DEFAULT 180,
	NpcMaxAggroDist         int            `yaml:"npc_max_aggro_dist" db:"npc_max_aggro_dist"`               // int(11) NOT NULL DEFAULT 600,
	MaxMovementUpdateRange  int            `yaml:"max_movement_update_range" db:"max_movement_update_range"` // int(11) unsigned NOT NULL DEFAULT 600,
	MinExpansion            int            `yaml:"min_expansion" db:"min_expansion"`                         // tinyint(4) NOT NULL DEFAULT -1,
	MaxExpansion            int            `yaml:"max_expansion" db:"max_expansion"`                         // tinyint(4) NOT NULL DEFAULT -1,
	ContentFlags            sql.NullString `yaml:"content_flags" db:"content_flags"`                         // varchar(100) DEFAULT NULL,
	ContentFlagsDisabled    sql.NullString `yaml:"content_flags_disabled" db:"content_flags_disabled"`       // varchar(100) DEFAULT NULL,
	UnderworldTeleportIndex int            `yaml:"underworld_teleport_index" db:"underworld_teleport_index"` // int(4) NOT NULL DEFAULT 0,
	LavaDamage              int            `yaml:"lava_damage" db:"lava_damage"`                             // int(11) DEFAULT 50,
	MinLavaDamage           int            `yaml:"min_lava_damage" db:"min_lava_damage"`                     // int(11) NOT NULL DEFAULT 10,
	MapFileName             sql.NullString `yaml:"map_file_name" db:"map_file_name"`                         // varchar(100) DEFAULT NULL,
	Points                  []*ZonePoint   `yaml:"points" db:"points"`
}

type ZonePoint struct {
	Id                   int            `yaml:"id" db:"id"`                                         // int(11) NOT NULL AUTO_INCREMENT,
	Zone                 sql.NullString `yaml:"zone" db:"zone"`                                     // varchar(32) DEFAULT NULL,
	Number               int            `yaml:"number" db:"number"`                                 // smallint(4) unsigned NOT NULL DEFAULT 1,
	TargetZoneId         int            `yaml:"target_zone_id" db:"target_zone_id"`                 // int(10) unsigned NOT NULL DEFAULT 0,
	Zoneinst             int            `yaml:"zoneinst" db:"zoneinst"`                             // smallint(5) unsigned DEFAULT 0,
	TargetX              float32        `yaml:"target_x" db:"target_x"`                             // float NOT NULL DEFAULT 0,
	TargetY              float32        `yaml:"target_y" db:"target_y"`                             // float NOT NULL DEFAULT 0,
	TargetZ              float32        `yaml:"target_z" db:"target_z"`                             // float NOT NULL DEFAULT 0,
	TargetHeading        float32        `yaml:"target_heading" db:"target_heading"`                 // float NOT NULL DEFAULT 0,
	TargetInstance       int            `yaml:"target_instance" db:"target_instance"`               // int(10) unsigned NOT NULL DEFAULT 0,
	Version              int            `yaml:"version" db:"version"`                               // int(11) NOT NULL DEFAULT 0,
	Buffer               float32        `yaml:"buffer" db:"buffer"`                                 // float DEFAULT 0,
	ClientVersionMask    int            `yaml:"client_version_mask" db:"client_version_mask"`       // int(10) unsigned NOT NULL DEFAULT 4294967295,
	MinExpansion         int            `yaml:"min_expansion" db:"min_expansion"`                   // tinyint(4) NOT NULL DEFAULT -1,
	MaxExpansion         int            `yaml:"max_expansion" db:"max_expansion"`                   // tinyint(4) NOT NULL DEFAULT -1,
	ContentFlags         sql.NullString `yaml:"content_flags" db:"content_flags"`                   // varchar(100) DEFAULT NULL,
	ContentFlagsDisabled sql.NullString `yaml:"content_flags_disabled" db:"content_flags_disabled"` // varchar(100) DEFAULT NULL,
	IsVirtual            int            `yaml:"is_virtual" db:"is_virtual"`                         // tinyint(4) NOT NULL DEFAULT 0,
	Height               int            `yaml:"height" db:"height"`                                 // int(11) NOT NULL DEFAULT 0,
	Width                int            `yaml:"width" db:"width"`                                   // int(11) NOT NULL DEFAULT 0,
	Y                    float32        `yaml:"zone_y" db:"y"`                                      // float NOT NULL DEFAULT 0,
	X                    float32        `yaml:"zone_x" db:"x"`                                      // float NOT NULL DEFAULT 0,
	Z                    float32        `yaml:"zone_z" db:"z"`                                      // float NOT NULL DEFAULT 0,
	Heading              float32        `yaml:"heading" db:"heading"`                               // float NOT NULL DEFAULT 0,
}

func (e *ZoneYaml) sanitize() error {
	for _, zone := range e.Zones {

		if !zone.ShortName.Valid {
			return fmt.Errorf("short name must has valid flagged")
		}

		if zone.ShortName.String == "" {
			return fmt.Errorf("short name must not be empty")
		}
	}
	return nil
}
