package password

import (
	"bufio"
	"io"
)

// CommonList holds a list of strings containing common passwords
type CommonList struct {
	mapping map[string]bool
}

// NewCommonList list returns a CommonList with a list of sorted strings from the contents
// of the supplied filename
func NewCommonList(r io.Reader) (CommonList, error) {
	c := CommonList{mapping: map[string]bool{}}

	reader := bufio.NewScanner(r)

	for reader.Scan() {
		if text := reader.Text(); len(text) > 0 {
			c.mapping[text] = true
		}
	}

	if err := reader.Err(); err != nil && err != io.EOF {
		return c, err
	}
	return c, nil
}

// Matches determines if the supplied string it a match to the list of common passwords
func (c *CommonList) Matches(input string) bool {
	_, ok := c.mapping[input]
	return ok
}
