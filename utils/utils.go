package utils

import (
	"encoding/json"
	"net/http"

	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

func ValidateToken(r *http.Request) (bool) {
	var IFTTT_SERVICE_KEY = "J-0JA9dOUROSWiOutUGYOR6wLVzIKKXv22MBPVoRqc4x0C1fuckOaZPJl7O9A320"
	return r.Header.Get("IFTTT-Channel-Key") != IFTTT_SERVICE_KEY
} 

func Message(status bool, message string) (map[string]interface{}) {
	return map[string]interface{} {"status" : status, "message" : message}
}

func Respond(w http.ResponseWriter, data interface{})  {
	w.Header().Add("Content-Type", "application/json; charset=utf-8")
	json.NewEncoder(w).Encode(data)
}

func ErrorRespond(w http.ResponseWriter, data string)  {
	w.Header().Add("Content-Type", "application/json; charset=utf-8")
	http.Error(w, string(data), 401)
}