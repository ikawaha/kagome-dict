package data

import(
	"os"
	"time"
)


func dictUnidictBtBytes() ([]byte, error) {
	return _dictUnidictBt, nil
}

func dictUnidictBt() (*asset, error) {
	bytes, err := dictUnidictBtBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "dict/unidict.bt", size: 524288, mode: os.FileMode(420), modTime: time.Unix(1596854809, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}
