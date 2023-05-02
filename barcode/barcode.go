package barcode

import (
	"bytes"
	"fmt"
	"image"
	"image/color"
	"image/draw"
	"image/png"
	"io"
	"os"

	"github.com/boombuler/barcode"
	"github.com/boombuler/barcode/code128"
	"github.com/golang/freetype"
	"github.com/golang/freetype/truetype"
	"golang.org/x/image/font/gofont/goregular"
)

type Product struct {
	SKU   string
	Name  string
	Price string
}

type Barcode struct {
	Width    int
	Height   int
	Margin   int
	Products []Product
	FileName string
}

func (bc *Barcode) Generate() {
	// Set the size and spacing of each barcode

	// Calculate the total size of the image
	totalWidth := len(bc.Products)*(bc.Width+bc.Margin) - bc.Margin
	totalHeight := bc.Height + 60 // Add space for the headers

	// Create the image
	img := image.NewRGBA(image.Rect(0, 0, totalWidth, totalHeight))

	// Fill the image with a white background
	draw.Draw(img, img.Bounds(), &image.Uniform{color.White}, image.Point{}, draw.Src)

	// Draw the headers for each product
	headerX := 0
	footerY := bc.Height + 40
	for _, p := range bc.Products {
		headerText := p.Name
		footerText := p.Price + " - " + p.SKU
		drawData(img, headerText, headerX, 10)
		drawData(img, footerText, headerX, footerY)
		headerX += bc.Width + bc.Margin
	}

	// Generate each barcode and draw it on the image
	x := 0
	for _, p := range bc.Products {
		code128, err := code128.Encode(p.SKU)
		if err != nil {
			panic(err)
		}

		barcodeImg, err := barcode.Scale(code128, bc.Width, bc.Height)
		if err != nil {
			panic(err)
		}

		draw.Draw(img, image.Rect(x, 30, x+bc.Width, bc.Height+30), barcodeImg, image.Point{}, draw.Src)
		x += bc.Width + bc.Margin
	}

	// Save the image as a PNG file
	file, err := os.Create(bc.FileName)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	png.Encode(file, img)
}

func drawData(img *image.RGBA, headerText string, x, y int) {
	fmt.Printf("x:%d , y:%d\r\n", x, y)
	c := freetype.NewContext()
	c.SetDPI(100)
	fontBytes := goregular.TTF
	fontReader := bytes.NewReader(fontBytes)
	fontBytes, err := io.ReadAll(fontReader)
	if err != nil {
		panic(err)
	}
	font, err := truetype.Parse(fontBytes)
	if err != nil {
		panic(err)
	}

	c.SetFont(font)
	c.SetFontSize(14)
	c.SetClip(img.Bounds())
	c.SetDst(img)
	c.SetSrc(image.Black)

	pt := freetype.Pt(x+50, y+10)
	_, err = c.DrawString(headerText, pt)
	if err != nil {
		panic(err)
	}
}
