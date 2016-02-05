package cmd

import (

	"bytes"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"text/template"
	"errors"
)


type Project int

var projects = map[string]Project{
	"web": WEB,
}

const (
	WEB Project = iota
)

func errDie(msg interface{}) {
	fmt.Println("Error:", msg)
	os.Exit(-1)
}

func WriteProject(project Project) {
	switch project{
	case WEB:
	}
}

func WriteTemplateToFile(filePath string, name string, templateName string, data interface{}) error {
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

// Creates a new io.Reader that contains the contents of the template and the data that was passed into it
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

// write the file that is already in the io.Reader at the path
func writeFile(path string, reader io.Reader) error {
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
	return nil
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

// attempt to make the dir if the specified dir does not exist
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

func joinPath(base string, addition string) string {
	return filepath.Join(base, addition)
}

func getType(name string) (Project, error) {
	var result Project
	result, ok := projects[name]
	if !ok {
		return result, errors.New("could not find project type")
	}
	return result, nil
}
