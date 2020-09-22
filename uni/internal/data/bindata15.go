package data

import(
	"os"
	"time"
)


func dictUnidictApBytes() ([]byte, error) {
	return _dictUnidictAp, nil
}

func dictUnidictAp() (*asset, error) {
	bytes, err := dictUnidictApBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "dict/unidict.ap", size: 524288, mode: os.FileMode(420), modTime: time.Unix(1596854809, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}
