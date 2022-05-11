package parsinglogfiles

import "regexp"

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
	panic("Please implement the CountQuotedPasswords function")
}

func RemoveEndOfLineText(text string) string {
	panic("Please implement the RemoveEndOfLineText function")
}

func TagWithUserName(lines []string) []string {
	panic("Please implement the TagWithUserName function")
}
