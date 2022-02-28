package backend

import "net/http"

func main() {
	http.ServeFile(nil, nil, "../client/pages/index.html")
}
