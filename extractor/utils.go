package extractor

import (
	"strconv"
	"strings"
)

func isDateTime(s string) bool {
	return len(s) >= 19 && s[2] == '/' && s[5] == '/' && s[10] == ' '
}

func isMoney(s string) bool {
	for _, r := range s {
		if (r < '0' || r > '9') && r != ',' && r != '.' {
			return false
		}
	}
	return true
}

func parseMoney(s string) float64 {
	s = strings.TrimSpace(s)
	s = strings.ReplaceAll(s, ",", "")
	v, err := strconv.ParseFloat(s, 64)
	if err != nil {
		return 0
	}
	return v
}

func isBalance(line string) bool {
	// Saldo BNI selalu mengandung koma atau desimal (.00)
	return isMoney(line) &&
		(strings.Contains(line, ",") || strings.Contains(line, "."))
}
