package techpalace
import (
  "strings"
	"regexp"
)

// WelcomeMessage returns a welcome message for the customer.
func WelcomeMessage(customer string) string {
	return "Welcome to the Tech Palace, " + strings.ToUpper(customer)
}

// AddBorder adds a border to a welcome message.
func AddBorder(welcomeMsg string, numStarsPerLine int) string {
	stars  := strings.Repeat("*", numStarsPerLine)
	return stars + "\n" + welcomeMsg + "\n" + stars
}

// CleanupMessage cleans up an old marketing message.
func CleanupMessage(oldMsg string) string {
	pattern := "\\*" 
	oldMsg = regexp.MustCompile(pattern).ReplaceAllString(oldMsg, "")
	pattern2 := "\\n" 
	oldMsg = regexp.MustCompile(pattern2).ReplaceAllString(oldMsg, "")
	pattern3 := "^\\s*" 
	oldMsg = regexp.MustCompile(pattern3).ReplaceAllString(oldMsg, "")
	pattern4 := "\\s*$" 
	return regexp.MustCompile(pattern4).ReplaceAllString(oldMsg, "")
}
