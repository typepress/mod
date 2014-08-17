package mod

import (
	"net/http"
)

/**
Interface for mod
*/
type Interface interface {
	// 占位签名, 表示采用 mod 的接口风格.
	GitHubTypepressMod()
}

type Finish func(http.ResponseWriter, *http.Request) bool

func (h Finish) GitHubTypepressMod() {}

func (h Finish) ServeHTTP(
	rw http.ResponseWriter, req *http.Request) bool {

	return h(rw, req)
}

type Dir func(http.Dir, http.ResponseWriter, *http.Request) bool

func (h Dir) GitHubTypepressMod() {}

func (h Dir) ServeHTTP(dir http.Dir,
	rw http.ResponseWriter, req *http.Request) bool {

	return h(dir, rw, req)
}

type String func(string, http.ResponseWriter, *http.Request) bool

func (h String) GitHubTypepressMod() {}

func (h String) ServeHTTP(v string,
	rw http.ResponseWriter, req *http.Request) bool {

	return h(v, rw, req)
}

type Uint func(uint, http.ResponseWriter, *http.Request) bool

func (h Uint) GitHubTypepressMod() {}

func (h Uint) ServeHTTP(v uint,
	rw http.ResponseWriter, req *http.Request) bool {

	return h(v, rw, req)
}
