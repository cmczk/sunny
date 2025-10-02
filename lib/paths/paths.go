package paths

import (
	"log"
	"os"
	"path/filepath"
	"strings"
)

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
