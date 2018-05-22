package password

import (
	"io/ioutil"
	"sort"
	"strings"
)

// CommonList holds a list of strings containing common passwords
type CommonList struct {
	list []string
}

// NewCommonList list returns a CommonList with a list of sorted strings from the contents
// of the supplied filename
func NewCommonList(filename string) (CommonList, error) {
	c := CommonList{}

	data, err := ioutil.ReadFile(filename)
	if err != nil {
		return c, err
	}

	// Split on newline
	c.list = strings.Split(string(data), "\n")

	// Sort strings. This uses quicksort
	// TODO: maybe try another soring algorithm for larger datasets
	sort.Strings(c.list)
	return c, nil
}

// Matches determines if the supplied string it a match to the list of common passwords
func (c *CommonList) Matches(input string) bool {
	// Can't search without the list being sorted, so just to be sure...
	if !sort.StringsAreSorted(c.list) {
		sort.Strings(c.list)
	}
	idx := sort.SearchStrings(c.list, input)
	return idx < len(c.list) && c.list[idx] == input
}
