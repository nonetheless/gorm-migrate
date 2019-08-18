package utils

import (
	"github.com/nonetheless/gorm-migrate/asset"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"runtime"
)

func getTemplateDirPath() string {
	if _, fileNameWithPath, _, ok := runtime.Caller(1); ok {
		return fileNameWithPath + "/../template"
	}
	return ""
}

func ReadOrCreate(path string) ([]os.FileInfo, error) {
	dir, err := ioutil.ReadDir(path)
	if err != nil {
		err := os.MkdirAll(path, os.ModePerm)
		if err != nil {
			return nil, err
		}
		dir, _ = ioutil.ReadDir(path)
	}
	return dir, nil
}

func WriteTemplate(outputFile string, writeFunc func(writer io.Writer) error) error {
	var w io.Writer
	if outputFile == "" {
		w = os.Stdout
	} else {
		os.MkdirAll(filepath.Dir(outputFile), 0755)
		f, err := os.OpenFile(outputFile, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0666)
		defer f.Close()
		w = io.Writer(f)
		if err != nil {
			return err
		}
	}
	// ToDo: c.RenderContext
	err := writeFunc(w)
	if err != nil {
		return err
	}
	return nil
}

func AssetTemplate(path string) ([]byte, error) {
	data, err := asset.Asset(path)
	if err != nil {
		// Asset was not found.
		return nil, err
	}
	return data, nil
}
