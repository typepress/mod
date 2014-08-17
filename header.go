package mod

import (
	"net/http"
	"path"
)

/**
SetCacheControl 以扩展名设置 Cache-Control 头
参数:
	cc 举例
		map[string]string{
			".jpg": "max-age=31536000",
		}
返回:
	总是 false
*/
func SetCacheControl(cc map[string]string) Finish {

	return Finish(func(
		rw http.ResponseWriter, req *http.Request) bool {

		s, ok := cc[path.Ext(req.URL.Path)]
		if ok {
			rw.Header().Set("CacheControl", s)
		}
		return false
	})
}
