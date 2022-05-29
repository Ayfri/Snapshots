package web

import (
	"html/template"
	"io/fs"
	"net/http"
)

type Site struct {
	fs         fs.FS
	fileServer http.Handler
	funcs      template.FuncMap
}

func NewSite(fs fs.FS) *Site {
	return &Site{
		fs:         fs,
		fileServer: http.FileServer(http.FS(fs)),
	}
}

func (s Site) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	s.fileServer.ServeHTTP(writer, request)
}
