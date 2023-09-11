package startingitems

import (
	"database/sql"
	"fmt"
)

type StartingItemYaml struct {
	StartingItems []*StartingItem `yaml:"starting_items,omitempty"`
}

type StartingItem struct {
	ID                   int            `yaml:"id,omitempty" db:"id"`                                         // int(11) unsigned NOT NULL AUTO_INCREMENT,
	Race                 int            `yaml:"race" db:"race"`                                               // int(11) NOT NULL DEFAULT 0,
	Class                int            `yaml:"class" db:"class"`                                             // int(11) NOT NULL DEFAULT 0,
	Deityid              int            `yaml:"deity_id,omitempty" db:"deityid"`                              // int(11) NOT NULL DEFAULT 0,
	ZoneID               int            `yaml:"zone_id,omitempty" db:"zoneid"`                                // int(11) NOT NULL DEFAULT 0,
	ItemID               int            `yaml:"item_id,omitempty" db:"itemid"`                                // int(11) NOT NULL DEFAULT 0,
	Name                 string         `yaml:"name,omitempty" db:"name"`                                     // varchar(64) NOT NULL DEFAULT '',
	ItemCharges          int            `yaml:"item_charges,omitempty" db:"item_charges"`                     // tinyint(3) unsigned NOT NULL DEFAULT 1,
	Gm                   int            `yaml:"gm,omitempty" db:"gm"`                                         // tinyint(1) NOT NULL DEFAULT 0,
	Slot                 int            `yaml:"slot,omitempty" db:"slot"`                                     // mediumint(9) NOT NULL DEFAULT -1,
	MinExpansion         int            `yaml:"min_expansion,omitempty" db:"min_expansion"`                   // tinyint(4) NOT NULL DEFAULT -1,
	MaxExpansion         int            `yaml:"max_expansion,omitempty" db:"max_expansion"`                   // tinyint(4) NOT NULL DEFAULT -1,
	ContentFlags         sql.NullString `yaml:"content_flags,omitempty" db:"content_flags"`                   // varchar(100) DEFAULT NULL,
	ContentFlagsDisabled sql.NullString `yaml:"content_flags_disabled,omitempty" db:"content_flags_disabled"` // varchar(100) DEFAULT NULL,
}

func (e *StartingItemYaml) sanitize() error {
	for _, item := range e.StartingItems {
		if item.ID == 0 {
			return fmt.Errorf("starting_items.id cannot be 0")
		}
	}
	return nil
}

func (e *StartingItemYaml) omitEmpty() error {

	return nil
}
