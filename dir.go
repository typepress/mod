package mod

import (
	"fmt"
	"html"
	"net/http"
	"os"
	"strings"
)

/**
List 使用简洁风格列表目录, root 路径由调用者传递.

参数:
	n    限制目录列表的个数, n <= 0 表示不限制.
返回:
	true  目录被列表
	false 目录未被列表
*/
func List(n int) Dir {

	return Dir(func(root http.Dir,
		w http.ResponseWriter, req *http.Request) bool {

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

/**
DirList 调用 List 使用简洁风格列表目录.

参数:
	root 指定文档根路径
	n    限制目录列表的个数, n <= 0 表示不限制.
返回:
	true  目录被列表
	false 目录未被列表
*/
func DirList(root http.Dir, n int) Finish {
	finish := List(n)
	return Finish(func(
		rw http.ResponseWriter, req *http.Request) bool {

		return finish.ServeHTTP(root, rw, req)
	})
}

const indexRedirect = "/index.html"

/**
Redirect 遵守 http 1.1 规范对目录进行 301 重定向.

依据:
	URL.Path 以 "/index.html" 结尾.
	URL.Path 是目录, 没有以 "/" 结尾.

参数:
	root 指定文档根路径
*/
func Redirect(root http.Dir,
	rw http.ResponseWriter, req *http.Request) bool {

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
}

/**
DirRedirect 调用 Redirect 对目录进行 301 重定向.

参数:
	root 指定文档根路径
*/
func DirRedirect(root http.Dir) Finish {

	return Finish(func(
		rw http.ResponseWriter, req *http.Request) bool {

		return Redirect(root, rw, req)
	})
}
