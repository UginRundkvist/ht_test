package hw10programoptimization

import (
	"bufio"
	"io"
	"strings"
)

type DomainStat map[string]int

func GetDomainStat(r io.Reader, domain string) (DomainStat, error) {
	stat := make(DomainStat)
	scanner := bufio.NewScanner(r)

	for scanner.Scan() {
		line := scanner.Text()
		email := extractEmail(line)

		if strings.HasSuffix(email, "."+domain) {
			stat[email]++
		}
	}
	return stat, nil
}

func extractEmail(jsonStr string) string {
	start := strings.Index(jsonStr, `"Email":`)
	start += len(`"Email":"`)

	end := strings.Index(jsonStr[start:], `"`)
	email := jsonStr[start : start+end]

	domain := strings.Split(email, "@")

	return strings.ToLower(domain[1])
}
