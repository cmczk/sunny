package cmd

import (
	"log"
	"os"
	"path"

	"github.com/cmczk/sunny/lib/build_lua"
	"github.com/cmczk/sunny/lib/download"
	"github.com/cmczk/sunny/lib/gz"
	"github.com/cmczk/sunny/lib/config"
	"github.com/cmczk/sunny/lib/urls"
	"github.com/spf13/cobra"
)

var installLuaCmd = &cobra.Command{
	Use:   "install [version]",
	Short: "Install a specified version of Lua",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		version := args[0]
		url := urls.LuaURLByVersion(version)
		if url == "" {
			log.Fatalf("cannot find version: %s\n", version)
		}

		log.Printf("fetch lua archive from %s", url)

		dest := config.DownloadLuaArchivePath(path.Base(url))
		if err := download.Archive(url, dest); err != nil {
			log.Fatalln(err.Error())
		}

		log.Println("lua archive was fetched")

		homeDir := config.MustHomeDir()
		if err := gz.Unpack(dest, homeDir); err != nil {
			log.Fatalf("cannot unpack lua archive to %s", homeDir)
		}

		unpackedLua := config.LuaUnpackedDir(path.Base(url))

		if err := build_lua.Run(
			unpackedLua,
			config.InstallLuaDir(version),
			version,
		); err != nil {
			log.Fatalf("cannot build lua: %s", err.Error())
		}

		if err := os.Remove(dest); err != nil {
			log.Printf("cannot remove downloaded archive: %s\n", dest)
		}

		if err := os.RemoveAll(unpackedLua); err != nil {
			log.Println("cannot remove unpacked lua directory")
		}

		log.Printf("lua %s was installed!\n\nnext steps:\n\nsource ~/.bashrc\nlua -v", version)
	},
}
