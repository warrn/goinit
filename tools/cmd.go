package tools

import (
	"errors"
	"os"
	"path"
	"strings"
)

var (
	NoDownloaderError = errors.New("no downloader available in path")
	NoSHAError        = errors.New("no sha256 hasher available in path")
	NoGitError        = errors.New("no git binary available in path")
	NoTarError        = errors.New("no tar binary available in path")
	NoMakeError       = errors.New("no make binary available in path")
)

func CheckGimmeDependencies() (bool, error) {
	if !(CheckBinaryExistsInPath("wget") || CheckBinaryExistsInPath("curl") || CheckBinaryExistsInPath("fetch")) {
		return false, NoDownloaderError
	}
	if !(CheckBinaryExistsInPath("sha256sum") || CheckBinaryExistsInPath("gsha256sum") ||
		CheckBinaryExistsInPath("shasum")) {
		return false, NoSHAError
	}
	if !(CheckBinaryExistsInPath("git")) {
		return false, NoGitError
	}
	if !(CheckBinaryExistsInPath("tar")) {
		return false, NoTarError
	}
	if !(CheckBinaryExistsInPath("make")) {
		return false, NoMakeError
	}
	return true, nil
}

func CheckBinaryExistsInPath(binary string) bool {
	pathEnvVar := os.Getenv("PATH")
	var delim string

	switch {
	case strings.Contains(pathEnvVar, ":"):
		delim = ":"
	case strings.Contains(pathEnvVar, ";"):
		delim = ";"
	case strings.Contains(pathEnvVar, " "):
		delim = " "
	default:
		delim = ""
	}

	var paths []string

	if delim != "" {
		paths = strings.Split(pathEnvVar, delim)
	} else {
		paths = []string{pathEnvVar}
	}

	for _, p := range paths {
		testPath := path.Join(p, binary)
		if _, err := os.Stat(testPath); err == nil {
			return true
		}
	}

	return false
}
