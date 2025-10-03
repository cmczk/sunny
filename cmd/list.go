package cmd

import (
	"fmt"
	"log"

	"github.com/cmczk/sunny/lib/lua"
	"github.com/spf13/cobra"
)

var listLuaCmd = &cobra.Command{
	Use:   "list [version]",
	Short: "List installed versions of Lua",
	Args:  cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		luaVersions := lua.MustInstalledLuaVersions()
		if len(luaVersions) == 0 {
			log.Fatalln(`No Lua versions installed.
To install Lua version use sunny install [version].`)
			return
		}

		for _, v := range luaVersions {
			fmt.Println(v)
		}
	},
}
