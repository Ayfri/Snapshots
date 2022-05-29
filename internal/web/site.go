package web

import (
	"errors"
	"fmt"
	"html/template"
	"io/fs"
	"net/http"
	"path"
	"strings"
)

const BaseTemplate = "base.tmpl"

type Site struct {
	fs         fs.FS
	fileServer http.Handler
	funcs      template.FuncMap
	verbose    bool
}

func NewSite(fs fs.FS) *Site {
	return &Site{
		fs:         fs,
		fileServer: http.FileServer(http.FS(fs)),
	}
}

func (s *Site) SetVerbose(verbose bool) {
	s.verbose = verbose
}

func (s *Site) Funcs(funcs template.FuncMap) {
	if s.funcs == nil {
		s.funcs = funcs
	} else {
		for k, v := range funcs {
			s.funcs[k] = v
		}
	}
}

func (s *Site) readFile(dir, file string) ([]byte, error) {
	if strings.HasPrefix(file, "/") {
		file = path.Clean(file)
	} else {
		file = path.Join(dir, file)
	}

	file = strings.Trim(file, "/")
	if file == "" {
		file = "."
	}

	return fs.ReadFile(s.fs, file)
}

func (s *Site) ServeError(w http.ResponseWriter, r *http.Request, err error) {
	s.ServeErrorStatus(w, r, err, http.StatusInternalServerError)
}

func (s *Site) ServeErrorStatus(w http.ResponseWriter, r *http.Request, err error, status int) {
	w.WriteHeader(status)
	w.Write([]byte(err.Error()))
}

func (s *Site) ServePage(w http.ResponseWriter, r *http.Request, page Page) {
	html, err := s.renderHTML(page, BaseTemplate, r)
	if err != nil {
		s.ServeErrorStatus(w, r, fmt.Errorf("template execution: %v", err), http.StatusInternalServerError)
		return
	}

	w.Write(html)
}

func (s *Site) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	absolutePath := r.URL.Path
	realPath := path.Clean(strings.TrimPrefix(absolutePath, "/"))

	if page, err := s.openPage(realPath); err == nil {
		s.ServePage(w, r, *page)
		return
	}

	info, err := fs.Stat(s.fs, realPath)
	if err != nil {
		status := http.StatusInternalServerError
		if errors.Is(err, fs.ErrNotExist) {
			status = http.StatusNotFound
		}

		s.ServeErrorStatus(w, r, err, status)
		return
	}

	if info != nil && info.IsDir() {
		s.serveDirectory(w, r, realPath)
		return
	}

	s.fileServer.ServeHTTP(w, r)
}

func (s *Site) serveDirectory(w http.ResponseWriter, r *http.Request, realPath string) {
	list, err := fs.ReadDir(s.fs, realPath)
	if err != nil {
		s.ServeError(w, r, err)
		return
	}

	var info []fs.FileInfo
	for _, f := range list {
		i, err := f.Info()
		if err == nil {
			info = append(info, i)
		}
	}

	s.ServePage(w, r, Page{
		"URL":   r.URL.Path,
		"Files": realPath,
		"dir":   info,
	})
}
