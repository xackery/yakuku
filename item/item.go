package item

import (
	"database/sql"
	"fmt"
	"time"
)

type ItemYaml struct {
	Items []*Item `yaml:"items,omitempty"`
}

type Item struct {
	ID                  int            `yaml:"id,omitempty" db:"id"`                                   // int(11) NOT NULL DEFAULT 0,
	Minstatus           int            `yaml:"minstatus,omitempty" db:"minstatus"`                     // smallint(5) NOT NULL DEFAULT 0,
	Name                string         `yaml:"name,omitempty" db:"Name"`                               // varchar(64) NOT NULL DEFAULT '',
	Aagi                int            `yaml:"aagi,omitempty" db:"aagi"`                               // int(11) NOT NULL DEFAULT 0,
	Ac                  int            `yaml:"ac,omitempty" db:"ac"`                                   // int(11) NOT NULL DEFAULT 0,
	Accuracy            int            `yaml:"accuracy,omitempty" db:"accuracy"`                       // int(11) NOT NULL DEFAULT 0,
	Acha                int            `yaml:"acha,omitempty" db:"acha"`                               // int(11) NOT NULL DEFAULT 0,
	Adex                int            `yaml:"adex,omitempty" db:"adex"`                               // int(11) NOT NULL DEFAULT 0,
	Aint                int            `yaml:"aint,omitempty" db:"aint"`                               // int(11) NOT NULL DEFAULT 0,
	Artifactflag        int            `yaml:"artifactflag,omitempty" db:"artifactflag"`               // tinyint(3) unsigned NOT NULL DEFAULT 0,
	Asta                int            `yaml:"asta,omitempty" db:"asta"`                               // int(11) NOT NULL DEFAULT 0,
	Astr                int            `yaml:"astr,omitempty" db:"astr"`                               // int(11) NOT NULL DEFAULT 0,
	Attack              int            `yaml:"attack,omitempty" db:"attack"`                           // int(11) NOT NULL DEFAULT 0,
	Augrestrict         int            `yaml:"augrestrict,omitempty" db:"augrestrict"`                 // int(11) NOT NULL DEFAULT 0,
	Augslot1type        int            `yaml:"augslot1type,omitempty" db:"augslot1type"`               // tinyint(3) NOT NULL DEFAULT 0,
	Augslot1visible     int            `yaml:"augslot1visible,omitempty" db:"augslot1visible"`         // tinyint(3) NOT NULL DEFAULT 0,
	Augslot2type        int            `yaml:"augslot2type,omitempty" db:"augslot2type"`               // tinyint(3) NOT NULL DEFAULT 0,
	Augslot2visible     int            `yaml:"augslot2visible,omitempty" db:"augslot2visible"`         // tinyint(3) NOT NULL DEFAULT 0,
	Augslot3type        int            `yaml:"augslot3type,omitempty" db:"augslot3type"`               // tinyint(3) NOT NULL DEFAULT 0,
	Augslot3visible     int            `yaml:"augslot3visible,omitempty" db:"augslot3visible"`         // tinyint(3) NOT NULL DEFAULT 0,
	Augslot4type        int            `yaml:"augslot4type,omitempty" db:"augslot4type"`               // tinyint(3) NOT NULL DEFAULT 0,
	Augslot4visible     int            `yaml:"augslot4visible,omitempty" db:"augslot4visible"`         // tinyint(3) NOT NULL DEFAULT 0,
	Augslot5type        int            `yaml:"augslot5type,omitempty" db:"augslot5type"`               // tinyint(3) NOT NULL DEFAULT 0,
	Augslot5visible     int            `yaml:"augslot5visible,omitempty" db:"augslot5visible"`         // tinyint(3) NOT NULL DEFAULT 0,
	Augslot6type        int            `yaml:"augslot6type,omitempty" db:"augslot6type"`               // tinyint(3) NOT NULL DEFAULT 0,
	Augslot6visible     int            `yaml:"augslot6visible,omitempty" db:"augslot6visible"`         // tinyint(3) NOT NULL DEFAULT 0,
	Augtype             int            `yaml:"augtype,omitempty" db:"augtype"`                         // int(11) NOT NULL DEFAULT 0,
	Avoidance           int            `yaml:"avoidance,omitempty" db:"avoidance"`                     // int(11) NOT NULL DEFAULT 0,
	Awis                int            `yaml:"awis,omitempty" db:"awis"`                               // int(11) NOT NULL DEFAULT 0,
	Bagsize             int            `yaml:"bagsize,omitempty" db:"bagsize"`                         // int(11) NOT NULL DEFAULT 0,
	Bagslots            int            `yaml:"bagslots,omitempty" db:"bagslots"`                       // int(11) NOT NULL DEFAULT 0,
	Bagtype             int            `yaml:"bagtype,omitempty" db:"bagtype"`                         // int(11) NOT NULL DEFAULT 0,
	Bagwr               int            `yaml:"bagwr,omitempty" db:"bagwr"`                             // int(11) NOT NULL DEFAULT 0,
	Banedmgamt          int            `yaml:"banedmgamt,omitempty" db:"banedmgamt"`                   // int(11) NOT NULL DEFAULT 0,
	Banedmgraceamt      int            `yaml:"banedmgraceamt,omitempty" db:"banedmgraceamt"`           // int(11) NOT NULL DEFAULT 0,
	Banedmgbody         int            `yaml:"banedmgbody,omitempty" db:"banedmgbody"`                 // int(11) NOT NULL DEFAULT 0,
	Banedmgrace         int            `yaml:"banedmgrace,omitempty" db:"banedmgrace"`                 // int(11) NOT NULL DEFAULT 0,
	Bardtype            int            `yaml:"bardtype,omitempty" db:"bardtype"`                       // int(11) NOT NULL DEFAULT 0,
	Bardvalue           int            `yaml:"bardvalue,omitempty" db:"bardvalue"`                     // int(11) NOT NULL DEFAULT 0,
	Book                int            `yaml:"book,omitempty" db:"book"`                               // int(11) NOT NULL DEFAULT 0,
	Casttime            int            `yaml:"casttime,omitempty" db:"casttime"`                       // int(11) NOT NULL DEFAULT 0,
	Casttime_           int            `yaml:"casttime_,omitempty" db:"casttime_"`                     // int(11) NOT NULL DEFAULT 0,
	Charmfile           string         `yaml:"charmfile,omitempty" db:"charmfile"`                     // varchar(32) NOT NULL DEFAULT '',
	Charmfileid         string         `yaml:"charmfileid,omitempty" db:"charmfileid"`                 // varchar(32) NOT NULL DEFAULT '',
	Classes             int            `yaml:"classes,omitempty" db:"classes"`                         // int(11) NOT NULL DEFAULT 0,
	Color               int            `yaml:"color,omitempty" db:"color"`                             // int(10) unsigned NOT NULL DEFAULT 0,
	Combateffects       string         `yaml:"combateffects,omitempty" db:"combateffects"`             // varchar(10) NOT NULL DEFAULT '',
	Extradmgskill       int            `yaml:"extradmgskill,omitempty" db:"extradmgskill"`             // int(11) NOT NULL DEFAULT 0,
	Extradmgamt         int            `yaml:"extradmgamt,omitempty" db:"extradmgamt"`                 // int(11) NOT NULL DEFAULT 0,
	Price               int            `yaml:"price,omitempty" db:"price"`                             // int(11) NOT NULL DEFAULT 0,
	Cr                  int            `yaml:"cr,omitempty" db:"cr"`                                   // int(11) NOT NULL DEFAULT 0,
	Damage              int            `yaml:"damage,omitempty" db:"damage"`                           // int(11) NOT NULL DEFAULT 0,
	Damageshield        int            `yaml:"damageshield,omitempty" db:"damageshield"`               // int(11) NOT NULL DEFAULT 0,
	Deity               int            `yaml:"deity,omitempty" db:"deity"`                             // int(11) NOT NULL DEFAULT 0,
	Delay               int            `yaml:"delay,omitempty" db:"delay"`                             // int(11) NOT NULL DEFAULT 0,
	Augdistiller        int            `yaml:"augdistiller,omitempty" db:"augdistiller"`               // int(11) NOT NULL DEFAULT 0,
	Dotshielding        int            `yaml:"dotshielding,omitempty" db:"dotshielding"`               // int(11) NOT NULL DEFAULT 0,
	Dr                  int            `yaml:"dr,omitempty" db:"dr"`                                   // int(11) NOT NULL DEFAULT 0,
	Clicktype           int            `yaml:"clicktype,omitempty" db:"clicktype"`                     // int(11) NOT NULL DEFAULT 0,
	Clicklevel2         int            `yaml:"clicklevel2,omitempty" db:"clicklevel2"`                 // int(11) NOT NULL DEFAULT 0,
	Elemdmgtype         int            `yaml:"elemdmgtype,omitempty" db:"elemdmgtype"`                 // int(11) NOT NULL DEFAULT 0,
	Elemdmgamt          int            `yaml:"elemdmgamt,omitempty" db:"elemdmgamt"`                   // int(11) NOT NULL DEFAULT 0,
	Endur               int            `yaml:"endur,omitempty" db:"endur"`                             // int(11) NOT NULL DEFAULT 0,
	Factionamt1         int            `yaml:"factionamt1,omitempty" db:"factionamt1"`                 // int(11) NOT NULL DEFAULT 0,
	Factionamt2         int            `yaml:"factionamt2,omitempty" db:"factionamt2"`                 // int(11) NOT NULL DEFAULT 0,
	Factionamt3         int            `yaml:"factionamt3,omitempty" db:"factionamt3"`                 // int(11) NOT NULL DEFAULT 0,
	Factionamt4         int            `yaml:"factionamt4,omitempty" db:"factionamt4"`                 // int(11) NOT NULL DEFAULT 0,
	Factionmod1         int            `yaml:"factionmod1,omitempty" db:"factionmod1"`                 // int(11) NOT NULL DEFAULT 0,
	Factionmod2         int            `yaml:"factionmod2,omitempty" db:"factionmod2"`                 // int(11) NOT NULL DEFAULT 0,
	Factionmod3         int            `yaml:"factionmod3,omitempty" db:"factionmod3"`                 // int(11) NOT NULL DEFAULT 0,
	Factionmod4         int            `yaml:"factionmod4,omitempty" db:"factionmod4"`                 // int(11) NOT NULL DEFAULT 0,
	Filename            string         `yaml:"filename,omitempty" db:"filename"`                       // varchar(32) NOT NULL DEFAULT '',
	Focuseffect         int            `yaml:"focuseffect,omitempty" db:"focuseffect"`                 // int(11) NOT NULL DEFAULT 0,
	Fr                  int            `yaml:"fr,omitempty" db:"fr"`                                   // int(11) NOT NULL DEFAULT 0,
	Fvnodrop            int            `yaml:"fvnodrop,omitempty" db:"fvnodrop"`                       // int(11) NOT NULL DEFAULT 0,
	Haste               int            `yaml:"haste,omitempty" db:"haste"`                             // int(11) NOT NULL DEFAULT 0,
	Clicklevel          int            `yaml:"clicklevel,omitempty" db:"clicklevel"`                   // int(11) NOT NULL DEFAULT 0,
	Hp                  int            `yaml:"hp,omitempty" db:"hp"`                                   // int(11) NOT NULL DEFAULT 0,
	Regen               int            `yaml:"regen,omitempty" db:"regen"`                             // int(11) NOT NULL DEFAULT 0,
	Icon                int            `yaml:"icon,omitempty" db:"icon"`                               // int(11) NOT NULL DEFAULT 0,
	Idfile              string         `yaml:"idfile,omitempty" db:"idfile"`                           // varchar(30) NOT NULL DEFAULT '',
	Itemclass           int            `yaml:"itemclass,omitempty" db:"itemclass"`                     // int(11) NOT NULL DEFAULT 0,
	Itemtype            int            `yaml:"itemtype,omitempty" db:"itemtype"`                       // int(11) NOT NULL DEFAULT 0,
	Ldonprice           int            `yaml:"ldonprice,omitempty" db:"ldonprice"`                     // int(11) NOT NULL DEFAULT 0,
	Ldontheme           int            `yaml:"ldontheme,omitempty" db:"ldontheme"`                     // int(11) NOT NULL DEFAULT 0,
	Ldonsold            int            `yaml:"ldonsold,omitempty" db:"ldonsold"`                       // int(11) NOT NULL DEFAULT 0,
	Light               int            `yaml:"light,omitempty" db:"light"`                             // int(11) NOT NULL DEFAULT 0,
	Lore                string         `yaml:"lore,omitempty" db:"lore"`                               // varchar(80) NOT NULL DEFAULT '',
	Loregroup           int            `yaml:"loregroup,omitempty" db:"loregroup"`                     // int(11) NOT NULL DEFAULT 0,
	Magic               int            `yaml:"magic,omitempty" db:"magic"`                             // int(11) NOT NULL DEFAULT 0,
	Mana                int            `yaml:"mana,omitempty" db:"mana"`                               // int(11) NOT NULL DEFAULT 0,
	Manaregen           int            `yaml:"manaregen,omitempty" db:"manaregen"`                     // int(11) NOT NULL DEFAULT 0,
	Enduranceregen      int            `yaml:"enduranceregen,omitempty" db:"enduranceregen"`           // int(11) NOT NULL DEFAULT 0,
	Material            int            `yaml:"material,omitempty" db:"material"`                       // int(11) NOT NULL DEFAULT 0,
	Herosforgemodel     int            `yaml:"herosforgemodel,omitempty" db:"herosforgemodel"`         // int(11) NOT NULL DEFAULT 0,
	Maxcharges          int            `yaml:"maxcharges,omitempty" db:"maxcharges"`                   // int(11) NOT NULL DEFAULT 0,
	Mr                  int            `yaml:"mr,omitempty" db:"mr"`                                   // int(11) NOT NULL DEFAULT 0,
	Nodrop              int            `yaml:"nodrop,omitempty" db:"nodrop"`                           // int(11) NOT NULL DEFAULT 0,
	Norent              int            `yaml:"norent,omitempty" db:"norent"`                           // int(11) NOT NULL DEFAULT 0,
	Pendingloreflag     int            `yaml:"pendingloreflag,omitempty" db:"pendingloreflag"`         // tinyint(3) unsigned NOT NULL DEFAULT 0,
	Pr                  int            `yaml:"pr,omitempty" db:"pr"`                                   // int(11) NOT NULL DEFAULT 0,
	Procrate            int            `yaml:"procrate,omitempty" db:"procrate"`                       // int(11) NOT NULL DEFAULT 0,
	Races               int            `yaml:"races,omitempty" db:"races"`                             // int(11) NOT NULL DEFAULT 0,
	Range               int            `yaml:"range,omitempty" db:"range"`                             // int(11) NOT NULL DEFAULT 0,
	Reclevel            int            `yaml:"reclevel,omitempty" db:"reclevel"`                       // int(11) NOT NULL DEFAULT 0,
	Recskill            int            `yaml:"recskill,omitempty" db:"recskill"`                       // int(11) NOT NULL DEFAULT 0,
	Reqlevel            int            `yaml:"reqlevel,omitempty" db:"reqlevel"`                       // int(11) NOT NULL DEFAULT 0,
	Sellrate            float32        `yaml:"sellrate,omitempty" db:"sellrate"`                       // float NOT NULL DEFAULT 0,
	Shielding           int            `yaml:"shielding,omitempty" db:"shielding"`                     // int(11) NOT NULL DEFAULT 0,
	Size                int            `yaml:"size,omitempty" db:"size"`                               // int(11) NOT NULL DEFAULT 0,
	Skillmodtype        int            `yaml:"skillmodtype,omitempty" db:"skillmodtype"`               // int(11) NOT NULL DEFAULT 0,
	Skillmodvalue       int            `yaml:"skillmodvalue,omitempty" db:"skillmodvalue"`             // int(11) NOT NULL DEFAULT 0,
	Slots               int            `yaml:"slots,omitempty" db:"slots"`                             // int(11) NOT NULL DEFAULT 0,
	Clickeffect         int            `yaml:"clickeffect,omitempty" db:"clickeffect"`                 // int(11) NOT NULL DEFAULT 0,
	Spellshield         int            `yaml:"spellshield,omitempty" db:"spellshield"`                 // int(11) NOT NULL DEFAULT 0,
	Strikethrough       int            `yaml:"strikethrough,omitempty" db:"strikethrough"`             // int(11) NOT NULL DEFAULT 0,
	Stunresist          int            `yaml:"stunresist,omitempty" db:"stunresist"`                   // int(11) NOT NULL DEFAULT 0,
	Summonedflag        int            `yaml:"summonedflag,omitempty" db:"summonedflag"`               // tinyint(3) unsigned NOT NULL DEFAULT 0,
	Tradeskills         int            `yaml:"tradeskills,omitempty" db:"tradeskills"`                 // int(11) NOT NULL DEFAULT 0,
	Favor               int            `yaml:"favor,omitempty" db:"favor"`                             // int(11) NOT NULL DEFAULT 0,
	Weight              int            `yaml:"weight,omitempty" db:"weight"`                           // int(11) NOT NULL DEFAULT 0,
	UNK012              int            `yaml:"unk012,omitempty" db:"UNK012"`                           // int(11) NOT NULL DEFAULT 0,
	UNK013              int            `yaml:"unk013,omitempty" db:"UNK013"`                           // int(11) NOT NULL DEFAULT 0,
	Benefitflag         int            `yaml:"benefitflag,omitempty" db:"benefitflag"`                 // int(11) NOT NULL DEFAULT 0,
	UNK054              int            `yaml:"unk054,omitempty" db:"UNK054"`                           // int(11) NOT NULL DEFAULT 0,
	UNK059              int            `yaml:"unk059,omitempty" db:"UNK059"`                           // int(11) NOT NULL DEFAULT 0,
	Booktype            int            `yaml:"booktype,omitempty" db:"booktype"`                       // int(11) NOT NULL DEFAULT 0,
	Recastdelay         int            `yaml:"recastdelay,omitempty" db:"recastdelay"`                 // int(11) NOT NULL DEFAULT 0,
	Recasttype          int            `yaml:"recasttype,omitempty" db:"recasttype"`                   // int(11) NOT NULL DEFAULT 0,
	Guildfavor          int            `yaml:"guildfavor,omitempty" db:"guildfavor"`                   // int(11) NOT NULL DEFAULT 0,
	UNK123              int            `yaml:"unk123,omitempty" db:"UNK123"`                           // int(11) NOT NULL DEFAULT 0,
	UNK124              int            `yaml:"unk124,omitempty" db:"UNK124"`                           // int(11) NOT NULL DEFAULT 0,
	Attuneable          int            `yaml:"attuneable,omitempty" db:"attuneable"`                   // int(11) NOT NULL DEFAULT 0,
	Nopet               int            `yaml:"nopet,omitempty" db:"nopet"`                             // int(11) NOT NULL DEFAULT 0,
	Updated             time.Time      `yaml:"updated,omitempty" db:"updated"`                         // datetime NOT NULL DEFAULT '0000-00-00 00:00:00',
	Comment             string         `yaml:"comment,omitempty" db:"comment"`                         // varchar(255) NOT NULL DEFAULT '',
	UNK127              int            `yaml:"unk127,omitempty" db:"UNK127"`                           // int(11) NOT NULL DEFAULT 0,
	Pointtype           int            `yaml:"pointtype,omitempty" db:"pointtype"`                     // int(11) NOT NULL DEFAULT 0,
	Potionbelt          int            `yaml:"potionbelt,omitempty" db:"potionbelt"`                   // int(11) NOT NULL DEFAULT 0,
	Potionbeltslots     int            `yaml:"potionbeltslots,omitempty" db:"potionbeltslots"`         // int(11) NOT NULL DEFAULT 0,
	Stacksize           int            `yaml:"stacksize,omitempty" db:"stacksize"`                     // int(11) NOT NULL DEFAULT 0,
	Notransfer          int            `yaml:"notransfer,omitempty" db:"notransfer"`                   // int(11) NOT NULL DEFAULT 0,
	Stackable           int            `yaml:"stackable,omitempty" db:"stackable"`                     // int(11) NOT NULL DEFAULT 0,
	UNK134              string         `yaml:"unk134,omitempty" db:"UNK134"`                           // varchar(255) NOT NULL DEFAULT '',
	UNK137              int            `yaml:"unk137,omitempty" db:"UNK137"`                           // int(11) NOT NULL DEFAULT 0,
	Proceffect          int            `yaml:"proceffect,omitempty" db:"proceffect"`                   // int(11) NOT NULL DEFAULT 0,
	Proctype            int            `yaml:"proctype,omitempty" db:"proctype"`                       // int(11) NOT NULL DEFAULT 0,
	Proclevel2          int            `yaml:"proclevel2,omitempty" db:"proclevel2"`                   // int(11) NOT NULL DEFAULT 0,
	Proclevel           int            `yaml:"proclevel,omitempty" db:"proclevel"`                     // int(11) NOT NULL DEFAULT 0,
	UNK142              int            `yaml:"unk142,omitempty" db:"UNK142"`                           // int(11) NOT NULL DEFAULT 0,
	Worneffect          int            `yaml:"worneffect,omitempty" db:"worneffect"`                   // int(11) NOT NULL DEFAULT 0,
	Worntype            int            `yaml:"worntype,omitempty" db:"worntype"`                       // int(11) NOT NULL DEFAULT 0,
	Wornlevel2          int            `yaml:"wornlevel2,omitempty" db:"wornlevel2"`                   // int(11) NOT NULL DEFAULT 0,
	Wornlevel           int            `yaml:"wornlevel,omitempty" db:"wornlevel"`                     // int(11) NOT NULL DEFAULT 0,
	UNK147              int            `yaml:"unk147,omitempty" db:"UNK147"`                           // int(11) NOT NULL DEFAULT 0,
	Focustype           int            `yaml:"focustype,omitempty" db:"focustype"`                     // int(11) NOT NULL DEFAULT 0,
	Focuslevel2         int            `yaml:"focuslevel2,omitempty" db:"focuslevel2"`                 // int(11) NOT NULL DEFAULT 0,
	Focuslevel          int            `yaml:"focuslevel,omitempty" db:"focuslevel"`                   // int(11) NOT NULL DEFAULT 0,
	UNK152              int            `yaml:"unk152,omitempty" db:"UNK152"`                           // int(11) NOT NULL DEFAULT 0,
	Scrolleffect        int            `yaml:"scrolleffect,omitempty" db:"scrolleffect"`               // int(11) NOT NULL DEFAULT 0,
	Scrolltype          int            `yaml:"scrolltype,omitempty" db:"scrolltype"`                   // int(11) NOT NULL DEFAULT 0,
	Scrolllevel2        int            `yaml:"scrolllevel2,omitempty" db:"scrolllevel2"`               // int(11) NOT NULL DEFAULT 0,
	Scrolllevel         int            `yaml:"scrolllevel,omitempty" db:"scrolllevel"`                 // int(11) NOT NULL DEFAULT 0,
	UNK157              int            `yaml:"unk157,omitempty" db:"UNK157"`                           // int(11) NOT NULL DEFAULT 0,
	Serialized          sql.NullTime   `yaml:"serialized,omitempty" db:"serialized"`                   // datetime DEFAULT NULL,
	Verified            sql.NullTime   `yaml:"verified,omitempty" db:"verified"`                       // datetime DEFAULT NULL,
	Serialization       sql.NullString `yaml:"serialization,omitempty" db:"serialization"`             // text DEFAULT NULL,
	Source              string         `yaml:"source,omitempty" db:"source"`                           // varchar(20) NOT NULL DEFAULT '',
	UNK033              int            `yaml:"unk033,omitempty" db:"UNK033"`                           // int(11) NOT NULL DEFAULT 0,
	Lorefile            string         `yaml:"lorefile,omitempty" db:"lorefile"`                       // varchar(32) NOT NULL DEFAULT '',
	UNK014              int            `yaml:"unk014,omitempty" db:"UNK014"`                           // int(11) NOT NULL DEFAULT 0,
	Svcorruption        int            `yaml:"svcorruption,omitempty" db:"svcorruption"`               // int(11) NOT NULL DEFAULT 0,
	Skillmodmax         int            `yaml:"skillmodmax,omitempty" db:"skillmodmax"`                 // int(11) NOT NULL DEFAULT 0,
	UNK060              int            `yaml:"unk060,omitempty" db:"UNK060"`                           // int(11) NOT NULL DEFAULT 0,
	Augslot1unk2        int            `yaml:"augslot1unk2,omitempty" db:"augslot1unk2"`               // int(11) NOT NULL DEFAULT 0,
	Augslot2unk2        int            `yaml:"augslot2unk2,omitempty" db:"augslot2unk2"`               // int(11) NOT NULL DEFAULT 0,
	Augslot3unk2        int            `yaml:"augslot3unk2,omitempty" db:"augslot3unk2"`               // int(11) NOT NULL DEFAULT 0,
	Augslot4unk2        int            `yaml:"augslot4unk2,omitempty" db:"augslot4unk2"`               // int(11) NOT NULL DEFAULT 0,
	Augslot5unk2        int            `yaml:"augslot5unk2,omitempty" db:"augslot5unk2"`               // int(11) NOT NULL DEFAULT 0,
	Augslot6unk2        int            `yaml:"augslot6unk2,omitempty" db:"augslot6unk2"`               // int(11) NOT NULL DEFAULT 0,
	UNK120              int            `yaml:"unk120,omitempty" db:"UNK120"`                           // int(11) NOT NULL DEFAULT 0,
	UNK121              int            `yaml:"unk121,omitempty" db:"UNK121"`                           // int(11) NOT NULL DEFAULT 0,
	Questitemflag       int            `yaml:"questitemflag,omitempty" db:"questitemflag"`             // int(11) NOT NULL DEFAULT 0,
	UNK132              sql.NullString `yaml:"unk132,omitempty" db:"UNK132"`                           // text CHARACTER SET utf8 DEFAULT NULL,
	Clickunk5           int            `yaml:"clickunk5,omitempty" db:"clickunk5"`                     // int(11) NOT NULL DEFAULT 0,
	Clickunk6           string         `yaml:"clickunk6,omitempty" db:"clickunk6"`                     // varchar(32) NOT NULL DEFAULT '',
	Clickunk7           int            `yaml:"clickunk7,omitempty" db:"clickunk7"`                     // int(11) NOT NULL DEFAULT 0,
	Procunk1            int            `yaml:"procunk1,omitempty" db:"procunk1"`                       // int(11) NOT NULL DEFAULT 0,
	Procunk2            int            `yaml:"procunk2,omitempty" db:"procunk2"`                       // int(11) NOT NULL DEFAULT 0,
	Procunk3            int            `yaml:"procunk3,omitempty" db:"procunk3"`                       // int(11) NOT NULL DEFAULT 0,
	Procunk4            int            `yaml:"procunk4,omitempty" db:"procunk4"`                       // int(11) NOT NULL DEFAULT 0,
	Procunk6            string         `yaml:"procunk6,omitempty" db:"procunk6"`                       // varchar(32) NOT NULL DEFAULT '',
	Procunk7            int            `yaml:"procunk7,omitempty" db:"procunk7"`                       // int(11) NOT NULL DEFAULT 0,
	Wornunk1            int            `yaml:"wornunk1,omitempty" db:"wornunk1"`                       // int(11) NOT NULL DEFAULT 0,
	Wornunk2            int            `yaml:"wornunk2,omitempty" db:"wornunk2"`                       // int(11) NOT NULL DEFAULT 0,
	Wornunk3            int            `yaml:"wornunk3,omitempty" db:"wornunk3"`                       // int(11) NOT NULL DEFAULT 0,
	Wornunk4            int            `yaml:"wornunk4,omitempty" db:"wornunk4"`                       // int(11) NOT NULL DEFAULT 0,
	Wornunk5            int            `yaml:"wornunk5,omitempty" db:"wornunk5"`                       // int(11) NOT NULL DEFAULT 0,
	Wornunk6            string         `yaml:"wornunk6,omitempty" db:"wornunk6"`                       // varchar(32) NOT NULL DEFAULT '',
	Wornunk7            int            `yaml:"wornunk7,omitempty" db:"wornunk7"`                       // int(11) NOT NULL DEFAULT 0,
	Focusunk1           int            `yaml:"focusunk1,omitempty" db:"focusunk1"`                     // int(11) NOT NULL DEFAULT 0,
	Focusunk2           int            `yaml:"focusunk2,omitempty" db:"focusunk2"`                     // int(11) NOT NULL DEFAULT 0,
	Focusunk3           int            `yaml:"focusunk3,omitempty" db:"focusunk3"`                     // int(11) NOT NULL DEFAULT 0,
	Focusunk4           int            `yaml:"focusunk4,omitempty" db:"focusunk4"`                     // int(11) NOT NULL DEFAULT 0,
	Focusunk5           int            `yaml:"focusunk5,omitempty" db:"focusunk5"`                     // int(11) NOT NULL DEFAULT 0,
	Focusunk6           string         `yaml:"focusunk6,omitempty" db:"focusunk6"`                     // varchar(32) NOT NULL DEFAULT '',
	Focusunk7           int            `yaml:"focusunk7,omitempty" db:"focusunk7"`                     // int(11) NOT NULL DEFAULT 0,
	Scrollunk1          int            `yaml:"scrollunk1,omitempty" db:"scrollunk1"`                   // int(11) NOT NULL DEFAULT 0,
	Scrollunk2          int            `yaml:"scrollunk2,omitempty" db:"scrollunk2"`                   // int(11) NOT NULL DEFAULT 0,
	Scrollunk3          int            `yaml:"scrollunk3,omitempty" db:"scrollunk3"`                   // int(11) NOT NULL DEFAULT 0,
	Scrollunk4          int            `yaml:"scrollunk4,omitempty" db:"scrollunk4"`                   // int(11) NOT NULL DEFAULT 0,
	Scrollunk5          int            `yaml:"scrollunk5,omitempty" db:"scrollunk5"`                   // int(11) NOT NULL DEFAULT 0,
	Scrollunk6          string         `yaml:"scrollunk6,omitempty" db:"scrollunk6"`                   // varchar(32) NOT NULL DEFAULT '',
	Scrollunk7          int            `yaml:"scrollunk7,omitempty" db:"scrollunk7"`                   // int(11) NOT NULL DEFAULT 0,
	UNK193              int            `yaml:"unk193,omitempty" db:"UNK193"`                           // int(11) NOT NULL DEFAULT 0,
	Purity              int            `yaml:"purity,omitempty" db:"purity"`                           // int(11) NOT NULL DEFAULT 0,
	Evoitem             int            `yaml:"evoitem,omitempty" db:"evoitem"`                         // int(11) NOT NULL DEFAULT 0,
	Evoid               int            `yaml:"evoid,omitempty" db:"evoid"`                             // int(11) NOT NULL DEFAULT 0,
	Evolvinglevel       int            `yaml:"evolvinglevel,omitempty" db:"evolvinglevel"`             // int(11) NOT NULL DEFAULT 0,
	Evomax              int            `yaml:"evomax,omitempty" db:"evomax"`                           // int(11) NOT NULL DEFAULT 0,
	Clickname           string         `yaml:"clickname,omitempty" db:"clickname"`                     // varchar(64) NOT NULL DEFAULT '',
	Procname            string         `yaml:"procname,omitempty" db:"procname"`                       // varchar(64) NOT NULL DEFAULT '',
	Wornname            string         `yaml:"wornname,omitempty" db:"wornname"`                       // varchar(64) NOT NULL DEFAULT '',
	Focusname           string         `yaml:"focusname,omitempty" db:"focusname"`                     // varchar(64) NOT NULL DEFAULT '',
	Scrollname          string         `yaml:"scrollname,omitempty" db:"scrollname"`                   // varchar(64) NOT NULL DEFAULT '',
	Dsmitigation        int            `yaml:"dsmitigation,omitempty" db:"dsmitigation"`               // smallint(6) NOT NULL DEFAULT 0,
	HeroicStr           int            `yaml:"heroic_str,omitempty" db:"heroic_str"`                   // smallint(6) NOT NULL DEFAULT 0,
	HeroicInt           int            `yaml:"heroic_int,omitempty" db:"heroic_int"`                   // smallint(6) NOT NULL DEFAULT 0,
	HeroicWis           int            `yaml:"heroic_wis,omitempty" db:"heroic_wis"`                   // smallint(6) NOT NULL DEFAULT 0,
	HeroicAgi           int            `yaml:"heroic_agi,omitempty" db:"heroic_agi"`                   // smallint(6) NOT NULL DEFAULT 0,
	HeroicDex           int            `yaml:"heroic_dex,omitempty" db:"heroic_dex"`                   // smallint(6) NOT NULL DEFAULT 0,
	HeroicSta           int            `yaml:"heroic_sta,omitempty" db:"heroic_sta"`                   // smallint(6) NOT NULL DEFAULT 0,
	HeroicCha           int            `yaml:"heroic_cha,omitempty" db:"heroic_cha"`                   // smallint(6) NOT NULL DEFAULT 0,
	HeroicPr            int            `yaml:"heroic_pr,omitempty" db:"heroic_pr"`                     // smallint(6) NOT NULL DEFAULT 0,
	HeroicDr            int            `yaml:"heroic_dr,omitempty" db:"heroic_dr"`                     // smallint(6) NOT NULL DEFAULT 0,
	HeroicFr            int            `yaml:"heroic_fr,omitempty" db:"heroic_fr"`                     // smallint(6) NOT NULL DEFAULT 0,
	HeroicCr            int            `yaml:"heroic_cr,omitempty" db:"heroic_cr"`                     // smallint(6) NOT NULL DEFAULT 0,
	HeroicMr            int            `yaml:"heroic_mr,omitempty" db:"heroic_mr"`                     // smallint(6) NOT NULL DEFAULT 0,
	HeroicSvcorrup      int            `yaml:"heroic_svcorrup,omitempty" db:"heroic_svcorrup"`         // smallint(6) NOT NULL DEFAULT 0,
	Healamt             int            `yaml:"healamt,omitempty" db:"healamt"`                         // smallint(6) NOT NULL DEFAULT 0,
	Spelldmg            int            `yaml:"spelldmg,omitempty" db:"spelldmg"`                       // smallint(6) NOT NULL DEFAULT 0,
	Clairvoyance        int            `yaml:"clairvoyance,omitempty" db:"clairvoyance"`               // smallint(6) NOT NULL DEFAULT 0,
	Backstabdmg         int            `yaml:"backstabdmg,omitempty" db:"backstabdmg"`                 // smallint(6) NOT NULL DEFAULT 0,
	Created             string         `yaml:"created,omitempty" db:"created"`                         // varchar(64) NOT NULL DEFAULT '',
	Elitematerial       int            `yaml:"elitematerial,omitempty" db:"elitematerial"`             // smallint(6) NOT NULL DEFAULT 0,
	Ldonsellbackrate    int            `yaml:"ldonsellbackrate,omitempty" db:"ldonsellbackrate"`       // smallint(6) NOT NULL DEFAULT 0,
	Scriptfileid        int            `yaml:"scriptfileid,omitempty" db:"scriptfileid"`               // smallint(6) NOT NULL DEFAULT 0,
	Expendablearrow     int            `yaml:"expendablearrow,omitempty" db:"expendablearrow"`         // smallint(6) NOT NULL DEFAULT 0,
	Powersourcecapacity int            `yaml:"powersourcecapacity,omitempty" db:"powersourcecapacity"` // smallint(6) NOT NULL DEFAULT 0,
	Bardeffect          int            `yaml:"bardeffect,omitempty" db:"bardeffect"`                   // smallint(6) NOT NULL DEFAULT 0,
	Bardeffecttype      int            `yaml:"bardeffecttype,omitempty" db:"bardeffecttype"`           // smallint(6) NOT NULL DEFAULT 0,
	Bardlevel2          int            `yaml:"bardlevel2,omitempty" db:"bardlevel2"`                   // smallint(6) NOT NULL DEFAULT 0,
	Bardlevel           int            `yaml:"bardlevel,omitempty" db:"bardlevel"`                     // smallint(6) NOT NULL DEFAULT 0,
	Bardunk1            int            `yaml:"bardunk1,omitempty" db:"bardunk1"`                       // smallint(6) NOT NULL DEFAULT 0,
	Bardunk2            int            `yaml:"bardunk2,omitempty" db:"bardunk2"`                       // smallint(6) NOT NULL DEFAULT 0,
	Bardunk3            int            `yaml:"bardunk3,omitempty" db:"bardunk3"`                       // smallint(6) NOT NULL DEFAULT 0,
	Bardunk4            int            `yaml:"bardunk4,omitempty" db:"bardunk4"`                       // smallint(6) NOT NULL DEFAULT 0,
	Bardunk5            int            `yaml:"bardunk5,omitempty" db:"bardunk5"`                       // smallint(6) NOT NULL DEFAULT 0,
	Bardname            string         `yaml:"bardname,omitempty" db:"bardname"`                       // varchar(64) NOT NULL DEFAULT '',
	Bardunk7            int            `yaml:"bardunk7,omitempty" db:"bardunk7"`                       // smallint(6) NOT NULL DEFAULT 0,
	UNK214              int            `yaml:"unk214,omitempty" db:"UNK214"`                           // smallint(6) NOT NULL DEFAULT 0,
	Subtype             int            `yaml:"subtype,omitempty" db:"subtype"`                         // int(11) NOT NULL DEFAULT 0,
	UNK220              int            `yaml:"unk220,omitempty" db:"UNK220"`                           // int(11) NOT NULL DEFAULT 0,
	UNK221              int            `yaml:"unk221,omitempty" db:"UNK221"`                           // int(11) NOT NULL DEFAULT 0,
	Heirloom            int            `yaml:"heirloom,omitempty" db:"heirloom"`                       // int(11) NOT NULL DEFAULT 0,
	UNK223              int            `yaml:"unk223,omitempty" db:"UNK223"`                           // int(11) NOT NULL DEFAULT 0,
	UNK224              int            `yaml:"unk224,omitempty" db:"UNK224"`                           // int(11) NOT NULL DEFAULT 0,
	UNK225              int            `yaml:"unk225,omitempty" db:"UNK225"`                           // int(11) NOT NULL DEFAULT 0,
	UNK226              int            `yaml:"unk226,omitempty" db:"UNK226"`                           // int(11) NOT NULL DEFAULT 0,
	UNK227              int            `yaml:"unk227,omitempty" db:"UNK227"`                           // int(11) NOT NULL DEFAULT 0,
	UNK228              int            `yaml:"unk228,omitempty" db:"UNK228"`                           // int(11) NOT NULL DEFAULT 0,
	UNK229              int            `yaml:"unk229,omitempty" db:"UNK229"`                           // int(11) NOT NULL DEFAULT 0,
	UNK230              int            `yaml:"unk230,omitempty" db:"UNK230"`                           // int(11) NOT NULL DEFAULT 0,
	UNK231              int            `yaml:"unk231,omitempty" db:"UNK231"`                           // int(11) NOT NULL DEFAULT 0,
	UNK232              int            `yaml:"unk232,omitempty" db:"UNK232"`                           // int(11) NOT NULL DEFAULT 0,
	UNK233              int            `yaml:"unk233,omitempty" db:"UNK233"`                           // int(11) NOT NULL DEFAULT 0,
	UNK234              int            `yaml:"unk234,omitempty" db:"UNK234"`                           // int(11) NOT NULL DEFAULT 0,
	Placeable           int            `yaml:"placeable,omitempty" db:"placeable"`                     // int(11) NOT NULL DEFAULT 0,
	UNK236              int            `yaml:"unk236,omitempty" db:"UNK236"`                           // int(11) NOT NULL DEFAULT 0,
	UNK237              int            `yaml:"unk237,omitempty" db:"UNK237"`                           // int(11) NOT NULL DEFAULT 0,
	UNK238              int            `yaml:"unk238,omitempty" db:"UNK238"`                           // int(11) NOT NULL DEFAULT 0,
	UNK239              int            `yaml:"unk239,omitempty" db:"UNK239"`                           // int(11) NOT NULL DEFAULT 0,
	UNK240              int            `yaml:"unk240,omitempty" db:"UNK240"`                           // int(11) NOT NULL DEFAULT 0,
	UNK241              int            `yaml:"unk241,omitempty" db:"UNK241"`                           // int(11) NOT NULL DEFAULT 0,
	Epicitem            int            `yaml:"epicitem,omitempty" db:"epicitem"`                       // int(11) NOT NULL DEFAULT 0,
}

