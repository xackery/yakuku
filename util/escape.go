package util

// EscapeSQL escapes a string for use in SQL statements.
func EscapeSQL(source string) string {
	var j int = 0
	if len(source) == 0 {
		return ""
	}
	tempStr := source[:]
	desc := make([]byte, len(tempStr)*2)
	for i := 0; i < len(tempStr); i++ {
		isFlagEscape := false
		var ecapeRune byte
		switch tempStr[i] {
		case '\r':
			isFlagEscape = true
			ecapeRune = '\r'
		case '\n':
			isFlagEscape = true
			ecapeRune = '\n'
		case '\\':
			isFlagEscape = true
			ecapeRune = '\\'
		case '\'':
			isFlagEscape = true
			ecapeRune = '\''
		case '"':
			isFlagEscape = true
			ecapeRune = '"'
		case '\032':
			isFlagEscape = true
			ecapeRune = 'Z'
		}
		if isFlagEscape {
			desc[j] = '\\'
			desc[j+1] = ecapeRune
			j = j + 2
		} else {
			desc[j] = tempStr[i]
			j = j + 1
		}
	}
	return string(desc[0:j])
}
