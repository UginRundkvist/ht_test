package hw10programoptimization

import (
	"bufio"
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

func getUsers(r io.Reader) (*bufio.Reader, error) {
	bufferedReader := bufio.NewReader(r)
	return bufferedReader, nil
}

func countDomains(u *bufio.Reader, domain string) (DomainStat, error) {
	stat := make(DomainStat)

	for {
		line, err := u.ReadString('\n')
		if err != nil {
			if err != io.EOF {
				panic(err)
			}
			if len(line) > 0 {
				stat = countChanges(stat, domain, line)
			}
			break
		}

		stat = countChanges(stat, domain, line)

	}
	return stat, nil
}

func countChanges(stat DomainStat, domain string, line string) DomainStat {
	name := stringChange(line)
	domainName := strings.Split(name, ".")
	if domainName[1] == domain {
		if _, exists := stat[name]; exists {
			stat[name]++
		} else {
			stat[name] = 1
		}
	}
	return stat
}

func stringChange(line string) string {
	line = strings.ToLower(line)
	//user := strings.Split(line, ",")
	//email := strings.Split(strings.Split(line, ",")[3], ":")
	//newEmail := strings.ReplaceAll(strings.Split(strings.Split(line, ",")[3], ":")[1], "\"", "")
	//emailName := strings.Split(strings.ReplaceAll(strings.Split(strings.Split(line, ",")[3], ":")[1], "\"", ""), "@")
	return strings.Split(strings.ReplaceAll(strings.Split(strings.Split(line, ",")[3], ":")[1], "\"", ""), "@")[1]
}
