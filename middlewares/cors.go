package middlewares

import (
	"net/http"
	"github.com/julienschmidt/httprouter"
)

func Cors(next httprouter.Handle) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		header := w.Header()
		header.Set("Access-Control-Allow-Origin", "*")
		header.Set("Access-Control-Allow-Credentials", "true")
		header.Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, Accept, Origin, Cache-Control, X-Requested-With")
		header.Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")

		next(w, r, ps)
	}
}