package route_api

import (
	"encoding/json"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"mysms4you.com/app/src/repo/das/access"
)

/**
 * APIs Methods
 */
func GetV1Main() func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	return func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
		w.Header().Set("Content-Type", "application/json")
		statusList, _ := access.GetHomeList(r.Context())
		json.NewEncoder(w).Encode(statusList)
	}
}

func GetHelloName() func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode("Hello " + ps.ByName("name"))
	}
}
