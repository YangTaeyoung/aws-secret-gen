package aws_secret_gen

import (
	"os"
	"path/filepath"
	"runtime"
)

func HomeDirPath() string {
	// reference: https://stackoverflow.com/questions/7922270/obtain-users-home-directory
	if runtime.GOOS == "windows" {
		home := os.Getenv("HOMEDRIVE") + os.Getenv("HOMEPATH")
		if home == "" {
			home = os.Getenv("USERPROFILE")
		}
		return home
	} else if runtime.GOOS == "linux" {
		home := os.Getenv("XDG_CONFIG_HOME")
		if home != "" {
			return home
		}
	}
	return os.Getenv("HOME")
}

func WriteFile(content string, path string) error {
	var err error

	// create directory
	dir := filepath.Dir(path)
	if _, err = os.Stat(dir); os.IsNotExist(err) {
		if err = os.MkdirAll(dir, 0755); err != nil {
			return err
		}
	} else if err != nil {
		return err
	}

	// create file
	file, err := os.Create(path)
	if err != nil {
		return err
	}

	// write file
	if _, err = file.WriteString(content); err != nil {
		return err
	}

	return nil
}
