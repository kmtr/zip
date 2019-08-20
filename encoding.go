package zip

import (
	"io/ioutil"
	"strings"

	"golang.org/x/text/transform"
)

func encodeText(str string, t transform.Transformer) (string, error) {
	iostr := strings.NewReader(str)
	rio := transform.NewReader(iostr, t)
	ret, err := ioutil.ReadAll(rio)
	if err != nil {
		return "", err
	}
	return string(ret), nil
}
