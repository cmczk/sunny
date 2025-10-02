package build_lua

import (
	"fmt"
	"log"
	"os"
	"os/exec"

	"github.com/cmczk/sunny/lib/paths"
)

func Run(buildDir, installDir, version string) error {
	if err := os.Chdir(buildDir); err != nil {
		return fmt.Errorf("cannot change dir to %s\n[ERROR] %w", buildDir, err)
	}

	log.Printf("%s\n", buildDir)

	cmd := exec.Command("make", "all", "test")
	if err := cmd.Run(); err != nil {
		return fmt.Errorf("cannot run make all test: %w", err)
	}

	cmd = exec.Command("make", fmt.Sprintf("INSTALL_TOP=%s", installDir), "install")
	if err := cmd.Run(); err != nil {
		return fmt.Errorf("cannot run make install: %w", err)
	}

	// TODO: extract this logic to another package
	envVar := paths.ProfileExportPathLuaStmt(version)

	profileCfgPath := paths.ProfileConfigPath()
	file, err := os.OpenFile(profileCfgPath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return fmt.Errorf("cannot open %s: %w", profileCfgPath, err)
	}

	if _, err = file.WriteString(envVar); err != nil {
		return fmt.Errorf("cannot add .sunny to PATH")
	}

	return nil
}
