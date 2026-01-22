package global_router

import "net/http"

// Preflight Request Reusable Function
func HandlePreflightReq(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodOptions {
		w.WriteHeader(200)
	}
}
