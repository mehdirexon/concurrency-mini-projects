package handlers

import (
	"final-project/internal/database"
	"final-project/internal/helpers"
	"final-project/internal/render"
	"final-project/internal/shared"
	"fmt"
	"net/http"
	"strconv"
)

func (a *Repository) ChooseSubscription(w http.ResponseWriter, r *http.Request) {
	user, ok := a.App.Session.Get(r.Context(), "user").(*database.User)
	if ok {
		print(user)
	}
	plans, err := a.Store.GetAllPlans()
	if err != nil {
		a.App.ErrorLogger.Println(err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	dataMap := make(map[string]any)
	dataMap["plans"] = plans

	render.Template(w, r, shared.Plans, &render.TemplateData{
		Data: dataMap,
	})

}

func (a *Repository) Subscribe(w http.ResponseWriter, r *http.Request) {
	// get the id
	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil {
		helpers.ClientError(w, http.StatusBadRequest)
		a.App.Session.Put(r.Context(), shared.Error, "can't get the id")
		http.Redirect(w, r, "/member/plans", http.StatusSeeOther)
		return
	}

	// get the plan
	plans, err := a.Store.GetPlanByID(id)
	if err != nil {
		helpers.ClientError(w, http.StatusBadRequest)
		a.App.Session.Put(r.Context(), shared.Error, "can't get the plan")
		http.Redirect(w, r, "/member/plans", http.StatusSeeOther)
		return
	}

	// get the user
	user, ok := a.App.Session.Get(r.Context(), "user").(database.User)
	if !ok {
		a.App.Session.Put(r.Context(), shared.Warning, "You are not authorized to access this resource")
		//helpers.ClientError(w, http.StatusForbidden)
		http.Redirect(w, r, "/login", http.StatusSeeOther)
		return
	}
	fmt.Print(user)
	fmt.Print(plans)
	// generate invoice

	// send an email with an invoice attached

	// generate a manual

	// send an email with the manual attached

	// subscribe the user to an account

	// redirect

}
