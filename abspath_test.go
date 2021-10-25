package abspath

import (
	"fmt"
	"os/user"
	"testing"
)

func TestGetAbsoluteDirPath(t *testing.T) {
	usr, _ := user.Current()

	paths := []struct {
		RelativePath    string
		ExpectedAbsPath string
	}{
		{
			RelativePath:    "~/.m2",
			ExpectedAbsPath: fmt.Sprintf("%s/.m2", usr.HomeDir),
		},
		{
			RelativePath:    ".",
			ExpectedAbsPath: fmt.Sprintf("%s/abs-path", usr.HomeDir),
		},
		{
			RelativePath:    "./go.mod",
			ExpectedAbsPath: fmt.Sprintf("%s/abs-path/go.mod", usr.HomeDir),
		},
		{
			RelativePath:    "./",
			ExpectedAbsPath: fmt.Sprintf("%s/abs-path", usr.HomeDir),
		},
		{
			RelativePath:    "~.m2",
			ExpectedAbsPath: fmt.Sprintf("%s/.m2", usr.HomeDir),
		},
	}

	for _, p := range paths {
		absp := GetAbsoluteDirPath(p.RelativePath)
		if absp != p.ExpectedAbsPath {
			t.Errorf("Expected: %s Got: %s", p.ExpectedAbsPath, absp)
			t.Fail()
		}
	}
}

func TestGetAbsoluteFilePath(t *testing.T) {
	usr, _ := user.Current()

	paths := []struct {
		ParentDir       string
		Filename        string
		ExpectedAbsPath string
	}{
		{
			ParentDir:       "~/.m2",
			Filename:        "settings.xml",
			ExpectedAbsPath: fmt.Sprintf("/%s/.m2/settings.xml", usr.HomeDir),
		},
		{
			ParentDir:       ".",
			Filename:        "go.mod",
			ExpectedAbsPath: fmt.Sprintf("%s/abs-path/go.mod", usr.HomeDir),
		},
		{
			ParentDir:       "./",
			Filename:        "go.mod",
			ExpectedAbsPath: fmt.Sprintf("%s/abs-path/go.mod", usr.HomeDir),
		},
		{
			ParentDir:       "~",
			Filename:        ".m2/settings.xml",
			ExpectedAbsPath: fmt.Sprintf("%s/.m2/settings.xml", usr.HomeDir),
		},
		{
			ParentDir:       "~.m2",
			Filename:        "settings.xml",
			ExpectedAbsPath: fmt.Sprintf("%s/.m2/settings.xml", usr.HomeDir),
		},
	}

	for _, p := range paths {
		absp := GetAbsoluteFilePath(p.ParentDir, p.Filename)
		if absp != p.ExpectedAbsPath {
			t.Errorf("Expected: %s Got: %s", p.ExpectedAbsPath, absp)
			t.Fail()
		}
	}
}
