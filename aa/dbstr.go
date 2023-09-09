package aa

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/xackery/yakuku/util"
)

type dbReader struct {
	SIDs              map[string]string
	lastLine          string
	lastSID           int
	lastCategory      int
	changedDBStrCount int
}

func (d *dbReader) line(scanner *bufio.Scanner, sid int, category int) string {

	line, ok := d.SIDs[fmt.Sprintf("%d^%d^", sid, category)]
	if ok {
		d.changedDBStrCount++
		return line
	}

	var err error
	//fmt.Println("current", sid, category, "| last", d.lastSID, d.lastCategory)
	if d.lastLine == "" || d.lastSID < sid || (d.lastSID == sid && d.lastCategory < category) {
		if !scanner.Scan() {
			return ""
		}

		d.lastLine = scanner.Text()
		//fmt.Println("grabbing new line")
		lineBreakdown := strings.Split(d.lastLine, "^")
		d.lastSID, err = strconv.Atoi(lineBreakdown[0])
		if err != nil {
			fmt.Printf("line atoi %s: %s\n", d.lastLine, err)
			os.Exit(1)
		}
		d.lastCategory, err = strconv.Atoi(lineBreakdown[1])
		if err != nil {
			fmt.Printf("line atoi %s: %s\n", d.lastLine, err)
			os.Exit(1)
		}
	}

	if d.lastSID > sid {
		return ""
	}
	if d.lastCategory > category {
		return ""
	}

	line = d.lastLine

	//fmt.Println("inserting line, grabbing next")
	if !scanner.Scan() {
		return line
	}
	d.lastLine = scanner.Text()
	lineBreakdown := strings.Split(d.lastLine, "^")
	d.lastSID, err = strconv.Atoi(lineBreakdown[0])
	if err != nil {
		fmt.Printf("line atoi %s: %s\n", d.lastLine, err)
		os.Exit(1)
	}
	d.lastCategory, err = strconv.Atoi(lineBreakdown[1])
	if err != nil {
		fmt.Printf("line atoi %s: %s\n", d.lastLine, err)
		os.Exit(1)
	}
	return line
}

func modifyDBStr(db *dbReader, aa *AAYaml) error {
	var err error

	err = util.PrepFile("dbstr_us", ".txt")
	if err != nil {
		return err
	}

	defer func() {
		os.Remove("dbstr_us_tmp.txt")
	}()

	r, err := os.Open("dbstr_us_tmp.txt")
	if err != nil {
		return err
	}
	defer r.Close()
	scanner := bufio.NewScanner(r)

	key := ""
	db.SIDs = map[string]string{}
	for _, skill := range aa.Skills {
		key = fmt.Sprintf("%d^1^", skill.NameSID)
		db.SIDs[key] = fmt.Sprintf("%s%s^0", key, skill.Name)

		for _, rank := range skill.Ranks {
			if rank.TitleSID != 0 {
				key = fmt.Sprintf("%d^1^", rank.TitleSID)
				db.SIDs[key] = fmt.Sprintf("%s%s^0", key, rank.Title)
			}

			key = fmt.Sprintf("%d^4^", rank.DescriptionSID)
			if db.SIDs[key] != "" {
				return fmt.Errorf("sid %s already used, description sid", key)
			}
			db.SIDs[key] = fmt.Sprintf("%s%s^0", key, rank.Description)
		}
	}

	w, err := os.Create("dbstr_us.txt")
	if err != nil {
		return err
	}
	defer w.Close()

	for sid := 0; sid < 1458120306; sid++ { //< 1458120306; sid++ {
		if sid > 10000 {
			if scanner.Scan() {
				_, err = w.WriteString(scanner.Text() + "\n")
				if err != nil {
					return fmt.Errorf("line %s bypass writeString: %w", scanner.Text(), err)
				}
			}
			continue
		}
		for category := 0; category < 40; category++ {
			line := db.line(scanner, sid, category)
			if line == "" {
				continue
			}

			_, err = w.WriteString(line + "\n")
			if err != nil {
				return fmt.Errorf("line %s writeString: %w", line, err)
			}
		}
	}

	return nil
}
