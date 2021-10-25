package abspath

import (
	"fmt"
	"os"
	"os/user"
	"strings"
)

func GetAbsoluteDirPath(relativePath string) string {
	return GetAbsoluteFilePath(relativePath, "")
}

func GetAbsoluteFilePath(parentDir string, fileName string) string {
	if strings.HasPrefix(parentDir, "~") {
		usr, _ := user.Current()

		if strings.HasPrefix(parentDir, "~/") {
			parentDir = strings.TrimPrefix(parentDir, "~/")
		}

		if strings.HasPrefix(parentDir, "~") {
			parentDir = strings.TrimPrefix(parentDir, "~")
		}

		if len(fileName) == 0 {
			return strings.Join([]string{
				usr.HomeDir,
				parentDir,
			}, "/")
		}

		if len(parentDir) == 0 {
			return strings.Join([]string{
				usr.HomeDir,
				fileName,
			}, "/")
		}

		return strings.Join([]string{
			usr.HomeDir,
			parentDir,
			fileName,
		}, "/")
	} else if strings.HasPrefix(fileName, "/") && len(parentDir) == 0 {
		return fileName
	} else if strings.HasPrefix(parentDir, "/") && len(fileName) == 0 {
		return strings.TrimSuffix(parentDir, "/")
	} else if strings.HasSuffix(parentDir, "/") && !strings.HasPrefix(parentDir, ".") {
		return parentDir + fileName
	} else if strings.HasPrefix(parentDir, ".") {
		if prefix, err := os.Getwd(); err != nil {
			fmt.Errorf("Could not get current working directory. Reason: %s ", err.Error())
			return ""
		} else {
			parentDir = strings.TrimPrefix(parentDir, ".")
			parentDir = strings.TrimPrefix(parentDir, "/")
			return GetAbsoluteFilePath(strings.Join([]string{prefix, parentDir}, "/"), fileName)
		}
	} else {
		return strings.Join([]string{
			parentDir,
			fileName,
		}, "/")
	}
}
