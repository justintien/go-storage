package util

// Md5PathName 转换MD5为路径文件名。
func Md5PathName(name string) string {
	pathName := ""
	i := 0
	for ; i < 32; i += 2 {
		pathName += "/" + name[i:i+2]
	}
	if len(name) > 32 {
		pathName += name[i:]
	}

	return pathName
}
