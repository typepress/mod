package mod

import (
	"strings"
)

/**
ExtIndex 辅助函数, 模仿 strings.Index 风格.
返回 ext 在 exts 中第一个完整匹配的位置.
参数:
	exts 举例 ".jpg.html", 不要有空格
	ext  举例 ".js"

返回:
	-1 表示匹配失败, ext 为空也返回 -1.
	其它值表示匹配到的位置.
*/
func ExtIndex(exts string, ext string) int {

	if ext == "" {
		return -1
	}

	pos := strings.Index(exts, ext)

	if pos == -1 ||
		(pos+len(ext) == len(exts) ||
			exts[pos+len(ext)] == '.') {

		return -1
	}

	return pos
}
