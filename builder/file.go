package builder

import (
	"bytes"
	"fmt"
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
	buffer.WriteString(f.name)
	return ioutil.WriteFile(buffer.String(), f.contents, f.permission)
}

func BuildFile(key string, name string, path string, permission os.FileMode) *File {
	return &File{
		permission: permission,
		name:       name,
		path:       path,
		contents:   getFileContents(key),
	}
}

func getFileContents(key string) []byte {
	result := new(bytes.Buffer)
	fileMap := viper.GetStringMap(key)
	fmt.Println(fileMap)
	for k, v := range fileMap {
		switch k {
		case "package":
			result.WriteString("package")
			result.WriteString(v.(string))
		case "imports":
			result.WriteString("import (\n")
			result.WriteString(v.(string))
		default:
			getFuncContents(key, result)
		}
	}
	return result.Bytes()
}

func getFuncContents(baseKey string, b *bytes.Buffer) {
	viper := viper.New()
	funcPath := bytes.NewBufferString(baseKey)
	funcPath.WriteString(".functions")
	funcsMap := viper.GetStringMap(funcPath.String())
	for _, v := range funcsMap {
		b.WriteString("\n")
		buildFunc(v.(map[string]string), b)
	}
}

func buildFunc(funcMap map[string]string, b *bytes.Buffer) {
	b.WriteString("func ")
	for _, v := range funcMap {
		b.WriteString(v)
	}
	b.WriteString("}\n")
}
