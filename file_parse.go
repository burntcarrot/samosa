package samosa

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

// read the file in lines and load in to map
func readCoverReport(f string, covReport map[string]int) error {
	file, err := os.Open(f)
	if err != nil {
		return err
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	for scanner.Scan() {
		// split file name from line
		data := scanner.Text()
		b, af, found := cut(data, ":")
		if found {
			isCovered := string(af[len(af)-1])
			c, _ := strconv.Atoi(isCovered)
			if c > 0 {
				coveredLines := getCoverage(af)
				increment(covReport, b, coveredLines)
			}
		}
	}
	return nil

}

func getCoverage(stat string) int {
	lines := string(stat[len(stat)-3])
	i, _ := strconv.Atoi(lines)
	return i

}

func cut(s string, sep string) (string, string, bool) {
	if i := strings.Index(s, sep); i >= 0 {
		return s[:i], s[i+len(sep):], true
	}
	return s, "", false
}

func increment(m map[string]int, b string, i int) {
	v, ok := m[b]
	if ok {
		m[b] = v + i
		return
	}
	m[b] = 1
}
