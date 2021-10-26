package main

import (
	"github.com/signintech/gopdf"
)

func main() {
	// Create object of gopdf
	pdf := gopdf.GoPdf{}
	// Pdf size
	pdf.Start(gopdf.Config{PageSize: *gopdf.PageSizeA4})
	// Import Font
	err := pdf.AddTTFFont("ipaexg", "./ipaexg.ttf")
	if err != nil {
		panic(err)
	}
	err = pdf.SetFont("ipaexg", "", 26)
	if err != nil {
		panic(err)
	}
	// Write string in Pdf
	pdf.AddPage()
	pdf.SetX(80)
	pdf.SetY(120)
	pdf.Cell(nil, "おはよう！今日も一日がんばろうね！")
	pdf.SetX(80)
	pdf.SetY(180)
	pdf.Cell(nil, "コラショ")
	// Extract pdf
	pdf.WritePdf("benesse.pdf")
}