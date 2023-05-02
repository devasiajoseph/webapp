package barcode

import "testing"

func TestGenerate(t *testing.T) {
	products := []Product{
		Product{Name: "Product 1",
			SKU: "SKU112", Price: "$100"},
		Product{Name: "Product 2",
			SKU: "SKU2232", Price: "$150"},
		Product{Name: "Product 3",
			SKU: "SKU333", Price: "$200"},
	}

	bc := Barcode{Height: 100, Width: 300, FileName: "tmp/test.png", Products: products, Margin: 10}
	bc.Generate()
}
