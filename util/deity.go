package util

func DeityIDToName(in int) string {
	switch in {
	case 140:
		return "agnostic"
	case 396:
		return "agnostic2"
	case 201:
		return "bertoxxulous"
	case 202:
		return "brellserilis"
	case 203:
		return "cazicthule"
	case 204:
		return "erollsimarr"
	case 205:
		return "bristlebane"
	case 206:
		return "innoruuk"
	case 207:
		return "karana"
	case 208:
		return "mithanielmarr"
	case 209:
		return "prexus"
	case 210:
		return "quellious"
	case 211:
		return "ralloszek"
	case 212:
		return "rodcetnife"
	case 213:
		return "solusekro"
	case 214:
		return "thetribunal"
	case 215:
		return "tunare"
	case 216:
		return "veeshan"
	default:
		return "unknown"
	}
}

func DeityNameToID(in string) int {
	switch in {
	case "agnostic":
		return 140
	case "agnostic2":
		return 396
	case "bertoxxulous":
		return 201
	case "brellserilis":
		return 202
	case "cazicthule":
		return 203
	case "erollsimarr":
		return 204
	case "bristlebane":
		return 205
	case "innoruuk":
		return 206
	case "karana":
		return 207
	case "mithanielmarr":
		return 208
	case "prexus":
		return 209
	case "quellious":
		return 210
	case "ralloszek":
		return 211
	case "rodcetnife":
		return 212
	case "solusekro":
		return 213
	case "thetribunal":
		return 214
	case "tunare":
		return 215
	case "veeshan":
		return 216
	default:
		return -1
	}
}
