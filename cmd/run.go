package cmd

import (
	"errors"
	"log"
	"os"
	"os/exec"
	"path/filepath"

	"github.com/spf13/cobra"
)

var (
	srcMain = filepath.Join("src", "main.lua")
	main    = "main.lua"
)

var runLuaProjectCmd = &cobra.Command{
	Use:   "run",
	Short: "Run Lua project",
	Run: func(cmd *cobra.Command, args []string) {
		if _, err := os.Stat(srcMain); errors.Is(err, os.ErrNotExist) {
			if _, err := os.Stat(main); errors.Is(err, os.ErrNotExist) {
				log.Fatalln("Lua project could not be found.")
			}

			c := exec.Command("lua", main)
			c.Stdin = os.Stdin
			c.Stdout = os.Stdout
			c.Stderr = os.Stderr

			if err := c.Run(); err != nil {
				log.Fatalf("[ERROR]: %s\n", err.Error())
			}
		} else {
			c := exec.Command("lua", srcMain)
			c.Stdin = os.Stdin
			c.Stdout = os.Stdout
			c.Stderr = os.Stderr

			if err := c.Run(); err != nil {
				log.Fatalf("[ERROR]: %s\n", err.Error())
			}

		}
	},
}
