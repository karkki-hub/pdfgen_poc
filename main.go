package main

import (
	"fmt"
	"log"
	"os"

	"time"

	cdpdf "pdf_poc/chromedp"
	gapdf "pdf_poc/gpdf"
	mapdf "pdf_poc/maroto"
	pcpdf "pdf_poc/pdfcpu"
)

func main() {

	os.MkdirAll("output", os.ModePerm)

	fmt.Println("Generating maroto PDF...")
	nowm := time.Now()

	// ── 1. Invoice ────────────────────────────────────────────────────────────
	invoiceData := mapdf.InvoiceData{
		Number: "123456",
		Date:   time.Date(2030, 5, 24, 0, 0, 0, 0, time.UTC),
		Provider: mapdf.PartyInfo{
			Name:    "STUDIO SHODWE",
			Address: "123 Anywhere St., Any City",
			City:    "ST 12345",
			Phone:   "+123-456-7890",
			Email:   "hello@reallygreatsite.com",
		},
		Client: mapdf.PartyInfo{
			Name:    "Rachel Beaudry",
			Address: "123 Anywhere St., Any City",
			City:    "ST 12345",
			Phone:   "+123-456-7890",
			Email:   "hello@reallygreatsite.com",
		},
		Items: []mapdf.InvoiceItem{
			{"Service 1", 100.00, 1},
			{"Service 2", 150.00, 1},
			{"Service 3", 200.00, 1},
		},
		TaxRate: 0.06,
		Notes:   "Payment is due within 15 days\nof receiving this invoice.",
		PaymentMethod: mapdf.PaymentInfo{
			Bank:          "Borcelle Bank",
			AccountName:   "Studio Shodwe",
			AccountNumber: "1234567890",
		},
		PreparedBy: mapdf.PreparedByInfo{
			Name:  "Benjamin Shah",
			Title: "Sales Administrator, Studio Shodwe",
		},
	}
	if err := mapdf.GenerateInvoice("output/invoice-maroto.pdf", invoiceData); err != nil {
		log.Fatal("invoice:", err)
	}

	// ── 2 & 3. Services Agreement (p.1) + Payment Plan (p.2) — single PDF ────
	fullAgreement := mapdf.FullAgreementData{
		// Page 1 — Services Agreement
		State:           "California",
		Day:             "1st",
		Month:           "January",
		Year:            "25",
		ProviderName:    "Acme Services Inc.",
		ProviderAddress: "123 Main Street, Los Angeles, CA 90001",
		BuyerName:       "Globex Corp.",
		BuyerAddress:    "456 Market Street, San Francisco, CA 94105",
		Services: []mapdf.ServiceItem{
			{"Web Development", "2", "5,000"},
			{"UI/UX Design", "1", "3,000"},
			{"SEO Package", "3", "1,200"},
			{"Maintenance (monthly)", "6", "800"},
		},
		PurchasePrice: "20,000",
		Notes:         "All work will be delivered digitally. Revisions are limited to 3 rounds per project.",

		// Page 2 — Payment Plan
		Payer:           "John Doe",
		Payee:           "Jane Smith",
		Product:         "Web Development Services",
		AmountPerPeriod: "$500",
		Interval:        "month",
		TotalAmount:     "$3,000",
		Payments: []mapdf.PaymentEntry{
			{"1 Feb 2025", "$500"},
			{"1 Mar 2025", "$500"},
			{"1 Apr 2025", "$500"},
			{"1 May 2025", "$500"},
			{"1 Jun 2025", "$500"},
			{"1 Jul 2025", "$500"},
		},
		LateFee:         "$50",
		BounceFee:       "$75",
		LenderAction:    "contact a debt collection service",
		TermsConditions: "No refunds after work commencement. Disputes subject to California jurisdiction.",
	}
	if err := mapdf.GenerateFullAgreement("output/agreement-maroto.pdf", fullAgreement); err != nil {
		log.Fatal("full agreement:", err)
	}

	badgeData := mapdf.GPayBadgeData{
		BusinessName: "Your Business Name",
		PhoneNumber:  "+91 12345 67890",
		UPIHandle:    "12345 67890@yhh",
		QRCodePath:   "qr.png", // Set to your QR image path, e.g. "qr.png"
		// QRContent: "upi://pay?pa=1234567890@yhh&pn=Your+Business+Name&cu=INR",
	}
	if err := mapdf.GenerateGPayBadge("output/badge-maroto.pdf", badgeData); err != nil {
		log.Fatal("gpay badge:", err)
	}
	runtimem := time.Since(nowm)
	fmt.Printf("✅ maroto PDFs generated successfully in %v!\n", runtimem)

	fmt.Println("Generating pdfcpu PDF...")

	nowp := time.Now()

	Idata := pcpdf.InvoiceData{
		InvoiceNumber: "#123456",
		InvoiceDate:   "24/05/2030",
		CompanyName:   "STUDIO SHODWE",
		CompanyAddr:   "123 Anywhere St., Any City, ST 12345",
		CompanyPhone:  "+123-456-7890",
		CompanyEmail:  "hello@reallygreatsite.com",
		BillToName:    "Rachel Beaudry",
		BillToAddr:    "123 Anywhere St., Any City, ST 12345",
		BillToPhone:   "+123-456-7890",
		BillToEmail:   "hello@reallygreatsite.com",
		Items: []pcpdf.InvoiceItem{
			{"Service 1", 100.00, 1, 100.00},
			{"Service 2", 150.00, 1, 150.00},
			{"Service 3", 200.00, 1, 200.00},
		},
		SubTotal:      450.00,
		TaxRate:       6,
		TaxAmount:     36.00,
		Total:         486.00,
		Notes:         "Payment is due within 15 days\nof receiving this invoice.",
		BankName:      "Borcelle Bank",
		AccountName:   "Studio Shodwe",
		AccountNumber: "1234567890",
		PreparedBy:    "Benjamin Shah",
		PreparedTitle: "Sales Administrator, Studio Shodwe",
	}

	if err := pcpdf.GenerateInvoice("output/invoice-pdfcpu.pdf", Idata); err != nil {
		log.Fatalf("Failed to generate invoice: %v", err)
	}
	Bdata := pcpdf.BadgeData{
		BusinessName: "Your Business Name",
		PhoneNumber:  "+91 12345 67890",
		UPIHandle:    "12345 67890@yhh",
		QRContent:    "upi://pay?pa=1234567890@yhh&pn=Your+Business+Name&cu=INR",
	}

	// Generate Payment Badge PDF
	if err := pcpdf.GenerateGPayBadge("output/badge-pdfcpu.pdf", Bdata); err != nil {
		log.Fatalf("Failed to generate payment badge: %v", err)
	}

	Adata := pcpdf.AgreementData{
		State:           "California",
		Day:             "1st",
		Month:           "January",
		Year:            "2030",
		ProviderName:    "Studio Shodwe LLC",
		ProviderAddress: "123 Anywhere St., Any City, ST 12345",
		BuyerName:       "Rachel Beaudry",
		BuyerAddress:    "456 Other St., Other City, ST 67890",
		Services: []pcpdf.ServiceItem{
			{"Brand Identity Design", "2", "$500"},
			{"Website Mockup", "1", "$800"},
			{"Social Media Kit", "3", "$250"},
		},
		TotalPrice:     "2,300.00",
		TaxResponsible: "Service Provider",
		PaymentMethod:  "Wire transfer",
		PayerName:      "Rachel Beaudry",
		PayeeName:      "Studio Shodwe LLC",
		PayerUPI:       "studioghodwe@bank",
		Interval:       "monthly",
		TotalAmount:    "$2,300.00",
		PaymentDates: []string{
			"Feb 1, 2030 — $384",
			"Mar 1, 2030 — $384",
			"Apr 1, 2030 — $383",
			"May 1, 2030 — $383",
			"Jun 1, 2030 — $383",
			"Jul 1, 2030 — $383",
		},
		LateFee:      "$25",
		BounceFee:    "$50",
		LenderAction: "contacting a debt collection service",
		Terms:        "Governing law: State of California. Dispute resolution: Arbitration.",
	}

	// Generate Two-Page Services Agreement PDF
	if err := pcpdf.GenerateAgreement("output/agreement-pdfcpu.pdf", Adata); err != nil {
		log.Fatalf("Failed to generate agreement: %v", err)
	}

	runtimep := time.Since(nowp)
	fmt.Printf("✅ pdfcpu PDFs generated successfully in %v!\n", runtimep)

	fmt.Println("Generating gpdf PDF...")

	datab := gapdf.FullAgreementData{
		// ── Page 1: Services Agreement ────────────────────────────────────────
		State:           "California",
		Day:             "1st",
		Month:           "January",
		Year:            "25",
		ProviderName:    "Acme Services Inc.",
		ProviderAddress: "123 Main Street, Los Angeles, CA 90001",
		BuyerName:       "Globex Corporation",
		BuyerAddress:    "456 Market Street, San Francisco, CA 94105",
		Services: []gapdf.ServiceItem{
			{Description: "Web Development", NumProjects: "2", PricePerProject: "5,000"},
			{Description: "UI/UX Design", NumProjects: "1", PricePerProject: "3,000"},
			{Description: "SEO Optimization", NumProjects: "3", PricePerProject: "1,200"},
			{Description: "Monthly Maintenance", NumProjects: "6", PricePerProject: "800"},
		},
		PurchasePrice: "20,000",
		Notes:         "All deliverables will be provided digitally. Revisions are limited to 3 rounds per project.",

		// ── Page 2: Payment Plan ──────────────────────────────────────────────
		Payer:           "Globex Corporation",
		Payee:           "Acme Services Inc.",
		Product:         "Web Development & Design Services",
		AmountPerPeriod: "$3,334",
		Interval:        "month",
		TotalAmount:     "$10,000",
		Payments: []gapdf.PaymentEntry{
			{Date: "1 February 2025", Amount: "$3,334"},
			{Date: "1 March 2025", Amount: "$3,334"},
			{Date: "1 April 2025", Amount: "$3,332"},
		},
		LateFee:         "$50",
		BounceFee:       "$75",
		LenderAction:    "engage a debt collection service and pursue legal remedies",
		TermsConditions: "No refunds after commencement of work. All disputes are subject to California jurisdiction.",
	}

	nowg := time.Now()

	out := "output/agreement-gpdf.pdf"
	if err := gapdf.GenerateFullAgreement(out, datab); err != nil {
		log.Fatalf("error: %v", err)
	}
	runtimeg := time.Since(nowg)
	fmt.Printf("✅ gpdf PDFs generated successfully in %v!\n", runtimeg)

	cdinvoiceData := cdpdf.InvoiceData{
		Number: "123456",
		Date:   time.Date(2030, 5, 24, 0, 0, 0, 0, time.UTC),
		Provider: cdpdf.PartyInfo{
			Name:    "STUDIO SHODWE",
			Address: "123 Anywhere St., Any City",
			City:    "ST 12345",
			Phone:   "+123-456-7890",
			Email:   "hello@reallygreatsite.com",
		},
		Client: cdpdf.PartyInfo{
			Name:    "Rachel Beaudry",
			Address: "123 Anywhere St., Any City",
			City:    "ST 12345",
			Phone:   "+123-456-7890",
			Email:   "hello@reallygreatsite.com",
		},
		Items: []cdpdf.InvoiceItem{
			{"Service 1", 100.00, 1},
			{"Service 2", 150.00, 1},
			{"Service 3", 200.00, 1},
		},
		TaxRate: 0.06,
		Notes:   "Payment is due within 15 days\nof receiving this invoice.",
		PaymentMethod: cdpdf.PaymentInfo{
			Bank:          "Borcelle Bank",
			AccountName:   "Studio Shodwe",
			AccountNumber: "1234567890",
		},
		PreparedBy: cdpdf.PreparedByInfo{
			Name:  "Benjamin Shah",
			Title: "Sales Administrator, Studio Shodwe",
		},
	}

	cdbadgeData := cdpdf.GPayBadgeData{
		BusinessName: "Your Business Name",
		PhoneNumber:  "+91 12345 67890",
		UPIHandle:    "12345 67890@yhh",
		QRContent:    "upi://pay?pa=1234567890@yhh&pn=Your+Business+Name&cu=INR",
	}

	cdfullAgreement := cdpdf.FullAgreementData{
		AgreementData: cdpdf.AgreementData{
			State:           "California",
			Day:             "1st",
			Month:           "January",
			Year:            "25",
			ProviderName:    "Acme Services Inc.",
			ProviderAddress: "123 Main Street, Los Angeles, CA 90001",
			BuyerName:       "Globex Corp.",
			BuyerAddress:    "456 Market Street, San Francisco, CA 94105",
			Services: []cdpdf.ServiceItem{
				{"Web Development", "2", "5,000"},
				{"UI/UX Design", "1", "3,000"},
				{"SEO Package", "3", "1,200"},
				{"Maintenance (monthly)", "6", "800"},
			},
			PurchasePrice: "20,000",
			Notes:         "All work will be delivered digitally. Revisions are limited to 3 rounds per project.",
		},
		PaymentPlanData: cdpdf.PaymentPlanData{
			Payer:           "John Doe",
			Payee:           "Jane Smith",
			Product:         "Web Development Services",
			AmountPerPeriod: "$500",
			Interval:        "month",
			TotalAmount:     "$3,000",
			Payments: []cdpdf.PaymentEntry{
				{"1 Feb 2025", "$500"},
				{"1 Mar 2025", "$500"},
				{"1 Apr 2025", "$500"},
				{"1 May 2025", "$500"},
				{"1 Jun 2025", "$500"},
				{"1 Jul 2025", "$500"},
			},
			LateFee:         "$50",
			BounceFee:       "$75",
			LenderAction:    "contact a debt collection service",
			TermsConditions: "No refunds after work commencement. Disputes subject to California jurisdiction.",
		},
	}

	fmt.Println("Generating chromedp PDF...")

	now := time.Now()

	invoiceHTML, err := cdpdf.RenderHTMLTemplate("chromedp/template/invoice.html", cdinvoiceData)
	if err != nil {
		log.Fatal("chromedp invoice template:", err)
	}
	if err := cdpdf.GeneratePDF(invoiceHTML, "output/invoice-chromedp.pdf"); err != nil {
		log.Fatal("chromedp invoice pdf:", err)
	}

	badgeHTML, err := cdpdf.RenderHTMLTemplate("chromedp/template/badge.html", cdbadgeData)
	if err != nil {
		log.Fatal("chromedp badge template:", err)
	}
	if err := cdpdf.GeneratePDF(badgeHTML, "output/badge-chromedp.pdf"); err != nil {
		log.Fatal("chromedp badge pdf:", err)
	}

	agreementHTML, err := cdpdf.RenderHTMLTemplate("chromedp/template/agreement.html", cdfullAgreement)
	if err != nil {
		log.Fatal("chromedp agreement template:", err)
	}
	if err := cdpdf.GeneratePDF(agreementHTML, "output/agreement-chromedp.pdf"); err != nil {
		log.Fatal("chromedp agreement pdf:", err)
	}

	runtime := time.Since(now)
	fmt.Printf("✅ chromedp PDFs generated successfully in %v!\n", runtime)

	fmt.Println("✅ All PDFs generated successfully!")
}
