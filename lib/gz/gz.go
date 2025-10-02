package gz

import (
	"archive/tar"
	"compress/gzip"
	"errors"
	"fmt"
	"io"
	"os"
	"path/filepath"
)

func Unpack(pathToFile, destDir string) error {
	file, err := os.Open(pathToFile)
	if err != nil {
		return fmt.Errorf("cannot open file: %s", pathToFile)
	}
	defer file.Close()

	gzReader, err := gzip.NewReader(file)
	if err != nil {
		return fmt.Errorf("cannot process file: %w", err)
	}

	tarReader := tar.NewReader(gzReader)

	if _, err := os.Stat(destDir); errors.Is(err, os.ErrNotExist) {
		if err := os.MkdirAll(destDir, os.ModePerm); err != nil {
			return fmt.Errorf("cannot create directory for archive unpacking: %s", destDir)
		}
	}

	for {
		header, err := tarReader.Next()
		if err == io.EOF {
			break
		}
		if err != nil {
			return fmt.Errorf("cannot unpack archive: %w", err)
		}

		targetPath := filepath.Join(destDir, header.Name)

		switch header.Typeflag {
		case tar.TypeDir:
			if err := os.MkdirAll(targetPath, os.ModePerm); err != nil {
				return fmt.Errorf("cannot create directory: %w", err)
			}

		case tar.TypeReg:
			if err := os.MkdirAll(filepath.Dir(targetPath), os.ModePerm); err != nil {
				return fmt.Errorf("cannot create parent dir: %w", err)
			}

			outFile, err := os.Create(targetPath)
			if err != nil {
				return fmt.Errorf("cannot save unpacked file: %w", err)
			}

			if _, err := io.Copy(outFile, tarReader); err != nil {
				outFile.Close()
				return fmt.Errorf("cannot write file: %w", err)
			}
			outFile.Close()
		}
	}

	return nil
}
