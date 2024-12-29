package utils

import "regexp"

// IsTemplateVariable returns true if the text is a template variable.
// The text must start with a dot and be a valid template.
func IsTemplateVariable(text string) bool {
	match, _ := regexp.MatchString(`(?m)^{{(\.\w+)+}}$`, text)
	return match
}

// TemplateVariable returns the variable name of the template.
func TemplateVariable(text string) string {
	if IsTemplateVariable(text) {
		return text[3 : len(text)-2]
	}
	return ""
}
