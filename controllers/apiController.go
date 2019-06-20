package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/jinzhu/gorm"
	
	"github.com/duyetdev/go_ifttt/models"
	u "github.com/duyetdev/go_ifttt/utils"

)

var Btn1 = func(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	event := models.Event{}

	err := json.NewDecoder(r.Body).Decode(&event)
	if err != nil {
		u.Respond(w, u.Message(false, "Error while decoding request body"))
		return
	}
	defer r.Body.Close()

	if err := db.Save(&event).Error; err != nil {
		u.Respond(w, u.Message(false, err.Error()))
		return
	}

	response, err := json.Marshal(event)
	u.Respond(w, u.Message(true, string(response)))
}