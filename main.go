package main

import (
    "bytes"
    "fmt"
    "io/ioutil"
    "log"

    "github.com/jung-kurt/gofpdf"
    // "golang.org/x/image/font/gofont/goregular" // replace with your desired TTF font
    // "golang.org/x/image/font/opentype"
)

func main() {
    // Load TTF font file
    ttfBytes, err := ioutil.ReadFile("Gashington Classy.ttf") // replace with your desired TTF font file path
    if err != nil {
        log.Fatalf("failed to read font file: %v", err)
    }

    // Create PDF document
    pdf := gofpdf.New("P", "mm", "A4", "")

    // Add a page
    pdf.AddPage()

    // Set font
    pdf.AddUTF8FontFromBytes("GoRegular", "", ttfBytes) // register the font in PDF document
    pdf.SetFont("GoRegular", "", 24) // replace with your desired font name and size

    // Set row and column count
    rowCount := 13
    colCount := 8

    _, lineHt := pdf.GetFontSize()                                 // get the line height of the font
    cellWidth, cellHeight := 20.0, 20.0                            // set cell height and width
    // 210 * 297 are the A4 Dimensions
    xMargin := (210.0 - (float64(colCount) * cellWidth)) / 2.0      // calculate horizontal margin
    yMargin := (297.0 - (float64(rowCount) * cellHeight)) / 2.0 - lineHt // calculate vertical margin
    pdf.SetMargins(xMargin, yMargin, xMargin)
    pdf.SetDrawColor(0, 0, 0)
    pdf.SetFillColor(255, 255, 255)

    // Set opacity
    pdf.SetAlpha(1.0, "Normal")

    // Set opacity and add letter to worksheet for each row and column
    pdf.Ln(-1)
    for i := 0; i < rowCount; i++ {
        for j := 0; j < colCount; j++ {
            // Set opacity for rows 2-12
            if i > 0 {
                pdf.SetAlpha(0.5, "Normal")
            }

            // Calculate x and y position for cell
            xPos := float64(j) * cellWidth
            yPos := float64(i) * cellHeight

            // Add letter to cell
            letter := "A" // replace with your desired letter
            pdf.CellFormat(cellWidth, cellHeight, letter, "1", 0, "C", false, 0, fmt.Sprintf("x=%.2f, y=%.2f", xPos, yPos))
        }
        pdf.Ln(-1) // move to next row
    }

    // Save PDF document
    var buffer bytes.Buffer
    err = pdf.Output(&buffer)
    if err != nil {
        log.Fatalf("failed to generate PDF: %v", err)
    }

    err = ioutil.WriteFile("worksheet.pdf", buffer.Bytes(), 0644) // replace with your desired PDF file path
    if err != nil {
        log.Fatalf("failed to write PDF file: %v", err)
    }

    fmt.Println("Worksheet created successfully!")
}