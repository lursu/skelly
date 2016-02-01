package cmd

import (
	"path/filepath"
	"text/template"
	"io"
	"bytes"
	"fmt"
	"os"
)


func WriteTemplateToFile(filePath string, name string, templateName string, data interface{}) (error) {
	fileName := filepath.Join(filePath, name)

	reader, err := getTemplateReader(templateName, data)
	if err != nil {
		return err
	}

	err = writeFile(fileName, reader)

	if err != nil {
		return err
	}
	return nil
}

func getTemplateReader(templateName string, data interface{}) (io.Reader, error) {
	result := template.New("")
	result, err := result.ParseFiles(templateName)
	if err != nil {
		return nil, err
	}

	buffer := new(bytes.Buffer)
	err = result.Execute(buffer, data)
	return buffer, err
}

func writeFile(path string, reader io.Reader) (error) {
	ensureDirExists(path)
	exists, err := fileExists(path)
	if err != nil {
		return err
	}
	if exists {
		return fmt.Errorf("%v already exists", path)
	}

	file, err := os.Create(path)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = io.Copy(file, reader)
	return
}

// Check to see if the filepath exists if it doesn't exist we make it
func fileExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, nil
}

func ensureDirExists(path string) {
	dir, _ := filepath.Split(path)
	osPath := filepath.FromSlash(dir)

	if osPath != "" {
		err := os.Mkdir(osPath, 0777)
		if err != nil {
			return
		}

	}
}