package main

import (
	"fmt"
	"log_segment/args"
	"log_segment/util"
	"os"
	"time"
)

// 日志切分, 遍历日志
// 依赖 fs_util_oil
// nohup ./log_segment catalina.out 10 10 10 > /dev/null 2>&1 &
// ./logSegment log.info 	10				10				10
// ./logSegment 文件名称		切分大小10MB		保存数量 10个		检查时间10s
// @author 红烧鲈鱼
func main() {
	// 获取参数
	fileArgs := args.InitByArgs()

	fmt.Println("启动参数如下")
	fmt.Println("执行路径", fileArgs.ExePrefixPath())
	fmt.Println("文件名", fileArgs.ExeName())
	fmt.Println("文件名", fileArgs.FileName())
	fmt.Println("切分大小", fileArgs.MaxSize())
	fmt.Println("保存数量", fileArgs.SaveNumber())
	fmt.Println("检查间隔", fileArgs.CheckInterval())

	duration := time.Duration(fileArgs.CheckInterval())

	// 遍历循环
	for {
		// 获取指定文件信息
		logFile, err := os.Stat(fileArgs.FileName())
		if nil == err {
			fmt.Println("当前文件大小:", logFile.Size())
			// 这个方法是以字节为单位的
			if logFile.Size() > fileArgs.MaxSize() {
				util.CopyFileAndClear(fileArgs)
			}
		}

		time.Sleep(duration * time.Second)
	}

}
