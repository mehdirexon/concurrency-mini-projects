package render

import (
	"final-project/internal/config"
	"final-project/internal/helpers"
	"final-project/internal/shared"
	"net/http"
	"time"
)

var PathToTemplate = "./templates"
var app *config.AppConfig

func Register(a *config.AppConfig) {
	app = a
}

type TemplateData struct {
	StringMap     map[string]string
	IntMap        map[string]int
	FloatMap      map[string]float64
	Data          map[string]interface{}
	Flash         string
	Error         string
	Warning       string
	Authenticated bool
	Now           time.Time
	//User	*data.user
}

func AddDefaultData(td *TemplateData, r *http.Request) *TemplateData {
	td.Flash = app.Session.PopString(r.Context(), shared.Flash)
	td.Warning = app.Session.PopString(r.Context(), shared.Warning)
	td.Error = app.Session.PopString(r.Context(), shared.Error)

	if helpers.IsAuthenticated(r) {
		td.Authenticated = true
		// TODO - get more user information
	}
	td.Now = time.Now()

	return td
}
