package data

import(
	"os"
	"time"
)


func dictIpadictAcBytes() ([]byte, error) {
	return _dictIpadictAc, nil
}

func dictIpadictAc() (*asset, error) {
	bytes, err := dictIpadictAcBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "dict/ipadict.ac", size: 524288, mode: os.FileMode(420), modTime: time.Unix(1596545184, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}
