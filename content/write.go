package content

import (
	"os"
	"path/filepath"

	"github.com/spf13/afero"
)

func Write(fs afero.Fs, filename string, content []byte) error {

	fileDir := filepath.Dir(filename)
	if _, err := os.Stat(fileDir); os.IsNotExist(err) {
		fs.MkdirAll(fileDir, os.ModePerm)
	}

	return afero.WriteFile(fs, filename, content, os.FileMode(int(0777)))
}
