package resistorcolor

var resistors = [...]string{
	"black",
	"brown",
	"red",
	"orange",
	"yellow",
	"green",
	"blue",
	"violet",
	"grey",
	"white",
}

// Colors return different band colors.
func Colors() []string {
	return resistors[:]
}

// ColorCode returns the resistance value of the given color.
func ColorCode(color string) int {
	for index, value := range resistors {
		if value == color {
			return index
		}
	}

	return -1
}
