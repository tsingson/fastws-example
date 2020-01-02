package vtils

import (
	"bytes"
	"io/ioutil"

	"emperror.dev/errors"

	"github.com/tsingson/chardet"
	"golang.org/x/text/encoding/simplifiedchinese"
	"golang.org/x/text/encoding/traditionalchinese"
	"golang.org/x/text/transform"
)

// Detect detect charset
func Detect(b []byte) (*chardet.Result, error) {
	textDetector := chardet.NewTextDetector()
	return textDetector.DetectBest(b)
}

// DecodeGBK convert GBK to UTF-8
func DecodeGBK(s []byte) ([]byte, error) {
	I := bytes.NewReader(s)
	O := transform.NewReader(I, simplifiedchinese.GBK.NewDecoder())
	d, e := ioutil.ReadAll(O)
	if e != nil {
		return nil, e
	}
	return d, nil
}

// DecodeBIG5 convert BIG5 to UTF-8
func DecodeBIG5(s []byte) ([]byte, error) {
	I := bytes.NewReader(s)
	O := transform.NewReader(I, traditionalchinese.Big5.NewDecoder())
	d, e := ioutil.ReadAll(O)
	if e != nil {
		return nil, e
	}
	return d, nil
}

// EncodeBIG5  convert UTF-8 to BIG5
func EncodeBIG5(s []byte) ([]byte, error) {
	I := bytes.NewReader(s)
	O := transform.NewReader(I, traditionalchinese.Big5.NewEncoder())
	d, e := ioutil.ReadAll(O)
	if e != nil {
		return nil, e
	}
	return d, nil
}

// TransToUTF  translate  code to simplechinese
func TransToUTF(input []byte) (output []byte, err error) {
	var code *chardet.Result

	code, err = Detect(input)
	if err != nil {
		// log.Error().Err(err).Msg("ioutil ReadAll error")
		return nil, err
	}

	switch code.Charset {
	case "GB-18030":
		{
			output, err = DecodeGBK(input)
			if err != nil {
				// 	log.Error().Err(err).Msg("Error")
				return nil, err
			}
		}
		// break
	case "Big5":
		{
			output, err = DecodeBIG5(input)
			if err != nil {
				// 	log.Error().Err(err).Msg("Error")
				return nil, err
			}
		}
		// break
	case "UTF-8":
		output = input
		// return output, nil
		// break
	default:
		err = errors.New("unknow code type")
	}
	return output, err
}
