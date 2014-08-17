package mod

import (
	"fmt"
	"html"
	"net/http"
	"os"
	"strings"
)

/**
DirList 对目录进行简洁风格列表

参数:
	root 指定文档根路径
	n    限制目录列表的个数, n <= 0 表示不限制.
返回:
	true  目录被列表
	false 目录未被列表
*/
func DirList(root http.Dir, n int) Finish {

	return Finish(func(w http.ResponseWriter, req *http.Request) bool {

		f, e := root.Open(req.URL.Path)
		if e != nil {
			return false
		}

		defer f.Close()

		fi, e := f.Stat()

		if e != nil || !fi.IsDir() {
			return false
		}

		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		fmt.Fprintf(w, "<pre>\n")
		fmt.Fprint(w, "<a href=\"../\">..</a>\n")

		var dirs []os.FileInfo

		for {

			dirs, _ = f.Readdir(n)
			if len(dirs) == 0 {
				break
			}

			for _, fi = range dirs {
				if !fi.IsDir() {
					continue
				}

				fmt.Fprintf(w, "<a href=\"%[1]s/\">%[1]s/</a>\n",
					html.EscapeString(fi.Name()))
			}

			for _, fi = range dirs {
				if fi.IsDir() {
					continue
				}

				fmt.Fprintf(w, "<a href=\"%[1]s\">%[1]s</a>\n",
					html.EscapeString(fi.Name()))
			}
		}
		fmt.Fprintf(w, "</pre>\n")
		return true
	})
}

const indexRedirect = "/index.html"

/**
DirRedirect 遵守 http 1.1 规范对目录进行 301 重定向.

依据:
	URL.Path 以 "/index.html" 结尾.
	URL.Path 是目录, 没有以 "/" 结尾.

依据被肯定, 301 重定向到以 "/" 结尾.
参数:
	root 指定文档根路径
*/
func DirRedirect(root http.Dir) Finish {

	return Finish(func(rw http.ResponseWriter,
		req *http.Request) bool {

		name := req.URL.Path

		pos := strings.LastIndex(name, "/")

		if len(name) == pos+1 {
			return false
		}

		if name[pos:] == indexRedirect {

			name = name[:pos+1]
			if q := req.URL.RawQuery; q != "" {
				name += "?" + q
			}

			rw.Header().Set("Location", name)
			rw.WriteHeader(http.StatusMovedPermanently)
			return true
		}

		f, e := root.Open(name)
		if e != nil {
			return false
		}

		defer f.Close()

		fi, e := f.Stat()

		if e != nil || !fi.IsDir() {
			return false
		}

		if q := req.URL.RawQuery; q != "" {
			name += "/?" + q
		} else {
			name += "/"
		}

		rw.Header().Set("Location", name)
		rw.WriteHeader(http.StatusMovedPermanently)
		return true
	})
}
