package hw10programoptimization

import (
	"fmt"
	"io"
	"strings"
)

type DomainStat map[string]int

func GetDomainStat(r io.Reader, domain string) (DomainStat, error) {
	u, err := getUsers(r)
	if err != nil {
		return nil, fmt.Errorf("get users error: %w", err)
	}
	return countDomains(u, domain)
}

func getUsers(r io.Reader) ([]string, error) {
	content, err := io.ReadAll(r) // изменить
	//line, err := r.ReadString('\n')
	if err != nil {
		return nil, err
	}
	lines := strings.Split(string(content), "\n")
	return lines, nil
}

func countDomains(u []string, domain string) (DomainStat, error) {
	stat := make(DomainStat)
	for _, line := range u {
		name := stringChange(line)
		domainName := strings.Split(name, ".")
		if domainName[1] == domain {
			if _, exists := stat[name]; exists {
				stat[name]++
			} else {
				stat[name] = 1
			}
		}
	}
	return stat, nil
}

func stringChange(line string) string {
	line = strings.ToLower(line)
	user := strings.Split(line, ",")
	email := strings.Split(user[3], ":")
	newEmail := strings.ReplaceAll(email[1], "\"", "")
	emailName := strings.Split(newEmail, "@")
	return emailName[1]
}
