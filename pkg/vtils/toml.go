package vtils

import (
	"bytes"

	"github.com/BurntSushi/toml"
)

// TomlEncode  encode
func TomlEncode(v interface{}) ([]byte, error) {
	var b bytes.Buffer
	e := toml.NewEncoder(&b)
	er2 := e.Encode(v)
	if er2 != nil {
		return nil, er2
	}

	return b.Bytes(), nil
}

// TomlDecode  decode
func TomlDecode(data []byte, v interface{}) error {
	_, err := toml.Decode(B2S(data), v)

	return err
}
