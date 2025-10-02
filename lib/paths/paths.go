package paths

import (
	"fmt"
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

func ProfileConfigPath() string {
	return filepath.Join(MustHomeDir(), ".bashrc")
}

func ProfileExportPathLuaStmt(version string) string {
	return fmt.Sprintf("\nexport PATH=\"$HOME/.sunny/lua/%s/bin:$PATH\"\n", version)
}
