package data

import(
	"os"
	"time"
)


func dictUnidictCvBytes() ([]byte, error) {
	return _dictUnidictCv, nil
}

func dictUnidictCv() (*asset, error) {
	bytes, err := dictUnidictCvBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "dict/unidict.cv", size: 524288, mode: os.FileMode(420), modTime: time.Unix(1596854809, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}
