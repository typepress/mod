package mod

import (
	"net/http"
	"strings"
)

/**
DenyPrefix, URL.Path 含有指定 prefix, 403 完结.
*/
func DenyPrefix(prefix string) Finish {

	return Finish(func(rw http.ResponseWriter,
		req *http.Request) bool {

		if !strings.HasPrefix(req.URL.Path, prefix) {
			return false
		}
		rw.WriteHeader(http.StatusForbidden)
		return true
	})
}

/**
DenyExtIn, URL.Path 扩展名在 exts 内, 403 完结.
*/
func DenyExtIn(exts string) Finish {

	return Finish(func(rw http.ResponseWriter,
		req *http.Request) bool {

		if ExtIndex(exts, req.URL.Path) == -1 {
			return false
		}
		rw.WriteHeader(http.StatusForbidden)
		return true
	})
}
