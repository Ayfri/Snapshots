package main

import (
	"VersionCraft2/internal/web"
	"flag"
	"io/fs"
	"log"
	"net/http"
	"os"
	"strconv"
)

var (
	contentDir = flag.String("content", "content", "content directory")
	port       = flag.Int("port", 8080, "port to listen on")
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
	log.Printf("Listening on port %d", *port)
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
	return site
}

func newSite(mux *http.ServeMux, contentFs fs.FS) *web.Site {
	site := web.NewSite(contentFs)

	mux.Handle("/", site)

	return site
}
