package fileManage

import (
	"io/ioutil"
	"path/filepath"
)

func GetAllFileNames(dirPath string) ([]string, error) {
	var fileNames []string

	// 读取目录
	files, err := ioutil.ReadDir(dirPath)
	if err != nil {
		return nil, err
	}

	// 遍历目录下的文件和子目录
	for _, file := range files {
		if file.IsDir() {
			subDirPath := filepath.Join(dirPath, file.Name())
			subFileNames, err := GetAllFileNames(subDirPath)
			if err != nil {
				return nil, err
			}
			fileNames = append(fileNames, subFileNames...)
		} else {
			fileNames = append(fileNames, file.Name())
		}
	}

	return fileNames, nil
}
