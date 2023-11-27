package utils

import (
	"bufio"
	"log"
	"regexp"
	"strings"
)

func Import(input string, re string) []map[string]string {
	var rex = regexp.MustCompile(re)
	var results []map[string]string
	scanner := bufio.NewScanner(strings.NewReader(input))
	for scanner.Scan() {
		match := rex.FindStringSubmatch(scanner.Text())
		if len(match) == 0 {
			log.Println("import: regexp problem", re)
			break
		}
		result := make(map[string]string)
		for i, name := range rex.SubexpNames() {
			if i != 0 && name != "" {
				result[name] = match[i]
			}
		}
		results = append(results, result)
	}
	if err := scanner.Err(); err != nil {
		log.Println("import: scanner problem", err)
	}
	return results
}
