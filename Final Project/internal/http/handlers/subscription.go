package handlers

import (
	"final-project/internal/helpers"
	render2 "final-project/internal/http/render"
	"final-project/internal/models"
	pdf2 "final-project/internal/service/pdf"
	"final-project/internal/shared"
	"final-project/internal/store"
	"fmt"
	"net/http"
	"strconv"
)

func (a *Repository) ChooseSubscription(w http.ResponseWriter, r *http.Request) {
	user, ok := a.App.Session.Get(r.Context(), "user").(*store.User)
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

	render2.Template(w, r, shared.Plans, &render2.TemplateData{
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
	plan, err := a.Store.GetPlanByID(id)
	if err != nil {
		helpers.ClientError(w, http.StatusBadRequest)
		a.App.Session.Put(r.Context(), shared.Error, "can't get the plan")
		http.Redirect(w, r, "/member/plans", http.StatusSeeOther)
		return
	}

	// get the user
	user, ok := a.App.Session.Get(r.Context(), "user").(store.User)
	if !ok {
		a.App.Session.Put(r.Context(), shared.Warning, "You are not authorized to access this resource")
		//helpers.ClientError(w, http.StatusForbidden)
		http.Redirect(w, r, "/login", http.StatusSeeOther)
		return
	}

	// generate invoice & send an email with an invoice attached
	a.App.Wait.Add(1)

	go func() {
		defer a.App.Wait.Done()

		invoice, err := getInvoice(user, plan)
		if err != nil {
			a.App.ErrorChan <- err
		}

		msg := models.Message{
			To:       user.Email,
			Subject:  "Your invoice",
			Data:     invoice,
			Template: "invoice",
		}

		helpers.SendEmail(msg)
	}()

	// send an email with the manual attached
	a.App.Wait.Add(1)

	go func() {
		defer a.App.Wait.Done()

		// generate a manual
		pdf := pdf2.GenerateManual(user, plan)

		err := pdf.OutputFileAndClose(fmt.Sprintf("./tmp/%d_manual.pdf", user.ID))
		if err != nil {
			a.App.ErrorChan <- err
			return
		}

		msg := models.Message{
			To:            user.Email,
			Subject:       "Your manual",
			Data:          "Your user manual is attached.",
			AttachmentMap: map[string]string{"Manual.PDF": fmt.Sprintf("./tmp/%d_manual.pdf", user.ID)},
		}

		helpers.SendEmail(msg)

	}()

	// subscribe the user to an account
	err = a.Store.SubscribeUserToPlan(user, *plan)
	if err != nil {
		a.App.Session.Put(r.Context(), shared.Error, "Error subscribing to plan")
		http.Redirect(w, r, "/member/plans", http.StatusSeeOther)
		return
	}

	// update the user to update the session
	user.Plan = plan
	a.App.Session.Put(r.Context(), "user", user)

	// redirect
	a.App.Session.Put(r.Context(), shared.Flash, "Subscribed!")
	http.Redirect(w, r, "/member/plans", http.StatusSeeOther)
}

func getInvoice(user store.User, plan *store.Plan) (string, error) {
	return plan.PlanAmountFormatted, nil
}
