package classpath

import (
	"os"
	"strings"
)

const pathListSeparator = string(os.PathListSeparator)

// define Entry interface: class directory items
type Entry interface {
	readClass(className string) ([]byte, Entry, error)
	String() string
}

// initialize Entry type depending on path
func newEntry(path string) Entry {
	if strings.Contains(path, pathListSeparator) {
		return newCompositeEntry(path)
	}
	if strings.HasSuffix(path, "*") {
		//return newWildcardEntry(path)
		return nil
	}
	if strings.HasSuffix(path, ".jar") || strings.HasSuffix(path, ".JAR") ||
		strings.HasSuffix(path, ".zip") || strings.HasSuffix(path, ".ZIP") {
		return newZipEntry(path)
	}
	return newDirEntry(path)
}