func (e *ItemYaml) sanitize() error {
	for _, item := range e.Items {
		if item.Name == "" {
			return fmt.Errorf("item name must not be empty")
		}

		for _, item := range e.Items {
			if item.Augslot1visible == -1 {
				item.Augslot1visible = 0
			}
			if item.Augslot1visible == 0 {
				item.Augslot1visible = 1
			}
			if item.Augslot2visible == -1 {
				item.Augslot2visible = 0
			}
			if item.Augslot2visible == 0 {
				item.Augslot2visible = 1
			}
			if item.Augslot3visible == -1 {
				item.Augslot3visible = 0
			}
			if item.Augslot3visible == 0 {
				item.Augslot3visible = 1
			}
			if item.Augslot4visible == -1 {
				item.Augslot4visible = 0
			}
			if item.Augslot4visible == 0 {
				item.Augslot4visible = 1
			}
			if item.Augslot5visible == -1 {
				item.Augslot5visible = 0
			}
			if item.Augslot5visible == 0 {
				item.Augslot5visible = 1
			}
			if item.Augslot6visible == -1 {
				item.Augslot6visible = 0
			}
			if item.Augslot6visible == 0 {
				item.Augslot6visible = 1
			}

			if item.Combateffects == "" {
				item.Combateffects = "0"
			}
			if item.Charmfileid == "" {
				item.Charmfileid = "0"
			}

			if item.Focuseffect == 0 {
				item.Focuseffect = -1
			}
			if item.Skillmodtype == 0 {
				item.Skillmodtype = -1
			}
			if item.Clickeffect == 0 {
				item.Clickeffect = -1
			}
			if item.Proceffect == 0 {
				item.Proceffect = -1
			}
			if item.Worneffect == 0 {
				item.Worneffect = -1
			}
			if item.Scrolleffect == 0 {
				item.Scrolleffect = -1
			}
			if item.UNK120 == 0 {
				item.UNK120 = -1
			}
			if item.Clickunk7 == 0 {
				item.Clickunk7 = -1
			}
			if item.Procunk7 == 0 {
				item.Procunk7 = -1
			}
			if item.Scrollunk7 == 0 {
				item.Scrollunk7 = -1
			}
			if item.Wornunk7 == 0 {
				item.Wornunk7 = -1
			}
			if item.Focusunk7 == 0 {
				item.Focusunk7 = -1
			}
			if item.Bardeffect == 0 {
				item.Bardeffect = -1
			}
			if item.Bardunk7 == 0 {
				item.Bardunk7 = -1
			}
			if item.UNK231 == 0 {
				item.UNK231 = -1
			}
			if item.UNK233 == 0 {
				item.UNK233 = -256
			}
			if item.UNK234 == 0 {
				item.UNK234 = 255
			}

		}
	}
	return nil
}

