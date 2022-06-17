package api

import (
	"VersionCraft2/internal"
	"VersionCraft2/internal/snapshots"
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

func (h *handler) CacheSiteData(data versions.Versions) {
	h.cache.Set("data", data, 12*time.Hour)
}

func RegisterHandlers(mux *http.ServeMux, site *web.Site, content fs.FS) {
	data, err := versions.FetchVersions()
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

func (h *handler) getData() (versions.Versions, error) {
	data, found := h.cache.Get("data")

	if found {
		return data.(versions.Versions), nil
	} else {
		data, err := versions.FetchVersions()
		if err != nil {
			return nil, err
		}

		h.CacheSiteData(data)
		return data, nil
	}
}

func (h *handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Path[len("/versions/"):]
	if name == "" {
		name = "none"
	}

	data, err := h.getData()
	if err != nil {
		h.site.ServeError(w, r, err)
		return
	}

	versionsMenu, order := data.SplitByMenu()
	versionsPage := web.Page{
		"layout":        "versions",
		"title":         "Versions - VersionCraft",
		"versionsOrder": order,
		"versions":      versionsMenu,
	}

	if name == "none" {
		h.site.ServePage(w, r, versionsPage)
		return
	}

	var version *versions.Version
	if name == "latest" {
		version = data.GetLatestVersion()
	} else {
		version = data.GetVersion(name)
	}

	if version == nil {
		if name == "none" {
			h.site.ServePage(w, r, versionsPage)
			return
		}

		h.site.ServeError(w, r, errors.New("version not found"))
	}

	h.site.ServePage(w, r, web.Page{
		"layout":  "version",
		"orders":  snapshots.Orders,
		"version": version,
		"title":   version.Name + " - VersionCraft",
	})
}
