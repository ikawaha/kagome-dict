package data

import(
	"os"
	"time"
)


func dictIpadictAiBytes() ([]byte, error) {
	return _dictIpadictAi, nil
}

func dictIpadictAi() (*asset, error) {
	bytes, err := dictIpadictAiBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "dict/ipadict.ai", size: 524288, mode: os.FileMode(420), modTime: time.Unix(1596545184, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}
