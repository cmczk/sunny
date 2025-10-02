/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"os"
	"path"
	"path/filepath"

	build_lua "github.com/cmczk/sunny/lib/build-lua"
	"github.com/cmczk/sunny/lib/download"
	"github.com/cmczk/sunny/lib/gz"
	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "sunny",
	Short: "A brief description of your application",
	Long: `A longer description that spans multiple lines and likely contains
examples and usage of using your application. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		homeDir, err := os.UserHomeDir()
		if err != nil {
			fmt.Printf("cannot find user's home directory: %s", err.Error())
			os.Exit(1)
		}

		url := "https://lua.org/ftp/lua-5.4.8.tar.gz"
		dest := filepath.Join(homeDir, path.Base(url))

		err = download.Archive(url, dest)
		if err != nil {
			fmt.Println(err.Error())
			os.Exit(1)
		}

		if err := gz.Unpack(dest, homeDir); err != nil {
			fmt.Println(err.Error())
			os.Exit(1)
		}

		if err := os.Remove(dest); err != nil {
			fmt.Println(err.Error())
			os.Exit(1)
		}

		if err := build_lua.Run(
			filepath.Join(homeDir, "lua-5.4.8"),
			filepath.Join(homeDir, ".sunny", "lua", "5.4"),
		); err != nil {
			fmt.Printf("%s\n", err.Error())
			os.Exit(1)
		}

		if err := os.RemoveAll(filepath.Join(homeDir, "lua-5.4.8")); err != nil {
			fmt.Println(err.Error())
			os.Exit(1)
		}
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.sunny.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
