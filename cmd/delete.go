package cmd

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/cmczk/sunny/lib/config"
	"github.com/spf13/cobra"
)

var deleteLuaCmd = &cobra.Command{
	Use:   "delete [version]",
	Short: "Delete a specified version of Lua",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		version := args[0]

		installLuaDir := config.InstallLuaDir(version)

		if _, err := os.Stat(installLuaDir); errors.Is(err, os.ErrNotExist) {
			log.Fatalf("lua %s is not installed", version)
		}

		if err := deleteLuaDir(installLuaDir); err != nil {
			log.Fatalf("cannot delete lua %s: %s", version, err.Error())
		}

		if err := deleteExportPath(config.ProfileConfigPath(), version); err != nil {
			log.Fatalf("%s", err.Error())
		}

		log.Printf("lua %s has been deleted", version)
	},
}

func deleteLuaDir(dirPath string) error {
	if err := os.RemoveAll(dirPath); err != nil {
		return fmt.Errorf("cannot delete lua installation: %w", err)
	}

	return nil
}

func deleteExportPath(profileConfigPath, version string) error {
	lineToDelete := config.ProfileExportPathLuaStmt(version)

	input, err := os.Open(profileConfigPath)
	if err != nil {
		return fmt.Errorf("cannot open profile config file: %s", profileConfigPath)
	}
	defer input.Close()

	tmpFilePath := profileConfigPath + ".tmp"
	output, err := os.Create(tmpFilePath)
	if err != nil {
		return fmt.Errorf("cannot create tmp file to change profile config")
	}
	defer output.Close()

	scanner := bufio.NewScanner(input)
	writer := bufio.NewWriter(output)

	for scanner.Scan() {
		line := scanner.Text()
		if strings.TrimSpace(line) == strings.TrimSpace(lineToDelete) {
			continue
		}

		fmt.Fprintln(writer, line)
	}

	if err := scanner.Err(); err != nil {
		return nil
	}

	writer.Flush()

	if err := os.Rename(tmpFilePath, profileConfigPath); err != nil {
		return fmt.Errorf("cannot save new profile config")
	}

	return nil
}
