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

func AddLuaInstallationToProfile(version string) error {
	envVar := ProfileExportPathLuaStmt(version)

	profileCfgPath := ProfileConfigPath()
	file, err := os.OpenFile(profileCfgPath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return fmt.Errorf("cannot open %s: %w", profileCfgPath, err)
	}

	if _, err = file.WriteString(envVar); err != nil {
		return fmt.Errorf("cannot add .sunny to PATH")
	}

	return nil
}

func DeleteLuaInstallationFromProfile() error {
	lineToDelete := filepath.Join(".sunny", "lua")
	profileConfigPath := ProfileConfigPath()

	input, err := os.Open(profileConfigPath)
	if err != nil {
		return fmt.Errorf("cannot open profile config file: %s", profileConfigPath)
	}
	defer input.Close()

	tmpFilePath := profileConfigPath + ".tmp"
	output, err := os.Create(tmpFilePath)
	if err != nil {
		return fmt.Errorf("cannot create tmp file to change profile config")
	}
	defer output.Close()

	scanner := bufio.NewScanner(input)
	writer := bufio.NewWriter(output)

	for scanner.Scan() {
		line := scanner.Text()
		if strings.Contains(strings.TrimSpace(line), strings.TrimSpace(lineToDelete)) {
			continue
		}

		fmt.Fprintln(writer, line)
	}

	if err := scanner.Err(); err != nil {
		return nil
	}

	writer.Flush()

	if err := os.Rename(tmpFilePath, profileConfigPath); err != nil {
		return fmt.Errorf("cannot save new profile config")
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
