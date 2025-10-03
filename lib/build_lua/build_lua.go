package build_lua

import (
	"fmt"
	"log"
	"os"
	"os/exec"
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

	return nil
}
