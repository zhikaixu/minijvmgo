package classpath

import (
	"os"
	"path/filepath"
	"strings"
)

func newWildcardEntry(path string) *CompositeEntry {
	baseDir := path[:len(path)-1] // remove *
	entryList := make([]Entry, 5)
	walkFn := func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		// 返回SkipDir跳过子目录，通配符类路径不能递归匹配子目录下的JAR文件
		if info.IsDir() && path != baseDir {
			return filepath.SkipDir
		}
		if strings.HasSuffix(path, ".jar") || strings.HasSuffix(path, ".JAR") {
			jarEntry := newZipEntry(path)
			entryList = append(entryList, jarEntry)
		}
		return nil
	}
	filepath.Walk(path, walkFn)
	return &CompositeEntry{entries: entryList}
}
