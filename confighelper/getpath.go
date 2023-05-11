package confighelper

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)
import (
	"runtime"
)

func GetCurrentPath() (string, error) {

	file, err := exec.LookPath(os.Args[0])
	if err != nil {
		return "", err
	}
	path, err := filepath.Abs(file)
	if err != nil {
		return "", err
	}

	//去降exe文件名
	i := strings.LastIndex(path, "/")
	if i < 0 {
		i = strings.LastIndex(path, "\\")
	}

	path = string(path[0 : i+1])

	//fmt.Println("path111:", path)
	if runtime.GOOS == "windows" {
		path = strings.Replace(path, "/", "\\", -1)
		path += "config\\"
	} else {
		path = strings.Replace(path, "\\", "/", -1)
		path += "config/"
	}

	fmt.Println("通过当前目录获取的 config路径为:", path)
	return path, nil
}
