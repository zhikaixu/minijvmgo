package classpath

import (
	"errors"
	"strings"
)

type CompositeEntry struct {
	entries []Entry
}

func newCompositeEntry(pathList string) *CompositeEntry {
	entryList := make([]Entry, 0)
	for _, path := range strings.Split(pathList, pathListSeparator) {
		entry := newEntry(path)
		entryList = append(entryList, entry)
	}
	return &CompositeEntry{entries: entryList}
}

func (self *CompositeEntry) readClass(className string) ([]byte, Entry, error) {
	for _, entry := range self.entries {
		data, from, err := entry.readClass(className)
		if err == nil {
			return data, from, nil
		}
	}
	return nil, nil, errors.New("class not found: " + className)
}

func (self *CompositeEntry) String() string {
	strs := make([]string, len(self.entries))
	for i, entry := range self.entries {
		strs[i] = entry.String()
	}
	return strings.Join(strs, pathListSeparator)
}
