package partyrobot
import (
	"fmt"
	"strings"
	"strconv"
)

// Welcome greets a person by name.
func Welcome(name string) string {
	return fmt.Sprintf("Welcome to my party, %s!", name)
}

// HappyBirthday wishes happy birthday to the birthday person and exclaims their age.
func HappyBirthday(name string, age int) string {
	return fmt.Sprintf("Happy birthday %s! You are now %d years old!", name, age)
}

// AssignTable assigns a table to each guest.
func AssignTable(name string, table int, neighbor, direction string, distance float64) string {
	var b strings.Builder

	workingWord := strconv.Itoa(table)
	length := len(workingWord)
	finalTable := strings.Repeat("0", (3 - length)) + workingWord

	b.WriteString(fmt.Sprintf("Welcome to my party, %s!\n", name))
	b.WriteString(fmt.Sprintf("You have been assigned to table %s. Your table is %s, exactly %.1f meters from here.\n", finalTable, direction, distance))
	b.WriteString(fmt.Sprintf("You will be sitting next to %s.", neighbor))

	return b.String()
}
