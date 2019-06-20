package controllers

import (
	"encoding/json"
	"net/http"
	
	"github.com/jinzhu/gorm"

	u "github.com/duyetdev/go_ifttt/utils"
)

var TestSetup = func(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	if u.ValidateToken(r) != false {
		u.ErrorRespond(w, string("IFTTT-Channel-Key is invalid"))
		return
	}

	byt := []byte(`{
		"data": {
			"samples": {
				"actionRecordSkipping": {
					"create_new_thing": {
						"invalid": "true"
					}
				}
			}
		}
	}`)

	var data interface{}
	json.Unmarshal(byt, &data)

	u.Respond(w, data)
}

var Status = func(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	if u.ValidateToken(r) != false {
		u.ErrorRespond(w, string("IFTTT-Channel-Key is invalid"))
		return
	}

	u.Respond(w, u.Message(true, "ok"))
}