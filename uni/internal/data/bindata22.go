package data

import(
	"os"
	"time"
)


func dictUnidictAwBytes() ([]byte, error) {
	return _dictUnidictAw, nil
}

func dictUnidictAw() (*asset, error) {
	bytes, err := dictUnidictAwBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "dict/unidict.aw", size: 524288, mode: os.FileMode(420), modTime: time.Unix(1596854809, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}
