package cmd

import (
	"errors"
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/cmczk/sunny/lib/config"
	"github.com/spf13/cobra"
)

var versionFilePath = filepath.Join(config.HomeDir, ".sunny", ".lua-version.global")

var sunnyPathCmd = &cobra.Command{
	Use:   "path",
	Short: "Add global Lua version to PATH",
	Run: func(cmd *cobra.Command, args []string) {
		file, err := os.Open(versionFilePath)
		if err != nil {
			if errors.Is(err, os.ErrNotExist) {
				if _, err := os.Create(versionFilePath); err != nil {
					log.Fatalf("cannot create .lua-version.global: %s\n", err.Error())
				}

				return
			}

			log.Fatalf("cannot find current Lua version: %s\n", err.Error())
		}
		defer file.Close()

		version, err := io.ReadAll(file)
		if err != nil {
			log.Fatalf("cannot read current Lua version")
		}

		fmt.Println(
			config.ProfileExportPathLuaStmt(
				strings.TrimSpace((string(version))),
			),
		)
	},
}
