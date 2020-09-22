package data

import(
	"os"
	"time"
)


func dictIpadictAgBytes() ([]byte, error) {
	return _dictIpadictAg, nil
}

func dictIpadictAg() (*asset, error) {
	bytes, err := dictIpadictAgBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "dict/ipadict.ag", size: 524288, mode: os.FileMode(420), modTime: time.Unix(1596545184, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}
