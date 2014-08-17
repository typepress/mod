package mod

import (
	"mime"
	"net/http"
	"path"
	"strings"
)

/**
GzipPre 发送静态预压缩文件, root 路径由调用者传递.

依据:
	Request.Method 为 GET/HEAD
	Request.Header 含 "Accept-Encoding" 支持 "gzip".
	URL.Path 有扩展名.
	exts 非空, 扩展名在 exts 内.
	对应的 ".gz" 预压缩文件存在.

参数:
	exts 指定文件扩展名, 如果为空, 表示尝试所有文件.

返回:
	文件被发送返回 true, 否则返回 false

注意:
	确保相应文件类型已注册 mime Content-Type.
*/
func GzipPreInExts(exts string) Dir {

	return Dir(func(root http.Dir,
		rw http.ResponseWriter, req *http.Request) bool {

		if req.Method != "GET" && req.Method != "HEAD" {
			return false
		}
		// Accept-Encoding:gzip,deflate,sdch
		if strings.Index(
			req.Header.Get("Accept-Encoding"), "gzip") == -1 {

			return false
		}

		ext := path.Ext(req.URL.Path)

		if ext == "" || len(exts) != 0 && ExtIndex(exts, ext) == -1 {

			return false
		}

		name := req.URL.Path
		basename := name[strings.LastIndex(name, "/")+1:]

		if ext == ".gz" {

			oext := path.Ext(name[:len(name)-3])
			if oext != "" {
				ext = oext
				basename = basename[:len(basename)-3]
			}
		} else {

			name += ".gz"
		}

		f, e := root.Open(name)
		if e != nil {
			return false
		}

		defer f.Close()

		fi, e := f.Stat()

		if e != nil || fi.IsDir() {
			return false
		}

		ctype := mime.TypeByExtension(ext)

		if ctype != "" {
			rw.Header().Set("Content-Type", ctype)
		} else if ext == ".gz" {
			rw.Header().Set("Content-Type", "application/gzip")
		}

		rw.Header().Set("Content-Encoding", "gzip")

		http.ServeContent(rw, req, basename, fi.ModTime(), f)
		return true
	})
}

/**
GzipPre 调用 GzipPreInExts 发送静态预压缩文件.
参数:
	root 指定文档根路径
	exts 指定文件扩展名, 如果为空, 表示尝试所有文件.
返回:
	文件被发送返回 true, 否则返回 false
*/
func GzipPre(root http.Dir, exts string) Finish {

	finish := GzipPreInExts(exts)
	return Finish(func(
		rw http.ResponseWriter, req *http.Request) bool {

		return finish.ServeHTTP(root, rw, req)
	})
}

/**
StaticInExts 发送静态文件, root 路径由调用者传递.

依据:
	Request.Method 为 GET/HEAD
	URL.Path 有扩展名.
	exts 非空, 扩展名在 exts 内.
	文件存在.

参数:
	exts 指定文件扩展名, 如果为空, 表示尝试所有文件.
返回:
	文件被发送返回 true, 否则返回 false
*/
func StaticInExts(exts string) Dir {

	return Dir(func(root http.Dir,
		rw http.ResponseWriter, req *http.Request) bool {

		if req.Method != "GET" && req.Method != "HEAD" {
			return false
		}

		ext := path.Ext(req.URL.Path)

		if ext == "" || len(exts) != 0 && ExtIndex(exts, ext) == -1 {
			return false
		}

		name := req.URL.Path
		f, e := root.Open(name)
		if e != nil {
			return false
		}

		defer f.Close()

		fi, e := f.Stat()

		if e != nil || fi.IsDir() {
			return false
		}

		http.ServeContent(rw, req,
			name[strings.LastIndex(name, "/")+1:], fi.ModTime(), f)

		return true
	})
}

/**
Static 调用 StaticInExts 发送静态文件.
参数:
	root 指定文档根路径
	exts 指定文件扩展名, 如果为空, 表示尝试所有文件.
返回:
	文件被发送返回 true, 否则返回 false
*/
func Static(root http.Dir, exts string) Finish {

	finish := StaticInExts(exts)
	return Finish(func(
		rw http.ResponseWriter, req *http.Request) bool {

		return finish.ServeHTTP(root, rw, req)
	})
}
