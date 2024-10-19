package utils

import (
	"regexp"
	"strings"
)

// SanitizeName sanitizes a name by removing any character that is not a letter or underscore.
func SanitizeName(name string) string {
	// Regular expression that selects only letters (uppercase and lowercase) and underscores
	reg := regexp.MustCompile(`[^a-zA-Z_]+`)

	// Replaces anything that is not a letter or underscore with a space
	sanitized := reg.ReplaceAllString(name, " ")

	// Splits the string into words and joins them with an underscore
	parts := strings.Fields(sanitized)
	return strings.Join(parts, "_")
}
