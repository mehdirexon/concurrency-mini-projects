package pdf

import (
	"final-project/internal/store"
	"fmt"

	"github.com/phpdave11/gofpdf"
	"github.com/phpdave11/gofpdf/contrib/gofpdi"
)

func GenerateManual(user store.User, plan *store.Plan) *gofpdf.Fpdf {
	pdf := gofpdf.New("P", "mm", "Letter", "")
	pdf.SetMargins(10, 13, 10)

	importer := gofpdi.NewImporter()

	t := importer.ImportPage(pdf, "internal/service/pdf/manual.pdf", 1, "/MediaBox")

	pdf.AddPage()

	importer.UseImportedTemplate(pdf, t, 0, 0, 215.9, 0)

	pdf.SetX(75)
	pdf.SetY(150)

	pdf.SetFont("Arial", "", 12)
	pdf.MultiCell(0, 4, fmt.Sprintf("%s %s", user.FirstName, user.LastName), "", "C", false)

	pdf.Ln(5)

	pdf.MultiCell(0, 4, fmt.Sprintf("%s User Guide", plan.PlanName), "", "C", false)

	return pdf
}
