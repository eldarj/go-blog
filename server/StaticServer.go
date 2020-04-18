package server

import (
	"net/http"
)

type StaticServer struct{}

func (ss *StaticServer) ServeResources(uri string, serveFromPath string) {
	dir := http.Dir(serveFromPath)
	fileServer := http.FileServer(dir)
	http.Handle(uri, http.StripPrefix(uri, fileServer))
}

func (ss *StaticServer) ServeStylesheets(uri string, serveFromPath string) {
	dir := http.Dir(serveFromPath)
	fileServer := http.FileServer(dir)
	http.Handle(uri, http.StripPrefix(uri, fileServer))
}
