package parsinglogfiles

import (
	"fmt"
	"regexp"
)

func IsValidLine(text string) bool {

	logsToMatch := []*regexp.Regexp{
		regexp.MustCompile(`\[TRC\]`),
		regexp.MustCompile(`\[DBG\]`),
		regexp.MustCompile(`\[INF\]`),
		regexp.MustCompile(`\[WRN\]`),
		regexp.MustCompile(`\[ERR\]`),
		regexp.MustCompile(`\[FTL\]`)}

		for _,exp := range logsToMatch {
			result := exp.FindStringIndex(text)

			if(len(result) > 0 && result[0] == 0){
				return true
			}
		}

	return false
}

func SplitLogLine(text string) []string {
	result := regexp.MustCompile(`\<\W*[^a-z]|[^a-z]\>`).Split(text,-1)

	if len(result) > 0 {
		return result
	}

	return []string{text}
}

func CountQuotedPasswords(lines []string) int {
	passwordPattern := regexp.MustCompile(`(\"passWord\")|\".+(password)\"`)
	var count int

	for _,v := range lines {
		if(passwordPattern.MatchString(v)) {
			count++
		}
	}

	return count
}

func RemoveEndOfLineText(text string) string {
	lineToDelete := regexp.MustCompile(`end-of-line\d+`)

	return lineToDelete.ReplaceAllString(text,"")
}

func TagWithUserName(lines []string) []string {
	nameToTagWith := regexp.MustCompile(`(User\s+)(\w+)`)
	var finalResult []string

	for _,v := range lines {
		matches := nameToTagWith.FindStringSubmatch(v)
		var nameToPrefix string

		if(len(matches) > 0) {
			nameToPrefix =  fmt.Sprintf("[USR] %s %s", matches[2], v)
		} else {
			nameToPrefix = v
		}

		finalResult = append(finalResult,nameToPrefix)
	}

	return finalResult
}
