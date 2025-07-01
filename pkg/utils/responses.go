package utils

import (
	"encoding/json"
	"net/http"

	mod "github.com/smartineztri_meli/W17-G2-Bootcamp/pkg/models"
)

func BadResponse(w http.ResponseWriter, code int, message string) {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(code)
	resp, _ := json.Marshal(mod.Response{Success: false, Message: message})
	w.Write(resp)
}
func GoodResponse(w http.ResponseWriter, code int, message string, data interface{}) {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(code)
	resp, _ := json.Marshal(mod.Response{Success: true, Message: message, Data: data})
	w.Write(resp)
}
