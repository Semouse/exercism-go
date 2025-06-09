package sorting

import (
	"fmt"
	"strconv"
)

type NumberBox interface {
	Number() int
}

type FancyNumberBox interface {
	Value() string
}

type FancyNumber struct {
	n string
}

func (i FancyNumber) Value() string {
	return i.n
}

// DescribeNumber should return a string describing the number.
func DescribeNumber(f float64) string {
	return fmt.Sprintf("This is the number %.1f", f)
}

// DescribeNumberBox should return a string describing the NumberBox.
func DescribeNumberBox(nb NumberBox) string {
	i := nb.Number()
	value := float64(i)
	return fmt.Sprintf("This is a box containing the number %.1f", value)
}

// ExtractFancyNumber should return the integer value for a FancyNumber
// and 0 if any other FancyNumberBox is supplied.
func ExtractFancyNumber(fnb FancyNumberBox) int {
	if value, ok := fnb.(FancyNumber); ok {
		i, _ := strconv.Atoi(value.n)
		return i
	}

	return 0
}

// DescribeFancyNumberBox should return a string describing the FancyNumberBox.
func DescribeFancyNumberBox(fnb FancyNumberBox) string {
	number := ExtractFancyNumber(fnb)
	return fmt.Sprintf("This is a fancy box containing the number %.1f", float64(number))
}

// DescribeAnything should return a string describing whatever it contains.
func DescribeAnything(i interface{}) string {
	switch v := i.(type) {
	case int:
		return DescribeNumber(float64(v))
	case float64:
		return DescribeNumber(v)
	case NumberBox:
		return DescribeNumberBox(v)
	case FancyNumberBox:
		return DescribeFancyNumberBox(v)
	default:
		return fmt.Sprintf("Return to sender")
	}
}