func (e *ItemYaml) omitEmpty() error {
	for _, item := range e.Items {
		if item.Augslot1visible == -1 {
			item.Augslot1visible = 0
		}
		if item.Augslot1visible == 0 {
			item.Augslot1visible = 1
		}
		if item.Augslot2visible == -1 {
			item.Augslot2visible = 0
		}
		if item.Augslot2visible == 0 {
			item.Augslot2visible = 1
		}
		if item.Augslot3visible == -1 {
			item.Augslot3visible = 0
		}
		if item.Augslot3visible == 0 {
			item.Augslot3visible = 1
		}
		if item.Augslot4visible == -1 {
			item.Augslot4visible = 0
		}
		if item.Augslot4visible == 0 {
			item.Augslot4visible = 1
		}
		if item.Augslot5visible == -1 {
			item.Augslot5visible = 0
		}
		if item.Augslot5visible == 0 {
			item.Augslot5visible = 1
		}
		if item.Augslot6visible == -1 {
			item.Augslot6visible = 0
		}
		if item.Augslot6visible == 0 {
			item.Augslot6visible = 1
		}

		if item.Combateffects == "" {
			item.Combateffects = "0"
		}
		if item.Charmfileid == "" {
			item.Charmfileid = "0"
		}

		if item.Focuseffect == 0 {
			item.Focuseffect = -1
		}
		if item.Skillmodtype == 0 {
			item.Skillmodtype = -1
		}
		if item.Clickeffect == 0 {
			item.Clickeffect = -1
		}
		if item.Proceffect == 0 {
			item.Proceffect = -1
		}
		if item.Worneffect == 0 {
			item.Worneffect = -1
		}
		if item.Scrolleffect == 0 {
			item.Scrolleffect = -1
		}
		if item.UNK120 == 0 {
			item.UNK120 = -1
		}
		if item.Clickunk7 == 0 {
			item.Clickunk7 = -1
		}
		if item.Procunk7 == 0 {
			item.Procunk7 = -1
		}
		if item.Scrollunk7 == 0 {
			item.Scrollunk7 = -1
		}
		if item.Wornunk7 == 0 {
			item.Wornunk7 = -1
		}
		if item.Focusunk7 == 0 {
			item.Focusunk7 = -1
		}
		if item.Bardeffect == 0 {
			item.Bardeffect = -1
		}
		if item.Bardunk7 == 0 {
			item.Bardunk7 = -1
		}
		if item.UNK231 == 0 {
			item.UNK231 = -1
		}
		if item.UNK233 == 0 {
			item.UNK233 = -256
		}
		if item.UNK234 == 0 {
			item.UNK234 = 255
		}

	}

	return nil
}
