package data

import(
	"os"
	"time"
)


func dictIpadictAbBytes() ([]byte, error) {
	return _dictIpadictAb, nil
}

func dictIpadictAb() (*asset, error) {
	bytes, err := dictIpadictAbBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "dict/ipadict.ab", size: 524288, mode: os.FileMode(420), modTime: time.Unix(1596545184, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}
