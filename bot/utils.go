package bot

import (
	"regexp"
)

func IsRegexMatch(expr, text string) bool {
	m, _ := regexp.Match(expr, []byte(text))
	return m
}

func GetRegexSubMatch(regex, text string) map[string]string {
	expr := regexp.MustCompile(regex)
	match := expr.FindStringSubmatch(text)
	result := make(map[string]string)
	if len(match) == 0 {
		return result
	}

	for i, name := range expr.SubexpNames() {
		if i != 0 && name != "" {
			result[name] = match[i]
		}
	}

	return result
}
