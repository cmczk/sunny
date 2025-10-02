package build_lua

import (
	"fmt"
	"log"
	"os"
	"os/exec"
)

func Run(buildDir, installDir string) error {
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

	envVars := `
for dir in $HOME/.sunny/lua/5.*/bin; do
  if [ -d "$dir" ]; then
    export PATH="$dir:$PATH"
  fi
done`

	file, err := os.OpenFile("/home/cmaczok/.bashrc", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return fmt.Errorf("cannot open .bashrc: %w", err)
	}

	if _, err = file.WriteString(envVars); err != nil {
		return fmt.Errorf("cannot add .sunny to PATH")
	}

	return nil
}
