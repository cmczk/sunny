package urls

import (
	"fmt"
	"slices"
)

const urlTemplate = "https://lua.org/ftp/lua-%s.tar.gz"

var versions = []string{
	"5.4.8",
	"5.4.7",
	"5.4.6",
	"5.4.5",
	"5.4.4",
	"5.4.3",
	"5.4.2",
	"5.4.1",
	"5.4.0",
}

func LuaURLByVersion(version string) string {
	if slices.Contains(versions, version) {
		return fmt.Sprintf(urlTemplate, version)
	}

	return ""
}
