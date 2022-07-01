package fs_stringhandle

// PathSeparator 删除文件末尾 for example:/var/log/log.info -> /var/log log.info
func PathSeparator(str string) (prefix string, suffix string) {
	// 获取系统的文件分隔下标
	Separator := "/"

	lastIndex := LastStringIndex(str, Separator)

	prefix = SubString(str, 0, lastIndex)
	suffix = SubString(str, lastIndex+1, len(str))

	return
}
