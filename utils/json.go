package utils

import (
	"encoding/json"
	"log"
	"net/http"
)

func SendMessageAsJsonResponse(w http.ResponseWriter, msg string, code int) {
	rsp := map[string]string {
		"message": msg,
	}

	w.WriteHeader(code)
	//This error can be ignored because a map[string]string will always be converted to json
	if err := json.NewEncoder(w).Encode(rsp); err != nil {
		log.Println("ERROR: Failed to send response:", err)
	}
}
