package data

import(
	"os"
	"time"
)


func dictUnidictBrBytes() ([]byte, error) {
	return _dictUnidictBr, nil
}

func dictUnidictBr() (*asset, error) {
	bytes, err := dictUnidictBrBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "dict/unidict.br", size: 524288, mode: os.FileMode(420), modTime: time.Unix(1596854809, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}
