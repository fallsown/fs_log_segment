package fs_file

import (
	"io/ioutil"
	"os"
)

// GetFilesAndDirs 获取指定目录下的所有文件和目录
func GetFilesAndDirs(dirPth string) (files []string, dirs []string, err error) {
	dir, err := ioutil.ReadDir(dirPth)
	if err != nil {
		return nil, nil, err
	}

	PthSep := string(os.PathSeparator)
	//suffix = strings.ToUpper(suffix) //忽略后缀匹配的大小写

	for _, fi := range dir {
		if fi.IsDir() { // 目录, 递归遍历
			dirs = append(dirs, dirPth+PthSep+fi.Name())
			GetFilesAndDirs(dirPth + PthSep + fi.Name())
		} else {
			files = append(files, dirPth+PthSep+fi.Name())

			// 过滤指定格式
			// ok := strings.HasSuffix(fi.Name(), ".go")
			// if ok {
			//     files = append(files, dirPth+PthSep+fi.Name())
			// }
		}
	}

	return files, dirs, nil
}

// GetFiles 获取指定目录下的所有文件
// files 文件名(没有前面啥的东西)
func GetFiles(dirPth string) (files []string, err error) {
	dir, err := ioutil.ReadDir(dirPth)
	if err != nil {
		return nil, err
	}

	for _, fi := range dir {
		if fi.IsDir() {

		} else {
			files = append(files, fi.Name())
		}
	}

	return files, nil
}
