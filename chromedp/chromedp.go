package pdf

import (
	"context"
	"fmt"
	"net/url"
	"os"
	"time"

	"github.com/chromedp/cdproto/page"
	"github.com/chromedp/chromedp"
)

func readHTML(path string) (string, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return "", err
	}
	return string(data), nil
}

func generatePDF(html string, output string) error {

	ctx, cancel := chromedp.NewContext(context.Background())
	defer cancel()

	ctx, cancel = context.WithTimeout(ctx, 20*time.Second)
	defer cancel()

	var pdfBuf []byte

	formathtml := `
	<html>
	<head>
		<meta charset="UTF-8">
		<style>
			body { margin: 0; }
		</style>
	</head>
	<body>
		` + html + `
	</body>
	</html>`

	htmlURL := "data:text/html," + url.PathEscape(formathtml)

	err := chromedp.Run(ctx,
		chromedp.Navigate(htmlURL),
		chromedp.Sleep(2*time.Second), // wait for rendering

		chromedp.ActionFunc(func(ctx context.Context) error {
			buf, _, err := page.PrintToPDF().
				WithPrintBackground(true).
				WithPaperWidth(8.27).   // A4 width
				WithPaperHeight(11.69). // A4 height
				Do(ctx)

			pdfBuf = buf
			return err
		}),
	)

	if err != nil {
		return err
	}

	return os.WriteFile(output, pdfBuf, 0644)
}

func cdnew() {

	os.MkdirAll("output", os.ModePerm)

	invoiceHTML, err := readHTML("invoice.html")
	if err != nil {
		panic(err)
	}
	err = generatePDF(invoiceHTML, "output/invoice.pdf")
	if err != nil {
		panic(err)
	}

	badgeHTML, err := readHTML("badge.html")
	if err != nil {
		panic(err)
	}
	err = generatePDF(badgeHTML, "output/badge.pdf")
	if err != nil {
		panic(err)
	}

	agreementHTML, err := readHTML("agreement.html")
	if err != nil {
		panic(err)
	}
	err = generatePDF(agreementHTML, "output/agreement.pdf")
	if err != nil {
		panic(err)
	}

	fmt.Println("✅ All PDFs generated successfully!")
}
