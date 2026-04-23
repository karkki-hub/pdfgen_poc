package main

import (
	"fmt"
	"os"

	pdf "karkki-hub/pdfgen_poc/chromedp/pdf"
)

func main() {

	os.MkdirAll("output", os.ModePerm)

	invoiceHTML, err := pdf.ReadHTML("invoice.html")
	if err != nil {
		panic(err)
	}
	err = pdf.GeneratePDF(invoiceHTML, "output/invoice.pdf")
	if err != nil {
		panic(err)
	}

	badgeHTML, err := pdf.ReadHTML("badge.html")
	if err != nil {
		panic(err)
	}
	err = pdf.GeneratePDF(badgeHTML, "output/badge.pdf")
	if err != nil {
		panic(err)
	}

	agreementHTML, err := pdf.ReadHTML("agreement.html")
	if err != nil {
		panic(err)
	}
	err = pdf.GeneratePDF(agreementHTML, "output/agreement.pdf")
	if err != nil {
		panic(err)
	}

	fmt.Println("✅ All PDFs generated successfully!")
}
