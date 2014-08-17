package mod

import (
	"net/http"
	"path"
	"strings"
)

/**
AllowPrefix, URL.Path 不含有指定 prefix, 403 完结.
*/
func AllowPrefix(prefix string) Finish {

	return Finish(func(
		rw http.ResponseWriter, req *http.Request) bool {

		if strings.HasPrefix(req.URL.Path, prefix) {
			return false
		}
		rw.WriteHeader(http.StatusForbidden)
		return true
	})
}

/**
AllowExtIn, URL.Path 扩展名不在 exts 内, 403 完结.
*/
func AllowExtIn(exts string) Finish {

	return Finish(func(
		rw http.ResponseWriter, req *http.Request) bool {

		if ExtIndex(exts, path.Ext(req.URL.Path)) != -1 {
			return false
		}
		rw.WriteHeader(http.StatusForbidden)
		return true
	})
}
