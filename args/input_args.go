package args

import (
	"log_segment/fs_stringhandle"
	"os"
	"strconv"
)

type InputArgs struct {
	exePath       string `docs:"exe文件所在位置"`
	exeName       string `docs:"exe文件名称"`
	exePrefixPath string `docs:"前缀路径"`
	fileName      string `docs:"文件名称"`
	maxSize       int64  `docs:"大小(MB"`
	saveNumber    int    `docs:"保存数量"`
	checkInterval int    `docs:"查询间隔"`
}

func (i InputArgs) ExePath() string {
	return i.exePath
}

func (i InputArgs) ExeName() string {
	return i.exeName
}

func (i InputArgs) ExePrefixPath() string {
	return i.exePrefixPath
}

func (i InputArgs) FileName() string {
	return i.fileName
}

func (i InputArgs) MaxSize() int64 {
	return i.maxSize
}

func (i InputArgs) SaveNumber() int {
	return i.saveNumber
}

func (i InputArgs) CheckInterval() int {
	return i.checkInterval
}

// InitByArgs 返回一个InputArgs对象 根据配置的参数来
// log_segmentation.go info.log 10 10 10
func InitByArgs() *InputArgs {
	args := os.Args

	// 获取当前执行文件的全路径
	fileStr, err := os.Executable()
	if err != nil {
		panic(err)
	}
	//fileStr := filepath.Dir(ex)

	// 获取当前路径的前缀和后缀
	prefixPath, name := fs_stringhandle.PathSeparator(fileStr)

	if len(args) < 5 {
		return &InputArgs{
			exePath:       fileStr,
			exePrefixPath: prefixPath,
			exeName:       name,
			fileName:      "info.log",
			//maxSize:       10 * 1024 * 1024,
			maxSize:       512,
			saveNumber:    10,
			checkInterval: 10,
		}
	}

	maxSizeInt, _ := strconv.Atoi(args[2])
	saveNumberInt, _ := strconv.Atoi(args[3])
	checkIntervalInt, _ := strconv.Atoi(args[4])

	return &InputArgs{
		exePath:       fileStr,
		exePrefixPath: prefixPath,
		exeName:       name,
		fileName:      args[1],
		maxSize:       int64(1024 * maxSizeInt * 1024),
		saveNumber:    saveNumberInt,
		checkInterval: checkIntervalInt,
	}
}
