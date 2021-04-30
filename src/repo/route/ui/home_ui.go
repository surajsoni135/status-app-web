package route_ui

import (
	"bufio"
	"encoding/json"
	"net/http"
	"os"
	"strings"

	"github.com/julienschmidt/httprouter"
)

/**
 * UI Methods
 */
func GetHome() func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	return func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode("Hello World")
	}
}

func ServerStaticResource() func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	return func(w http.ResponseWriter, req *http.Request, _ httprouter.Params) {
		path := "public" + req.URL.Path
		var contentType string
		if strings.HasSuffix(path, ".css") {
			contentType = "text/css"
		} else if strings.HasSuffix(path, ".png") {
			contentType = "image/png"
		} else if strings.HasSuffix(path, ".jpg") {
			contentType = "image/jpg"
		} else if strings.HasSuffix(path, ".svg") {
			contentType = "image/svg+xml"
		} else if strings.HasSuffix(path, ".js") {
			contentType = "application/javascript"
		} else {
			contentType = "text/plain"
		}

		f, err := os.Open(path)

		if err == nil {
			defer f.Close()
			w.Header().Add("Content Type", contentType)
			br := bufio.NewReader(f)
			br.WriteTo(w)
		} else {
			w.WriteHeader(404)
		}
	}
}
