package main

import (
	"encoding/csv"
	"os"
)

type Product struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Attributes  string `json:"attributes"` // JSON string
}

func main() {
	// Sample product data
	products := []Product{
		{"Product A", "Description of Product A", `{"category": "books", "language": "english"}`},
		{"Product B", "Description of Product B", `{"category": "electronics", "language": "english"}`},
		{"Product C", "Description of Product C", `{"category": "books", "language": "french"}`},
		{"Product D", "Description of Product D", `{"category": "clothing", "size": "M"}`},
		{"Product E", "Description of Product E", `{"category": "electronics", "language": "spanish"}`},
		{"Product F", "Description of Product F", `{"category": "books", "language": "german"}`},
		{"Product G", "Description of Product G", `{"category": "clothing", "size": "L"}`},
		{"Product H", "Description of Product H", `{"category": "electronics", "features": ["wireless", "bluetooth"]}`},
		{"Product I", "Description of Product I", `{"category": "books", "language": "italian"}`},
		{"Product J", "Description of Product J", `{"category": "clothing", "size": "S"}`},

		// Additional products
		{"Product K", "Description of Product K", `{"category": "home appliances", "type": "kitchen"}`},
		{"Product L", "Description of Product L", `{"category": "toys", "age_range": "3-5 years"}`},
		{"Product M", "Description of Product M", `{"category": "books", "language": "spanish"}`},
		{"Product N", "Description of Product N", `{"category": "electronics", "features": ["4K resolution"]}`},
		{"Product O", "Description of Product O", `{"category": "clothing", "size": "XL"}`},
		{"Product P", "Description of Product P", `{"category": "furniture", "material": ["wood","metal"]}`},
		{"Product Q", "Description of Product Q", `{"category": "sports equipment","type":"basketball"}`},
		{"Product R", "Description of Product R", `{"category": "gardening","type":"tools"}`},
		{"Product S", "Description of Product S", `{"category": "health","type":"vitamins"}`},
		{"Product T", "Description of Product T", `{"category": "automotive","type":"accessories"}`},

        // More diverse products
        {"Product U",  "Description of Product U","{\"category\": \"music\", \"type\": \"vinyl\"}"},
        {"Product V","Description of Product V","{\"category\": \"software\", \"platform\": \"windows\"}"},
        {"Product W","Description of Product W","{\"category\": \"outdoor\", \"type\": \"camping gear\"}"},
        {"Product X","Description of Product X","{\"category\": \"pet supplies\", \"animal\": \"dog\"}"},
        {"Product Y","Description of Product Y","{\"category\": \"office supplies\", \"type\": \"stationery\"}"},
        {"Product Z","Description of Product Z","{\"category\": \"kitchenware\", \"type\": \"utensils\"}"},
    }

    // Create the CSV file
    file, err := os.Create("products.csv")
    if err != nil {
        panic(err)
    }
    defer file.Close()

    writer := csv.NewWriter(file)
    defer writer.Flush()

    // Write the header
    err = writer.Write([]string{"name", "description", "attributes"})
    if err != nil {
        panic(err)
    }

    // Write the product data
    for _, product := range products {
        record := []string{product.Name, product.Description, product.Attributes}
        err := writer.Write(record)
        if err != nil {
            panic(err)
        }
    }

    println("products.csv generated successfully!")
}