package cmd

import (
	"log"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"
)

var initLuaProjectCmd = &cobra.Command{
	Use:   "init",
	Short: "Init Lua project",
	Args:  cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		if err := os.Mkdir("src", os.ModePerm); err != nil {
			log.Fatalf("cannot init new Lua project: %s", err.Error())
		}

		mainFilePath := filepath.Join("src", "main.lua")

		file, err := os.Create(mainFilePath)
		if err != nil {
			log.Fatalf("cannot init new Lua project: %s", err.Error())
		}
		defer file.Close()

		mainText := []byte("local function main()\n    print(\"Hello, world!\")\nend\n\nmain()")
		if err := os.WriteFile(mainFilePath, mainText, 0644); err != nil {
			log.Fatalf("cannot init new Lua project: %s", err.Error())
		}

		readme, err := os.Create("README.md")
		if err != nil {
			log.Fatalf("cannot init new Lua project: %s", err.Error())
		}
		defer readme.Close()

		log.Println("Project created. Use `lua src/main.lua` or `sunny run` to run it.")
	},
}
