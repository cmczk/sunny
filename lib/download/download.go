package download

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func Archive(url, dest string) error {
	resp, err := http.Get(url)
	if err != nil {
		return fmt.Errorf("cannot download archive: %s\n[ERROR] %w", url, err)
	}
	defer resp.Body.Close()

	out, err := os.Create(dest)
	if err != nil {
		return fmt.Errorf(
			"cannot create file to save downloaded archive: %s\n[ERROR] %w", dest, err,
		)
	}
	defer out.Close()

	if _, err := io.Copy(out, resp.Body); err != nil {
		return fmt.Errorf(
			"cannot write downloaded archive: %s\n [ERROR] %w", out.Name(), err,
		)
	}

	return nil
}
