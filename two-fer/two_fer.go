// Package twofer determine what you will say as you give away the extra cookie.
package twofer

import "fmt"

// ShareWith returns "One for {name}, one for me." string.
func ShareWith(name string) string {
	if name == "" {
		name = "you"
	}
	return fmt.Sprintf("One for %s, one for me.", name)
}
