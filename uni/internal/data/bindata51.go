package data

import(
	"os"
	"time"
)


func dictUnidictBzBytes() ([]byte, error) {
	return _dictUnidictBz, nil
}

func dictUnidictBz() (*asset, error) {
	bytes, err := dictUnidictBzBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "dict/unidict.bz", size: 524288, mode: os.FileMode(420), modTime: time.Unix(1596854809, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}
