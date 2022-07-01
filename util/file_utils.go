package util

import (
	"fmt"
	"io"
	"log_segment/args"
	"log_segment/fs_file"
	"os"
	"time"
)

// CopyFileAndClear 对文件进行赋值并清空文件
func CopyFileAndClear(fileArgs *args.InputArgs) {

	// 创建新文件
	newFile, err := os.Create(fileArgs.FileName() + "." + time.Now().Format("20060102150405"))
	defer newFile.Close()

	if nil == err {
		// 写入文件 - 固态硬盘一个区为4096
		fileRead, error := os.Open(fileArgs.FileName())
		defer fileRead.Close()
		if error == nil {
			buf := make([]byte, 4096)
			for {
				n, err2 := fileRead.Read(buf)
				if err2 != nil && err2 == io.EOF {
					break
				}
				// 读多少写多少
				newFile.Write(buf[:n])
			}
		}

		// 清空原来文件
		fmt.Println("准备清空文件")
		os.Truncate(fileArgs.FileName(), 0)
		fmt.Println("清空完毕")

		// 获取当前文件夹下所有的文件名
		fmt.Println("准备获取文件夹下所有文件名")
		files, err := fs_file.GetFiles(fileArgs.ExePrefixPath())
		if err == nil {
			// 过滤出前缀为日志前缀的文件(不包括自己
			files = FilterPrefixList(&files, fileArgs.FileName())
			// 删除最老的文件
			if fileArgs.SaveNumber() < len(files) {
				fmt.Println("超出限制 准备删除文件")
				os.Remove(files[0])
			}
		} else {
			fmt.Println(err.Error())
		}
	}

}
