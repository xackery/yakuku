package util

import (
	"fmt"
	"io"
	"os"
)

// PrepFile takes an actual file baseName.extension and copies it to baseName_tmp.extension, or takes baseName_original.extension and copies it to baseName_tmp.extension
func PrepFile(baseName string, extension string) error {
	tempFile := baseName + "_tmp" + extension
	originalFile := baseName + "_original" + extension
	actualFile := baseName + extension

	_, err := os.Stat(actualFile)
	if err == nil {
		err = os.Rename(actualFile, tempFile)
		if err != nil {
			return err
		}
		return nil
	}
	_, err = os.Stat(originalFile)
	if err != nil {
		return fmt.Errorf("please copy %s into this path, and rename it to %s", baseName+extension, originalFile)
	}

	err = CopyFile(originalFile, tempFile)
	return err
}

// CopyFile copies a file from src to dst
func CopyFile(src string, dst string) error {
	fi, err := os.Stat(src)
	if err != nil {
		return err
	}

	if !fi.Mode().IsRegular() {
		return fmt.Errorf("%s is not a regular file", src)
	}

	r, err := os.Open(src)
	if err != nil {
		return err
	}
	defer r.Close()

	w, err := os.Create(dst)
	if err != nil {
		return err
	}
	defer w.Close()
	_, err = io.Copy(w, r)
	if err != nil {
		return fmt.Errorf("copy: %w", err)
	}
	return nil
}
