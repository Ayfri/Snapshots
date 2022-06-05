package api

import (
	"VersionCraft2/internal"
	"VersionCraft2/internal/versions"
	"VersionCraft2/internal/web"
	"errors"
	"github.com/patrickmn/go-cache"
	"io/fs"
	"net/http"
	"time"
)

type handler struct {
	cache   *cache.Cache
	content fs.FS
	site    *web.Site
}

func (h *handler) CacheSiteData(data Data) {
	h.cache.Set("data", data, 12*time.Hour)
}

func RegisterHandlers(mux *http.ServeMux, site *web.Site, content fs.FS) {
	data, err := FetchData()
	if err != nil {
		internal.Fatal(err)
	}

	handler := newHandler(site, content)
	handler.CacheSiteData(data)

	mux.Handle("/versions/", handler)
}

func newHandler(site *web.Site, content fs.FS) *handler {
	const CacheDuration = 24 * time.Hour
	const CleanupInterval = 12 * time.Hour

	return &handler{
		site:    site,
		content: content,
		cache:   cache.New(CacheDuration, CleanupInterval),
	}
}

func (h *handler) getData() (Data, error) {
	data, found := h.cache.Get("data")
	if found {
		return data.(Data), nil
	}

	return nil, errors.New("data not found")
}

func (h *handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Path[len("/versions/"):]
	if name == "" {
		name = "latest"
	}

	data, err := h.getData()
	if err != nil {
		h.site.ServeError(w, r, err)
	}

	var version *versions.Version
	if name == "latest" {
		version = data.GetLatestVersion()
	} else {
		version = data.GetVersion(name)
	}

	if version == nil {
		h.site.ServeError(w, r, errors.New("version not found"))
	}

	layout := "version"
	if name == "latest" {
		layout = "versions"
	}

	h.site.ServePage(w, r, web.Page{
		"name":    version.Name,
		"layout":  layout,
		"version": version,
		"title":   version.Name + " - VersionCraft",
	})
}
