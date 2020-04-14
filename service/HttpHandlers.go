package service

import (
	"net/http"
	"path"
	"strings"
)

var FileHandler = http.FileServer(http.Dir("./www/"))

type myFileHandler struct {
	endpoints map[string]string
}

func (f myFileHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	upath := r.URL.Path
	if !strings.HasPrefix(upath, "/") {
		upath = "/" + upath
	}
	upath = f.endpoints[path.Clean(upath)]
	http.ServeFile(w, r, upath)
	//f.serveFile(w, r, upath)
}

func (f myFileHandler) serveFile(w http.ResponseWriter, r *http.Request, filename string) {
	http.ServeFile(w, r, filename)
}

func GetMyUrlHandler() http.Handler {
	return myFileHandler{
		endpoints: map[string]string{
			"/hello": "./www/index.html",
		},
	}
}
