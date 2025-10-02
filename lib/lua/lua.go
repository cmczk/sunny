package lua

import (
	"log"
	"os"
	"path/filepath"

	"github.com/cmczk/sunny/lib/paths"
)

var (
	sunnyLuaDir = filepath.Join(paths.HomeDir, ".sunny", "lua")
)

func MustInstalledLuaVersions() []string {
	entries, err := os.ReadDir(sunnyLuaDir)
	if err != nil {
		log.Fatalf("cannot read directory with Lua installations: %s\n", err.Error())
	}

	luaVersions := make([]string, 0, len(entries))
	for _, en := range entries {
		luaVersions = append(luaVersions, en.Name())
	}

	return luaVersions
}
