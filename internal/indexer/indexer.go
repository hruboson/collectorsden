package indexer

import (
	"fmt"
	"path/filepath"
	"strings"
	"syscall"

	"github.com/karrick/godirwalk"
	"hrubos.dev/collectorsden/internal/logger"
)

func TreeDir(directory string, maxDepth int) error {
	baseDepth := len(strings.Split(filepath.Clean(directory), string(filepath.Separator)))

	err := godirwalk.Walk(directory, &godirwalk.Options{
		Callback: func(osPathname string, de *godirwalk.Dirent) error {
			if de.IsDir() {
				currentDepth := len(strings.Split(filepath.Clean(osPathname), string(filepath.Separator)))
				hiddenFolder, err := isHidden(osPathname)
				if err != nil || (currentDepth - baseDepth > maxDepth || hiddenFolder) {
					return godirwalk.SkipThis
				}
			}

			entry := fmt.Sprintf("%s %s", de.ModeType(), osPathname)
			logger.Log(entry, logger.CatIndexer)
			return nil
		},
	})
	
	return err
}

func isHidden(path string) (bool, error) {
    p, err := syscall.UTF16PtrFromString(path)
    if err != nil {
        return false, err
    }
    attrs, err := syscall.GetFileAttributes(p)
    if err != nil {
        return false, err
    }
    const FILE_ATTRIBUTE_HIDDEN = 0x2
    return attrs&FILE_ATTRIBUTE_HIDDEN != 0, nil
}
