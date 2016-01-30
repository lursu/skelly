package builder

import (
	"bytes"
	"github.com/spf13/viper"
	"io/ioutil"
	"os"
)

type File struct {
	permission os.FileMode
	name       string
	path       string
	contents   []byte
}

func (f File) Write() error {
	buffer := bytes.NewBufferString(f.path)
	fullName := buffer.WriteString(f.name)
	return ioutil.WriteFile(fullName.String(), contents, permission)
}

func BuildFile(key string, name string, path string, permission os.FileMode) *File {
	return File{
		permission: permission,
		name:       name,
		path:       path,
		contents:   getFileContents(key),
	}
}

func getFileContents(key string) []byte {
	result := new(bytes.Buffer)
	funcPath := bytes.NewBufferString(key)
	viper := viper.New()
	fileMap := viper.GetStringMap(key)
	for k, v := range fileMap {
		switch k {
		case "package":
			result.WriteString("package")
			result.WriteString(v)
		case "imports":
			result.WriteString("import (\n")
			result.WriteString(v)
		default:
			buildFuncContents(key, result)
		}
	}
	return result.Bytes()
}

func getFuncContents(baseKey string, b *Buffer) {
	viper := viper.New()
	funcPath := bytes.NewBufferString(baseKey)
	funcPath.WriteString(".functions")
	funcsMap := viper.GetStringMap(funcPath.String())
	for k, v := range funcsMap {
		b.WriteString("\n")
		buildFunc(v.(map[string]string), &b)
	}
}

func buildFunc(funcMap map[string]string, b *Buffer) {
	b.WriteString("func ")
	for k, v := range funcMap {
		b.WriteString(v)
	}
	b.WriteString("}\n")
}
