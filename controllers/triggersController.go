package controllers

import (
	"net/http"
	"encoding/json"

	"github.com/jinzhu/gorm"
	// "github.com/gorilla/mux"

	"github.com/duyetdev/go_ifttt/models"
	u "github.com/duyetdev/go_ifttt/utils"
)

type IFTTTPayload struct {
	trigger_identity 	string
	triggerFields 		interface{}
	limit 				int
	user 				interface{}
	ifttt_source  		interface{}
}

var TriggerBtn1 = func(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	var body IFTTTPayload
	err := json.NewDecoder(r.Body).Decode(&body)
	if err != nil {
		u.Respond(w, u.Message(false, "Error while decoding request body"))
		return
	}
	defer r.Body.Close()


	if u.ValidateToken(r) != false {
		u.ErrorRespond(w, string("IFTTT-Channel-Key is invalid"))
		return
	}

	events := []models.Event{}
	db.Limit(body.limit).Find(&events)
	resp := map[string]interface{} { "data": events }

	u.Respond(w, resp)
}