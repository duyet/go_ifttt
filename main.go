package main
import (
  "net/http"
  "log"

  "github.com/gorilla/mux"
  "github.com/jinzhu/gorm"

  "github.com/duyetdev/go_ifttt/controllers"
  "github.com/duyetdev/go_ifttt/models"
)


type App struct {
	Router 				*mux.Router
	DB     				*gorm.DB
}

// App initialize with predefined configuration
func (a *App) Initialize() {
	db, err := gorm.Open("sqlite3", "test.db")
	if err != nil {
		log.Fatal("Could not connect database")
	}

	a.DB = models.DBMigrate(db)
	a.Router = mux.NewRouter()
	a.setRouters()
}

// Set all required routers
func (a *App) setRouters() {
	a.Router.HandleFunc("/ifttt/v1/status", a.handleRequest(controllers.Status)).Methods("GET")
	a.Router.HandleFunc("/ifttt/v1/test/setup", a.handleRequest(controllers.TestSetup)).Methods("POST")
	a.Router.HandleFunc("/ifttt/v1/triggers/btn1", a.handleRequest(controllers.TriggerBtn1)).Methods("POST")
	
	a.Router.HandleFunc("/api/v1/btn1", a.handleRequest(controllers.Btn1)).Methods("POST")
}

type RequestHandlerFunction func(db *gorm.DB, w http.ResponseWriter, r *http.Request)

func (a *App) handleRequest(handler RequestHandlerFunction) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		handler(a.DB, w, r)
	}
}

// Run the app on it's router
func (a *App) Run(host string) {
	log.Print("Listen on http://", host)
	log.Fatal(http.ListenAndServe(host, a.Router))
}

// Main
func main() {
	app := &App{}
	app.Initialize()
	app.Run("0.0.0.0:3000")
}

