package paths

import (
	"errors"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"runtime"
	"strings"
)

const (
	linux   = "linux"
	darwin  = "darwin"
	windows = "windows"
)

const (
	bashrc = ".bashrc"
	zshrc  = ".zshrc"
)

var HomeDir = MustHomeDir()

var env = runtime.GOOS

func MustHomeDir() string {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		log.Fatalf("cannot find user's home directory: %s", err.Error())
	}

	return homeDir
}

func DownloadLuaArchivePath(archiveName string) string {
	return filepath.Join(MustHomeDir(), archiveName)
}

func LuaUnpackedDir(archiveName string) string {
	return filepath.Join(
		MustHomeDir(),
		strings.TrimSuffix(
			strings.TrimSuffix(archiveName, ".gz"),
			".tar",
		),
	)
}

func InstallLuaDir(version string) string {
	return filepath.Join(MustHomeDir(), ".sunny", "lua", version)
}

func ProfileConfigPath() string {
	switch env {
	case linux, darwin:
		return unixProfileConfig()
	case windows:
		return filepath.Join(MustHomeDir(), ".bashrc")
	default:
		return ""
	}
}

func ProfileExportPathLuaStmt(version string) string {
	return fmt.Sprintf("\nexport PATH=\"$HOME/.sunny/lua/%s/bin:$PATH\"\n", version)
}

func unixProfileConfig() string {
	if _, err := os.Stat(filepath.Join(HomeDir, bashrc)); !errors.Is(err, os.ErrNotExist) {
		return filepath.Join(HomeDir, bashrc)
	} else {
		log.Println(".bashrc not found")
	}

	if _, err := os.Stat(filepath.Join(HomeDir, zshrc)); !errors.Is(err, os.ErrNotExist) {
		return filepath.Join(HomeDir, bashrc)
	} else {
		log.Println(".zshrc not found")
	}

	return ""
}
