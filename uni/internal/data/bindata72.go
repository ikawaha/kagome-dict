package data

import(
	"os"
	"time"
)


func dictUnidictCuBytes() ([]byte, error) {
	return _dictUnidictCu, nil
}

func dictUnidictCu() (*asset, error) {
	bytes, err := dictUnidictCuBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "dict/unidict.cu", size: 524288, mode: os.FileMode(420), modTime: time.Unix(1596854809, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}
