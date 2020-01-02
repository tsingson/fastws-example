package vtils

import (
	"bytes"
	"io/ioutil"

	"github.com/BurntSushi/toml"
)

// SaveTOML  save interface to TOML file
func SaveTOML(v interface{}, filename string) error {
	// currentPath+"/toml1.toml"

	var b bytes.Buffer
	e := toml.NewEncoder(&b)
	err := e.Encode(v)
	if err != nil {
		return err
	}
	err = WriteToFile(b.Bytes(), filename)
	return err
}

// WriteToFile  write []byte to file
func WriteToFile(c []byte, filename string) error {
	// 将指定内容写入到文件中

	err := ioutil.WriteFile(filename, c, 0666)
	return err
}
