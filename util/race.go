package util

func RaceIDToName(in int) string {
	switch in {
	case 1:
		return "human"
	case 2:
		return "barbarian"
	case 3:
		return "erudite"
	case 4:
		return "woodelf"
	case 5:
		return "highelf"
	case 6:
		return "darkelf"
	case 7:
		return "halfelf"
	case 8:
		return "dwarf"
	case 9:
		return "troll"
	case 10:
		return "ogre"
	case 11:
		return "halfling"
	case 12:
		return "gnome"
	case 128:
		return "iksar"
	case 130:
		return "vahshir"
	case 330:
		return "froglok"
	case 522:
		return "drakkin"
	default:
		return "unknown"
	}
}

func RaceNameToID(in string) int {
	switch in {
	case "human":
		return 1
	case "barbarian":
		return 2
	case "erudite":
		return 3
	case "woodelf":
		return 4
	case "highelf":
		return 5
	case "darkelf":
		return 6
	case "halfelf":
		return 7
	case "dwarf":
		return 8
	case "troll":
		return 9
	case "ogre":
		return 10
	case "halfling":
		return 11
	case "gnome":
		return 12
	case "iksar":
		return 128
	case "vahshir":
		return 130
	case "froglok":
		return 330
	case "drakkin":
		return 522
	default:
		return -1
	}
}
