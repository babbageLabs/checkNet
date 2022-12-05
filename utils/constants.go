package utils

import (
	"regexp"
	"strings"
)

const WinMatchExp = `(?m)^\s+Key\b.*$`
const WinReplaceExp = `^\s+Key\s+:\s+`

const SQLiteDbName = "./foo.db"

func MustGetWinMatchExpForKey(key string) *regexp.Regexp {
	pattern := strings.Replace(WinMatchExp, "Key", key, 1)
	return regexp.MustCompile(pattern)
}

func MustGetWinReplaceExpForKey(key string) *regexp.Regexp {
	pattern := strings.Replace(WinReplaceExp, "Key", key, 1)
	return regexp.MustCompile(pattern)
}
