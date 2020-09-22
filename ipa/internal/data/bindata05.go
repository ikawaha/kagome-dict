package data

import(
	"os"
	"time"
)


func dictIpadictAfBytes() ([]byte, error) {
	return _dictIpadictAf, nil
}

func dictIpadictAf() (*asset, error) {
	bytes, err := dictIpadictAfBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "dict/ipadict.af", size: 524288, mode: os.FileMode(420), modTime: time.Unix(1596545184, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}
