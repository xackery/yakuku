package util

func ClassIDToName(in int) string {
	switch in {
	case 1:
		return "warrior"
	case 2:
		return "cleric"
	case 3:
		return "paladin"
	case 4:
		return "ranger"
	case 5:
		return "shadowknight"
	case 6:
		return "druid"
	case 7:
		return "monk"
	case 8:
		return "bard"
	case 9:
		return "rogue"
	case 10:
		return "shaman"
	case 11:
		return "necromancer"
	case 12:
		return "wizard"
	case 13:
		return "magician"
	case 14:
		return "enchanter"
	case 15:
		return "beastlord"
	case 16:
		return "berserker"
	default:
		return "unknown"
	}
}

func ClassNameToID(in string) int {
	switch in {
	case "warrior":
		return 1
	case "cleric":
		return 2
	case "paladin":
		return 3
	case "ranger":
		return 4
	case "shadowknight":
		return 5
	case "druid":
		return 6
	case "monk":
		return 7
	case "bard":
		return 8
	case "rogue":
		return 9
	case "shaman":
		return 10
	case "necromancer":
		return 11
	case "wizard":
		return 12
	case "magician":
		return 13
	case "enchanter":
		return 14
	case "beastlord":
		return 15
	case "berserker":
		return 16
	default:
		return -1
	}
}
