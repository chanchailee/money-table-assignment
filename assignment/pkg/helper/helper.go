package helper

import (
	"net/http"
)

// HandleSuccessResp :
func HandleSuccessResp(w http.ResponseWriter, data []byte) {
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	w.Write(data)
}
