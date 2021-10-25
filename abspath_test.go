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
			RelativePath:    fmt.Sprintf("~%s.m2", PathSeparator),
			ExpectedAbsPath: fmt.Sprintf("%s%s.m2", usr.HomeDir, PathSeparator),
		},
		{
			RelativePath:    ".",
			ExpectedAbsPath: fmt.Sprintf("%s%sabs-path", usr.HomeDir, PathSeparator),
		},
		{
			RelativePath:    fmt.Sprintf(".%sgo.mod", PathSeparator),
			ExpectedAbsPath: fmt.Sprintf("%s%sabs-path%sgo.mod", usr.HomeDir, PathSeparator, PathSeparator),
		},
		{
			RelativePath:    fmt.Sprintf(".%s", PathSeparator),
			ExpectedAbsPath: fmt.Sprintf("%s%sabs-path", usr.HomeDir, PathSeparator),
		},
		{
			RelativePath:    "~.m2",
			ExpectedAbsPath: fmt.Sprintf("%s%s.m2", usr.HomeDir, PathSeparator),
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
			ParentDir:       fmt.Sprintf("~%s.m2", PathSeparator),
			Filename:        "settings.xml",
			ExpectedAbsPath: fmt.Sprintf("%s%s%s.m2%ssettings.xml", PathSeparator, usr.HomeDir, PathSeparator, PathSeparator),
		},
		{
			ParentDir:       ".",
			Filename:        "go.mod",
			ExpectedAbsPath: fmt.Sprintf("%s%sabs-path%sgo.mod", usr.HomeDir, PathSeparator, PathSeparator),
		},
		{
			ParentDir:       fmt.Sprintf(".%s", PathSeparator),
			Filename:        "go.mod",
			ExpectedAbsPath: fmt.Sprintf("%s%sabs-path%sgo.mod", usr.HomeDir, PathSeparator, PathSeparator),
		},
		{
			ParentDir:       "~",
			Filename:        fmt.Sprintf(".m2%ssettings.xml", PathSeparator),
			ExpectedAbsPath: fmt.Sprintf("%s%s.m2%ssettings.xml", usr.HomeDir, PathSeparator, PathSeparator),
		},
		{
			ParentDir:       "~.m2",
			Filename:        "settings.xml",
			ExpectedAbsPath: fmt.Sprintf("%s%s.m2%ssettings.xml", usr.HomeDir, PathSeparator, PathSeparator),
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
