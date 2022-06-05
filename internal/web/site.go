package web

import (
	"VersionCraft2/internal"
	"errors"
	"fmt"
	"github.com/patrickmn/go-cache"
	"html/template"
	"io/fs"
	"net/http"
	"path"
	"strings"
	"time"
)

const BaseTemplate = "base.tmpl"

type Site struct {
	cache      *cache.Cache
	fs         fs.FS
	fileServer http.Handler
	funcs      template.FuncMap
	verbose    bool
}

func NewSite(fs fs.FS) *Site {
	const CacheDuration = 10 * time.Minute
	const CleanupInterval = 15 * time.Minute

	return &Site{
		cache:      cache.New(CacheDuration, CleanupInterval),
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
	if s.verbose {
		internal.LogError(err)
	}

	p := Page{
		"URL":    r.URL.Path,
		"layout": "error",
		"status": status,
		"error":  err.Error(),
	}

	s.ServePage(w, r, p)
}

func (s *Site) ServePage(w http.ResponseWriter, r *http.Request, page Page) {
	if r.Method == "GET" {
		if page, found := s.cache.Get(r.URL.Path); found {
			w.Write(page.([]byte))
			return
		}
	}

	html, err := s.renderHTML(page, BaseTemplate, r)
	if err != nil {
		s.ServeErrorStatus(w, r, fmt.Errorf("template execution: %v", err), http.StatusInternalServerError)
		return
	}

	if code, ok := page["status"].(int); ok {
		w.WriteHeader(code)
	}

	w.Write(html)
	if s.verbose {
		s.cache.Set(r.URL.Path, html, 5*time.Second)
	} else {
		s.cache.SetDefault(r.URL.Path, html)
	}
}

func (s *Site) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	absolutePath := r.URL.Path
	realPath := path.Clean(strings.TrimPrefix(absolutePath, "/"))

	if page, err := s.openPage(realPath); err == nil {
		s.ServePage(w, r, *page)
		return
	}

	_, err := fs.Stat(s.fs, realPath)
	if err != nil {
		status := http.StatusInternalServerError
		if errors.Is(err, fs.ErrNotExist) {
			status = http.StatusNotFound
		}

		s.ServeErrorStatus(w, r, err, status)
		return
	}

	s.fileServer.ServeHTTP(w, r)
}
