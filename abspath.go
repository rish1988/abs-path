package abspath

import (
	"fmt"
	"os"
	"os/user"
	"strings"
)

const PathSeparator = string(os.PathSeparator)

func GetAbsoluteDirPath(relativePath string) string {
	return GetAbsoluteFilePath(relativePath, "")
}

func GetAbsoluteFilePath(parentDir string, fileName string) string {
	if strings.HasPrefix(parentDir, "~") {
		usr, _ := user.Current()

		if strings.HasPrefix(parentDir, "~"+PathSeparator) {
			parentDir = strings.TrimPrefix(parentDir, "~"+PathSeparator)
		}

		if strings.HasPrefix(parentDir, "~") {
			parentDir = strings.TrimPrefix(parentDir, "~")
		}

		if len(fileName) == 0 {
			return strings.Join([]string{
				usr.HomeDir,
				parentDir,
			}, PathSeparator)
		}

		if len(parentDir) == 0 {
			return strings.Join([]string{
				usr.HomeDir,
				fileName,
			}, PathSeparator)
		}

		return strings.Join([]string{
			usr.HomeDir,
			parentDir,
			fileName,
		}, PathSeparator)
	} else if strings.HasPrefix(fileName, PathSeparator) && len(parentDir) == 0 {
		return fileName
	} else if strings.HasPrefix(parentDir, PathSeparator) && len(fileName) == 0 {
		return strings.TrimSuffix(parentDir, PathSeparator)
	} else if strings.HasSuffix(parentDir, PathSeparator) && !strings.HasPrefix(parentDir, ".") {
		return parentDir + fileName
	} else if strings.HasPrefix(parentDir, ".") {
		if prefix, err := os.Getwd(); err != nil {
			fmt.Errorf("Could not get current working directory. Reason: %s ", err.Error())
			return ""
		} else {
			parentDir = strings.TrimPrefix(parentDir, ".")
			parentDir = strings.TrimPrefix(parentDir, PathSeparator)
			return GetAbsoluteFilePath(strings.Join([]string{prefix, parentDir}, PathSeparator), fileName)
		}
	} else {
		return strings.Join([]string{
			parentDir,
			fileName,
		}, PathSeparator)
	}
}
