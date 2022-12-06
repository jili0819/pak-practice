package api

import "net/http"

// Write http common headers
func setCommonHeaders(w http.ResponseWriter) {
	// Set the "Server" http header.
	w.Header().Set(ServerInfo, "api_server")
}
