package util

import (
	"database/sql"
	"fmt"
	"reflect"
	"time"

	"github.com/fatih/structs"
)

// FieldParse
func FieldParse(field *structs.Field) string {
	switch field.Kind() {
	case reflect.String:
		return fmt.Sprintf("`%s` = '%s'", field.Tag("db"), escapeSQL(fmt.Sprintf("%s", field.Value())))
	case reflect.Int:
		return fmt.Sprintf("`%s` = %d", field.Tag("db"), field.Value())
	case reflect.Float64:
		return fmt.Sprintf("`%s` = %f", field.Tag("db"), field.Value())
	case reflect.Float32:
		return fmt.Sprintf("`%s` = %f", field.Tag("db"), field.Value())
	case reflect.Bool:
		return fmt.Sprintf("`%s` = %t", field.Tag("db"), field.Value())
	case reflect.Struct:
		switch val := field.Value().(type) {
		case time.Time:
			if field.Tag("db") == "updated" {
				return fmt.Sprintf("`%s` = NOW()", field.Tag("db"))
			} else {
				return fmt.Sprintf("`%s` = CAST('%s' as DATETIME)", field.Tag("db"), val.Format("2006-01-02 15:04:05"))
			}
		case sql.NullString:
			if val.Valid {
				return fmt.Sprintf("`%s` = '%s'", field.Tag("db"), escapeSQL(val.String))
			} else {
				return fmt.Sprintf("`%s` = NULL", field.Tag("db"))
			}
		case sql.NullTime:
			if val.Valid {
				return fmt.Sprintf("`%s` = CAST('%s' AS DATETIME)", field.Tag("db"), val.Time.Format("2006-01-02 15:04:05"))
			} else {
				return fmt.Sprintf("`%s` = NULL", field.Tag("db"))
			}
		default:
			return ""
		}
	}
	return ""
}

func escapeSQL(sql string) string {
	dest := make([]byte, 0, 2*len(sql))
	var escape byte
	for i := 0; i < len(sql); i++ {
		c := sql[i]

		escape = 0

		switch c {
		case 0: /* Must be escaped for 'mysql' */
			escape = '0'
		case '\n': /* Must be escaped for logs */
			escape = 'n'
		case '\r':
			escape = 'r'
		case '\\':
			escape = '\\'
		case '\'':
			escape = '\''
		case '"': /* Better safe than sorry */
			escape = '"'
		case '\032':
			escape = 'Z'
		}

		if escape != 0 {
			dest = append(dest, '\\', escape)
		} else {
			dest = append(dest, c)
		}
	}
	return string(dest)
}
