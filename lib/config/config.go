package config

import (
	"bufio"
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

var VersionFilePath = filepath.Join(HomeDir, ".sunny", ".lua-version.global")

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
	return fmt.Sprintf("export PATH=\"$HOME/.sunny/lua/%s/bin:$PATH\"", version)
}

func WriteGlobalLuaVersion(version string) error {
	file, err := os.OpenFile(VersionFilePath, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		return fmt.Errorf("cannot open .lua-version.global: %w", err)
	}
	defer file.Close()

	writer := bufio.NewWriter(file)
	if _, err := writer.WriteString(version); err != nil {
		return fmt.Errorf("cannot write global Lua version to .lua-version.global: %w", err)
	}

	if err := writer.Flush(); err != nil {
		return fmt.Errorf("cannot write global Lua version to .lua-version.global: %w", err)
	}

	return nil
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
