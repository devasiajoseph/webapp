package pdf

import (
	"github.com/jung-kurt/gofpdf"
)

func ToPdf() {
	// Create a new PDF document with portrait orientation and millimeter units
	pdf := gofpdf.New("P", "mm", "A4", "")

	// Add a new page to the document
	pdf.AddPage()

	// Set the font family and size
	pdf.SetFont("Arial", "B", 16)

	// Add the invoice title
	pdf.Cell(40, 10, "Invoice")

	// Set the font size and style for the table headers
	pdf.SetFont("Arial", "B", 12)

	// Define the table headers
	headers := []string{"Item", "Description", "Quantity", "Price", "Total"}

	// Set the width and height of the table cells
	w := []float64{30.0, 80.0, 30.0, 30.0, 30.0}
	h := 7.0

	// Add the table headers
	for _, header := range headers {
		pdf.CellFormat(w[0], h, header, "1", 0, "C", false, 0, "")
	}
	pdf.Ln(-1)

	// Set the font size and style for the table data
	pdf.SetFont("Arial", "", 10)

	// Define the table data
	data := [][]string{
		{"1", "Item 1", "1", "$10.00", "$10.00"},
		{"2", "Item 2", "2", "$20.00", "$40.00"},
		{"3", "Item 3", "3", "$30.00", "$90.00"},
	}

	// Add the table data
	for _, row := range data {
		for i, cell := range row {
			pdf.CellFormat(w[i], h, cell, "1", 0, "", false, 0, "")
		}
		pdf.Ln(-1)
	}

	// Add the invoice total
	pdf.CellFormat(w[0]+w[1]+w[2]+w[3], h, "Total", "1", 0, "R", false, 0, "")
	pdf.CellFormat(w[4], h, "$140.00", "1", 0, "", false, 0, "")

	// Output the PDF document to a file
	err := pdf.OutputFileAndClose("invoice.pdf")
	if err != nil {
		panic(err)
	}
}
