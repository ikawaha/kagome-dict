package data

import(
	"os"
	"time"
)


func dictUnidictBcBytes() ([]byte, error) {
	return _dictUnidictBc, nil
}

func dictUnidictBc() (*asset, error) {
	bytes, err := dictUnidictBcBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "dict/unidict.bc", size: 524288, mode: os.FileMode(420), modTime: time.Unix(1596854809, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}
