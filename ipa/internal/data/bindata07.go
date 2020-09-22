package data

import(
	"os"
	"time"
)


func dictIpadictAhBytes() ([]byte, error) {
	return _dictIpadictAh, nil
}

func dictIpadictAh() (*asset, error) {
	bytes, err := dictIpadictAhBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "dict/ipadict.ah", size: 524288, mode: os.FileMode(420), modTime: time.Unix(1596545184, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}
