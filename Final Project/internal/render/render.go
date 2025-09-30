package render

import (
	"fmt"
	"html/template"
	"net/http"
)

func Template(w http.ResponseWriter, r *http.Request, t string, td *TemplateData) {
	partials := []string{
		fmt.Sprintf("%s/base.layout.gohtml", PathToTemplate),
		fmt.Sprintf("%s/header.partial.gohtml", PathToTemplate),
		fmt.Sprintf("%s/navbar.partial.gohtml", PathToTemplate),
		fmt.Sprintf("%s/footer.partial.gohtml", PathToTemplate),
		fmt.Sprintf("%s/alerts.partial.gohtml", PathToTemplate),
	}

	var templateSlice []string
	templateSlice = append(templateSlice, fmt.Sprintf("%s/%s", PathToTemplate, t))

	for _, partial := range partials {
		templateSlice = append(templateSlice, partial)
	}

	if td == nil {
		td = &TemplateData{}
	}

	tmpl, err := template.ParseFiles(templateSlice...)
	if err != nil {
		// error log
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err := tmpl.Execute(w, AddDefaultData(td, r)); err != nil {
		// error log
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

}
