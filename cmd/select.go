package cmd

import (
	"log"
	"slices"

	"github.com/cmczk/sunny/lib/config"
	"github.com/cmczk/sunny/lib/lua"
	"github.com/spf13/cobra"
)

var selectLuaCmd = &cobra.Command{
	Use:   "select [version]",
	Short: "Select active version of Lua",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		version := args[0]

		luaVersions := lua.MustInstalledLuaVersions()
		if len(luaVersions) == 0 {
			log.Fatalln(`No Lua versions installed.
To install Lua version use sunny install [version].`)
		}

		if !slices.Contains(luaVersions, version) {
			log.Fatalf(`Lua %s is not installed.
To install it use sunny install %s
`, version, version)
		}

		if err := config.WriteGlobalLuaVersion(version); err != nil {
			log.Fatalln(err)
		}

		log.Printf("Lua %s selected. Type `source ~/.bashrc` and `lua -v` to see changes\n", version)
	},
}
