package main

import (
	"VersionCraft2/internal"
	"VersionCraft2/internal/api"
	"VersionCraft2/internal/web"
	"flag"
	"io/fs"
	"log"
	"net/http"
	"os"
	"strconv"
)

var (
	contentDir = flag.String("content", "../content", "content directory")
	port       = flag.Int("port", 8080, "port to listen on")
	verbose    = flag.Bool("verbose", false, "verbose logging")
)

func usage() {
	flag.PrintDefaults()
	os.Exit(2)
}

func main() {
	flag.Usage = usage
	flag.Parse()

	if *contentDir == "" {
		usage()
	}

	if *port == 0 {
		usage()
	}

	httpAddress := ":" + strconv.Itoa(*port)

	handler := NewHandler(*contentDir)
	if *verbose {
		log.Printf("Listening on %s", httpAddress)
		handler = loggingHandler(handler)
	}
	log.Fatal(http.ListenAndServe(httpAddress, handler))
}

func NewHandler(contentDir string) http.Handler {
	siteMux := http.NewServeMux()

	var contentFs fs.FS
	if contentDir != "" {
		contentFs = os.DirFS(contentDir)
	} else {
		contentFs = os.DirFS("../content")
	}

	site := newSite(siteMux, contentFs)
	data, err := api.FetchData()
	if err != nil {
		internal.Fatal(err)
	}

	site.CacheSiteData(data)

	return site
}

func newSite(mux *http.ServeMux, contentFs fs.FS) *web.Site {
	site := web.NewSite(contentFs)
	site.SetVerbose(*verbose)

	mux.Handle("/", site)

	return site
}

func loggingHandler(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("%s %s %s", r.RemoteAddr, r.Method, r.URL.RequestURI())
		h.ServeHTTP(w, r)
	})
}
